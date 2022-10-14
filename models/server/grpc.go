package server

import (
	"context"
	"fmt"
	"github.com/NEKETSKY/mnemosyne/internal/handler"
	"github.com/NEKETSKY/mnemosyne/pkg/api"
	"github.com/NEKETSKY/mnemosyne/pkg/log"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"net"
)

type Grpc struct{}

func (g *Grpc) Run(ctx context.Context, port int, handler *handler.Handler) (err error) {
	logger := log.LoggerFromContext(ctx)
	grpcServer := grpc.NewServer()
	api.RegisterMnemosyneServer(grpcServer, handler)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}
	logger.Info(fmt.Sprintf("gRPC server is listening on: %d", port))
	if err = grpcServer.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to serve")
	}

	return
}
