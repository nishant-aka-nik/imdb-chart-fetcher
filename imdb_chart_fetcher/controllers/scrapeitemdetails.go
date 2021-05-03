package controllers

import (
	"IMDb_Chart_Fetcher/models"
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

type MovieAttributes models.MovieAttributes

func GetAllItemDetails(links []string) ([]string, error) {

	var itemsWithAttributes []string

	c := colly.NewCollector(
		colly.CacheDir("./models/imdb_cache"),
	)

	for index, link := range links {

		MovieAttributes := MovieAttributes{}

		jsonString, err := MovieAttributes.ScrapeItemDetails(link, c)
		// jsonString, err := MovieAttributes.ScrapeItemDetails(link)
		if index+1 == 1 {
			log.Printf("Scraping %dst movie details.....", index+1)
		} else if index+1 == 2 {
			log.Printf("Scraping %dnd movie details.....", index+1)

		} else if index+1 == 3 {
			log.Printf("Scraping %drd movie details.....", index+1)
		} else {
			log.Printf("Scraping %dth movie details.....", index+1)
		}

		if err != nil {
			log.Println(err)
			return itemsWithAttributes, errors.New("error Occurred : missing command line arguments")
		}

		itemsWithAttributes = append(itemsWithAttributes, string(jsonString))

	}

	return itemsWithAttributes, nil
}

//, c *colly.Collector
func (MovieAttributesObj *MovieAttributes) ScrapeItemDetails(link string, c *colly.Collector) ([]byte, error) {

	var summarylinks []string

	// c := colly.NewCollector()

	// fmt.Println("Type of C ", reflect.TypeOf(c))

	c.OnHTML("#wrapper", func(e *colly.HTMLElement) {

		title := e.ChildText("div.vital > div.title_block > div > div.titleBar > div.title_wrapper > h1")
		MovieAttributesObj.Title = title[:len(title)-8]

		releaseDate := e.ChildText("#titleYear > a")
		MovieAttributesObj.Movie_release_year, _ = strconv.Atoi(releaseDate)

		ImdbRating := e.ChildText("#title-overview-widget > div.vital > div.title_block > div > div.ratings_wrapper > div.imdbRating > div.ratingValue > strong > span")

		MovieAttributesObj.Imdb_rating, _ = strconv.ParseFloat(ImdbRating, 64)

		summarylinks = append(summarylinks, "https://www.imdb.com"+e.ChildAttr("#titleStoryLine > span.see-more.inline > a:nth-child(1)", "href"))

		subTextBar := e.ChildText("div.vital > div.title_block > div > div.titleBar > div.title_wrapper > div.subtext")

		subTextBarArr := strings.Split(subTextBar, "|")

		trimmedGenre := strings.TrimSpace(subTextBarArr[2])
		genre := strings.ReplaceAll(trimmedGenre, "\n", " ")
		MovieAttributesObj.Genre = genre

		duration := strings.ReplaceAll(subTextBarArr[1], "\n", " ")
		duration = strings.TrimSpace(duration)
		MovieAttributesObj.Duration = duration

	})

	c.Visit(link)

	clonedObj := c.Clone()

	clonedObj.OnHTML("#main > section", func(e *colly.HTMLElement) {
		e.ForEachWithBreak("#plot-summaries-content > li", func(i int, h *colly.HTMLElement) bool {

			summary := h.ChildText(".ipl-zebra-list__item > p")
			MovieAttributesObj.Summary = summary

			return i < 0
		})
	})

	for _, summarylink := range summarylinks {
		clonedObj.Visit(summarylink)

	}

	jsonData, err := json.Marshal(MovieAttributesObj)

	if err != nil {
		return jsonData, err
	}

	return jsonData, nil
}
