package main

import (
	"desktop/internal"
	"fmt"
)

var (
	INTERVAL_SEC = 10
)

// type App struct {
// 	appConf *AppConfig
// }

// Get the config interval duration
// Get the images path
// Set the desktop wallpaper
func main() {
	conf := internal.AppConfig{}
	r, _ := conf.Get("image.selected_abs_path")
	fmt.Println(r)

	// t := time.NewTicker(time.Duration(INTERVAL_SEC) * time.Second)
	// for v := range t.C {
	// 	fmt.Printf("%v value here", v)
	// }
}

func scheduleDownloadImages() error {
	return nil
}

func scheduleSetDesktopWallpaper() error {
	return nil
}
