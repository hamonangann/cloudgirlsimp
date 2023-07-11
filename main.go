package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"os"
	"strings"
)

const ResultDir = "./result/"

func main() {
	c := colly.NewCollector()

	c.OnHTML(`img[class="article-image"]`, func(e *colly.HTMLElement) {
		link := e.Attr("src")
		e.Request.Visit(link)
	})

	c.OnError(func(r *colly.Response, e error) {
		log.Println("error:", e, r.Request.URL, string(r.Body))
	})

	c.OnResponse(func(r *colly.Response) {
		if strings.Index(r.Headers.Get("Content-Type"), "image") > -1 {
			_, err := os.Stat(ResultDir + r.FileName())
			if os.IsNotExist(err) { // download if file doesn't exist
				fmt.Println("downloading image:", r.FileName())
				r.Save(ResultDir + r.FileName())
				return
			}
		}
	})

	c.Visit("https://thecloudgirl.dev/sketchnote.html")
}
