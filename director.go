package main

import (
	"net/http"
)

func NewDirector() func(req *http.Request){
	return func(req *http.Request){
		req.URL.Host = opts.Host
		req.URL.Scheme = opts.Scheme
		req.Header.Set("X-Forwarded-Host", opts.Host)
		req.Host = opts.Host
	}
}