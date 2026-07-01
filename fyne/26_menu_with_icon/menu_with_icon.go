package main

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Fyne Menu with Shortcuts & Symbols")
	w.Resize(fyne.NewSize(500, 300))

	// Label to show actions
	status := widget.NewLabel("Ready")

	mainMenu := createMenuItem(status)
	// Example content
	content := canvas.NewText("Hello, Fyne!", color.Black)
	content.TextSize = 20

	w.SetMainMenu(mainMenu)

	// Layout
	w.SetContent(container.NewVBox(
		content,
		status,
	))

	// Error handling
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recovered from panic:", r)
		}
	}()

	w.ShowAndRun()
}

func createMenuItem(status *widget.Label) *fyne.MainMenu {
	// Menu items with icons and shortcuts
	newItem := fyne.NewMenuItem("New", func() {
		status.SetText("New File Created")
	})
	newItem.Icon = theme.FileIcon()
	newItem.Shortcut = &desktop.CustomShortcut{KeyName: fyne.KeyN, Modifier: fyne.KeyModifierControl}

	copyItem := fyne.NewMenuItem("Copy", func() {
		status.SetText("Copy Action Triggered")
	})
	copyItem.Icon = theme.ContentCopyIcon()
	copyItem.Shortcut = &desktop.CustomShortcut{KeyName: fyne.KeyC, Modifier: fyne.KeyModifierControl}

	saveItem := fyne.NewMenuItem("Save", func() {
		status.SetText("Save Action Triggered")
	})
	saveItem.Icon = theme.DocumentSaveIcon()
	saveItem.Shortcut = &desktop.CustomShortcut{KeyName: fyne.KeyS, Modifier: fyne.KeyModifierControl}

	openItem := fyne.NewMenuItem("Open…", func() {
		status.SetText("Open File Dialog")
	})
	openItem.Icon = theme.FolderOpenIcon()
	openItem.Shortcut = &desktop.CustomShortcut{KeyName: fyne.KeyO, Modifier: fyne.KeyModifierControl}

	quitItem := fyne.NewMenuItem("Quit", func() {
		status.SetText("Quit Action Triggered")
	})
	quitItem.Icon = theme.CancelIcon()
	quitItem.Shortcut = &desktop.CustomShortcut{KeyName: fyne.KeyQ, Modifier: fyne.KeyModifierControl}

	// File menu
	fileMenu := fyne.NewMenu("File", newItem, openItem, copyItem, saveItem, fyne.NewMenuItemSeparator(), quitItem)

	bitmapItem := fyne.NewMenuItem("Bitmap", func() {
		status.SetText("Bitmap Action Triggered")
	})
	bitmapItem.Icon = theme.StorageIcon()
	bitmapItem.Shortcut = &desktop.CustomShortcut{KeyName: fyne.KeyB, Modifier: fyne.KeyModifierControl}

	bimapAlertItem := fyne.NewMenuItem("Bitmap Alert", func() {
		status.SetText("Bitmap Alert Triggered")
	})
	bimapAlertItem.Icon = theme.InfoIcon()
	bimapAlertItem.Shortcut = &desktop.CustomShortcut{KeyName: fyne.KeyB, Modifier: fyne.KeyModifierControl | fyne.KeyModifierShift}

	modbusItem := fyne.NewMenuItem("Modbus", func() {
		status.SetText("Modbus Action Triggered")
	})
	modbusItem.Icon = theme.UploadIcon()
	modbusItem.Shortcut = &desktop.CustomShortcut{KeyName: fyne.KeyM, Modifier: fyne.KeyModifierControl}

	modbusAlertItem := fyne.NewMenuItem("Modbus Alert", func() {
		status.SetText("Modbus Alert Triggered")
	})
	modbusAlertItem.Icon = theme.WarningIcon()
	modbusAlertItem.Shortcut = &desktop.CustomShortcut{KeyName: fyne.KeyM, Modifier: fyne.KeyModifierControl | fyne.KeyModifierShift}

	deviceItem := fyne.NewMenuItem("Device", func() {
		status.SetText("Device Action Triggered")
	})
	deviceItem.Icon = theme.ComputerIcon()
	deviceItem.Shortcut = &desktop.CustomShortcut{KeyName: fyne.KeyD, Modifier: fyne.KeyModifierControl}
	// Maps menu
	mapsMenu := fyne.NewMenu("Maps", bitmapItem, bimapAlertItem, modbusItem, modbusAlertItem, deviceItem)
	// Help menu
	aboutItem := fyne.NewMenuItem("About ℹ️", func() {
		status.SetText("Fyne Menu Example v1.0")
	})
	helpMenu := fyne.NewMenu("Help", aboutItem)

	// Create main menu bar
	mainMenu := fyne.NewMainMenu(fileMenu, mapsMenu, helpMenu)
	return mainMenu
}
