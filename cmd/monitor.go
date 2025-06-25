package cmd

import (
	"context"

	"cthulhu-cli/internal/kube"
	"github.com/spf13/cobra"
)

var monitorCmd = &cobra.Command{
	Use:   "monitor [deployment]",
	Short: "Tail pod logs & show Prometheus metrics for a deployment",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		return kube.Monitor(ctx, args[0])
	},
}

func init() { rootCmd.AddCommand(monitorCmd) }
