/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"main/pkg"
	"main/pkg/services"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Update wallpaper config",
	Long:  `Update wallpaper config data.`,
	Run: func(cmd *cobra.Command, args []string) {
		interval := cmd.Flags().Lookup("interval").Value.String()
		accessKey := cmd.Flags().Lookup("access_key").Value.String()
		secretKey := cmd.Flags().Lookup("secret_key").Value.String()
		query := cmd.Flags().Lookup("query").Value.String()
		maxImage := cmd.Flags().Lookup("max_image").Value.String()

		if interval == "" && accessKey == "" && secretKey == "" && query == "" && maxImage == "" {
			fmt.Println("Wallpaper: nothing to update")
			return
		}
		config := services.NewConfigService()
		if interval != "" {
			config.Set("config.interval", interval)
		}
		if accessKey != "" {
			config.Set("config.access_key", accessKey)
		}
		if secretKey != "" {
			config.Set("config.secret_key", secretKey)
		}
		if query != "" {
			config.Set("api.query", query)
		}
		if maxImage != "" {
			if pkg.HasLetters(maxImage) {
				log.Fatal("max Image requires a number. e.g 5, 10")
			}
			config.Set("config.max_image", maxImage)
		}
		fmt.Println("Wallpaper: config updated successfully")
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
	configCmd.Flags().StringP("interval", "i", "", "set interval time to change wallpaper, default is 5m")
	configCmd.Flags().StringP("access_key", "a", "", "set unsplash access key")
	configCmd.Flags().StringP("secret_key", "s", "", "set unsplash secret key")
	configCmd.Flags().StringP("query", "q", "", "set unsplash image query")
	configCmd.Flags().StringP("max_image", "m", "", "set max image to download from unsplash")
}
