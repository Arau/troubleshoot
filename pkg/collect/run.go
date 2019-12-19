package collect

import (
	"encoding/json"
	"errors"
	"time"

	troubleshootv1beta1 "github.com/replicatedhq/troubleshoot/pkg/apis/troubleshoot/v1beta1"
	"github.com/replicatedhq/troubleshoot/pkg/logger"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type RunOutput map[string][]byte

func Run(ctx *Context, runCollector *troubleshootv1beta1.Run) ([]byte, error) {
	client, err := kubernetes.NewForConfig(ctx.ClientConfig)
	if err != nil {
		return nil, err
	}

	pod, err := runPod(client, runCollector)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := client.CoreV1().Pods(pod.Namespace).Delete(pod.Name, &metav1.DeleteOptions{}); err != nil {
			logger.Printf("Failed to delete pod %s: %v\n", pod.Name, err)
		}
	}()

	if runCollector.Timeout == "" {
		return runWithoutTimeout(ctx, pod, runCollector)
	}

	timeout, err := time.ParseDuration(runCollector.Timeout)
	if err != nil {
		return nil, err
	}

	errCh := make(chan error, 1)
	resultCh := make(chan []byte, 1)
	go func() {
		b, err := runWithoutTimeout(ctx, pod, runCollector)
		if err != nil {
			errCh <- err
		} else {
			resultCh <- b
		}
	}()

	select {
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	case result := <-resultCh:
		return result, nil
	case err := <-errCh:
		return nil, err
	}
}

func runWithoutTimeout(ctx *Context, pod *corev1.Pod, runCollector *troubleshootv1beta1.Run) ([]byte, error) {
	client, err := kubernetes.NewForConfig(ctx.ClientConfig)
	if err != nil {
		return nil, err
	}

	for {
		status, err := client.CoreV1().Pods(pod.Namespace).Get(pod.Name, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}
		if status.Status.Phase == corev1.PodRunning ||
			status.Status.Phase == corev1.PodFailed ||
			status.Status.Phase == corev1.PodSucceeded {
			break
		}
		time.Sleep(time.Second * 1)
	}

	runOutput := RunOutput{}

	limits := troubleshootv1beta1.LogLimits{
		MaxLines: 10000,
	}
	podLogs, err := getPodLogs(client, *pod, runCollector.CollectorName, "", &limits, true)

	for k, v := range podLogs {
		runOutput[k] = v
	}

	if ctx.Redact {
		runOutput, err = runOutput.Redact()
		if err != nil {
			return nil, err
		}
	}

	b, err := json.MarshalIndent(runOutput, "", "  ")
	if err != nil {
		return nil, err
	}

	return b, nil
}

func runPod(client *kubernetes.Clientset, runCollector *troubleshootv1beta1.Run) (*corev1.Pod, error) {
	podLabels := make(map[string]string)
	podLabels["troubleshoot-role"] = "run-collector"

	pullPolicy := corev1.PullIfNotPresent
	if runCollector.ImagePullPolicy != "" {
		pullPolicy = corev1.PullPolicy(runCollector.ImagePullPolicy)
	}

	pod := corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      runCollector.CollectorName,
			Namespace: runCollector.Namespace,
			Labels:    podLabels,
		},
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Pod",
		},
		Spec: corev1.PodSpec{
			RestartPolicy: corev1.RestartPolicyNever,
			Containers: []corev1.Container{
				{
					Image:           runCollector.Image,
					ImagePullPolicy: pullPolicy,
					Name:            "collector",
					Command:         runCollector.Command,
					Args:            runCollector.Args,
				},
			},
		},
	}

	created, err := client.CoreV1().Pods(runCollector.Namespace).Create(&pod)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (r RunOutput) Redact() (RunOutput, error) {
	podLogs, err := redactMap(r)
	if err != nil {
		return nil, err
	}

	return podLogs, nil
}
