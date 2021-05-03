package controllers

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

func GetArguments(args []string) (string, int, error) {
	if len(os.Args) != 3 {
		// log.Println("Error Occurred : missing command line arguments")
		return "", 0, errors.New("error Occurred : missing command line arguments")
	}

	chart_url := os.Args[1]

	if !strings.Contains(chart_url, "www.imdb.com") {
		// log.Println("Error Occurred : Invalid Chart url")
		return "", 0, errors.New("error Occurred : Invalid Chart url")
	}

	chart_url = chart_url[1 : len(chart_url)-1]

	items_count, err := strconv.Atoi(os.Args[2])

	if err != nil {
		// log.Println("Error Occurred : Invalid item count")
		return "", 0, errors.New("error Occurred : Invalid item count")
	}

	return chart_url, items_count, nil
}
