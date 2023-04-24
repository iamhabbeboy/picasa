/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"main/pkg/services"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	API_URL    = "https://api.unsplash.com/"
	ACCESS_KEY = "Nw5jS2P4zr_oO_qbFt_39zyj7QTIMI49vYx5lCzxujY"
	SECRET_KEY = "pseMeAYqR4G1I8cx8vbwkm4HTs1o56NzW6ZiKGHCMNs"
	MAX_IMAGE  = "5"
	QUERY      = "wallpapers"
	IMAGE_DIR  = "images"
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
