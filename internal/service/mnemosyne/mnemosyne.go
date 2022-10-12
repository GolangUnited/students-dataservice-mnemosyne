package mnemosyne

import (
	"context"
	"github.com/NEKETSKY/mnemosyne/internal/repository"
	"github.com/NEKETSKY/mnemosyne/models/mnemosyne"
)

//go:generate mockgen -source=mnemosyne.go -destination=mocks/mnemosyne.go

type Service struct {
	mnemosyne repository.Mnemosyne
}

func NewService(mnemosyne repository.Mnemosyne) *Service {
	return &Service{
		mnemosyne: mnemosyne,
	}
}

func (s *Service) Test(ctx context.Context, req mnemosyne.Request) (resp mnemosyne.Response, err error) {
	_ = ctx
	_ = req
	resp = *mnemosyne.NewResponse()

	return
}
