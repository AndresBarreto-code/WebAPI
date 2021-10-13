package main

import (
	"net/http"
)

type Router struct {
	rules map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, done := r.FindHandler(request.URL.Path)
	if !done {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	handler(w, request)
}

func (r *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	handler, done := r.rules[path]
	return handler, done
}
