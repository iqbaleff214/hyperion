package main

import (
	"context"
	"github.com/wailsapp/wails/v2/pkg/application"
	m "github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"hyperion/backend/obfuscator"
	"reflect"
	"regexp"
)

type menu struct {
	ctx    context.Context
	config *obfuscator.Config
}

func (mn *menu) setContext(ctx context.Context) {
	mn.ctx = ctx
}

func (mn *menu) setConfig(config *obfuscator.Config) {
	mn.config = config
}

func (mn *menu) createNewMenu(app *application.Application) *m.Menu {
	appMenu := m.NewMenu()

	// File Menu
	fileMenu := appMenu.AddSubmenu("File")
	fileMenu.AddText("Open", keys.CmdOrCtrl("o"), func(_ *m.CallbackData) {
		runtime.EventsEmit(mn.ctx, "menu:open")
	})
	fileMenu.AddText("Open Folder", keys.Combo("o", keys.ShiftKey, keys.CmdOrCtrlKey), func(_ *m.CallbackData) {
		runtime.EventsEmit(mn.ctx, "menu:open-folder")
	})
	fileMenu.AddSeparator()
	fileMenu.AddText("Save", keys.CmdOrCtrl("s"), func(_ *m.CallbackData) {
		runtime.EventsEmit(mn.ctx, "menu:save")
	})
	fileMenu.AddText("Save As", keys.Combo("s", keys.ShiftKey, keys.CmdOrCtrlKey), func(_ *m.CallbackData) {
		runtime.EventsEmit(mn.ctx, "menu:save-as")
	})
	fileMenu.AddSeparator()
	fileMenu.AddText("Close", keys.CmdOrCtrl("w"), func(_ *m.CallbackData) {
		runtime.EventsEmit(mn.ctx, "menu:close")
	})
	fileMenu.AddText("Close Folder", keys.Combo("w", keys.ShiftKey, keys.CmdOrCtrlKey), func(_ *m.CallbackData) {
		runtime.EventsEmit(mn.ctx, "menu:close-folder")
	})
	fileMenu.AddSeparator()
	fileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *m.CallbackData) {
		app.Quit()
	})

	// Edit menu
	//if rt.GOOS == "darwin" {
	//	appMenu.Append(m.EditMenu()) // on macos platform, we should append EditMenu to enable Cmd+C,Cmd+V,Cmd+Z... shortcut
	//}

	// Run Menu
	runMenu := appMenu.AddSubmenu("Run")
	runMenu.AddText("Obfuscate", keys.CmdOrCtrl("r"), func(_ *m.CallbackData) {
		println("Obfuscate clicked")
	})
	runMenu.AddText("Obfuscate All", keys.Combo("r", keys.ShiftKey, keys.CmdOrCtrlKey), func(_ *m.CallbackData) {
		println("Obfuscate All clicked")
	})

	// Configuration Submenu
	configMenu := runMenu.AddSubmenu("Configuration")
	configValue := reflect.ValueOf(mn.config).Elem()
	configType := configValue.Type()
	for i := 0; i < configValue.NumField(); i++ {
		field := configType.Field(i)
		if field.Name == "path" {
			continue
		}

		configMenu.AddCheckbox(toCamelCaseLabel(field.Name), configValue.FieldByName(field.Name).Bool(), nil, func(_ *m.CallbackData) {
			currentValue := configValue.FieldByName(field.Name).Bool()
			configValue.FieldByName(field.Name).SetBool(!currentValue)

			mn.config.Save(*mn.config)
		})
	}

	// Help Menu
	helpMenu := appMenu.AddSubmenu("Help")
	helpMenu.AddText("How to Contribute", nil, func(_ *m.CallbackData) {
		println("How to Contribute clicked")
	})
	helpMenu.AddText("Report Issue", nil, func(_ *m.CallbackData) {
		runtime.BrowserOpenURL(mn.ctx, "https://github.com/404NotFoundIndonesia/hyperion/issues")
	})

	return appMenu
}

func toCamelCaseLabel(input string) string {
	re := regexp.MustCompile("([a-z])([A-Z])")
	return re.ReplaceAllString(input, "$1 $2")
}
