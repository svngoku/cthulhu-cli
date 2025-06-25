package cmd

import (
	"context"

	"cthulhu-cli/internal/docker"
	"cthulhu-cli/internal/kube"
	"github.com/spf13/cobra"
)

var (
	imageTag  string
	helmChart string
	namespace string
)

var deployCmd = &cobra.Command{
	Use:   "deploy [path]",
	Short: "Build, push and deploy a Docker image to Kubernetes",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		path := args[0] // Dockerfile directory

		// 1. Build & push image
		img, err := docker.BuildAndPush(ctx, path, imageTag)
		if err != nil {
			return err
		}

		// 2. Deploy via Helm template or raw manifests
		return kube.Deploy(ctx, img, helmChart, namespace)
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	deployCmd.Flags().StringVarP(&imageTag, "tag", "t", "latest", "docker image tag")
	deployCmd.Flags().StringVarP(&helmChart, "chart", "c", "", "helm chart directory")
	deployCmd.Flags().StringVarP(&namespace, "namespace", "n", "default", "kubernetes namespace")
}
