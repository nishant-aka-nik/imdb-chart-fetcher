package controllers

import (
	"errors"
	"strings"

	"github.com/gocolly/colly"
)

func ScrapeItemsLink(chart_url string, items_count int) ([]string, error) {
	var links []string

	c := colly.NewCollector(
		colly.CacheDir("./models/imdb_cache"),
	)

	c.OnHTML("#main > div > span > div > div > div.lister > table > tbody", func(e *colly.HTMLElement) {
		e.ForEachWithBreak("#main > div > span > div > div > div.lister > table > tbody > tr", func(i int, h *colly.HTMLElement) bool {

			link := strings.Split(h.ChildAttr("td.titleColumn > a", "href"), "?")

			uri := "https://www.imdb.com"

			links = append(links, uri+link[0])

			return (i + 1) != items_count

		})

	})

	c.Visit(chart_url)

	if len(links) == 0 {
		return links, errors.New("scrapping error")
	}

	return links, nil

}
