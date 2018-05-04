package servicehandlers

import (
	"net/http"
)

type HttpServiceHandler interface {
	Get(*http.Request) SrvcRes
	Put(*http.Request) SrvcRes
	Post(*http.Request) SrvcRes
}

func methodRouter(p HttpServiceHandler, w http.ResponseWriter, r *http.Request) interface{} {
	var response interface{}

	if r.Method == "GET" {
		response = p.Get(r)
	} else if r.Method == "PUT" {
		response = p.Put(r)
	} else if r.Method == "POST" {
		response = p.Post(r)
	}
	return response
}
