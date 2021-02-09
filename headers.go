package gemini

import (
	"bytes"
	"fmt"
)

type Header struct {
	Status uint8
	Meta   string
}

func (responseHeader *Header) Bytes() []byte {
	buf := new(bytes.Buffer)

	fmt.Fprintf(buf, "%d ", responseHeader.Status)

	if responseHeader.Meta != "" {
		fmt.Fprintf(buf, "%s", responseHeader.Meta)
	}

	return buf.Bytes()
}
