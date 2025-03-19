package webhooksecretkey

import "encoding/json"

// Represents a webhook secret key
type WebhookSecretKey struct {
	// The webhook secret key
	SecretKey *string `json:"secret_key,omitempty" required:"true" maxLength:"64" minLength:"8" pattern:"^.*$"`
}

func (w *WebhookSecretKey) GetSecretKey() *string {
	if w == nil {
		return nil
	}
	return w.SecretKey
}

func (w *WebhookSecretKey) SetSecretKey(secretKey string) {
	w.SecretKey = &secretKey
}

func (w WebhookSecretKey) String() string {
	jsonData, err := json.MarshalIndent(w, "", "  ")
	if err != nil {
		return "error converting struct: WebhookSecretKey to string"
	}
	return string(jsonData)
}
