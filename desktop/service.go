package main

import (
	"fmt"
	"os"
	"path"
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
