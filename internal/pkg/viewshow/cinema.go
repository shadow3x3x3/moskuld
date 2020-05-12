package viewshow

import (
	"errors"
	"moskuld/internal/pkg/util"
	"moskuld/pkg/cinema"
	"strconv"
	"strings"
)

// Cinema represents the cinema information
type Cinema struct {
	Name string `json:"strText"`
	ID   string `json:"strValue"`
}

const (
	getCinemasURL = "https://www.vscinemas.com.tw/vsweb/api/GetLstDicCinema"
)

func getAllCinema() ([]*cinema.Cinema, error) {
	respBody, err := util.GetBody(getCinemasURL)
	if err != nil {
		return nil, err
	}

	cinemas := []*Cinema{}
	if err := json.Unmarshal(respBody, &cinemas); err != nil {
		return nil, err
	}

	respCinemas := []*cinema.Cinema{}
	for _, c := range cinemas {
		keyText := strings.Split(c.ID, "|")[0]
		key, err := strconv.Atoi(keyText)
		if err != nil {
			return nil, errors.New("Fetching cinemas error.")
		}
		respCinemas = append(respCinemas, &cinema.Cinema{
			Name: c.Name,
			ID:   c.ID,
			Key:  key,
		})
	}

	return respCinemas, nil
}

func getCinemasTable() (map[int]*cinema.Cinema, error) {
	respBody, err := util.GetBody(getCinemasURL)
	if err != nil {
		return nil, err
	}

	cinemas := []*Cinema{}
	if err := json.Unmarshal(respBody, &cinemas); err != nil {
		return nil, err
	}

	respCinemas := make(map[int]*cinema.Cinema, len(cinemas))
	for _, c := range cinemas {
		keyText := strings.Split(c.ID, "|")[0]
		key, err := strconv.Atoi(keyText)
		if err != nil {
			return nil, errors.New("Fetching cinemas error.")
		}
		respCinemas[key] = &cinema.Cinema{
			Name: c.Name,
			ID:   c.ID,
			Key:  key,
		}
	}

	return respCinemas, nil
}
