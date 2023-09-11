package main

import (
	"changeme/internal/oneinstance"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	instance, err := oneinstance.Ensure()
	if err != nil {
		println("Another instance is already running. I sent it some args, everything is okay, you don't need to worry. Here's the error anyways:", err.Error())
		return
	}

	// Create an instance of the app structure
	app := NewApp(instance)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "sample-arg-pass",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
