package hop

import (
	"context"
)

type contextKey string

const (
	keyPathParams  contextKey = "hop.pathparams"
	keyPathInt                = "hop.path.int."
	keyContentType contextKey = "hop.content-type"
	keyBodyJSON    contextKey = "hop.body.json"
)

func setPathParams(c context.Context, params map[string]string) context.Context {
	return context.WithValue(c, keyPathParams, params)
}

// PathParams returns Path parameters
func PathParams(c context.Context) map[string]string {
	if params, ok := c.Value(keyPathParams).(map[string]string); ok {
		return params
	}
	return map[string]string{}
}

func setPathInt(c context.Context, key string, value int) context.Context {
	return context.WithValue(c, contextKey(keyPathInt+key), value)
}

// PathInt returns path parameter as int value. If value is not int,
// returns 0 instread.
func PathInt(c context.Context, key string) int {
	if v, ok := c.Value(contextKey(keyPathInt + key)).(int); ok {
		return v
	}
	return 0
}

func setContentType(c context.Context, contentType string) context.Context {
	return context.WithValue(c, keyContentType, contentType)
}

// ContentType returns Content-Type header
func ContentType(c context.Context) string {
	if v, ok := c.Value(keyContentType).(string); ok {
		return v
	}
	return ""
}

func setBodyJSON(c context.Context, body map[string]interface{}) context.Context {
	return context.WithValue(c, keyBodyJSON, body)
}

// BodyJSON returns request body as JSON(map[string]interface{}) format
func BodyJSON(c context.Context) map[string]interface{} {
	if v, ok := c.Value(keyBodyJSON).(map[string]interface{}); ok {
		return v
	}
	return map[string]interface{}{}
}
