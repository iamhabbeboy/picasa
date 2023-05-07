/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"main/internal"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop all services",
	Long:  `Stop all wallpaper cronjobs`,
	Run: func(cmd *cobra.Command, args []string) {
		if internal.RemoveCronTab("wallpaper") {
			fmt.Println("...Stopped wallpaper cronjob")
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
