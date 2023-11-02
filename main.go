package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"go-fyne/forms"
)

func main() {
	a := app.New()
	w := a.NewWindow("Запуск скриптов")
	w.Resize(fyne.NewSize(400, 300))

	labelMinutes, inputMinutes, textValidationMinutes, resultLabelShutdown, buttonShutdown := forms.FormShutDown()

	content := container.NewVBox(
		labelMinutes, inputMinutes, textValidationMinutes, buttonShutdown, resultLabelShutdown,
	)
	w.SetContent(content)
	w.ShowAndRun()
}
