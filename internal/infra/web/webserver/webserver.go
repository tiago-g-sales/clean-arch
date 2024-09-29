package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type MethodPath struct {
	Method string
	Path   string
}

type WebServer struct {
	Router        chi.Router
	Handlers      map[MethodPath]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[MethodPath]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method string, path string, handler http.HandlerFunc) {

	s.Handlers[MethodPath{Method: method, Path: path}] = handler

}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for methodpath, handler := range s.Handlers {
		s.Router.Method(methodpath.Method, methodpath.Path, handler)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
