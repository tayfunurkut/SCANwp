package scraper

import (
	"sync"
)

var (
	Url string
	wg  sync.WaitGroup
)

func Scrape(url string) {
	Url = url

	wg.Add(3)

	go getThemes()
	go getPlugins()
	go getVersion()

	wg.Wait()
}
