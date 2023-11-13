package scraper

import (
	"log"
	"regexp"

	"github.com/gocolly/colly/v2"
	"github.com/tayfun8/scanwp/utils"
)

var Version string

func getVersion() {
	c := colly.NewCollector()
	defer wg.Done()

	c.OnHTML("meta[name='generator']", func(h *colly.HTMLElement) {
		vers := h.Attr("content")
		re := regexp.MustCompile(`WordPress (\d+\.\d+(\.\d+)?)`)
		matches := re.FindStringSubmatch(vers)
		if len(matches) >= 2 {
			Version = matches[1]
		}

	})

	err := c.Visit(Url)
	if err != nil {
		log.Fatal("Can not connected")
	}

	utils.WriteVersion(Version)

}
