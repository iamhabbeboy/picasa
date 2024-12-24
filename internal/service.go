package internal

import (
	"desktop/internal/api"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
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
	apikey := conf.Apikey
	sp := conf.Path
	svc := api.NewImageDownload("unsplash", sp, apikey)
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

func GetAllFilesInDir(dir string) ([]string, error) {
	var images []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && isImageFile(path) {
			if isImageFile(info.Name()) {
				images = append(images, path)
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return images, nil
}

func isImageFile(file string) bool {
	extn := []string{".jpg", ".jpeg", ".png"}
	for _, ext := range extn {
		if strings.HasSuffix(strings.ToLower(file), ext) {
			return true
		}
	}
	return false
}
