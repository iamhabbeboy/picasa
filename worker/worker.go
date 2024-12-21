package main

import (
	"desktop/internal"
	"fmt"
	"log"
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
		tmx = 10 * time.Minute
	} else {
		tm := cint.(string)

		var seq rune
		var dur []rune

		rtm := []rune(tm)

		if tm != "" {
			indx := len(rtm) - 1
			seq = rtm[indx]
			dur = rtm[:indx]
		}

		cs := string(seq)
		var t time.Duration

		if cs == "m" {
			t = time.Minute
		} else if cs == "s" {
			t = time.Second
		} else if cs == "h" {
			t = time.Hour
		}

		idur := string(dur)
		n, _ := strconv.Atoi(idur)

		tmx = time.Duration(n) * t
	}

	//tf := 30 * time.Second

	deskw := time.NewTicker(tmx)
	// down := time.NewTicker(7 * 24 * time.Hour)
	//
	defer deskw.Stop()
	// defer down.Stop()

	// quit := make(chan struct{})

	for {
		select {
		case <-deskw.C:
			scheduleSetDesktopWallpaper(conf, tmx)
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
	return nil
}

func scheduleSetDesktopWallpaper(conf *internal.AppConfig, t time.Duration) error {
	cnf, _ := conf.Get("image.selected_abs_path")
	if cnf == nil {
		log.Fatal("Image directory not set")
	}
	fp := cnf.(string)

	fmt.Println(t)
	fmt.Println(fp)
	imgs := getImages(fp)

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
