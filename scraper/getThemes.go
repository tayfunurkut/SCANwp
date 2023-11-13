package scraper

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
	"github.com/tayfun8/scanwp/models"
	"github.com/tayfun8/scanwp/utils"
)

var themes []models.Themes

func getThemes() {

	c := colly.NewCollector()

	defer wg.Done()

	// Get Themes On The Wordpress Site
	c.OnHTML("link[rel='stylesheet']", func(t *colly.HTMLElement) {
		attr := t.Attr("href")
		if strings.Contains(attr, "/wp-content/themes/") {
			vers, err := getThemeVersion(attr)
			if err != nil {
				fmt.Println(err)
			}

			themeName := getThemeName(attr)

			newTheme := models.Themes{
				Name:    themeName,
				Version: vers,
			}
			if status := IsThemeContains(themes, newTheme); status == false {
				themes = append(themes, newTheme)
			}
		}
	})

	c.OnHTML("script[type='text/javascript']", func(h *colly.HTMLElement) {
		attr := h.Attr("src")
		if strings.Contains(attr, "/wp-content/themes/") {
			vers, err := getThemeVersion(attr)
			if err != nil {
				fmt.Println(err)
			}

			themeName := getThemeName(Url)

			newTheme := models.Themes{
				Name:    themeName,
				Version: vers,
			}

			if status := IsThemeContains(themes, newTheme); status == false {
				themes = append(themes, newTheme)
			}

			// themes append
		}
	})

	// -----------------------------

	if err := c.Visit(Url); err != nil {
		log.Fatal("Can not connected")
	}

	utils.WriteThemes(themes)
}

func getThemeName(url string) string {
	parts := strings.Split(url, "themes/")
	if len(parts) < 2 {
		return ""
	}

	themePath := parts[1]
	slashIndex := strings.Index(themePath, "/")
	if slashIndex >= 0 {
		themePath = themePath[:slashIndex]
	}

	return themePath
}

func getThemeVersion(url string) (string, error) {
	parts := strings.Split(url, "ver=")
	if len(parts) > 1 {
		// ver= sonrasındaki değeri al
		version := strings.Split(parts[1], "&")[0]
		return version, nil
	}
	return "", errors.New("Theme Version Not Found")
}

func IsThemeContains(themes []models.Themes, theme models.Themes) bool {
	for _, v := range themes {
		if v.Name == theme.Name && v.Version == theme.Version {
			return true
		}
	}
	return false
}
