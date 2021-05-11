package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

func HitsWebGETRequest(url string, params map[string]interface{}, header map[string]string) (statusCode *int, res map[string]interface{}, err error) {
	// Froms query url
	u, err := formQueryURL(url, params)
	if err != nil {
		return nil, nil, err
	}
	log.Println("URL - ", u)

	// Creates new http request
	request, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, nil, err
	}

	// Sets headers
	request.Header.Add("Accept", `*/*`)
	request.Header.Add("User-Agent", `PostmanRuntime/7.26.8`)
	if header != nil {
		for k, v := range header {
			request.Header.Set(k, v)
		}
	}

	// Hits request
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	// Forms response
	json.NewDecoder(resp.Body).Decode(&res)

	// Returns
	return &resp.StatusCode, res, nil
}

func formQueryURL(baseURL string, params map[string]interface{}) (u *url.URL, err error) {
	// Parse url
	u, err = url.Parse(baseURL)
	if err != nil {
		return nil, err
	}

	// Sets query string parameters
	q := u.Query()
	for key, value := range params {
		q.Set(key, value.(string))
	}
	u.RawQuery = q.Encode()

	// Returns
	return u, nil
}
