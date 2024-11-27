// Package cmd pei qin/*
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	//Use:   "nagios-plugins",
	//Short: "监控cpu使用率",
	//Long:  `监控cpu使用率 command nagios-plugins -w 20%`,
	//Uncomment the following line if your bare application
	//has an action associated with it:
	//Run: func(cmd *cobra.Command, args []string) {
	//	MonitorCpuLoad(w)
	//},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nagios-plugins.yaml)")

	// Cobra also supports local flags, which will only run
	//rootCmd.Flags().IntVarP(&w, "nagios-plugins", "w", 20, "监控cpu使用率")
}
