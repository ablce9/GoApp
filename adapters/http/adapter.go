package http

import (
	"gitlab.com/zenport.io/go-assignment/engine"
)

type HTTPAdapter struct{}

func (adapter *HTTPAdapter) Start() {
	// todo: start to listen
}

func (adapter *HTTPAdapter) Stop() {
	// todo: shutdown server
}

func NewHTTPAdapter(e engine.Engine) *HTTPAdapter {
	// todo: init your http server and routes

	return &HTTPAdapter{}
}
