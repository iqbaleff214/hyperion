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

const (
	menuItemSave         = "save"
	menuItemSaveAs       = "save_as"
	menuItemClose        = "close"
	menuItemCloseFolder  = "close_folder"
	menuItemObfuscate    = "obfuscate"
	menuItemObfuscateAll = "obfuscate_all"
)

type menu struct {
	ctx        context.Context
	config     *obfuscator.Config
	items      map[string]*m.MenuItem
	openedFile bool
}

func (mn *menu) startup(ctx context.Context) {
	mn.ctx = ctx
	runtime.EventsOn(mn.ctx, "files", func(data ...any) {
		mn.openedFile, _ = data[0].(bool)
		if mn.openedFile {
			mn.items[menuItemCloseFolder].Enable()
			mn.items[menuItemObfuscateAll].Enable()
		} else {
			mn.items[menuItemCloseFolder].Disable()
			mn.items[menuItemObfuscateAll].Disable()
		}
	})
}

func (mn *menu) setConfig(config *obfuscator.Config) {
	mn.config = config
	mn.items = make(map[string]*m.MenuItem)
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
	mn.items[menuItemSave] = fileMenu.AddText("Save", keys.CmdOrCtrl("s"), func(_ *m.CallbackData) {
		runtime.EventsEmit(mn.ctx, "menu:save")
	})
	//mn.items[menuItemSave].Disable()
	mn.items[menuItemSaveAs] = fileMenu.AddText("Save As", keys.Combo("s", keys.ShiftKey, keys.CmdOrCtrlKey), func(_ *m.CallbackData) {
		runtime.EventsEmit(mn.ctx, "menu:save-as")
	})
	//mn.items[menuItemSaveAs].Disable()
	fileMenu.AddSeparator()
	mn.items[menuItemClose] = fileMenu.AddText("Close", keys.CmdOrCtrl("w"), func(_ *m.CallbackData) {
		runtime.EventsEmit(mn.ctx, "menu:close")
	})
	//mn.items[menuItemClose].Disable()
	mn.items[menuItemCloseFolder] = fileMenu.AddText("Close Folder", keys.Combo("w", keys.ShiftKey, keys.CmdOrCtrlKey), func(_ *m.CallbackData) {
		runtime.EventsEmit(mn.ctx, "menu:close-folder")
	})
	//mn.items[menuItemCloseFolder].Disable()
	fileMenu.AddSeparator()
	fileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *m.CallbackData) {
		app.Quit()
	})

	// Run Menu
	runMenu := appMenu.AddSubmenu("Run")
	mn.items[menuItemObfuscate] = runMenu.AddText("Obfuscate", keys.CmdOrCtrl("r"), func(_ *m.CallbackData) {
		runtime.EventsEmit(mn.ctx, "run:obfuscate")
	})
	//mn.items[menuItemObfuscate].Disable()
	mn.items[menuItemObfuscateAll] = runMenu.AddText("Obfuscate All", keys.Combo("r", keys.ShiftKey, keys.CmdOrCtrlKey), func(_ *m.CallbackData) {
		runtime.EventsEmit(mn.ctx, "run:obfuscate-all")
	})
	//mn.items[menuItemObfuscateAll].Disable()
	runMenu.AddSeparator()
	runMenu.AddText("Cancel Obfuscation", keys.CmdOrCtrl("e"), func(_ *m.CallbackData) {
		runtime.EventsEmit(mn.ctx, "run:undo")
	})
	runMenu.AddText("Cancel All Obfuscation", keys.Combo("e", keys.ShiftKey, keys.CmdOrCtrlKey), func(_ *m.CallbackData) {
		runtime.EventsEmit(mn.ctx, "run:undo-all")
	})
	runMenu.AddSeparator()

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
