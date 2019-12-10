package viewshow

import (
	"encoding/json"
	"moskuld/internal/pkg/util"
	"moskuld/pkg/cinema"
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
		rc := cinema.Cinema(*c)
		respCinemas = append(respCinemas, &rc)
	}

	return respCinemas, nil
}
