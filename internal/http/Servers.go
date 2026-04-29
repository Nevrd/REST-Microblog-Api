package http

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	router       *mux.Router
	httpHandlers *HTTPHandlers
}

// context(ctx) for database(waitCancel)
func NewServer(ctx context.Context) (*HTTPServer, error) {
	handlerHTTP, err := NewHTTPHandlers(ctx)
	return &HTTPServer{router: mux.NewRouter(), httpHandlers: handlerHTTP}, err
}

func (s *HTTPServer) StartServer() {
	s.router.Path("/posts").Methods(http.MethodPost).HandlerFunc(s.httpHandlers.CreatePost) // created Post
	s.router.Path("/posts").Methods(http.MethodGet).HandlerFunc(s.httpHandlers.GetPosts)
}
