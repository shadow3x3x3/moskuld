package cinema

import (
	"encoding/json"
	"moskuld/internal/pkg/movie"
	"moskuld/internal/pkg/util"
)

// Cinema represents the cinema information
type Cinema struct {
	Name string `json:"strText"`
	ID   string `json:"strValue"`
}

type cinema struct {
	Name string `json:"strText"`
	ID   string `json:"strValue"`
}

const (
	getCinemasURL = "https://www.vscinemas.com.tw/vsweb/api/GetLstDicCinema"
)

// GetAll returns all Cinemas
func GetAll() ([]*Cinema, error) {
	respBody, err := util.GetBody(getCinemasURL)
	if err != nil {
		return nil, err
	}

	cinemas := []*Cinema{}
	if err := json.Unmarshal(respBody, &cinemas); err != nil {
		return nil, err
	}

	return cinemas, nil
}

func HaveMovie(cinemaID, movieID string) bool {
	movies, err := movie.GetAll(cinemaID)
	if err != nil {
		return false
	}

	for _, m := range movies {
		if m.ID == movieID {
			return true
		}
	}

	return false
}
