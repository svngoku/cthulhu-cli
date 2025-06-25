package cmd

import (
	"context"

	"cthulhu-cli/internal/aws"
	"cthulhu-cli/internal/kube"
	"github.com/spf13/cobra"
)

var (
	clusterName string
	createEKS   bool
)

var clusterCmd = &cobra.Command{
	Use:   "cluster",
	Short: "List, create, or switch Kubernetes clusters",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		if createEKS {
			return aws.CreateEKS(ctx, clusterName)
		}
		return kube.SwitchContext(clusterName)
	},
}

func init() {
	rootCmd.AddCommand(clusterCmd)
	clusterCmd.Flags().StringVarP(&clusterName, "name", "n", "", "cluster name")
	clusterCmd.Flags().BoolVar(&createEKS, "eks", false, "create AWS EKS cluster if not exists")
}
