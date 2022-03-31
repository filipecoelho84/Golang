package ui

func Setup(app *AppInit) {
	swatchsContainer := BuildSwatches(app)

	app.PixlWindow.SetContent(swatchsContainer)
}