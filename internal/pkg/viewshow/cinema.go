package viewshow

import (
	"encoding/json"
	"moskuld/internal/pkg/util"
)

// Cinema represents the cinema information
type Cinema struct {
	Name string `json:"strText"`
	ID   string `json:"strValue"`
}

const (
	getCinemasURL = "https://www.vscinemas.com.tw/vsweb/api/GetLstDicCinema"
)

func getAllCinema() ([]*Cinema, error) {
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
