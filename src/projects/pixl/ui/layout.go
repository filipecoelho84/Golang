package ui

import "fyne.io/fyne/v2/container"

func Setup(app *AppInit) {
	swatchsContainer := BuildSwatches(app)
	colorPicker := SetupColorPicker(app)

	appLayout := container.NewBorder(nil, swatchsContainer, nil, colorPicker)

	app.PixlWindow.SetContent(appLayout)
}
