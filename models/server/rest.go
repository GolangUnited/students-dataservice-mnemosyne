package server

import (
	"context"
	"fmt"
	"github.com/NEKETSKY/mnemosyne/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net/http"
)

type Rest struct {
	httpServer *http.Server
}

func (r *Rest) Run(ctx context.Context, restPort int, grpcPort int) (err error) {
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = proto.RegisterGreeterHandlerFromEndpoint(ctx, mux,
		fmt.Sprintf(":%d", grpcPort), opts)
	if err != nil {
		return errors.Wrap(err, "failed to register greeter handler")
	}
	if err = http.ListenAndServe(fmt.Sprintf(":%d", restPort), mux); err != nil {
		return errors.Wrap(err, "failed to serve")
	}

	return
}
