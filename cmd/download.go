/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

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
		fmt.Println("download called")
		processImage()
	},
}

func init() {
	rootCmd.AddCommand(downloadCmd)

	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
}

func processImage() {
	url := fmt.Sprintf("%s/search/photos?query=%s&per_page=%v&&client_id=%s", API_URL, QUERY, MAX_IMAGE, ACCESS_KEY)
	result := getImage(url)
	var wg sync.WaitGroup
	for key, v := range result.Results {
		wg.Add(1)
		go downloadImage(v.Urls.Full, key, &wg)
	}
	wg.Wait()
}

func getImage(url string) Image {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	var p Image
	if err := json.NewDecoder(resp.Body).Decode(&p); err != nil {
		log.Fatal(err)
	}
	return p
}

func downloadImage(image string, index int, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(image)
	if err != nil {
		log.Fatal(err)
	}
	info := fmt.Sprintf("Downloading: %s/%v", IMAGE_DIR, index)
	fmt.Println(info)

	defer resp.Body.Close()
	f, err := os.Create(fmt.Sprintf("%s/%v.jpg", IMAGE_DIR, index))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Downloaded to: ", f.Name())
}
