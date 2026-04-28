package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	router       *mux.Router
	httpHandlers *HTTPHandlers
}

func NewServer(router string) *HTTPServer {
	return &HTTPServer{router: mux.NewRouter(), httpHandlers: &HTTPHandlers{}}
}

func (s *HTTPServer) StartServer() {
	s.router.Path("/post").Methods(http.MethodPost).HandlerFunc(s.httpHandlers.CreatePost) // created Post
}
