package gemini

type Handler interface {
	ServeGemini(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

func (mux *ServeMux) Handle(pattern string, handler Handler) {
	mux.mu.Lock()
	defer mux.mu.Unlock()

	if pattern == "" {
		panic("http: invalid pattern")
	}
	if handler == nil {
		panic("http: nil handler")
	}
	if _, exist := mux.m[pattern]; exist {
		panic("http: multiple registrations for " + pattern)
	}

	if mux.m == nil {
		mux.m = make(map[string]muxEntry)
	}
	e := muxEntry{h: handler, pattern: pattern}
	mux.m[pattern] = e
	if pattern[len(pattern)-1] == '/' {
		mux.es = appendSorted(mux.es, e)
	}

	if pattern[0] != '/' {
		mux.hosts = true
	}
}
func (f HandlerFunc) ServeGemini(w ResponseWriter, r *Request) {
	f(w, r)
}

func (mux *ServeMux) ServeGemini(w ResponseWriter, r *Request) {
	// TODO: это что такое?
	if r.RequestURI == "*" {
		return
	}
	h, _ := mux.Handler(r)
	h.ServeGemini(w, r)
}

func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string) {
	host := stripHostPort(r.Host)
	path := cleanPath(r.URL.Path)

	// If the given path is /tree and its handler is not registered,
	// redirect for /tree/.
	if u, ok := mux.redirectToPathSlash(host, path, r.URL); ok {
		return RedirectHandler(u.String(), StatusRedirectTemporary), u.Path // TODO: check status code
	}

	if path != r.URL.Path {
		_, pattern = mux.handler(host, path)
		url := *r.URL
		url.Path = path
		return RedirectHandler(url.String(), StatusRedirectTemporary), pattern // TODO: check status code
	}

	return mux.handler(host, r.URL.Path)
}
