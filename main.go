package main

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	cal "github.com/mnogu/go-calculator"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Calculator")
	myWindow.Resize(fyne.NewSize(800, 600))

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter the equation")

	str := binding.NewString()
	content := container.NewVBox(input, widget.NewButton("Calculate", func() {
		go func() {
			f, err := calculator(input.Text)
			log.Printf("Calculate result: %v", f)
			if err != nil {
				err := str.Set("服务器开小差了...")
				if err != nil {
					return
				}
			} else {
				err := str.Set(fmt.Sprint(f))
				if err != nil {
					return
				}
			}
		}()
	}), widget.NewLabelWithData(str))
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}

func calculator(str string) (float64, error) {
	return cal.Calculate(str)
}
