package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gocolly/colly"
)

type Product struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	Subtitle string `json:"subtitle"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.nike.com"),
	)

	var products []Product

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Println("Received response with status code:", r.StatusCode)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Error:", r.Request.URL, "failed to respond with status code", r.StatusCode, "Error:", err)
	})

	c.OnHTML("div.product-card", func(h *colly.HTMLElement) {
		product := Product{
			Name:     h.ChildText("div.product-card__title"),
			Subtitle: h.ChildText("div.product-card__subtitle"),
			Price:    h.ChildText("div.product-card__price"),
		}
		products = append(products, product)
	})

	c.Visit("https://www.nike.com/in/w/lifestyle-shoes-13jrmzy7ok")

	jsonData, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile("products.json", jsonData, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
