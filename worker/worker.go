package main

import (
	"desktop/internal"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}

	conf := &internal.AppConfig{}
	conf.Init(wd)

	cint, _ := conf.Get("image.interval")

	var tmx time.Duration

	if cint == nil {
		tmx = 30 * time.Minute
	} else {

		tm := cint.(string)
		var seq, dur string

		rtm := []rune(tm)
		fmt.Println(rtm)
		// if tm != "" {
		// 	seq = rtm[len(rtm)-1]
		// 	dur = rtm[:len(tm)-1]
		// }
		//
		// var tdur time.Duration
		// if seq == "m" {
		// 	tdur = time.Minute
		// } else if seq == "s" {
		// 	tdur = time.Second
		// } else if seq == "h" {
		// 	tdur = time.Hour
		// }
		//
		// c, _ := strconv.Atoi(seq)
		// fmt.Println("this is the current time interval ", c)
		// tmx = 30 * tdur
	}
	// stopChan := make(chan bool)
	deskw := time.NewTicker(tmx)
	// down := time.NewTicker(7 * 24 * time.Hour)
	//
	defer deskw.Stop()
	// defer down.Stop()
	//
	// quit := make(chan struct{})
	//
	for {
		select {
		case <-deskw.C:
			scheduleSetDesktopWallpaper(conf)
			// 	// case <-down.C:
			// 	// scheduleDownloadImages()
			// case <-stopChan:
			// 	fmt.Println("[Scheduler] Stopping...")
			// 	return
		}
	}

	// time.Sleep(21 * time.Second)

	// Simulate running for some time (e.g., 1 hour)
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
