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
	Use:   "picasa",
	Short: "Manage your desktop wallpaper with random images from unsplash, pixabay, etc.",
	Long:  `Manage your desktop wallpaper with random images from unsplash, , pixabay, etc`,
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
