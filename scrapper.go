package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func getUserAvatarUrl(userID string) string {
	c := colly.NewCollector()

	image := ""

	// Find and visit all links
	c.OnHTML("body", func(e *colly.HTMLElement) {
		fmt.Println(e.DOM.Html())
		image = e.ChildAttr(".speakerinfo img", "src")
		fmt.Println(image)
	})

	c.Visit("https://www.opensouthcode.org/users/" + userID)
	return image
}

// func main() {
// 	fmt.Println(getUserAvatarUrl("562"))
// }
