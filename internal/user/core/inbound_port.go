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
	Serve(string) error
	HandleRoutes()
	// MethodGet(string, func(w http.ResponseWriter, r *http.Request))
	// FetchMany(string, http.ResponseWriter, *http.Request)
	// Update(string, http.ResponseWriter, *http.Request)
	// Create(string, http.ResponseWriter, *http.Request)
	// Delete(string, http.ResponseWriter, *http.Request)
}
