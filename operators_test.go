package hop

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error { return nil }

func TestOperators_GetContentType(t *testing.T) {
	h := GetContentType(func(w http.ResponseWriter, r *http.Request) {
		contentType := ContentType(r.Context())
		if contentType != "application/json" {
			t.Errorf("Wrong contentType %s", contentType)
		}
	})
	r := &http.Request{
		Header: http.Header{
			"Content-Type": {"application/json"},
		},
	}
	h(nil, r)
}

func TestOperators_GetContentType_None(t *testing.T) {
	h := GetContentType(func(w http.ResponseWriter, r *http.Request) {
		contentType := ContentType(r.Context())
		if contentType != "" {
			t.Errorf("Wrong contentType %s", contentType)
		}
	})
	r := &http.Request{
		Header: http.Header{},
	}
	h(nil, r)
}

func TestOperators_GetBodyAsJSON(t *testing.T) {
	h := GetBodyAsJSON(func(w http.ResponseWriter, r *http.Request) {
		body := BodyJSON(r.Context())
		name, ok := body["name"]
		if !ok {
			t.Errorf("Wrong body %s", body)
			return
		}
		if name != "moke" {
			t.Errorf("Wrong name value : %s", name)
		}
	})
	body := "{\"name\":\"moke\"}"
	r := &http.Request{
		Body: nopCloser{bytes.NewBufferString(body)},
	}
	h(nil, r)
}

func TestOperators_GetBodyAsJSON_BrokenInput(t *testing.T) {
	h := GetBodyAsJSON(func(w http.ResponseWriter, r *http.Request) {
		t.Errorf("This handler must not be executed")
	})
	body := "broken"
	r := &http.Request{
		Body: nopCloser{bytes.NewBufferString(body)},
	}
	writer := newWriter()
	h(writer, r)
	statusCode := writer.statusCode
	outBody := writer.body.String()

	if statusCode != 400 {
		t.Errorf("Wrong Status code : %d", statusCode)
	}
	if outBody != "{\"error_code\":\"INPUT_ERROR\",\"msg\":\"Input is not JSON Object format\"}" {
		t.Errorf("Wrong output body : %s", outBody)
	}
}
