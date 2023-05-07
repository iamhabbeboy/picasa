/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"main/internal"
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
		HandleDownloadProcess()
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)
}

func HandleDownloadProcess() {
	cronWeeklyTask := internal.CRON_WEEKLY
	r := fmt.Sprintf("/usr/local/bin/%s download", internal.APP_NAME)
	if !internal.HasCronjob(cronWeeklyTask) {
		internal.SetCronTab(cronWeeklyTask, r)
	}
	processImages()
}

func processImages() {
	svc := api.NewImageDownload("unsplash")
	svc.GetImages()
}
