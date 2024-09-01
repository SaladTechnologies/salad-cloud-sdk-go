package handlers

import (
	"errors"

	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/httptransport"
)

type ApiKeyHandler[T any] struct {
	nextHandler Handler[T]
}

func NewApiKeyHandler[T any]() *ApiKeyHandler[T] {
	return &ApiKeyHandler[T]{
		nextHandler: nil,
	}
}

func (h *ApiKeyHandler[T]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[T]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[T](err, nil)
	}

	nextRequest := request.Clone()

	if request.Config.ApiKey == nil {
		return h.nextHandler.Handle(nextRequest)
	}

	nextRequest.SetHeader("Salad-Api-Key", *request.Config.ApiKey)

	return h.nextHandler.Handle(nextRequest)
}

func (h *ApiKeyHandler[T]) SetNext(handler Handler[T]) {
	h.nextHandler = handler
}
