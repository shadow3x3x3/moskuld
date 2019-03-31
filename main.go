package main

import (
	"fmt"
	"log"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	file, err := os.OpenFile("test_data/test.html", os.O_RDONLY, 04444)
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		log.Fatalln(err)
	}

	doc.Find(".DivSeat").Each(func(i int, s *goquery.Selection) {
		labelTitle, found := s.Find(".label").Attr("title")
		if found {
			fmt.Println(labelTitle)
		}
	})

}
