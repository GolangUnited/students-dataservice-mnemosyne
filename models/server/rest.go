package server

import (
	"context"
	"fmt"
	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api"
	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/logger"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"net/http"
	"os"
)

const swaggerDir = "./swagger"

type Rest struct {
	ctx        context.Context
	restServer *http.Server
}

// NewRest created new rest server
func NewRest(ctx context.Context) *Rest {
	return &Rest{
		ctx:        ctx,
		restServer: &http.Server{},
	}
}

// Run rest server with handle swagger
func (r *Rest) Run(grpcPort, restPort int) (err error) {
	gwMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = api.RegisterMnemosyneHandlerFromEndpoint(r.ctx, gwMux,
		fmt.Sprintf(":%d", grpcPort), opts)
	if err != nil {
		return errors.Wrap(err, "failed to register greeter handler")
	}

	// Serve the swagger-ui and swagger file
	mux := http.NewServeMux()
	mux.Handle("/", gwMux)
	handleSwaggerFile(mux)

	// Register Swagger Handler
	fs := http.FileServer(http.Dir(swaggerDir))
	mux.Handle("/swagger/", http.StripPrefix("/swagger/", fs))

	logger.Info(fmt.Sprintf("Rest server is running on: %d", restPort))
	r.restServer.Addr = fmt.Sprintf(":%d", restPort)
	r.restServer.Handler = mux
	if err = r.restServer.ListenAndServe(); err != nil {
		return errors.Wrap(err, "serve rest server")
	}

	return
}

// Shutdown rest server
func (r Rest) Shutdown() error {
	return r.restServer.Shutdown(r.ctx)
}

// handleSwaggerFile bind swagger handler with json
func handleSwaggerFile(mux *http.ServeMux) {
	mux.HandleFunc("/swagger.json", func(w http.ResponseWriter, req *http.Request) {
		errorMessage := []byte("Failed to open swagger file.")
		f, err := os.Open("swagger/api/api.swagger.json")
		if err != nil {
			_, _ = w.Write(errorMessage)
			return
		}

		_, err = io.Copy(w, f)
		if err != nil {
			_, _ = w.Write(errorMessage)
			return
		}
	})
}
