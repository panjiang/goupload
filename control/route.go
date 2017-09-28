package control

import (
	"net/http"
	"upload/app"
	"upload/context"
)

type handler func(c *context.Context)

// interceptor for handler
func interceptor(hdl handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		context := context.NewContext(w, r)
		hdl(context)
	}
}

// map url to handler function
func handle(url string, hdl handler) {
	http.Handle(url, interceptor(hdl))
}

// handle static file request
func handleFile(url string, dir string) {
	fileServer := http.FileServer(http.Dir(dir))
	http.Handle(url, http.StripPrefix(url, fileServer))
}

// init url route rule
func InitRoute() {
	handleFile("/temp/", "temp")

	handle("/", app.Index)
	handle("/upload", app.Upload)
}
