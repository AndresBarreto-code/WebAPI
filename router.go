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

func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	handler, pathExists, methodExists := r.FindHandler(request.URL.Path, request.Method)
	if !pathExists {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if !methodExists {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	handler(w, request)
}

func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, pathExists := r.rules[path]
	handler, methodExists := r.rules[path][method]
	return handler, pathExists, methodExists
}
