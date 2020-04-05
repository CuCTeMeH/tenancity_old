package core

import "net/http"

type Routes struct {
	Handlers map[string]http.Handler
}
