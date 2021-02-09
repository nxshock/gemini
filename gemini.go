package gemini

// DefaultServeMux is the default ServeMux used by Serve.
var DefaultServeMux = &defaultServeMux

var defaultServeMux ServeMux

// Handle registers the handler for the given pattern
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
func Handle(pattern string, handler Handler) { DefaultServeMux.Handle(pattern, handler) }

// HandleFunc registers the handler function for the given pattern
// in the DefaultServeMux.
// The documentation for ServeMux explains how patterns are matched.
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
	DefaultServeMux.HandleFunc(pattern, handler)
}

// serverHandler delegates to either the server's Handler or
// DefaultServeMux
type serverHandler struct {
	srv *Server
}

func (sh serverHandler) ServeGemini(rw ResponseWriter, req *Request) {
	handler := sh.srv.Handler
	if handler == nil {
		handler = DefaultServeMux
	}

	handler.ServeGemini(rw, req)
}

// ListenAndServeTLS acts identically to ListenAndServe, except that it
// expects HTTPS connections. Additionally, files containing a certificate and
// matching private key for the server must be provided. If the certificate
// is signed by a certificate authority, the certFile should be the concatenation
// of the server's certificate, any intermediates, and the CA's certificate.
func ListenAndServeTLS(addr, certFile, keyFile string, handler Handler) error {
	server := &Server{Addr: addr, Handler: handler}
	return server.ListenAndServeTLS(certFile, keyFile)
}
