package main

import (
	"desktop/internal"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	INTERVAL_SEC = 10
)

// type App struct {
// appConf *internal.AppConfig
// }

// Get the config interval duration
// Get the images path
// Set the desktop wallpaper
func main() {
	conf := &internal.AppConfig{}
	conf.Init()

	deskw := time.NewTicker(1 * time.Minute)
	down := time.NewTicker(7 * 24 * time.Hour)

	defer deskw.Stop()
	defer down.Stop()

	quit := make(chan struct{})

	go func() {
		for {
			select {
			case <-deskw.C:
				scheduleSetDesktopWallpaper(conf)
			// case <-down.C:
			// scheduleDownloadImages()
			case <-quit:
				fmt.Println("Worker stopped.")
				return
			}
		}
	}()

	// Simulate running for some time (e.g., 1 hour)
	// time.Sleep(1 * time.Hour)
	close(quit)
}

func scheduleDownloadImages() error {
	// dir := conf
	//
	return nil
}

func scheduleSetDesktopWallpaper(conf *internal.AppConfig) error {
	cnf, _ := conf.Get("image.selected_abs_path")
	fp := cnf.(string)
	imgs := getImages(fp)
	fmt.Println(imgs)

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(imgs), func(i, j int) {
		imgs[i], imgs[j] = imgs[j], imgs[i]
	})
	// internal.WallpaperEvent("")
	// read the directory
	// get the files in the folder
	// randomize it and set as wallpaper
	return nil
}

func getImages(path string) []string {
	var fp string = path
	if strings.Contains(path, "picasa") {
		home, _ := os.UserHomeDir()
		fp = fmt.Sprintf("%s/%s", home, path)
	}
	img, err := internal.GetAllFilesInDir(fp)
	if err != nil {
		println(err.Error())
	}
	return img
}
