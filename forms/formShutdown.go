package forms

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"strconv"
	"time"
)

func FormShutDown() (*widget.Label, *widget.Entry, *canvas.Text, *widget.Label, *widget.Button) {
	label := widget.NewLabel("Завершение работы компьютера через (в минутах)")
	inputMinutes := widget.NewEntry()
	inputMinutes.SetPlaceHolder("Минуты")
	inputMinutes.SetText("1")

	textValidation := canvas.NewText("", color.Black)
	textValidation.Hide()
	resultLabelShutdown := widget.NewLabel("")

	button := widget.NewButton("Отправить", func() {
		resultLabelShutdown.Text = callScriptBat(shellScriptsFolder+"ShutDown.bat", inputMinutes.Text)
		resultLabelShutdown.Refresh()

		time.AfterFunc(3*time.Second, func() {
			resultLabelShutdown.Text = ""
			resultLabelShutdown.Refresh()
		})
	})

	inputMinutes.OnChanged = func(s string) {
		value, err := strconv.ParseInt(s, 10, 32)
		if err != nil || value < 0 {
			textValidation.Text = "Ошибка ввода, введите положительное число"
			textValidation.Color = color.RGBA{255, 0, 0, 255}
			textValidation.Show()
			button.Disable()
		} else {
			textValidation.Text = ""
			textValidation.Color = color.Black
			textValidation.Hide()
			button.Enable()
		}
		textValidation.Refresh()
	}
	return label, inputMinutes, textValidation, resultLabelShutdown, button
}
