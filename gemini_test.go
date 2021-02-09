package gemini

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGemini(t *testing.T) {
	HandleFunc("/test", func(w ResponseWriter, r *Request) {
		w.Write([]byte("test"))
	})
	go func() {
		err := ListenAndServeTLS(":1000", "localhost.crt", "localhost.key", nil)
		assert.NoError(t, err)
	}()
	time.Sleep(time.Second / 10)

	c, err := tls.Dial("tcp", ":1000", &tls.Config{InsecureSkipVerify: true})
	assert.NoError(t, err)

	_, err = fmt.Fprintf(c, "/test\r\n")
	assert.NoError(t, err)

	b, err := ioutil.ReadAll(c)
	assert.NoError(t, err)

	assert.Equal(t, "20 text/gemini\r\ntest", string(b))
}
