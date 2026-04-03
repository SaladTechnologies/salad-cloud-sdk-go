package handlers

import (
	"errors"

	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/httptransport"
)

// ApiKeyHandler injects an API key into the request header for authentication.
// It adds the configured API key to the request if present in the config.
// T is the response type, E is the error type.
type ApiKeyHandler[T any, E any] struct {
	nextHandler Handler[T, E]
}

// NewApiKeyHandler creates a new API key authentication handler.
// Returns a handler that will inject the API key header when processing requests.
func NewApiKeyHandler[T any, E any]() *ApiKeyHandler[T, E] {
	return &ApiKeyHandler[T, E]{
		nextHandler: nil,
	}
}

// Handle processes a regular request by adding the API key header if configured.
// Clones the request, adds the API key, and passes it to the next handler.
func (h *ApiKeyHandler[T, E]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[E]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	nextRequest := request.Clone()

	if request.Config.ApiKey == nil {
		return h.nextHandler.Handle(nextRequest)
	}

	nextRequest.SetHeader("Salad-Api-Key", *request.Config.ApiKey)

	return h.nextHandler.Handle(nextRequest)
}

// HandleStream processes a streaming request by adding the API key header if configured.
// Clones the request, adds the API key, and passes it to the next handler.
func (h *ApiKeyHandler[T, E]) HandleStream(request httptransport.Request) (*httptransport.Stream[T], *httptransport.ErrorResponse[E]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	nextRequest := request.Clone()

	if request.Config.ApiKey == nil {
		return h.nextHandler.HandleStream(nextRequest)
	}

	nextRequest.SetHeader("Salad-Api-Key", *request.Config.ApiKey)

	return h.nextHandler.HandleStream(nextRequest)
}

// SetNext sets the next handler in the chain.
// This method is called during chain construction to link handlers together.
func (h *ApiKeyHandler[T, E]) SetNext(handler Handler[T, E]) {
	h.nextHandler = handler
}
