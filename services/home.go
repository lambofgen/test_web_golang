package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/lambofgen/testweb/models"
)

func GetResponse() (*models.Response, error) {
	resp, err := http.Get("https://goo.gl/hc5UJ3")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var data models.Response
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
