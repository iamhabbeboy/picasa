package main

import (
	"context"
	"desktop/internal"
	"desktop/internal/api"
	"fmt"
	"log"
	"os"
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
	Apikey        string
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
	a.appConf.Init(".")
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
	img, err := internal.GetAllFilesInDir(fp)
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

	imgs, err := internal.GetAllFilesInDir(dir)

	if err != nil {
		println(err)
		return []string{""}
	}

	return imgs
}

func (a *App) DownloadImages() {
	apikey, _ := a.appConf.Get("api.unsplash_apikey")
	dp, _ := a.appConf.Get("image.selected_abs_path")
	tot, _ := a.appConf.Get("api.download_limit")
	cat, _ := a.appConf.Get("api.image_category")

	if apikey == nil || dp == nil {
		log.Fatal("Image path not set")
	}

	var ct int
	if tot == nil {
		ct = 10
	} else {
		ct = tot.(int)
	}

	var ccat string
	if cat == nil {
		ccat = "technology"
	} else {
		ccat = cat.(string)
	}

	c := api.ImageConfig{
		Category:           ccat,
		TotalDownloadImage: ct,
		Path:               dp.(string),
		Apikey:             apikey.(string),
	}

	internal.FetchImages(c)
}

func (a *App) SetWallpaper(path string) {
	internal.WallpaperEvent(path)
}

func (a *App) GetConfig() Conf {
	imgCat, _ := a.appConf.Get("api.image_category")
	totalImg, _ := a.appConf.Get("api.download_limit")
	intvl, _ := a.appConf.Get("image.interval")
	dp, _ := a.appConf.Get("image.selected_abs_path")
	akey, _ := a.appConf.Get("api.unsplash_apikey")

	var img, intv, d string
	var tot int
	if imgCat == nil {
		img = ""
	} else {
		img = imgCat.(string)
	}

	if totalImg == nil {
		tot = 0
	} else {
		tot = totalImg.(int)
	}

	if intvl == nil {
		intv = ""
	} else {
		intv = intvl.(string)
	}

	if dp == nil {
		d = "Nw5jS2P4zr_oO_qbFt_39zyj7QTIMI49vYx5lCzxujY" //TODO: hardcoding api key is bad, but
	} else {
		d = dp.(string)
	}

	var k string
	if akey == nil {
		k = ""
	} else {
		k = akey.(string)
	}

	c := Conf{
		ImageCategory: img,
		TotalImage:    tot,
		Interval:      intv,
		DefaultPath:   d,
		Apikey:        k,
	}

	return c
}

func (a *App) SetConfig(conf Conf) {
	a.appConf.Set("api.image_category", conf.ImageCategory)
	a.appConf.Set("api.download_limit", conf.TotalImage)
	a.appConf.Set("image.selected_abs_path", conf.DefaultPath)
	a.appConf.Set("image.interval", conf.Interval)
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
