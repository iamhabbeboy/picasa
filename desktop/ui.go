package main

import (
	"context"
	"runtime"

	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	run "github.com/wailsapp/wails/v2/pkg/runtime"
)

func SetMenuItem(ctx context.Context, app *App) *menu.Menu {
	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("File")

	FileMenu.AddText("About", keys.CmdOrCtrl("a"), func(cd *menu.CallbackData) {
		_, err := run.MessageDialog(ctx, run.MessageDialogOptions{
			Type:          run.InfoDialog,
			Title:         "Picasa Desktop",
			Message:       "Get random wallpaper for your desktop. \n \n © 2024-2024. Abiodun Azeez. \n All Right Reserved.",
			DefaultButton: "No",
		})
		if err != nil {
			println(err.Error())
		}
	})
	FileMenu.AddSeparator()

	FileMenu.AddText("Setting", keys.CmdOrCtrl("s"), func(cd *menu.CallbackData) {
		run.EventsEmit(ctx, "shortcut.page.setting")
	})
	FileMenu.AddSeparator()

	FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		run.Quit(app.ctx)
	})

	if runtime.GOOS == "darwin" {
		AppMenu.Append(menu.EditMenu()) // on macos platform, we should append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z... shortcut
	}
	return AppMenu
}

func SettingPage(mn *menu.CallbackData) {
	run.EventsOn(context.TODO(), "Hello world", func(optionalData ...interface{}) {
	})
}

func OpenNativeDir(ctx context.Context) (string, error) {
	return run.OpenDirectoryDialog(ctx, run.OpenDialogOptions{
		DefaultDirectory: "",
		Title:            "Select directory",
		Filters: []run.FileFilter{
			{
				DisplayName: "Images (*.png;*.jpg)",
				Pattern:     "*.png;*.jpg",
			},
		},
	})
}
