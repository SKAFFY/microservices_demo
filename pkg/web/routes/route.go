package routes

import "net/http"

type Route struct {
	Path    string
	Method  string
	Handler http.Handler
	Secure  bool
}

type Routes []Route
