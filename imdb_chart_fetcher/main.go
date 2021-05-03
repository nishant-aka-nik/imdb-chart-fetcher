package main

import (
	"IMDb_Chart_Fetcher/controllers"
	"fmt"
	"log"
	"os"
)

func main() {

	//Getting the data from the CLI
	chart_url, items_count, err := controllers.GetArguments(os.Args)

	if err != nil {
		log.Println(err)
		return
	}

	//Scraping the chart page for the link of movies
	links, err := controllers.ScrapeItemsLink(chart_url, items_count)

	if err != nil {
		log.Println(err)
		return
	}

	//Scraping the details of movies
	itemsWithAttributes, err := controllers.GetAllItemDetails(links)

	if err != nil {
		log.Println(err)
		return
	}

	//Displaying the final result
	fmt.Println(itemsWithAttributes)

}
