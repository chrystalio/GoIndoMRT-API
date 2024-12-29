package client

import (
	"errors"
	"io/ioutil"
	"net/http"
)

func DoRequest(Client *http.Client, url string) ([]byte, error) {
	resp, err := Client.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("Failed to fetch data: " + resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
