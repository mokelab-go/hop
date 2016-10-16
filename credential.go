package hop

import (
	"net/http"
	"strings"
)

type Cred struct {
	TokenType string
	Token     string
}

func GetCredential(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("authorization")
		vals := strings.SplitN(authorization, " ", 2)
		if len(vals) < 2 {
			next(w, r)
			return
		}
		c := setCredential(r.Context(), Cred{
			TokenType: vals[0],
			Token:     vals[1],
		})
		next(w, r.WithContext(c))
	}
}
