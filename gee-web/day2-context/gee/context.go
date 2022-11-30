package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	writer     http.ResponseWriter
	Req        *http.Request
	Path       string
	Method     string
	StatusCode int
}

func newContext(rw http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		writer: rw,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

func (context *Context) PostForm(key string) string {
	return context.Req.FormValue(key)
}

func (context *Context) Query(key string) string {
	return context.Req.URL.Query().Get(key)
}

func (context *Context) Status(code int) {
	context.StatusCode = code
	context.writer.WriteHeader(code)
}

func (context *Context) SetHeader(key string, value string) {
	context.writer.Header().Set(key, value)
}

func (context *Context) String(code int, format string, values ...interface{}) {
	context.SetHeader("Content-Type", "text/plain")
	context.Status(code)
	context.writer.Write([]byte(fmt.Sprintf(format, values...)))
}

func (context *Context) JSON(code int, obj interface{}) {
	context.SetHeader("Content-Type", "application/json")
	context.Status(code)
	encoder := json.NewEncoder(context.writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(context.writer, err.Error(), 500)
	}
}

func (context *Context) Data(code int, data []byte) {
	context.Status(code)
	context.writer.Write(data)
}

func (context *Context) HTML(code int, html string) {
	context.SetHeader("Content-Type", "text/html")
	context.Status(code)
	context.writer.Write([]byte(html))
}
