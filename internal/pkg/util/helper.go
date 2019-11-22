package util

import (
	"io/ioutil"
	"net/http"
)

// Get Body returns bytes of request body
func GetBody(URL string) ([]byte, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	rawByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return rawByte, nil
}
