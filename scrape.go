package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"strings"
	"github.com/gocolly/colly"
)


type Recipe struct {
	Title string `json:"title"`
	Ingredients []string `json:"ingredients"`
	Directions []string `json:"directions"`
}

// Gets info for a recipe
func getRecipe(url string) Recipe {
	// Instantiate default collector
	c := colly.NewCollector()

	// Set up selectors and storage
	title_sel := "h1.recipe-title"
	ingredients_list := "ul.recipe-ingredients__list li"
	directions_list := "ol.recipe-directions__list li span"
	var title string
	var ingredients []string
	var directions []string
	
	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Catch error
	c.OnError(func(_ *colly.Response, err error) {
		if err != nil {
			panic(err)
		}
	})

	// Get the title
	c.OnHTML(title_sel, func(e *colly.HTMLElement) {
        title = e.Text
	})

	// Get the ingredients
	c.OnHTML(ingredients_list, func(e *colly.HTMLElement) {
		ingredients = append(ingredients, e.Text)
	})

	// Get the directions
	c.OnHTML(directions_list, func(e *colly.HTMLElement) {
		directions = append(directions, e.Text)
	})

	// Start scraping on given url
	c.Visit(url)
	
	// Save as recipe
	recipe := Recipe{
		Title: title,
		Ingredients: ingredients,
		Directions: directions,
	}

	return recipe
}

// Extract url from text file
func getURL(filePath string) []string {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	urls := strings.Split(string(file), "\n")

	return urls
}

// Output json file
func save(recipes []Recipe) {
	file, err := json.MarshalIndent(recipes, "", " ")
	if err != nil {
		panic(err)
	}
	
	_ = ioutil.WriteFile("recipes.json", file, 0644)
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {

	var recipes []Recipe
	var filePath string

	if filePath == "" {
		filePath = "recipe-urls.txt"
	}

	urls := getURL(filePath)

	for _, r := range urls {
		recipes = append(recipes, getRecipe(r))
	}
	
	save(recipes)
}
