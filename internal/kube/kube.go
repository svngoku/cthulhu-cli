package kube

import (
	"context"

	"cthulhu-cli/pkg/utils"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func clientset() (*kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", clientcmd.RecommendedHomeFile)
	if err != nil {
		return nil, err
	}
	return kubernetes.NewForConfig(config)
}

func Deploy(ctx context.Context, image, chart, ns string) error {
	// Use Helm template or K8s API to create/update deployment
	utils.Log().Infow("deploying", "image", image, "ns", ns)
	return nil
}

func Rollback(ctx context.Context, deployment string) error {
	// Execute `kubectl rollout undo` logic via API or shell
	utils.Log().Infow("rolling back", "deployment", deployment)
	return nil
}

func Monitor(ctx context.Context, deployment string) error {
	// Tail pod logs & query Prometheus
	utils.Log().Infow("monitoring", "deployment", deployment)
	return nil
}

func SwitchContext(cluster string) error {
	// Call out to `kubectx` or manipulate kubeconfig
	utils.Log().Infow("switching context", "cluster", cluster)
	return nil
}
