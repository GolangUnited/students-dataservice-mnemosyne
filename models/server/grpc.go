package server

import (
	"context"
	"fmt"
	"github.com/NEKETSKY/mnemosyne/internal/handler"
	"github.com/NEKETSKY/mnemosyne/pkg/api"
	"github.com/NEKETSKY/mnemosyne/pkg/logger"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type Grpc struct {
	ctx         context.Context
	grpcService *grpc.Server
}

// NewGrpc created new grpc server
func NewGrpc(ctx context.Context) *Grpc {
	return &Grpc{
		ctx: ctx,
	}
}

// Run grpc on port with handler
func (g *Grpc) Run(port int, handler *handler.Handler) (err error) {
	// grpc middleware
	g.grpcService = grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_auth.UnaryServerInterceptor(handler.Auth),
		)),
	)
	api.RegisterMnemosyneServer(g.grpcService, handler)
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

func (g *Grpc) Service() *grpc.Server {
	return g.grpcService
}
