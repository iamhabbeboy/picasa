package main

import (
	"context"
	"desktop/internal"
	"desktop/internal/api"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

var appConf = internal.AppConfig{}

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

func startSchedulerWorker() {
	fmt.Println("Loading picasa scheduler.....")
	cmd := exec.Command("/usr/local/bin/picasa_scheduler")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		fmt.Println("Error starting worker:", err)
		return
	}

	err = cmd.Wait()
	if err != nil {
		fmt.Println("Worker process finished with error:", err)
	}

	fmt.Printf("[Main App] Scheduler started with PID: %d\n", cmd.Process.Pid)
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	menu := SetMenuItem(ctx, a)
	runtime.MenuSetApplicationMenu(ctx, menu)
	appConf.Init("$HOME/.picasa")
	startSchedulerWorker()
}

func (a *App) GetDownloadedImages() []string {
	selectedPath, err := appConf.Get("image.selected_abs_path")

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
	appConf.Set("image.selected_abs_path", dir)

	imgs, err := internal.GetAllFilesInDir(dir)

	if err != nil {
		println(err)
		return []string{""}
	}

	return imgs
}

func (a *App) DownloadImages() {
	apikey, _ := appConf.Get("api.unsplash_apikey")
	dp, _ := appConf.Get("image.selected_abs_path")
	tot, _ := appConf.Get("api.download_limit")
	cat, _ := appConf.Get("api.image_category")

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

	var imagePath = dp.(string)
	if strings.Contains(imagePath, ".picasa") {
		home, _ := os.UserHomeDir()
		fp := fmt.Sprintf("%s/.picasa/images", home)
		imagePath = fp
	}

	c := api.ImageConfig{
		Category:           ccat,
		TotalDownloadImage: ct,
		Path:               imagePath,
		Apikey:             apikey.(string),
	}

	_, err := deleteFilesWithPrefix(imagePath, "picasa_")

	if err != nil {
		log.Fatal("Error deleting images ", err.Error())
	}
	internal.FetchImages(c)
}

func (a *App) SetWallpaper(path string) {
	internal.WallpaperEvent(path)
}

func (a *App) GetConfig() Conf {
	imgCat, _ := appConf.Get("api.image_category")
	totalImg, _ := appConf.Get("api.download_limit")
	intvl, _ := appConf.Get("image.interval")
	dp, _ := appConf.Get("image.selected_abs_path")
	akey, _ := appConf.Get("api.unsplash_apikey")

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
		d = "Nw5jS2P4zr_oO_qbFt_39zyj7QTIMI49vYx5lCzxujY" //TODO: hardcoding api key is bad, but what can i say... user can be funny, this will help restore the key even when deleted
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
	appConf.Set("api.image_category", conf.ImageCategory)
	appConf.Set("api.download_limit", conf.TotalImage)
	appConf.Set("image.selected_abs_path", conf.DefaultPath)
	appConf.Set("image.interval", conf.Interval)
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

func deleteFilesWithPrefix(dir, prf string) error {
	// var deletedFiles []string

	imgs, err := internal.GetAllFilesInDir(dir)
	if err != nil {
		return err
	}

	for _, v := range imgs {
		if strings.Contains(v, prf) {
			// deletedFiles = append(deletedFiles, v)
			if err := os.Remove(v); err != nil {
				return err
			}
		}
	}

	return nil
}

// https://gist.github.com/stupidbodo/0db61fa874213a31dc57 - replacement for cronjob
// https://gist.github.com/harubaru/f727cedacae336d1f7877c4bbe2196e1#model-overview
