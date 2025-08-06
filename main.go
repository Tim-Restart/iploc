package main

import (
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func desktopLayout(input *widget.Entry, button, helpButton *widget.Button, fixedContainer, fixedLabels *fyne.Container) *fyne.Container {
	return container.NewGridWithRows(3,
		container.NewGridWithColumns(3, // 3 x 3 grid made
			layout.NewSpacer(), // First spacer on left, first column first row
			container.NewVBox( // Second Column in first row (actually two items)
				input,
				button,
				helpButton,
			),
			layout.NewSpacer(), // Third Column in first row
		),
		container.NewVBox(
			fixedLabels,
			fixedContainer,
		),
	)
}

var table *widget.Table

func createSlice(ips []string) [][]string {
	ipAdd := make([][]string, len(ips))
	for i := range ipAdd {
		ipAdd[i] = make([]string, 7)
		for _ = range ipAdd[i] {
			ipAdd[i][0] = ips[i]
		}
	}

	return ipAdd
}

//func refreshTable(ips [][]string) {
//	for i := range ips {
//		ips[i] = ips[i][:0]
//	}

//	table.Refresh()
//}

func main() {

	// GUI Code
	myApp := app.New()
	myWindow := myApp.NewWindow("IP Loc")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter IP Address/s...")
	input.Resize(fyne.NewSize(100, 20))

	var ips [][]string

	button := widget.NewButton("Search", func() {
		ips = checkInput(input.Text)
		for i := range ips {
			res, err := getIP(ips[i][0])
			if err != nil {
				log.Println(err)
			}

			resultsSlice(res, ips[i])
		}

		//refreshTable(ips)
	})

	helpButton := widget.NewButton("Help", func() {
		help()
	})

	table = widget.NewTable(
		func() (int, int) {
			return len(ips), 7
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("IP Results")
		},
		func(i widget.TableCellID, cell fyne.CanvasObject) {
			cell.(*widget.Label).SetText(ips[i.Row][i.Col])
		},
	)
	table.Resize(fyne.NewSize(900, 400))
	//for i := 0; i < len(results); i++ {
	//		table.SetColumnWidth(i, 80)
	//}

	table.SetColumnWidth(0, 110) // IP
	table.SetColumnWidth(1, 170) // City
	table.SetColumnWidth(2, 60)  // Region
	table.SetColumnWidth(3, 100) // Country
	table.SetColumnWidth(4, 190) // ISP
	table.SetColumnWidth(5, 50)  // Mobile
	table.SetColumnWidth(6, 50)  // VPN

	fixedContainer := container.NewWithoutLayout(table)
	fixedContainer.Resize(fyne.NewSize(700, 150))

	// Code relating to labels used as headers for the table
	// Manually create the headers, and then size and position the text
	// according the the table output
	ipLabel := widget.NewLabel("IP Address")
	ipLabel.Resize(fyne.NewSize(100, 20))
	ipLabel.Move(fyne.NewPos(10, 0))

	cityLabel := widget.NewLabel("City")
	cityLabel.Resize(fyne.NewSize(80, 20))
	cityLabel.Move(fyne.NewPos(170, 0))

	rgLabel := widget.NewLabel("Region")
	rgLabel.Resize(fyne.NewSize(60, 20))
	rgLabel.Move(fyne.NewPos(285, 0))

	CnLabel := widget.NewLabel("Country")
	CnLabel.Resize(fyne.NewSize(100, 20))
	CnLabel.Move(fyne.NewPos(370, 0))

	ispLabel := widget.NewLabel("ISP")
	ispLabel.Resize(fyne.NewSize(120, 20))
	ispLabel.Move(fyne.NewPos(530, 0))

	mobileLabel := widget.NewLabel("Mobile")
	mobileLabel.Resize(fyne.NewSize(50, 20))
	mobileLabel.Move(fyne.NewPos(645, 0))

	vpnLabel := widget.NewLabel("VPN")
	vpnLabel.Resize(fyne.NewSize(50, 20))
	vpnLabel.Move(fyne.NewPos(710, 0))

	spacer := canvas.NewRectangle(color.Transparent)
	spacer.SetMinSize(fyne.NewSize(1, 20))

	fixedLabels := container.NewMax(
		spacer,
		container.NewWithoutLayout(
			ipLabel, cityLabel, CnLabel, rgLabel, ispLabel, mobileLabel, vpnLabel,
		))

	// This is the layout for the box, setup is done here, then called in myWindow
	myWindow.SetContent(desktopLayout(input, button, helpButton, fixedContainer, fixedLabels))
	myWindow.Resize(fyne.NewSize(900, 200))
	myWindow.ShowAndRun()
}

/* Trying GUI logic
queries := len(os.Args)
switch {
case queries == 1:
	// Calls the help file with usage from reports.go
	help()
	return
case queries == 2:
	ip := commaTrim(os.Args[1])
	result, err := getIP(ip)
	if err != nil {
		fmt.Println("Error sending request")
		fmt.Println(err)
		return
	}
	err = report(result)
	if err != nil {
		fmt.Println("error writing report")
	}
case queries > 2:
	i := 1
	for i < len(os.Args) {
		ip := commaTrim(os.Args[i])
		result, err := getIP(ip)
		if err != nil {
			fmt.Println("Error sending request")
			fmt.Println(err)
			return
		}
		err = report(result)
		if err != nil {
			fmt.Println("error writing report")
		}
		i++
	}
default:
	fmt.Println("Incorrect queries, please try again")
}

*/
