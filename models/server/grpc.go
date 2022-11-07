package server

import (
	"context"
	"fmt"
	"github.com/NEKETSKY/mnemosyne/internal/handler"
	"github.com/NEKETSKY/mnemosyne/pkg/api"
	"github.com/NEKETSKY/mnemosyne/pkg/logger"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"net"
)

type Grpc struct {
	ctx         context.Context
	grpcService *grpc.Server
}

// NewGrpc created new grpc server
func NewGrpc(ctx context.Context, grpcService *grpc.Server) *Grpc {
	return &Grpc{
		ctx:         ctx,
		grpcService: grpcService,
	}
}

// Run grpc on port with handler
func (g *Grpc) Run(port int, handler *handler.Handler) (err error) {
	api.RegisterMnemosyneServer(g.grpcService, handler)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}
	logger.Info(fmt.Sprintf("gRPC server is listening on: %d", port))
	if err = g.grpcService.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to serve")
	}

	return
}
