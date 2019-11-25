package movie

import (
	"encoding/json"
	"fmt"
	"log"
	"moskuld/internal/pkg/util"
)

const (
	getMoviesURL     = "https://www.vscinemas.com.tw/vsweb/api/GetLstDicMovie"
	getMoviesTimeURL = "https://www.vscinemas.com.tw/vsweb/api/GetLstDicDate"
)

// MovieTime represents the time of movie information
type MovieTime struct {
	Text      string `json:"strText"`
	TimeValue string `json:"strValue"`
}

// Movie represents the movie information
type Movie struct {
	Name    string       `json:"strText"`
	ID      string       `json:"strValue"`
	Session []*MovieTime `json:",omitempty"`
}

// GetAll returns a list of all movies by cinemaID
func GetAll(cinemaID string) ([]*Movie, error) {
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

	return movies, nil
}
