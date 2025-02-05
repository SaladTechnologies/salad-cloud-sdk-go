package webhooksecretkey

import (
	"encoding/json"
)

// Represents a webhook secret key
type WebhookSecretKey struct {
	// The webhook secret key
	SecretKey *string `json:"secret_key,omitempty" required:"true"`
	touched   map[string]bool
}

func (w *WebhookSecretKey) GetSecretKey() *string {
	if w == nil {
		return nil
	}
	return w.SecretKey
}

func (w *WebhookSecretKey) SetSecretKey(secretKey string) {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["SecretKey"] = true
	w.SecretKey = &secretKey
}

func (w *WebhookSecretKey) SetSecretKeyNil() {
	if w.touched == nil {
		w.touched = map[string]bool{}
	}
	w.touched["SecretKey"] = true
	w.SecretKey = nil
}

func (w WebhookSecretKey) MarshalJSON() ([]byte, error) {
	data := make(map[string]any)

	if w.touched["SecretKey"] && w.SecretKey == nil {
		data["secret_key"] = nil
	} else if w.SecretKey != nil {
		data["secret_key"] = w.SecretKey
	}

	return json.Marshal(data)
}

func (w WebhookSecretKey) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WebhookSecretKey to string"
	}
	return string(jsonData)
}
