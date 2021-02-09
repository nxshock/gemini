package gemini

import (
	"net"
)

type ResponseWriter interface {
	Header() *Header
	Write([]byte) (int, error)
	WriteHeader(statusCode uint8, meta string)
}

type response struct {
	rw     net.Conn
	header *Header

	wroteHeader bool
}

func (response *response) Header() *Header {
	return response.header
}

func (response *response) WriteHeader(statusCode uint8, meta string) {
	response.Header().Status = statusCode
	response.Header().Meta = meta

	response.rw.Write(append(response.header.Bytes(), newLine...))
	response.wroteHeader = true
}

func (response *response) Write(p []byte) (int, error) {
	if !response.wroteHeader {
		response.WriteHeader(response.header.Status, response.header.Meta)
	}

	return response.rw.Write(p)
}
