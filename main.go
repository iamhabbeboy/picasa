package main

import (
	"embed"
	"fmt"
	"net/http"
	"os"
	"os/exec"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/logger"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
)

//go:embed all:frontend/dist
var assets embed.FS

type FileLoader struct {
	http.Handler
}

func NewDownloadedImageFileLoader() *FileLoader {
	return &FileLoader{}
}

func (h *FileLoader) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var err error
	requestPath := req.URL.Path
	println("Requesting file:", requestPath)
	fileData, err := os.ReadFile(requestPath)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(fmt.Sprintf("Could not load file %s", requestPath)))
	}

	res.Write(fileData)
}

func startSchedulerWorker() {
	cmd := exec.Command("./worker/picasa_worker")
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

func main() {

	app := NewApp()

	err := wails.Run(&options.App{
		Title:         "Picasa Desktop",
		Width:         990,
		Height:        768,
		DisableResize: true,
		AssetServer: &assetserver.Options{
			Assets:  assets,
			Handler: NewDownloadedImageFileLoader(),
		},
		LogLevel:           logger.DEBUG,
		LogLevelProduction: logger.ERROR,
		BackgroundColour:   &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:          app.startup,
		Bind: []interface{}{
			app,
		},
		DragAndDrop: &options.DragAndDrop{
			EnableFileDrop:     false,
			DisableWebViewDrop: false,
			CSSDropProperty:    "--wails-drop-target",
			CSSDropValue:       "drop",
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: false,
				HideTitle:                  true,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       false,
			},
			Appearance:           mac.DefaultAppearance,
			WebviewIsTransparent: false,
			WindowIsTranslucent:  false,
			About: &mac.AboutInfo{
				Title:   "Picasa Desktop",
				Message: "©2024 Abiodun Azeez",
			},
		},
	})

	if err != nil {
		fmt.Printf("[Main] Error starting Wails app: %v\n", err)
		os.Exit(1)
	}
}
