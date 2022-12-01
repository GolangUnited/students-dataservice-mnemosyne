package handler

import (
	"context"
	"github.com/NEKETSKY/mnemosyne/models/mnemosyne"
	"github.com/NEKETSKY/mnemosyne/pkg/api/helloworld"
	"github.com/NEKETSKY/mnemosyne/pkg/auth"
	"github.com/NEKETSKY/mnemosyne/pkg/file"
	"github.com/NEKETSKY/mnemosyne/pkg/operations"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

// SayHello implements api.MnemosyneServer
func (h *Handler) SayHello(ctx context.Context, in *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	user := auth.GetUser(ctx)
	_ = user

	access := operations.CheckAccess(ctx, "view_all_students")
	_ = access

	f, _ := file.Save("mytest.txt", []byte("mytest"))
	_ = f

	var req mnemosyne.Request
	resp, err := h.services.Mnemosyne.Test(ctx, req)
	if err != nil {
		return nil, errors.Wrap(err, "say hello")
	}

	log.Printf("Received: %v", in.GetName())
	log.Printf("Version: %v", resp.Version)

	//return &helloworld.HelloReply{}, nil
	return &helloworld.HelloReply{}, status.Error(codes.OK, "")
}
