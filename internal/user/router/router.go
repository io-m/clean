package router

type IHttpRouter interface {
	Handle(string)
}
