package wpscanapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/olekukonko/tablewriter"
	"github.com/tayfun8/scanwp/models"
	"github.com/tayfun8/scanwp/scraper"
)

var (
	data models.WPScan
)

func Scan() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file can not loading...")
	}

	key := os.Getenv("API_KEY")

	fVersion := strings.ReplaceAll(scraper.Version, ".", "")

	client := &http.Client{}
	url := "https://wpscan.com/api/v3/wordpresses/" + fVersion
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Http request could not be created")
	}
	token := fmt.Sprintf("Token token=%s", key)
	req.Header.Set("Authorization", token)
	req.Header.Set("accept", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Http request could not be created")
	}

	err = json.NewDecoder(resp.Body).Decode(&data)

	if err != nil {
		fmt.Println("JSON could not be decoded", err)
	}

	// VULN INFO STATUS VERSION

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Version", "Release Date", "Status"})

	table.Append([]string{scraper.Version, data[scraper.Version].ReleaseDate, data[scraper.Version].Status})

	// VULN INFO STATUS VERSION

	// WPSCAN API VULN

	vulnTable := tablewriter.NewWriter(os.Stdout)
	vulnTable.SetHeader([]string{"Title", "Description", "Vuln Type", "Cvss Score"})

	vulnTable.SetColumnColor(tablewriter.Colors{tablewriter.FgRedColor}, tablewriter.Colors{tablewriter.FgRedColor}, tablewriter.Colors{tablewriter.FgRedColor}, tablewriter.Colors{tablewriter.FgRedColor}) //  tablewriter.Colors{tablewriter.FgRedColor}, tablewriter.Colors{tablewriter.FgRedColor}, tablewriter.Colors{tablewriter.FgRedColor}, tablewriter.Colors{tablewriter.FgRedColor}, tablewriter.Colors{tablewriter.FgRedColor}, tablewriter.Colors{tablewriter.FgRedColor}, tablewriter.Colors{tablewriter.FgRedColor}, tablewriter.Colors{tablewriter.FgRedColor}, tablewriter.Colors{tablewriter.FgRedColor}

	for _, vulns := range data[scraper.Version].Vulns {
		vulnTable.Append([]string{
			// vulns.Id,
			vulns.Title,
			// vulns.PublishedDate,
			vulns.Description,
			vulns.VulnType,
			// vulns.References.Url[0],
			// strings.Join(vulns.References.Cve, ","),
			vulns.Cvss.Score,
			// vulns.Cvss.Vector,
			// vulns.Cvss.Severity,
			// vulns.Verified,
			// vulns.FixedIn,
			// vulns.IntroducedIn,
		})
	}

	if len(data[scraper.Version].Vulns) > 0 {
		table.Render()
		vulnTable.Render()
	} else {
		fmt.Println("No vulnerabilities found.")
	}

	// WPSCAN API VULN

}
