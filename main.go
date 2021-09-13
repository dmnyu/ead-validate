package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/nyudlts/go-aspace"
	"image/color"
	"io/ioutil"
	"os"
)

var black = color.Black

func main() {
	a := app.NewWithID("edu.nyu.dlts.ead-validate")
	validator := a.NewWindow("EAD Validator")
	validator.Resize(fyne.NewSize(480, 480))
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
		path := entry.Text
		terminal.Add(canvas.NewText(fmt.Sprintf("Validating `%s`", path), black))
		if fileExists(path) == true {
			terminal.Add(canvas.NewText(fmt.Sprintf("  `%s` exists", path), black))
			if isFile(path) == true {
				terminal.Add(canvas.NewText(fmt.Sprintf("  `%s` is a file", path), black))
				if isValid(path) == true {
					terminal.Add(canvas.NewText(fmt.Sprintf("  `%s` is a valid", path), black))
				} else {
					terminal.Add(canvas.NewText(fmt.Sprintf("  `%s` is not valid", path), black))
				}
			} else {
				terminal.Add(canvas.NewText(fmt.Sprintf("  `%s` is a directory, cannot validate", path), black))
			}
		} else {
			terminal.Add(canvas.NewText(fmt.Sprintf("`%s` does not exist", path), black))
		}
	})
	c.Add(button)
	return c
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
