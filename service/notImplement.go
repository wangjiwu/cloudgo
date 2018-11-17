package service

import (
	"net/http"
)

// NotImplemented .
func NotImplemented(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "501 Not Implemented", 501)
}
