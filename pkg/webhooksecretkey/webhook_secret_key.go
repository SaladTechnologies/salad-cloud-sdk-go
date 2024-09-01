package webhooksecretkey

// Represents a webhook secret key
type WebhookSecretKey struct {
	// The webhook secret key
	SecretKey *string `json:"secret_key,omitempty" required:"true"`
}

func (w *WebhookSecretKey) SetSecretKey(secretKey string) {
	w.SecretKey = &secretKey
}

func (w *WebhookSecretKey) GetSecretKey() *string {
	if w == nil {
		return nil
	}
	return w.SecretKey
}
