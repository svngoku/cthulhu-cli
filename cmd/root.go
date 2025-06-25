package cmd

import (
	"fmt"
	"os"

	"cthulhu-cli/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "cthulhu",
	Short: "Cthulhu – DevOps tentacles for Docker→K8s",
	Long:  `Cthulhu wraps common container‑to‑Kubernetes tasks into one intuitive CLI.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cthulhu.yaml)")
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "enable verbose logging")
	viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose"))
}

func initConfig() {
	utils.SetupLogger(viper.GetBool("verbose"))
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigName(".cthulhu")
	}
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		utils.Log().Infow("using config", "file", viper.ConfigFileUsed())
	}
}