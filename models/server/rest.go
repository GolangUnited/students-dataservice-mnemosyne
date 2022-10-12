package server

import (
	"context"
	"net/http"
	"time"
)

type Rest struct {
	httpServer *http.Server
}

func (r *Rest) Run(port string, handler http.Handler) error {
	r.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return r.httpServer.ListenAndServe()
}

func (r *Rest) Shutdown(ctx context.Context) error {
	return r.httpServer.Shutdown(ctx)
}
