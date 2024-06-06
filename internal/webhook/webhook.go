package webhook

import (
	"github.com/oklog/ulid/v2"
)

const secret = "secret"

type Payload map[string]interface{}

type example_payload struct {
	Scope     string `json:"scope"`
	StoreID   string `json:"store_id"`
	Data      data   `json:"data"`
	Hash      string `json:"hash"`
	CreatedAt int64  `json:"created_at"`
	Producer  string `json:"producer"`
}

type data struct {
	Type string `json:"type"`
	ID   int    `json:"id"`
}

// This was taken from the BigCommerce API documentation.
// https://developer.bigcommerce.com/docs/integrations/webhooks
//
// It would be interesting to implement build a webhook payload based on items in
// a database. For now though, it will be out of scope.
func NewWebhook() example_payload {
	return example_payload{
		Scope:     "store/order/created",
		StoreID:   "1025646",
		Data:      data{Type: "order", ID: 250},
		Hash:      ulid.Make().String(),
		CreatedAt: 1561479335,
		Producer:  "stores/{store_hash}",
	}
}
