package main

import (
	"log"
	"net/http"
	"time"
)

func CheckAuth(canPass bool) Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(res http.ResponseWriter, req *http.Request) {
			if canPass {
				next(res, req)
			} else {
				Unauthorized(res, req)
			}
		}
	}
}

func Log() Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(res http.ResponseWriter, req *http.Request) {
			start := time.Now()
			defer func() { log.Println(req.URL.Path, time.Since(start)) }()
			next(res, req)
		}
	}
}
