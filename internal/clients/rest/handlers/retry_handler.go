package handlers

import (
	"errors"
	"time"

	"github.com/saladtechnologies/salad-cloud-sdk-go/internal/clients/rest/httptransport"
)

const (
	maxRetries = 3
	retryDelay = 150 * time.Millisecond
)

// RetryHandler automatically retries failed requests with exponential backoff.
// It retries requests that fail or return 4xx/5xx status codes up to maxRetries times.
// T is the response type, E is the error type.
type RetryHandler[T any, E any] struct {
	nextHandler Handler[T, E]
}

// NewRetryHandler creates a new retry handler with configured max retries and delay.
// Returns a handler that will automatically retry failed requests with backoff.
func NewRetryHandler[T any, E any]() *RetryHandler[T, E] {
	return &RetryHandler[T, E]{
		nextHandler: nil,
	}
}

// Handle processes a request with automatic retry logic on failures or error status codes.
// Implements exponential backoff between retry attempts. Returns the first successful response or the final error.
func (h *RetryHandler[T, E]) Handle(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[E]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	var err *httptransport.ErrorResponse[E]
	for tryCount := 0; tryCount < maxRetries; tryCount++ {
		nextRequest := request.Clone()

		var resp *httptransport.Response[T]
		resp, err = h.nextHandler.Handle(nextRequest)
		if err != nil {
			return nil, err
		}

		if resp.StatusCode < 400 {
			return resp, nil
		}

		backoffDuration := time.Duration(tryCount) * retryDelay
		time.Sleep(backoffDuration)
	}
	if err != nil {
		return nil, err
	}
	return nil, httptransport.NewErrorResponse[E](errors.New("max retries exceeded"), nil)
}

// HandleStream processes a streaming request with automatic retry logic on failures.
// Retries failed stream connections with exponential backoff. Returns the first successful stream or the final error.
func (h *RetryHandler[T, E]) HandleStream(request httptransport.Request) (*httptransport.Stream[T], *httptransport.ErrorResponse[E]) {
	if h.nextHandler == nil {
		err := errors.New("Handler chain terminated without terminating handler")
		return nil, httptransport.NewErrorResponse[E](err, nil)
	}

	var errResp *httptransport.ErrorResponse[E]
	for tryCount := 0; tryCount < maxRetries; tryCount++ {
		nextRequest := request.Clone()

		stream, err := h.nextHandler.HandleStream(nextRequest)
		if err == nil {
			return stream, nil
		}

		errResp = err

		backoffDuration := time.Duration(tryCount) * retryDelay
		time.Sleep(backoffDuration)
	}
	return nil, errResp
}

// SetNext sets the next handler in the chain.
// This method is called during chain construction to link handlers together.
func (h *RetryHandler[T, E]) SetNext(handler Handler[T, E]) {
	h.nextHandler = handler
}
