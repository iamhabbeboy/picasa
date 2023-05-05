/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"main/internal/api"

	"github.com/spf13/cobra"
)

type Image struct {
	Results []struct {
		Urls struct {
			Raw  string `json:"raw"`
			Full string `json:"full"`
		}
		Links struct {
			Download string `json:"download"`
		}
	} `json:"results"`
}

// downloadCmd represents the download command
var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download images from unsplash.com",
	Long:  `Download images from unsplash.com`,
	Run: func(cmd *cobra.Command, args []string) {
		// cronWeeklyTask := "0 0 * * 0"
		// if !pkg.HasCronjob(cronWeeklyTask) {
		// 	pkg.SetCronTab(cronWeeklyTask)
		// }
		processImages()
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}

func processImages() {
	svc := api.NewImageDownload("unsplash")
	svc.GetImages("nature")
}
