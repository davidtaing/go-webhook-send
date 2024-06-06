package webhook

import (
	"net/http"
)

func BuildRequest(url string, p Payload) (*http.Request, error) {
	// buffer for our http request body
	bodyBuf, err := p.ToBuffer()

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bodyBuf)
	if err != nil {
		return nil, err
	}

	// buffer for generating signatures
	sigBuf, err := p.ToBuffer()

	if err != nil {
		return nil, err
	}

	signature, err := NewSignature(sigBuf.Bytes(), secret)
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
