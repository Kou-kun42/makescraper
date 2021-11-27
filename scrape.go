package main

import (
	"fmt"
	//"io/ioutil"
	//"strings"
	//"encoding/json"
	"github.com/gocolly/colly"
)


type recipe struct {
	Name string `json:"name"`
	Ingredients []string `json:"ingredients"`
	Directions []string `json:"directions"`
}

// main() contains code adapted from example found in Colly's docs:
// http://go-colly.org/docs/examples/basic/
func main() {
	// Instantiate default collector
	c := colly.NewCollector()

	//title := "h1.recipe-title"
	ingredients_list := "ul.recipe-ingredients__list li"
	directions_list := "ol.recipe-directions__list li span"
	var ingredients []string

	// On every a element which has href attribute call callback
	c.OnHTML(ingredients_list, func(e *colly.HTMLElement) {
                ing := e.Text

				ingredients = append(ingredients, ing)
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Start scraping on https://hackerspaces.org
	c.Visit("https://www.tasteofhome.com/recipes/rum-balls/")
	fmt.Println(ingredients)
}
