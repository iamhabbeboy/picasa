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

func test() {
	t := time.NewTicker(time.Duration(30 * time.Second))
	for _ = range t.C {
		println("Hello, world")
		// scheduleSetDesktopWallpaper(conf)
	}

}

// Get the config interval duration
// Get the images path
// Set the desktop wallpaper
func main() {
	// conf := &internal.AppConfig{}
	// conf.Init()
	test()
	//
	// deskw := time.NewTicker(30 * time.Second)
	// down := time.NewTicker(7 * 24 * time.Hour)
	//
	// defer deskw.Stop()
	// defer down.Stop()
	//
	// quit := make(chan struct{})
	//
	// for {
	// 	select {
	// 	case <-deskw.C:
	// 		currentTime := time.Now()
	// 		if !lastRun.IsZero() {
	// 			diff := currentTime.Sub(lastRun)
	// 			fmt.Printf("Time difference between runs: %v\n", diff)
	// 		} else {
	// 			fmt.Println("First run!")
	// 		}
	// 		println("Setting wallpaper event")
	// 		// 		scheduleSetDesktopWallpaper(conf)
	// 		// 	// case <-down.C:
	// 		// 	// scheduleDownloadImages()
	//
	// 		lastRun = currentTime
	// 	}
	// }

	// Simulate running for some time (e.g., 1 hour)
	// close(quit)
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
	fmt.Println(len(imgs))

	random := rand.Intn(len(imgs))
	f := imgs[random]
	fmt.Println(f)
	internal.WallpaperEvent(f)
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
