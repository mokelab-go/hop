package hop

import (
	"net/http"
	"testing"
)

func TestCredential_GetCredential(t *testing.T) {
	h := GetCredential(func(w http.ResponseWriter, r *http.Request) {
		cred := Credential(r.Context())
		if cred.TokenType != "bearer" {
			t.Errorf("Wrong Token Type : %s", cred.TokenType)
		}
		if cred.Token != "mytoken AAAA" {
			t.Errorf("Wrong Token : %s", cred.Token)
		}
	})
	r := &http.Request{
		Header: http.Header{
			"Authorization": {"bearer mytoken AAAA"},
		},
	}
	h(nil, r)
}
