package main

import (
	"log"
	"sort"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// Trim any comma's located after arguments

//func commaTrim(entered string) string {
//	s := strings.TrimSuffix(entered, ",")
//	return s
// }

func desktopLayout(input *widget.Entry, button, helpButton *widget.Button, fixedContainer, fixedLabels *fyne.Container) *fyne.Container {
	return container.NewGridWithRows(3,
		container.NewGridWithColumns(3, // 3 x 3 grid made
			layout.NewSpacer(), // First column
			container.NewVBox( // Second Column in first row (actually two items)
				input,
				button,
				helpButton,
			),
			layout.NewSpacer(), // Third Column in second row
		),
		fixedLabels,
		fixedContainer,
	)
}

var kvPairs [][]string

var resData = make(map[string]string)
var table *widget.Table

func refreshTable() {
	kvPairs = kvPairs[:0]
	for k, v := range resData {
		kvPairs = append(kvPairs, []string{k, v})
	}
	sort.Slice(kvPairs, func(i, j int) bool {
		return kvPairs[i][0] < kvPairs[j][0]
	})
	table.Refresh()
}

// Headers for table

func main() {

	// GUI Code
	myApp := app.New()
	myWindow := myApp.NewWindow("IP Loc")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter IP Address/s...")
	input.Resize(fyne.NewSize(100, 20))

	//results := IPResults{}

	button := widget.NewButton("Search", func() {
		res, err := getIP(input.Text)
		if err != nil {
			log.Println(err)
		}
		results := resultsToSlice(res, resData)
		for k, v := range results {
			kvPairs = append(kvPairs, []string{k, v})
		}
		sort.Slice(kvPairs, func(i, j int) bool {
			return kvPairs[i][0] < kvPairs[j][0]
		})

		refreshTable()
	})

	helpButton := widget.NewButton("Help", func() {
		help()
	})

	table = widget.NewTable(
		func() (int, int) {
			return len(kvPairs), 2
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("IP Results")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(kvPairs[i.Row][i.Col])
		})
	table.Resize(fyne.NewSize(500, 400))
	table.SetColumnWidth(0, 100)
	table.SetColumnWidth(1, 200)
	fixedContainer := container.NewWithoutLayout(table)
	fixedContainer.Resize(fyne.NewSize(500, 150))

	ipLabel := widget.NewLabel("IP Address")
	CnLabel := widget.NewLabel("Country")
	rgLabel := widget.NewLabel("Region")
	cityLabel := widget.NewLabel("City")
	ispLabel := widget.NewLabel("ISP")
	mobileLabel := widget.NewLabel("Mobile")
	vpnLabel := widget.NewLabel("VPN")

	fixedLabels := container.NewGridWithColumns(7,
		ipLabel, CnLabel, rgLabel, cityLabel, ispLabel, mobileLabel, vpnLabel,
	)

	// This is the layout for the box, setup is done here, then called in myWindow
	myWindow.SetContent(desktopLayout(input, button, helpButton, fixedContainer, fixedLabels))
	myWindow.Resize(fyne.NewSize(600, 450))
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
