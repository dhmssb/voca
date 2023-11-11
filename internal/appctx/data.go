package appctx

import "net/http"

type Data struct {
	Request *http.Request
	// Config      *Config
	ServiceType string
	BytesValue  []byte
}
