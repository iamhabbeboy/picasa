/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "main",
	Short: "A simple CLI tool that downloads random pictures from [unsplash] and use it as a wallpaper.",
	Long:  `A simple CLI tool that downloads random pictures from [unsplash] and use it as a wallpaper.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.main.yaml)")

	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
