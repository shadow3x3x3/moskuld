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

	doc.Find("div.DivSeat").Each(func(i int, s *goquery.Selection) {
		notBookedSeat, found := s.Find(".label-info").Attr("title")
		if found {
			fmt.Println("Can book seat:", notBookedSeat)
		}

		beBookedSeat, found := s.Find(".label-danger").Attr("title")
		if found {
			fmt.Println("Can't book seat:", beBookedSeat)
		}
	})

}
