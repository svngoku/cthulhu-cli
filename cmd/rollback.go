package cmd

import (
	"context"

	"cthulhu-cli/internal/kube"
	"github.com/spf13/cobra"
)

var rollbackCmd = &cobra.Command{
	Use:   "rollback [deployment]",
	Short: "Rollback a deployment to the previous revision",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		return kube.Rollback(ctx, args[0])
	},
}

func init() { rootCmd.AddCommand(rollbackCmd) }
