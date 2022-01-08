package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// RequestAPIGETWithBody make a GET request with body
func RequestAPIGETWithBody(url string, timeout time.Duration, body interface{}, queryParams map[string]string) (bodyResponse []byte, statusCode int, err error) {
	bodyResponse, statusCode, err = makeRequest(url, timeout, http.MethodGet, body, queryParams)
	if err != nil {
		return bodyResponse, statusCode, err
	}

	return bodyResponse, statusCode, nil
}

// RequestAPIPPOST make a POST request
func RequestAPIPPOST(url string, timeout time.Duration, body interface{}, queryParams map[string]string) (bodyResponse []byte, statusCode int, err error) {
	bodyResponse, statusCode, err = makeRequest(url, timeout, http.MethodPost, body, queryParams)
	if err != nil {
		return bodyResponse, statusCode, err
	}

	return bodyResponse, statusCode, nil
}

// RequestAPIPPUT make a PUT request
func RequestAPIPPUT(url string, timeout time.Duration, body interface{}, queryParams map[string]string) (bodyResponse []byte, statusCode int, err error) {
	bodyResponse, statusCode, err = makeRequest(url, timeout, http.MethodPut, body, queryParams)
	if err != nil {
		return bodyResponse, statusCode, err
	}

	return bodyResponse, statusCode, nil
}

// RequestAPIPDELETE make a DELETE request
func RequestAPIPDELETE(url string, timeout time.Duration, body interface{}, queryParams map[string]string) (bodyResponse []byte, statusCode int, err error) {
	bodyResponse, statusCode, err = makeRequest(url, timeout, http.MethodDelete, body, queryParams)
	if err != nil {
		return bodyResponse, statusCode, err
	}

	return bodyResponse, statusCode, nil
}

func makeRequest(url string, timeout time.Duration, method string, body interface{}, queryParams map[string]string) (bodyResponse []byte, statusCode int, err error) {
	client := http.Client{
		Timeout: timeout,
	}

	var requestBody []byte
	if body != nil {
		requestBody, err = json.Marshal(body)
		if err != nil {
			return bodyResponse, statusCode, err
		}
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	request.Header.Set("Content-type", "application/json")

	if err != nil {
		return bodyResponse, statusCode, err
	}

	query := request.URL.Query()
	if len(queryParams) > 0 {
		for key, value := range queryParams {
			query.Add(key, value)
		}

		request.URL.RawQuery = query.Encode()
	}

	resp, err := client.Do(request)
	if err != nil {
		return bodyResponse, statusCode, err
	}

	defer resp.Body.Close()

	bodyResponse, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return bodyResponse, statusCode, err
	}

	statusCode = resp.StatusCode

	return bodyResponse, statusCode, nil
}
