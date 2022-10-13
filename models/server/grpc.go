package server

import (
	"fmt"
	"github.com/NEKETSKY/mnemosyne/internal/handler"
	"github.com/NEKETSKY/mnemosyne/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"net"
)

type Grpc struct{}

func (g *Grpc) Run(port int, handler *handler.Handler) (err error) {
	grpcServer := grpc.NewServer()
	proto.RegisterGreeterServer(grpcServer, handler)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}
	if err = grpcServer.Serve(lis); err != nil {
		return errors.Wrap(err, "failed to serve")
	}

	return
}
