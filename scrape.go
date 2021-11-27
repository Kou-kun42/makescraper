package main

import (
	"fmt"
	//"io/ioutil"
	//"strings"
	//"encoding/json"
	"github.com/gocolly/colly"
)


type Recipe struct {
	Title string `json:"title"`
	Ingredients []string `json:"ingredients"`
	Directions []string `json:"directions"`
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
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

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.tasteofhome.com/recipes/rum-balls/")
	
	// Save as recipe
	recipe := Recipe{
		Title: title,
		Ingredients: ingredients,
		Directions: directions,
	}
	fmt.Println(recipe)
}
