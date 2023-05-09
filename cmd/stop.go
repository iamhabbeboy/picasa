/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"main/internal"
	"os"

	"github.com/spf13/cobra"
)

// stopCmd represents the stop command
var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop all services",
	Long:  `Stop all picasa cronjobs`,
	Run: func(cmd *cobra.Command, args []string) {
		if internal.RemoveCronTab(internal.APP_NAME) {
			app := internal.APP_NAME
			path := fmt.Sprintf("./.%s", app)
			if err := os.RemoveAll(path); err != nil {
				fmt.Println("...Deleting " + app + " config path, manually delete it here: " + path)
			}
			fmt.Print("...Stopped " + internal.APP_NAME + " cronjob")
		}
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
