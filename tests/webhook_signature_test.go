package tests

import (
	"testing"

	"github.com/davidtaing/go-webhook-send/internal/webhook"
)

func TestNewSignatureValidSignature(t *testing.T) {
	const secret = "secret"

	payload := webhook.Payload{
		"hello": "world",
	}

	want := "VswE8kWdwK-wh9tYr2ZZ33LH2oejS2JgCb55YvAKc3Q"
	signature, err := webhook.NewSignature(payload, secret)

	if err != nil {
		t.Fatalf("error creating signature. Err: %v", err)
	}

	if want != signature {
		t.Errorf("expected signature to be %s; got %s", want, signature)
	}
}

func TestNewSignatureDifferentSignaturesForDifferentPayloads(t *testing.T) {
	const secret = "secret"

	payload1 := webhook.Payload{
		"hello": "world",
	}

	payload2 := webhook.Payload{
		"hellos": "world",
	}

	sig1, err := webhook.NewSignature(payload1, secret)

	if err != nil {
		t.Fatalf("error creating signature from payload 1. Err: %v", err)
	}

	sig2, err := webhook.NewSignature(payload2, secret)

	if err != nil {
		t.Fatalf("error creating signature from payload 2. Err: %v", err)
	}

	if sig1 == sig2 {
		t.Errorf("expected signature to be different for different payloads; \n payload1: %s\n, sig1: %s\n, payload2: %s\n, sig2: %s", payload1, payload2, sig1, sig2)
	}
}
