package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "View or edit configuration",
	Run: func(cmd *cobra.Command, args []string) {
		// Print current configuration
		viper.Debug()
	},
}

func init() { rootCmd.AddCommand(configCmd) }
