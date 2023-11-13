package utils

import (
	"fmt"
	"os"
	"strings"

	"github.com/olekukonko/tablewriter"
	"github.com/tayfun8/scanwp/models"

	"github.com/fatih/color"
)

func WriteVersion(version string) {
	green := color.New(color.FgGreen)

	greenText := green.Sprint(version)
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Wordpress Version"})
	data := [][]string{
		{greenText},
	}
	for _, row := range data {
		table.Append(row)
	}

	table.Render()
}

func WriteThemes(themes []models.Themes) {

	fmt.Println()
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Wordpress Theme", "Theme Version"})

	for _, theme := range themes {
		name := color.New(color.FgBlue).Sprint(theme.Name)
		version := color.New(color.FgBlue).Sprint(theme.Version)
		table.Append([]string{name, version})
	}

	table.Render()

}

func WritePlugins(plugins []models.Plugins) {
	fmt.Println()
	table := tablewriter.NewWriter(os.Stdout)

	table.SetHeader([]string{"Wordpress Plugin", "Plugin Version"})

	for _, plugin := range plugins {
		name := color.New(color.FgMagenta).Sprint(plugin.Name)
		version := color.New(color.FgMagenta).Sprint(plugin.Version)
		table.Append([]string{name, version})
	}

	table.Render()

}

func WriteDirb(dirb []models.Dirb) {

	tableSuccess := tablewriter.NewWriter(os.Stdout)
	tableSuccess.SetHeader([]string{"Url", "Status Code"})
	tableSuccess.SetColumnColor(tablewriter.Colors{tablewriter.FgGreenColor}, tablewriter.Colors{tablewriter.FgGreenColor})

	tableFourHundred := tablewriter.NewWriter(os.Stdout)
	tableFourHundred.SetHeader([]string{"Url", "Status Code"})
	tableFourHundred.SetColumnColor(tablewriter.Colors{tablewriter.FgRedColor}, tablewriter.Colors{tablewriter.FgRedColor})

	tableServerError := tablewriter.NewWriter(os.Stdout)
	tableServerError.SetHeader([]string{"Url", "Status Code"})
	tableServerError.SetColumnColor(tablewriter.Colors{tablewriter.FgBlueColor}, tablewriter.Colors{tablewriter.FgBlueColor})

	for _, fuzz := range dirb {
		if strings.HasPrefix(fuzz.StatusCode, "2") {
			tableSuccess.Append([]string{
				fuzz.Path,
				fuzz.StatusCode,
			})
		} else if strings.HasPrefix(fuzz.StatusCode, "5") {
			tableServerError.Append([]string{
				fuzz.Path,
				fuzz.StatusCode,
			})
		} else {
			tableFourHundred.Append([]string{
				fuzz.Path,
				fuzz.StatusCode,
			})
		}
	}

	tableSuccess.Render()
	tableServerError.Render()
	tableFourHundred.Render()
}
