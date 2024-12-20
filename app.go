package main

import (
	"context"
	"desktop/internal"
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
	appConf *internal.AppConfig
}

type Conf struct {
	ImageCategory string
	TotalImage    int
	Interval      string
	DefaultPath   string
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
	selectedPath, err := a.appConf.Get("image.selected_abs_path")

	if err != nil {
		println(err)
	}
	path := selectedPath.(string)
	var fp string = path
	if strings.Contains(path, "picasa") {
		home, _ := os.UserHomeDir()
		fp = fmt.Sprintf("%s/%s", home, path)
	}
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
		return []string{""}
	}

	// store the path selected.
	a.appConf.Set("image.selected_abs_path", dir)

	imgs, err := a.GetAllFilesInDir(dir)

	if err != nil {
		println(err)
		return []string{""}
	}

	return imgs
}

func (a *App) GetAllFilesInDir(dir string) ([]string, error) {
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

func (a *App) DownloadImages(conf api.ImageConfig) {
	internal.FetchImages(conf)
}

func (a *App) SetWallpaper(path string) {
	internal.WallpaperEvent(path)
}

func (a *App) GetConfig() Conf {
	imgCat, _ := a.appConf.Get("api.image_category")
	totalImg, _ := a.appConf.Get("api.download_limit")
	intvl, _ := a.appConf.Get("api.interval")
	dp, _ := a.appConf.Get("image.selected_abs_path")

	c := Conf{
		ImageCategory: imgCat.(string),
		TotalImage:    totalImg.(int),
		Interval:      intvl.(string),
		DefaultPath:   dp.(string),
	}
	return c
}

func (a *App) SetConfig(conf Conf) {
	a.appConf.Set("api.image_category", conf.ImageCategory)
	a.appConf.Set("api.download_limit", conf.TotalImage)
	a.appConf.Set("image.selected_abs_path", conf.DefaultPath)
}

func (a *App) OpenDirDialogWindow() string {
	dir, err := OpenNativeDir(a.ctx)
	if err != nil {
		println(err.Error())
		return ""
	}
	return dir
}

func (a *App) MessageDialog(m string) (string, error) {
	return MessageBox(a.ctx, m)
}

// https://gist.github.com/stupidbodo/0db61fa874213a31dc57 - replacement for cronjob
// https://gist.github.com/harubaru/f727cedacae336d1f7877c4bbe2196e1#model-overview
