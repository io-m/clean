package core

import "net/http"

type IHttpHandler interface {
	FindOne(http.ResponseWriter, *http.Request)
	// FetchMany(http.ResponseWriter, *http.Request)
	// Update(http.ResponseWriter, *http.Request)
	// Create(http.ResponseWriter, *http.Request)
	// Delete(http.ResponseWriter, *http.Request)
}

type IHttpRouter interface {
	Handle(string) error
}
