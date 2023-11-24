package main

import (
	"encoding/base64"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	// "fyne.io/fyne/v2/app"
	// "fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/container/layout"
	// "fyne.io/fyne/v2/container/widget"
)

func main() {
	myApp := app.New()
	myWinsdow := myApp.NewWindow("Base64 Decoder")

	inputEntry := widget.NewEntry()
	outputLabel := widget.NewLabel("Decode Result:")

	decodeButton := widget.NewButton("decode", func() {
		input := inputEntry.Text
		if input == "" {
			outputLabel.SetText("Decode Result: [Empty Input]")
			return
		}

		decodeBytes, err := base64.StdEncoding.DecodeString(input)
		if err != nil {
			outputLabel.SetText("Decode Failed: " + err.Error())
			return
		}

		outputLabel.SetText("Decode Result: " + string(decodeBytes))
	})

	myWinsdow.SetContent(container.NewVBox(
		widget.NewLabel("Enter Base64 String:"),
		inputEntry,
		decodeButton,
		layout.NewSpacer(),
		outputLabel,
	))

	myWinsdow.ShowAndRun()
}
