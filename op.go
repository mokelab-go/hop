package hop

import (
	"net/http"
)

// Op is Operator function
type Op func(http.HandlerFunc) http.HandlerFunc

// Operations returns Operator
func Operations(list ...Op) Op {
	return func(next http.HandlerFunc) http.HandlerFunc {
		f := next
		// we want to apply operations by given order.
		for i := len(list) - 1; i >= 0; i-- {
			f = list[i](f)
		}
		return f
	}
}
