package main

import (
	"fmt"
	"os/exec"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	app := app.New()

	mainWindow := app.NewWindow("HIDER")

	img := canvas.NewImageFromResource(resourceOkSvg)
	img.SetMinSize(fyne.NewSize(40, 40))

	mainWindow.SetContent(container.NewCenter(img))

	// cmd := exec.Command("ls", "-ia")
	// data, err := cmd.Output()

	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(string(data))

	r := strings.NewReader(string(resourceTestPy.Content()))

	cmd := exec.Command("python3", "-")
	cmd.Stdin = r
	data, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(data))

	mainWindow.CenterOnScreen()
	mainWindow.Resize(fyne.NewSize(300, 400))
	mainWindow.Show()

	app.Run()
}
