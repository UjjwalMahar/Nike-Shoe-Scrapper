Sure, here's a breakdown of the code:

```go
package main
```
This line defines the package that this file is part of. In Go, every file must declare the package it belongs to at the top of the file. The `main` package is a special name that tells the Go compiler that this package contains an executable program, not a library.

```go
import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gocolly/colly"
)
```
This block imports other packages that this code depends on. `encoding/json` is used for encoding and decoding JSON data, `fmt` provides functions for formatted I/O, `io/ioutil` provides functions to read and write files, and `github.com/gocolly/colly` is a web scraping library.

```go
type Product struct {
	Name     string `json:"name"`
	Price    string `json:"price"`
	Subtitle string `json:"subtitle"`
}
```
This block defines a new type called `Product`, which is a struct with three fields: `Name`, `Price`, and `Subtitle`. Each field is of type string. The text within backticks are struct field tags, which provide metadata about the struct field. Here, they're used to define how each field should be encoded to or decoded from JSON.

```go
func main() {
```
This line defines the main function. The main function serves as the entry point for the program. When you run the program, this function gets executed first.

```go
c := colly.NewCollector(
	colly.AllowedDomains("www.nike.com"),
)
```
This line creates a new instance of colly's Collector. This object will be used to visit web pages and collect data. The argument passed to `NewCollector` specifies that this collector is only allowed to visit pages on "www.nike.com".

```go
var products []Product
```
This line declares a variable named `products`, which is a slice of `Product`. A slice is a dynamically-sized sequence of array elements. This slice will be used to store all the products that are scraped from the website.

```go
c.OnRequest(func(r *colly.Request) {
	fmt.Println("Visiting", r.URL)
})
```
This block sets up a callback function that will be called each time the collector makes a request. The function takes one argument, which is an object representing the request being made. The function prints out the URL being visited.

```go
c.OnResponse(func(r *colly.Response) {
	fmt.Println("Received response with status code:", r.StatusCode)
})
```
This block sets up another callback function that will be called each time the collector receives a response. The function takes one argument, which is an object representing the response received. The function prints out the status code of the response.

```go
c.OnError(func(r *colly.Response, err error) {
	fmt.Println("Error:", r.Request.URL, "failed to respond with status code", r.StatusCode, "Error:", err)
})
```
This block sets up yet another callback function that will be called if an error occurs during the request. The function takes two arguments: an object representing the response received (if any), and an error object representing the error that occurred. The function prints out details about the error.

```go
c.OnHTML("div.product-card", func(h *colly.HTMLElement) {
	product := Product{
		Name:     h.ChildText("div.product-card__title"),
		Subtitle: h.ChildText("div.product-card__subtitle"),
		Price:    h.ChildText("div.product-card__price"),
	}
	products = append(products, product)
})
```
This block sets up a callback function that will be called each time the collector encounters an HTML element that matches the selector "div.product-card". The function takes one argument, which is an object representing the HTML element. The function creates a new `Product` object using data from child elements of the product card, and appends it to the `products` slice.

```go
c.Visit("https://www.nike.com/in/w/lifestyle-shoes-13jrmzy7ok")
```
This line tells the collector to start visiting pages, beginning with the provided URL.

```go
jsonData, err := json.MarshalIndent(products, "", "  ")
if err != nil {
	fmt.Println(err)
	return
}
```
These lines convert the `products` slice into JSON format. If an error occurs during this process (for example, if one of the products contains data that can't be represented in JSON), it prints out the error and returns from the main function.

```go
err = ioutil.WriteFile("products.json", jsonData, 0644)
if err != nil {
	fmt.Println(err)
}
```
These lines write the JSON data to a file named "products.json". If an error occurs during this process (for example, if there's no permission to write to the file), it prints out the error.

```go
}
```
This line marks the end of the main function. When the main function returns, the program ends.