# HOP
Handler OPerators

You can add pre operations to your `http.HandlerFunc`

op1 - op2 - op3 - ... - your http.HandlerFunc

## Example

Gets Contenty-Type(MIME type) Header
```go
import (
    "github.com/mokelab-go/hop"
)

// main handler
handler := func(w http.ResponseWriter, r *http.Request) {
    contentType := hop.ContentType(r.Context())
    fmt.Fprintf(w, "Content type is %s", contentType)
}
handler = op.Operations(op.GetContentType)(handler)
```

