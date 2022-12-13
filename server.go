package Back

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(host string, port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           fmt.Sprintf("%s:%s", host, port),
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, //1 mb
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
	}

	return s.httpServer.ListenAndServe()

}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
