package hop

import (
	"context"
)

type contextKey string

const (
	keyPathParams  contextKey = "hop.pathparams"
	keyPathInt                = "hop.path.int."
	keyPathString             = "hop.path.str."
	keyContentType contextKey = "hop.content-type"
	keyCredential  contextKey = "hop.credential"
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

func setPathString(c context.Context, key, value string) context.Context {
	return context.WithValue(c, contextKey(keyPathString+key), value)
}

// PathString returns path parameter as int value. If value is not string,
// returns "" instread.
func PathString(c context.Context, key string) string {
	if v, ok := c.Value(contextKey(keyPathString + key)).(string); ok {
		return v
	}
	return ""
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

func setCredential(c context.Context, cred Cred) context.Context {
	return context.WithValue(c, keyCredential, cred)
}

// Credential returns TokenType and Token
func Credential(c context.Context) Cred {
	if v, ok := c.Value(keyCredential).(Cred); ok {
		return v
	}
	return Cred{}
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
