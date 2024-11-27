// Package cmd pei qin/*
package cmd

import (
	"github.com/spf13/cobra"
	"nagios-plugins/pkg"
)

var w int

// cpuCmd represents the nagios-plugins command
var cpuCmd = &cobra.Command{
	Use:   "nagios-plugins",
	Short: "监控cpu使用率",
	Long:  `监控cpu使用率 command nagios-plugins -w 20%`,
	Run: func(cmd *cobra.Command, args []string) {
		pkg.MonitorCpuLoad(w)
		//MonitorCpuLoad(w)
		//nagios - plugins.MonitorCpuLoad(w)
	},
}

func init() {
	rootCmd.AddCommand(cpuCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cpuCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cpuCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	cpuCmd.Flags().IntVarP(&w, "nagios-plugins", "w", 20, "监控cpu使用率")
}
