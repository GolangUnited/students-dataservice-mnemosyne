package handler

import (
	"context"
	"github.com/GolangUnited/students-dataservice-mnemosyne/models/database"
	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/certificate"
	"github.com/GolangUnited/students-dataservice-mnemosyne/pkg/api/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) CreateCertificate(ctx context.Context, in *certificate.CertificateRequest) (cert *certificate.CertificateResponse, err error) {
	cert = &certificate.CertificateResponse{}
	certificateDb, err := database.CertificateFromProto(in)
	if err != nil {
		return cert, status.Error(codes.InvalidArgument, err.Error())
	}
	certificateId, err := h.services.CreateCertificate(ctx, certificateDb)
	if err != nil {
		return cert, status.Error(codes.Internal, err.Error())
	}
	cert.Id = uint32(certificateId)
	return
}

func (h *Handler) GetCertificates(ctx context.Context, in *certificate.Filter) (certificates *certificate.Certificates, err error) {
	certificates = &certificate.Certificates{}

	if err != nil {
		return certificates, status.Error(codes.InvalidArgument, err.Error())
	}

	certificateDb, err := h.services.Mnemosyne.GetCertificates(ctx, in.GetUserId())

	if err != nil {
		return certificates, status.Error(codes.Internal, err.Error())
	}
	var slice []*certificate.CertificateResponse
	for _, i := range certificateDb {
		slice = append(slice, i.CertificateToProto())
	}
	return &certificate.Certificates{Certificates: slice}, err
}

func (h *Handler) UpdateCertificate(ctx context.Context, in *certificate.CertificateRequest) (*common.Empty, error) {
	certificateDb, _ := database.CertificateFromProto(in)
	err := h.services.Mnemosyne.UpdateCertificate(ctx, certificateDb)
	if err != nil {
		return emptyProto, status.Error(codes.Internal, err.Error())
	}
	return emptyProto, nil
}

func (h *Handler) DeactivateCertificate(ctx context.Context, in *certificate.Id) (*common.Empty, error) {
	err := h.services.Mnemosyne.DeactivateCertificate(ctx, in.GetId())
	if err != nil {
		return emptyProto, status.Error(codes.Internal, err.Error())
	}
	return emptyProto, nil
}

func (h *Handler) ActivateCertificate(ctx context.Context, in *certificate.Id) (*common.Empty, error) {
	err := h.services.Mnemosyne.ActivateCertificate(ctx, in.GetId())
	if err != nil {
		return emptyProto, status.Error(codes.Internal, err.Error())
	}
	return emptyProto, nil
}
