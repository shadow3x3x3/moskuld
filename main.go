package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	queryCinemaURL = "https://www.vscinemas.com.tw/api/GetLstDicCinema"
	queryMoviesURL = "https://www.vscinemas.com.tw/api/GetLstDicMovie"
	queryTimeURL   = "https://www.vscinemas.com.tw/api/GetLstDicDate"
	querySeatURL   = "https://sales.vscinemas.com.tw/VoucherTicketing/SessionSeats.aspx"
)

// Cinema represents the cinema information
type Cinema struct {
	Name string `json:"strText"`
	ID   string `json:"strValue"`
}

// Movie represents the cinema information
type Movie struct {
	Name string `json:"strText"`
	ID   string `json:"strValue"`
}

// MovieTime represents the time of movie information
type MovieTime struct {
	Text      string `json:"strText"`
	TimeValue string `json:"strValue"`
}

func main() {
	// getAllCinemaArray()
	getAllMovieArray("19|LKMP")
	// file, err := os.OpenFile("test_data/test.html", os.O_RDONLY, 04444)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// defer file.Close()

	// doc, err := goquery.NewDocumentFromReader(file)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// doc.Find("div.DivSeat").Each(func(i int, s *goquery.Selection) {
	// 	notBookedSeat, found := s.Find(".label-info").Attr("title")
	// 	if found {
	// 		fmt.Println("Can book seat:", notBookedSeat)
	// 	}

	// 	beBookedSeat, found := s.Find(".label-danger").Attr("title")
	// 	if found {
	// 		fmt.Println("Can't book seat:", beBookedSeat)
	// 	}
	// })

}

func getAllCinemaArray() []*Cinema {
	rawString, err := httpGetBodyBytes(queryCinemaURL)
	if err != nil {
		log.Println(err)
		return nil
	}

	allCinema := []*Cinema{}

	if err := json.Unmarshal(rawString, &allCinema); err != nil {
		log.Println(err)
		return nil
	}

	return allCinema
}

func getAllMovieArray(cinemaID string) []*Movie {
	targetURL := fmt.Sprintf("%s?cinema=%s", queryMoviesURL, cinemaID)
	rawString, err := httpGetBodyBytes(targetURL)
	if err != nil {
		log.Println(err)
		return nil
	}

	allMovie := []*Movie{}

	if err := json.Unmarshal(rawString, &allMovie); err != nil {
		log.Println(err)
		return nil
	}

	return allMovie
}

func getTimeArray(cinemaID string, movieID string) []*MovieTime {
	targetURL := fmt.Sprintf("%s?cinema=%s&movie=%s", queryTimeURL, cinemaID, movieID)
	rawString, err := httpGetBodyBytes(targetURL)
	if err != nil {
		log.Println(err)
		return nil
	}

	allMovieTime := []*MovieTime{}

	if err := json.Unmarshal(rawString, &allMovieTime); err != nil {
		log.Println(err)
		return nil
	}

	return allMovieTime
}

func httpGetBodyBytes(URL string) ([]byte, error) {
	resp, err := http.Get(URL)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	defer resp.Body.Close()

	rawByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return rawByte, nil
}
