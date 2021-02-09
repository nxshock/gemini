# Gemini

Golang server library for [Gemini protocol](https://en.wikipedia.org/wiki/Gemini_(protocol)) based on standart `net/http` library.

## Usage

```go
import "github.com/nxshock/gemini"

gemini.HandleFunc("/", func(w gemini.ResponseWriter, r *gemini.Request) {
	w.Write([]byte("hello"))
})

err := gemini.ListenAndServeTLS("", "localhost.crt", "localhost.key", nil)
if err != nil {
	panic(err)
}
```
