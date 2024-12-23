package main

import (
	"desktop/internal"
	"desktop/internal/api"
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
	imgs, _ := conf.Get("image.selected_abs_path")

	apikey, _ := conf.Get("api.unsplash_apikey")
	dp, _ := conf.Get("image.selected_abs_path")
	cat, _ := conf.Get("api.image_category")
	totimg, _ := conf.Get("api.download_limit")

	if apikey == nil || dp == nil {
		log.Fatal("Image path not set")
	}

	cimgs := imgs.(string)

	var ccat string
	var ctotimg int

	if cat == nil {
		ccat = "technology"
	} else {
		ccat = cat.(string)
	}

	if totimg == nil {
		ctotimg = 10
	} else {
		ctotimg = totimg.(int)
	}

	// c := api.ImageConfig{
	// 	Category:           ccat,
	// 	TotalDownloadImage: ctotimg,
	// 	Path:               dp.(string),
	// 	Apikey:             apikey.(string),
	// }

	tmx := getDuration(cint.(string))

	deskw := time.NewTicker(tmx)
	down := time.NewTicker(20 * time.Minute)
	//
	defer deskw.Stop()
	defer down.Stop()

	// quit := make(chan struct{})

	for {
		select {
		case <-deskw.C:
			scheduleSetDesktopWallpaper(cimgs)
			// case <-down.C:
			// 	scheduleDownloadImages(c, cimgs)
		}
	}

	// time.Sleep(21 * time.Second)

	// Simulate running for some time (e.g., 1 hour)
}

func scheduleDownloadImages(c api.ImageConfig, imgs string) error {
	if imgs == "" {
		log.Fatal("Image directory not set")
	}

	fmt.Println(c)

	internal.FetchImages(c)

	println("Hello, world my people")
	return nil
}

func scheduleSetDesktopWallpaper(cnf string) error {
	if cnf == "" {
		log.Fatal("Image directory not set")
	}
	fp := cnf

	imgs := getImages(fp)

	random := rand.Intn(len(imgs))
	f := imgs[random]
	fmt.Println(f)
	internal.WallpaperEvent(f)

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

func getDuration(tm string) time.Duration {
	if tm == "" {
		return 30 * time.Minute
	}

	var tmx time.Duration
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
	} else if cs == "w" {
		t = 7 * 24 * time.Hour
	}

	idur := string(dur)
	n, _ := strconv.Atoi(idur)

	tmx = time.Duration(n) * t

	return tmx
}
