package gemini

import (
	"net/url"
)

type Request struct {
	URL        *url.URL
	Host       string
	RequestURI string
	RemoteAddr string
}
