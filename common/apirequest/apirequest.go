package apirequest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

const (
	Timeout int64 = 10
)

var Identity = ""

/*
	Executes an API request and populates the data with the response
*/
func Do(req *http.Request, data interface{}) (int, error) {
	client := http.Client{
		Timeout: time.Duration(Timeout) * time.Second,
	}

	resp, err := client.Do(req)

	if err != nil {
		return -1, err
	}

	defer resp.Body.Close()

	if data != nil {
		json.NewDecoder(resp.Body).Decode(data)
	}

	return resp.StatusCode, nil
}

/*
	Executes an API GET request and populates the data with the response
*/
func Get(url string, data interface{}) (int, error) {
	return doRequest("GET", url, nil, data)
}

/*
	Executes an API POST request and populates the data with the response
*/
func Post(url string, payload interface{}, data interface{}) (int, error) {
	return doRequest("POST", url, payload, data)
}

/*
	Executes an API PUT request and populates the data with the response
*/
func Put(url string, payload interface{}, data interface{}) (int, error) {
	return doRequest("PUT", url, payload, data)
}

/*
	Executes an API DELETE request and populates the data with the response
*/
func Delete(url string, data interface{}) (int, error) {
	return doRequest("DELETE", url, nil, data)
}

/*
	Builds a request, executes it, and then decodes the response into data
*/
func doRequest(method string, url string, payload interface{}, data interface{}) (int, error) {
	var req *http.Request
	var err error

	if payload != nil {
		var body bytes.Buffer
		json.NewEncoder(&body).Encode(payload)

		req, err = http.NewRequest(method, url, &body)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	if err != nil {
		return -1, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("HackIllinois-Identity", Identity)

	return Do(req, data)
}

/*
	Set the identity to place into HackIllinois-Identity when making requests
*/
func SetIdentity(identity string) {
	Identity = identity
}
