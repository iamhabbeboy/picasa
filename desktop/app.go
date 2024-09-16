package main

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
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
}

func (a *App) GetDownloadedImages() []string {
	res := GetImagesFromDir()
	return res
}
