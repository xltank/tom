package model

import "net/http"

type Context struct {
	Res *http.ResponseWriter
	Req *http.Request
}

type Route struct {
	Method  string
	Path    string
	Handles []http.HandlerFunc
}
