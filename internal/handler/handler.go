package handler

import (
	"context"
	"github.com/GolangUnited/students-dataservice-mnemosyne/internal/service"
	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api"
	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/common"
)

// Handler struct with grpc api server
type Handler struct {
	ctx      context.Context
	services service.Service
	api.MnemosyneServer
}

// NewHandler creates a new handler
func NewHandler(ctx context.Context, services *service.Service) *Handler {
	return &Handler{
		ctx:      ctx,
		services: *services,
	}
}

var (
	emptyProto = &common.Empty{}
)
