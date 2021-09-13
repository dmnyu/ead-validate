package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"image/color"
)

var black = color.Black

func main() {
	a := app.NewWithID("edu.nyu.dlts.ead-validate")
	validator := a.NewWindow("EAD Validator")
	validator.Resize(fyne.NewSize(480,480))
	terminal := buildTerminal()
	fileEntry := buildFileEntry(terminal)
	c := container.New(layout.NewVBoxLayout(), fileEntry, terminal)
	validator.SetContent(c)
	validator.ShowAndRun()
}

func buildTerminal() *fyne.Container {
	c := container.New(layout.NewVBoxLayout())
	c.Add(canvas.NewText("Output", black))
	return c
}

func buildFileEntry(terminal *fyne.Container) *fyne.Container {
	c := container.New(layout.NewVBoxLayout())
	c.Add(canvas.NewText("Set the location of the EAD file to validate", black))
	entry := widget.NewEntry()
	c.Add(entry)
	button := widget.NewButton("Validate", func() {
		terminal.Add(canvas.NewText("Test", black))
	})
	c.Add(button)
	return c
}
