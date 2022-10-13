package handler

import (
	"context"
	"github.com/NEKETSKY/mnemosyne/models/mnemosyne"
	"github.com/NEKETSKY/mnemosyne/proto/helloworld"
	"github.com/pkg/errors"
	"log"
)

// SayHello implements helloworld.GreeterServer
func (h *Handler) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	_ = ctx

	var req mnemosyne.Request
	resp, err := h.services.Mnemosyne.Test(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "say hello")
	}

	log.Printf("Received: %v", in.GetName())
	log.Printf("Version: %v", resp.Version)
	return &helloworld.HelloReply{Message: "Hello " + in.GetName() + ". Version " + resp.Version}, nil
}
