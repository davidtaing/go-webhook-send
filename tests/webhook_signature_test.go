package tests

import (
	"bytes"
	"testing"

	"github.com/davidtaing/go-webhook-send/internal/webhook"
)

func TestNewSignature(t *testing.T) {
	const secret = "secret"

	payload := webhook.Payload{
		"hello": "world",
	}

	buf, err := payload.ToBuffer()

	if err != nil {
		t.Fatalf("error creating signature. Err: %v", err)
	}

	want := "VswE8kWdwK-wh9tYr2ZZ33LH2oejS2JgCb55YvAKc3Q"
	signature, err := webhook.NewSignature(buf.Bytes(), secret)

	if err != nil {
		t.Fatalf("error creating signature. Err: %v", err)
	}

	if want != signature {
		t.Errorf("expected signature to be %s; got %s", want, signature)
	}
}

func TestNewSignatureWithDifferentPayloads(t *testing.T) {
	const secret = "secret"

	p1 := webhook.Payload{
		"hello": "world",
	}
	p2 := webhook.Payload{
		"hellos": "world",
	}

	buf1, _ := p1.ToBuffer()
	buf2, _ := p2.ToBuffer()

	sig1, err := webhook.NewSignature(buf1.Bytes(), secret)

	if err != nil {
		t.Fatalf("error creating signature from payload 1. Err: %v", err)
	}

	sig2, err := webhook.NewSignature(buf2.Bytes(), secret)

	if err != nil {
		t.Fatalf("error creating signature from payload 2. Err: %v", err)
	}

	if sig1 == sig2 {
		t.Errorf("expected different signatures for different payloads; \n payload1: %s\n, sig1: %s\n, payload2: %s\n, sig2: %s", p1, p2, sig1, sig2)
	}
}

func TestNewSignatureWithReorderedKeys(t *testing.T) {
	const secret = "secret"

	// Maps will reorder the key/value pairs, so we're using raw JSON strings
	payload1 := "{\"hello\":\"hello\", \"world\":\"world\"}"
	payload2 := "{\"world\":\"world\", \"hello\":\"hello\"}"

	buf1 := bytes.NewBufferString(payload1)
	buf2 := bytes.NewBufferString(payload2)

	sig1, err := webhook.NewSignature(buf1.Bytes(), secret)

	if err != nil {
		t.Fatalf("error creating signature from payload 1. Err: %v", err)
	}

	sig2, err := webhook.NewSignature(buf2.Bytes(), secret)

	if err != nil {
		t.Fatalf("error creating signature from payload 2. Err: %v", err)
	}

	if sig1 == sig2 {
		t.Errorf("expected different signatures when reordered keys are reordered; \n payload1: %s\n, sig1: %s\n, payload2: %s\n, sig2: %s", payload1, payload2, sig1, sig2)
	}
}
