package server

import (
	"context"
	"fmt"
	"github.com/NEKETSKY/mnemosyne/pkg/api"
	"github.com/NEKETSKY/mnemosyne/pkg/log"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"net/http"
	"os"
)

const swaggerDir = "./swagger"

type Rest struct{}

func (r *Rest) Run(ctx context.Context, grpcPort, restPort int) (err error) {
	logger := log.LoggerFromContext(ctx)

	gwMux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = api.RegisterMnemosyneHandlerFromEndpoint(ctx, gwMux,
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
	restServer := &http.Server{Addr: fmt.Sprintf(":%d", restPort), Handler: mux}
	if err = restServer.ListenAndServe(); err != nil {
		return errors.Wrap(err, "failed to serve rest server")
	}

	return
}

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
