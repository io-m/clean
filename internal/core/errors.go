package core

import "errors"

var ErrorUserNotFound = errors.New("user is not found")
var ErrorInternal = errors.New("something went wrong")
var ErrorHandlerNotSet = errors.New("handler cannot be set")
