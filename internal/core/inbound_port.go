package core

import "net/http"


type IUserHttpHandler interface {
	FindOne(http.ResponseWriter, *http.Request)
	// FetchMany(http.ResponseWriter, *http.Request)
	// Update(http.ResponseWriter, *http.Request)
	// Create(http.ResponseWriter, *http.Request)
	// Delete(http.ResponseWriter, *http.Request)
}

type IAuthHttpHandler interface {
	Verify(http.ResponseWriter, *http.Request)
	// FetchMany(http.ResponseWriter, *http.Request)
	// Update(http.ResponseWriter, *http.Request)
	// Create(http.ResponseWriter, *http.Request)
	// Delete(http.ResponseWriter, *http.Request)
}
