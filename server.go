package gemini

import (
	"bytes"
	"crypto/tls"
	"errors"
	"log"
	"net"
	"net/url"
)

type Server struct {
	Addr    string
	Handler Handler
	//ServeMux *ServeMux

	ListenerAcceptErrorHandler func(error)

	//tlsConfig *tls.Config

	listener net.Listener
	stop     chan struct{}
	newConn  chan net.Conn
}

func NewServer() *Server {
	server := &Server{
		Handler: new(ServeMux)}

	return server
}

func (srv *Server) ListenAndServeTLS(certFile, keyFile string) error {
	addr := srv.Addr
	if addr == "" {
		addr = defaultListenAddress
	}

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return err
	}

	tlsConfig := &tls.Config{Certificates: []tls.Certificate{cert}}

	srv.listener, err = tls.Listen("tcp", addr, tlsConfig)
	if err != nil {
		return err
	}

	srv.stop = make(chan struct{})
	srv.newConn = make(chan net.Conn)

	for {
		rw, err := srv.listener.Accept()
		if err != nil {
			if srv.ListenerAcceptErrorHandler != nil {
				srv.ListenerAcceptErrorHandler(err)
			}
			break
		}

		go srv.serve(rw)
	}

	return nil
}

func (srv *Server) serve(rw net.Conn) {
	defer rw.Close()

	requestBytes := make([]byte, 1024)

	n, err := rw.Read(requestBytes)
	if err != nil {
		log.Println(err)
		return
	}
	requestBytes = requestBytes[:n]

	newLinePos := bytes.Index(requestBytes, newLine)
	if newLinePos == -1 {
		if srv.ListenerAcceptErrorHandler != nil {
			srv.ListenerAcceptErrorHandler(errors.New("request does not ends with \\r\\n"))
		}
		return
	}

	requestBytes = requestBytes[:newLinePos]

	requestUrl, err := url.Parse(string(requestBytes))
	if err != nil {
		log.Println(err)
		return
	}

	request := &Request{
		URL:        requestUrl,
		Host:       requestUrl.Host,
		RequestURI: requestUrl.Path,
		RemoteAddr: rw.RemoteAddr().String()}

	response := &response{
		rw:     rw,
		header: new(Header)}
	response.header.Status = StatusSuccess
	response.header.Meta = "text/gemini"

	serverHandler{srv}.ServeGemini(response, request)
}

func (server *Server) Stop() error {
	return server.listener.Close()
}
