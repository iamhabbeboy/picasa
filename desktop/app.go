package main

import (
	"context"
	"desktop/internal/api"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx     context.Context
	appConf *AppConfig
}

const APP_NAME = ".picasa"

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	menu := SetMenuItem(ctx, a)
	runtime.MenuSetApplicationMenu(ctx, menu)
	a.appConf.Init()
}

func (a *App) GetDownloadedImages() []string {
	selectedPath, err := a.appConf.Get("app.default_path")

	if err != nil {
		println(err)
	}
	var path string
	if selectedPath.(string) == "" {
		defaultPath, _ := a.appConf.Get("image.selected_abs_path")
		path = defaultPath.(string)
	} else {
		path = selectedPath.(string)
	}

	home, _ := os.UserHomeDir()
	fp := fmt.Sprintf("%s/%s", home, path)

	img, err := a.GetAllFilesInDir(fp)
	if err != nil {
		println(err.Error())
	}

	return img
}

func (a *App) SelectImageDir() []string {
	dir, err := OpenNativeDir(a.ctx)
	if err != nil {
		println(err.Error())
	}

	// store the path selected.
	a.appConf.Set("image.selected_abs_path", dir)

	imgs, err := a.GetAllFilesInDir(dir)

	if err != nil {
		println(err)
	}

	return imgs
}

func (a *App) GetAllFilesInDir(dir string) ([]string, error) {
	var images []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
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

func (a *App) DownloadImages(conf api.ImageConfig) {
	FetchImages(conf)
}

func (a *App) SetWallpaper(path string) {
	WallpaperEvent(path)
}

// https://gist.github.com/stupidbodo/0db61fa874213a31dc57 - replacement for cronjob
