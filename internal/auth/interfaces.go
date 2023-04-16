package auth

import "net/http"

type IAuthHttpHandler interface {
	Verify(http.ResponseWriter, *http.Request)
	// FetchMany(http.ResponseWriter, *http.Request)
	// Update(http.ResponseWriter, *http.Request)
	// Create(http.ResponseWriter, *http.Request)
	// Delete(http.ResponseWriter, *http.Request)
}
