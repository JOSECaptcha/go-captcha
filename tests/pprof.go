package main

import (
	"github.com/JOSECaptcha/go-captcha/tests/pprof"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	// Example: demo
	http.HandleFunc("/handler", pprof.Handler)
	http.ListenAndServe(":9999", nil)
}