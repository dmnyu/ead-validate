package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/nyudlts/go-aspace"
	"image/color"
	"io/ioutil"
	"os"
)

var (
	black = color.Black
	data  binding.ExternalStringList
)

func main() {

	a := app.NewWithID("edu.nyu.dlts.ead-validate")
	validator := a.NewWindow("EAD Validator")
	validator.Resize(fyne.NewSize(480, 480))
	fileEntry := buildFileEntry()
	buttons := buildButtons()
	data = binding.BindStringList(&[]string{"OUTPUT"})
	list := widget.NewListWithData(data,
		func() fyne.CanvasObject {
			return widget.NewLabel("template")
		},
		func(i binding.DataItem, o fyne.CanvasObject) {
			o.(*widget.Label).Bind(i.(binding.String))
		})
	list.Resize(fyne.NewSize(480, 240))
	c := container.New(layout.NewVBoxLayout(), fileEntry, list, buttons)
	validator.SetContent(c)
	validator.ShowAndRun()
}

func buildFileEntry() *fyne.Container {
	c := container.New(layout.NewVBoxLayout())
	c.Add(canvas.NewText("Set the location of the EAD file to validate", black))
	entry := widget.NewEntry()
	c.Add(entry)
	button := widget.NewButton("Validate", func() {
		path := entry.Text
		runValidation(path)
	})
	c.Add(button)
	return c
}

func buildButtons() *fyne.Container {
	c := container.New(layout.NewHBoxLayout())
	logButton := widget.NewButton("Write Log", func() {})
	c.Add(logButton)
	quitButton := widget.NewButton("Exit", func() {
		os.Exit(0)
	})
	c.Add(quitButton)
	return c
}

func runValidation(path string) {
	data.Append(fmt.Sprintf("Validating `%s`", path))
	if fileExists(path) == true {
		data.Append(fmt.Sprintf("`%s` exists", path))
		if isFile(path) == true {
			data.Append(fmt.Sprintf("`%s` is a file", path))
			if isValid(path) == true {
				data.Append(fmt.Sprintf("`%s` is a valid", path))
			} else {
				data.Append(fmt.Sprintf("`%s` is not valid", path))
			}
		} else {
			data.Append(fmt.Sprintf("`%s` is a directory, cannot validate", path))
		}
	} else {
		data.Append(fmt.Sprintf("  `%s` does not exist", path))
	}
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if err != nil {
		return false
	} else {
		return false
	}
}

func isFile(path string) bool {
	fi, _ := os.Stat(path)
	if fi.IsDir() {
		return false
	}
	return true
}

func isValid(path string) bool {
	f, _ := ioutil.ReadFile(path)
	err := aspace.ValidateEAD(f)
	if err != nil {
		return false
	}
	return true
}
