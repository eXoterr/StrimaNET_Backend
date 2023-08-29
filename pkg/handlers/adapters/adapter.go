package adapters

import "net/http"

type HTTPContext struct {
	ResponseWriter http.ResponseWriter
	Request        *http.Request
}

type Handler func(HTTPContext)
