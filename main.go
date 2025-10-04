package main

import (
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("BMI Calculator")

	heightEntry := widget.NewEntry()
	heightEntry.SetPlaceHolder("Enter height (m)")

	weightEntry := widget.NewEntry()
	weightEntry.SetPlaceHolder("Enter weight (kg)")

	resultLabel := widget.NewLabel("")

	calculateBMI := func() {
		height, errH := strconv.ParseFloat(heightEntry.Text, 64)
		weight, errW := strconv.ParseFloat(weightEntry.Text, 64)

		if errH != nil || errW != nil {
			resultLabel.SetText("Invalid input")
			return
		}

		if height <= 0 {
			resultLabel.SetText("Height must be greater than 0")
			return
		}

		bmi := weight / (height * height)
		resultLabel.SetText(fmt.Sprintf("BMI: %.1f (%.2f)", bmi, bmi))
	}

	heightEntry.OnSubmitted = func(string) { calculateBMI() }
	weightEntry.OnSubmitted = func(string) { calculateBMI() }

	calcButton := widget.NewButton("Calc BMI", calculateBMI)

	myWindow.SetContent(container.NewVBox(
		widget.NewLabel("Height (m):"),
		heightEntry,
		widget.NewLabel("Weight (kg):"),
		weightEntry,
		calcButton,
		resultLabel,
	))

	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.Canvas().Focus(heightEntry)
	myWindow.ShowAndRun()
}
