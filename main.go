package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func shutdown(minute string) string {
	cmd := exec.Command(".\\shellScripts\\ShutDown.bat", minute)
	out, err := cmd.CombinedOutput()
	sOut := string(out)
	sOut = strings.Replace(sOut, "\n", "", -1)
	if err != nil {
		log.Println(sOut)
		log.Println(err)
		return ""
	}
	return sOut
}

func main() {
	a := app.New()
	w := a.NewWindow("Запуск скриптов")
	w.Resize(fyne.NewSize(400, 300))
	label := widget.NewLabel("Завершение работы компьютера через (в минутах)")
	input := widget.NewEntry()
	input.SetPlaceHolder("Минуты")
	input.SetText("1")

	result := canvas.NewText("", color.Black)

	button := widget.NewButton("Отправить", func() {
		result.Text = shutdown(input.Text)
		result.Refresh()
	})

	input.OnChanged = func(s string) {
		_, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			result.Text = "Ошибка ввода, введите число"
			result.Color = color.RGBA{255, 0, 0, 255}
			button.Disable()
		} else {
			result.Text = ""
			result.Color = color.Black
			button.Enable()
		}
		result.Refresh()
	}

	content := container.NewVBox(label, input, button, result)
	w.SetContent(content)
	w.ShowAndRun()
}
