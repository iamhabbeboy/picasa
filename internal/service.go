package internal

import (
	"desktop/internal/api"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/labstack/gommon/log"
)

func GetImagesFromDir() []string {
	homedir, err := os.UserHomeDir()
	if err != nil {
		return []string{err.Error()}
	}
	const APP_NAME = ".picasa"
	base := fmt.Sprintf("%s/%s", homedir, APP_NAME)
	dir := path.Join(base, "/", "images")
	ent, err := os.ReadDir(dir)
	if err != nil {
		return []string{err.Error()}
	}
	images := []string{}
	for i, e := range ent {
		if i == 0 {
			continue
		}
		fpath := fmt.Sprintf("%s/%s", dir, e.Name())
		images = append(images, fpath)
	}
	return images

}

func FetchImages(conf api.ImageConfig) {
	svc := api.NewImageDownload("unsplash")
	err := svc.GetImages(conf)
	if err != nil {
		log.Fatal(err.Error())
	}
}

func WallpaperEvent(path string) {
	filepath := path
	cmd := exec.Command("osascript", "-e", fmt.Sprintf("tell application \"Finder\" to set desktop picture to POSIX file \"%s\"", filepath))

	_, err := cmd.Output()
	if err != nil {
		println(err.Error())
		return
	}
}
