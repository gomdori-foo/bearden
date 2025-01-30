package controller

type Route struct {
	method string
	path string
	handlerName string
}

func NewRoute(method string, path string, handlerName string) Route {
	return Route{
		method: method,
		path: path,
		handlerName: handlerName,
	}
}