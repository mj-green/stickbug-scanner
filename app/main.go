package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	// "fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func DarkPanel(content fyne.CanvasObject) fyne.CanvasObject {
	bg := canvas.NewRectangle(color.NRGBA{R: 20, G: 20, B: 20, A: 255})
	bg.StrokeColor = color.NRGBA{R: 60, G: 60, B: 60, A: 255}
	bg.StrokeWidth = 2
	return container.NewMax(bg, container.NewPadded(content))
}

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Stickbug Scanner")

	// management bar
	networkScan := widget.NewButton("Network Scan", func() {})

	username := widget.NewEntry()
	username.SetPlaceHolder("Username")

	password := widget.NewEntry()
	password.SetPlaceHolder("Password")

	networkRange := widget.NewEntry()
	networkRange.SetPlaceHolder("Network Range")

	profile := widget.NewButton("Profile", func() {})

	topBar := container.NewGridWithColumns(5, networkScan, username, password, networkRange, profile)

	// details panel (left)
	deviceScan := widget.NewButton("Device Scan", func() {})
	remoteShell := widget.NewButton("Open Remote Shell", func() {})
	saveDevice := widget.NewButton("Save Device", func() {})
	deleteDevice := widget.NewButton("Delete Device", func() {})

	leftActionButtons := container.NewGridWithColumns(4, deviceScan, remoteShell, saveDevice, deleteDevice,)

	detailsForm := container.NewVBox(
		widget.NewLabel("Device Name:"),
		widget.NewLabel("Description:"),
		widget.NewLabel("Tags:"),
		widget.NewLabel("Open Ports:"),
		widget.NewLabel("Trace:"),
		widget.NewLabel("Mac Address:"),
		widget.NewLabel("Vulnerabilities:"),
	)

	detailsPanel := DarkPanel(container.NewBorder(leftActionButtons, nil, nil, nil, detailsForm))
	shellPanel := DarkPanel(container.NewBorder(nil, nil, nil, nil, nil))
	leftPanel := container.NewVSplit(detailsPanel, shellPanel)

	// devices list (right)
	filter := widget.NewEntry()
	filter.SetPlaceHolder("Filter by name, tags or address")
	filterBar := container.NewBorder(nil, nil, nil, widget.NewButton("Save", func() {}), filter)
	
	device1 := DarkPanel(container.NewBorder(nil, nil, widget.NewLabel("Unknown Device\nOpen Ports: 22, 80"), widget.NewLabel("192.168.1.120"), nil))
	device2 := DarkPanel(container.NewBorder(nil, nil, widget.NewLabel("Gary's Laptop\nOpen Ports: "), widget.NewLabel("192.168.1.116"), nil))
	device3 := DarkPanel(container.NewBorder(nil, nil, widget.NewLabel("TP-Link WAP\nOpen Ports: 22, 80, 443"), widget.NewLabel("192.168.1.0"), nil))
	
	deviceList := container.NewVBox(device1, device2, device3)
	rightPanel := DarkPanel(container.NewBorder(filterBar, nil, nil, nil, container.NewVScroll(deviceList)))

	// assemble content on window and run
	content := container.NewHSplit(leftPanel, rightPanel)
	content.Offset = 0.6

	mainLayout := container.NewBorder(
		container.NewPadded(topBar), 
		nil, nil, nil, 
		content,
	)

	window.SetContent(mainLayout)
	window.Resize(fyne.NewSize(1200, 800))
	window.ShowAndRun()
}