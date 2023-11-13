package scraper

import (
	"errors"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/tayfun8/scanwp/models"
	"github.com/tayfun8/scanwp/utils"
)

var plugins []models.Plugins

func getPlugins() {
	c := colly.NewCollector()
	// plugins := []Plugins{}

	defer wg.Done()

	// Get Plugins On the Wordpress Site
	c.OnHTML("link[rel='stylesheet']", func(h *colly.HTMLElement) {
		attr := h.Attr("href")
		// attr = strings.ReplaceAll(attr, "./", "/")
		if strings.Contains(attr, "/wp-content/plugins/") {
			// fmt.Println(attr)
			vers, err := getPluginVersion(attr)
			if err != nil {

			}

			// getPluginName(URL)
			pluginName := getPluginName(attr)

			newPlugin := models.Plugins{
				Name:    pluginName,
				Version: vers,
			}

			if status := IsPluginContains(plugins, newPlugin); status == false {
				plugins = append(plugins, newPlugin)
			}
			// plugins = append(plugins, newPlugin)
		}
	})

	c.OnHTML("script[type='text/javascript']", func(h *colly.HTMLElement) {
		attr := h.Attr("src")
		// attr = strings.ReplaceAll(attr, "./", "/")
		if strings.Contains(attr, "/wp-content/plugins/") {
			// fmt.Println(attr)
			vers, err := getPluginVersion(attr)
			if err != nil {

			}

			pluginName := getPluginName(attr)

			newPlugin := models.Plugins{
				Name:    pluginName,
				Version: vers,
			}
			if status := IsPluginContains(plugins, newPlugin); status == false {
				plugins = append(plugins, newPlugin)
			}

			// plugins = append(plugins, newPlugin)
		}
	})

	// VISIT SITE
	if err := c.Visit(Url); err != nil {
		log.Fatal("Can not connected")
	}

	utils.WritePlugins(plugins)
}

// GET PLUGIN VERSION
func getPluginVersion(url string) (string, error) {
	parts := strings.Split(url, "ver=")
	if len(parts) > 1 {
		// ver= sonrasındaki değeri al
		version := strings.Split(parts[1], "&")[0]
		return version, nil
	}
	return "", errors.New("Plugin Version Not Found")
}

// GET PLUGIN NAME
func getPluginName(url string) string {

	parts := strings.Split(url, "plugins/")
	if len(parts) < 2 {
		return ""
	}

	pluginPath := parts[1]
	slashIndex := strings.Index(pluginPath, "/")
	if slashIndex >= 0 {
		pluginPath = pluginPath[:slashIndex]
	}

	return pluginPath

}

func IsPluginContains(plugins []models.Plugins, plugin models.Plugins) bool {
	for _, v := range plugins {
		if v.Name == plugin.Name && v.Version == plugin.Version {
			return true
		}
	}
	return false
}
