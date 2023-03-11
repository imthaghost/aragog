package api

import "net/http"


// Error messages
const (
	// ErrBodyBinding is an error message when we fail to bind the request body
	ErrBodyBinding = "could not bind body"
)

// Http response codes
const (
	// httpInternal is a http 500 response
	httpInternal = http.StatusInternalServerError
	// httpOK is a http 200 response
	httpOk = http.StatusOK
	// httpBadRequest is an http 400 response
	httpBadRequest = http.StatusBadRequest
)