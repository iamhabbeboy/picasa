/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"main/pkg"
	"main/pkg/services"

	log "github.com/sirupsen/logrus"
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
	Short: "A command to download images from unsplash.com",
	Long:  `A command to download images from unsplash.com`,
	Run: func(cmd *cobra.Command, args []string) {
		cronWeeklyTask := "0 0 * * 0"
		if !pkg.HasCronjob(cronWeeklyTask) {
			pkg.SetCronTab(cronWeeklyTask)
		}
		processImages()
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func processImages() {
	svc := services.NewImageServicer("unsplash")
	svc.GetImages("nigeria")
}
