package viewshow

import (
	"encoding/json"
	"fmt"
	"log"
	"moskuld/internal/pkg/util"
	"moskuld/pkg/movie"
	"net/http"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	getMoviesURL        = "https://www.vscinemas.com.tw/vsweb/api/GetLstDicMovie"
	getMoviesTimeURL    = "https://www.vscinemas.com.tw/vsweb/api/GetLstDicDate"
	getMoviesSessionURL = "https://www.vscinemas.com.tw/vsweb/api/GetLstDicSession"
	getSessionSeatsURL  = "https://sales.vscinemas.com.tw/VoucherTicketing/SessionSeats.aspx"
)

// MovieSession represents the providing session of the movie
type MovieSession struct {
	Value string `json:"strValue"`
	Text  string `json:"strText"`
}

// MovieDate represents the providing date of the movie
type MovieDate struct {
	Text      string `json:"strText"`
	TimeValue string `json:"strValue"`
}

// Movie represents the movie information
type Movie struct {
	Name string `json:"strText"`
	ID   string `json:"strValue"`
}

// Seat represents the seat information of movie
type Seat struct {
	Idle   []string
	Booked []string
}

// GetAll returns a list of all movies by cinemaID
func getAllMovie(cinemaID string) ([]*movie.Movie, error) {
	url := fmt.Sprintf("%s?cinema=%s", getMoviesURL, cinemaID)

	rawString, err := util.GetBody(url)
	if err != nil {
		return nil, err
	}

	movies := []*Movie{}
	if err := json.Unmarshal(rawString, &movies); err != nil {
		log.Println(err)
		return nil, err
	}

	var respMovies []*movie.Movie
	for _, m := range movies {
		respMovies = append(respMovies, &movie.Movie{
			Name: m.Name,
			ID:   m.ID,
		})
	}

	var wg sync.WaitGroup
	for _, rm := range respMovies {
		wg.Add(1)

		go func(m *movie.Movie) {
			defer wg.Done()
			dates, err := getMovieDate(cinemaID, m.ID)
			if err != nil {
				return
			}

			m.Dates = dates

			var dateWg sync.WaitGroup
			for _, d := range m.Dates {
				dateWg.Add(1)

				go func(cinemaID, movieID string, d *movie.Date) {
					defer dateWg.Done()
					sessions, err := getMovieSession(cinemaID, movieID, d.TimeValue)
					if err != nil {
					}
					d.Sessions = sessions
				}(cinemaID, m.ID, d)

				dateWg.Wait()
			}
		}(rm)
	}

	wg.Wait()
	return respMovies, nil
}

func getMovieDate(cinemaID, movieID string) ([]*movie.Date, error) {
	url := fmt.Sprintf("%s?cinema=%s&movie=%s", getMoviesTimeURL, cinemaID, movieID)

	rawString, err := util.GetBody(url)
	if err != nil {
		return nil, err
	}

	movieDates := []*MovieDate{}
	if err := json.Unmarshal(rawString, &movieDates); err != nil {
		return nil, err
	}

	respMovieDates := []*movie.Date{}
	for _, md := range movieDates {
		respMovieDates = append(respMovieDates, &movie.Date{
			Text:      md.Text,
			TimeValue: md.TimeValue,
		})
	}

	return respMovieDates, nil
}

func getMovieSession(cinemaID, movieID, timeValue string) ([]*movie.Session, error) {
	url := fmt.Sprintf("%s?cinema=%s&movie=%s&date=%s", getMoviesSessionURL, cinemaID, movieID, timeValue)

	rawString, err := util.GetBody(url)
	if err != nil {
		return nil, err
	}

	movieSessions := []*MovieSession{}
	if err := json.Unmarshal(rawString, &movieSessions); err != nil {
		return nil, err
	}

	respSessions := []*movie.Session{}
	for _, s := range movieSessions {
		respSessions = append(respSessions, &movie.Session{
			Value: s.Value,
			Text:  s.Text,
		})
	}

	return respSessions, nil

}

func getSeats(sessionValue string) (*Seat, error) {
	url := fmt.Sprintf("%s?%s", getSessionSeatsURL, sessionValue)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {

	}
	req.Host = `sales.vscinemas.com.tw`
	req.Header.Set("Referer", `https://www.vscinemas.com.tw/vsweb/`)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	ss := &Seat{}

	start := time.Now()
	doc.Find("div.DivSeat").Each(func(_ int, s *goquery.Selection) {
		notBookedSeat, found := s.Find(".label-info").Attr("title")
		if found {
			ss.Idle = append(ss.Idle, notBookedSeat)
		}

		beBookedSeat, found := s.Find(".label-danger").Attr("title")
		if found {
			ss.Booked = append(ss.Booked, beBookedSeat)
		}
	})
	log.Printf("Parse Seats took %s\n", time.Since(start))

	return ss, nil
}
