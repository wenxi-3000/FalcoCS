package connect

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

type resp struct {
	body []byte
	code int
}

func (c *connector) NewRequest(method string, url string, body []byte) (*resp, error) {

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Cookie", c.token)
	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("failed with status code %d", res.StatusCode)
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}
	return &resp{body: bodyBytes, code: res.StatusCode}, nil
}
