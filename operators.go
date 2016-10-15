package hop

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

func getPathParams(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		c := setPathParams(r.Context(), params)
		next(w, r.WithContext(c))
	}
}

// GetContentType returns handler, this handler gets Content-Type
func GetContentType(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")
		vals := strings.SplitN(contentType, ";", -1)

		c := setContentType(r.Context(), vals[0])
		next(w, r.WithContext(c))
	}
}

// GetPathInt returns Operator, this handler gets Path parameter as int value
func GetPathInt(name string) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			c := r.Context()
			params := PathParams(c)

			valStr := params[name]
			val, err := strconv.Atoi(valStr)
			if err != nil {
				return
			}

			c = setPathInt(c, name, val)
			next(w, r.WithContext(c))
		}
	}
}

// GetBodyAsJSON returns handler, this handler decodes request body as JSON
// format.
func GetBodyAsJSON(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		doc := json.NewDecoder(r.Body)
		var obj map[string]interface{}
		err := doc.Decode(&obj)
		if err != nil {
			w.WriteHeader(400)
			fmt.Fprintf(w, "{\"error_code\":\"INPUT_ERROR\",\"msg\":\"Input is not JSON Object format\"}")
			return
		}

		c := setBodyJSON(r.Context(), obj)
		next(w, r.WithContext(c))
	}
}
