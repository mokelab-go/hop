package hop

import (
	"bytes"
	"net/http"
)

type writer struct {
	statusCode int
	body       *bytes.Buffer
}

func newWriter() *writer {
	return &writer{
		body: new(bytes.Buffer),
	}
}

func (o *writer) Header() http.Header {
	return nil
}

func (o *writer) Write(data []byte) (int, error) {
	o.body.Write(data)
	return len(data), nil
}

func (o *writer) WriteHeader(code int) {
	o.statusCode = code
}
