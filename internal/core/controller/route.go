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

func (r *Route) Method() string {
	return r.method
}

func (r *Route) Path() string {
	return r.path
}

func (r *Route) HandlerName() string {
	return r.handlerName
}