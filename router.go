package main

import (
	"net/http"
)

type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool) {
	handler, exists := r.rules[path][method]
	return handler, exists
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, exists := r.FindHandler(request.URL.Path, request.Method)
	if !exists {
		NotFound(w, request)
		return
	}
	handler(w, request)
}
