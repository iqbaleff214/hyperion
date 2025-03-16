package main

import (
	"context"
	"embed"
	"github.com/wailsapp/wails/v2/pkg/application"
	fs "hyperion/backend/filesystem"
	ob "hyperion/backend/obfuscator"
	"log"

	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	config := ob.NewConfig()
	file := fs.NewFileManager()
	dialog := fs.NewDialog()
	obfuscator := ob.NewObfuscator(config)

	menuHelper := menu{}
	menuHelper.setConfig(config)
	app := application.NewWithOptions(&options.App{
		Title:     "hyperion",
		Width:     1024,
		Height:    768,
		MinWidth:  600,
		MinHeight: 400,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		OnStartup: func(ctx context.Context) {
			file.SetContext(ctx)
			dialog.SetContext(ctx)
			obfuscator.SetContext(ctx)
			menuHelper.startup(ctx)
		},
		Bind: []interface{}{
			dialog, file, obfuscator,
		},
		Windows: &windows.Options{
			WindowIsTranslucent: true,
			BackdropType:        windows.Mica,
		},
		Mac: &mac.Options{
			WindowIsTranslucent: true,
		},
		Linux: &linux.Options{
			WindowIsTranslucent: true,
		},
		Debug: options.Debug{
			OpenInspectorOnStartup: true,
		},
	})
	app.SetApplicationMenu(menuHelper.createNewMenu(app))

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
