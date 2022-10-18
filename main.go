package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	initConfig()
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:     getString("en", "WINDOW_TITLE"),
		Width:     Config.Window.MinWidth,
		Height:    Config.Window.MinHeight,
		MinWidth:  Config.Window.MinWidth,
		MinHeight: Config.Window.MinHeight,
		Assets:    assets,
		// BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
