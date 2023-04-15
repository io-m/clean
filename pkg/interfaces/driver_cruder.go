package common_interface

// Common driver interface for handler api requests
type CruderHandler interface {
	FindOne() error
	FetchMany() error
	Update() error
	Create() error
	Delete() error
}
