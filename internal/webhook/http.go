package webhook

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func BuildRequest(url string, payload Payload) (*http.Request, error) {
	p := bytes.Buffer{}
	err := json.NewEncoder(&p).Encode(payload)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, &p)
	if err != nil {
		return nil, err
	}

	signature, err := NewSignature(payload, secret)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Signature", signature)

	return req, nil
}

func Send(url string, payload Payload) (*http.Response, error) {
	req, err := BuildRequest(url, payload)

	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}

	defer resp.Body.Close()

	return resp, nil
}
