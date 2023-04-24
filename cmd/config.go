/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A CLI tool that update wallpaper config",
	Long:  `A CLI tool that update wallpaper config data.`,
	Run: func(cmd *cobra.Command, args []string) {
		// stopCron()
		fmt.Println("config called")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
