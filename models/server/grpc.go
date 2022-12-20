package server

import (
	"context"
	"fmt"
	"net"

	"github.com/NEKETSKY/mnemosyne/internal/handler"
	"github.com/NEKETSKY/mnemosyne/pkg/api"
	"github.com/NEKETSKY/mnemosyne/pkg/logger"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Grpc struct {
	ctx         context.Context
	grpcService *grpc.Server
	handler     *handler.Handler
}

// NewGrpc created new grpc server
func NewGrpc(ctx context.Context, handler *handler.Handler) *Grpc {
	return &Grpc{
		ctx: ctx,
		// grpc middleware
		grpcService: grpc.NewServer(
			grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			//grpc_auth.UnaryServerInterceptor(handler.Auth),
			)),
		),
		handler: handler,
	}
}

// Run grpc on port with handler
func (g *Grpc) Run(port int) (err error) {
	api.RegisterMnemosyneServer(g.grpcService, g.handler)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}
	logger.Info(fmt.Sprintf("gRPC server is listening on: %d", port))
	reflection.Register(g.grpcService)
	if err = g.grpcService.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to serve")
	}

	return
}

// GracefulStop grpc service
func (g *Grpc) GracefulStop() {
	g.grpcService.GracefulStop()
}
