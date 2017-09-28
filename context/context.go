package context

import (
	"encoding/json"
	"fmt"
	"net/http"
	"upload/cache"
)

// context for every request
type Context struct {
	W http.ResponseWriter
	R *http.Request
}

// common response result with json
type JsonResult struct {
	Status  int         `json:"status"`            // 0:success, !0:error
	Message string      `json:"message,omitempty"` // error message
	Data    interface{} `json:"data,omitempty"`    // success data
}

// create new context for every request
func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{W: w, R: r}
}

// fetch form value from request
func (c *Context) FormValue(key string) string {
	return c.R.FormValue(key)
}

// fetch form value from post request
func (c *Context) PostFormValue(key string) string {
	return c.R.PostFormValue(key)
}

// response text data
func (c *Context) Text(s string) {
	fmt.Fprint(c.W, s)
}

// response json data
func (c *Context) Json(status int, message string, data interface{}) {
	m := JsonResult{status, message, data}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Fprint(c.W, err)
	}
	fmt.Fprint(c.W, string(b))
}

// response json success data
func (c *Context) JsonSuccess(data interface{}) {
	c.Json(0, "", data)
}

// response json error data
func (c *Context) JsonError(status int, message string) {
	if status == 0 {
		panic("error status should not be 0")
	}
	c.Json(status, message, nil)
}

// response html page
func (c *Context) Html(name string, data interface{}) {
	page := name + ".html"
	tmpl := cache.RenderTemplate(page)
	err := tmpl.ExecuteTemplate(c.W, page, data)
	if err != nil {
		http.Error(c.W, err.Error(), http.StatusInternalServerError)
	}
}
