package tests

import (
	"net/http"
	"testing"

	"github.com/davidtaing/go-webhook-send/internal/webhook"
)

func TestBuildRequestSetsSignatureHeader(t *testing.T) {
	url := "http://localhost"
	payload := webhook.Payload{
		"hello": "world",
	}

	req, err := webhook.BuildRequest(url, payload)

	if err != nil {
		t.Fatalf("error building request. Err: %v", err)
	}

	signature := req.Header.Get("signature")

	if signature == "" {
		t.Errorf("expected signature header to be set; got %v", signature)
	}
}

func TestBuildRequestCreatesPostRequest(t *testing.T) {
	url := "http://localhost"
	payload := webhook.Payload{
		"hello": "world",
	}

	req, err := webhook.BuildRequest(url, payload)

	if err != nil {
		t.Fatalf("error building request. Err: %v", err)
	}

	if req.Method != http.MethodPost {
		t.Errorf("expected HTTP POST method; got %s", req.Method)
	}
}
