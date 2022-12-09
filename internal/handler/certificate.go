package handler

import (
	"context"
	"github.com/NEKETSKY/mnemosyne/models/database"
	"github.com/NEKETSKY/mnemosyne/pkg/api/certificate"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) CreateCertificate(ctx context.Context, in *certificate.CertificateRequest) (cert *certificate.CertificateResponse, err error) {
	cert = &certificate.CertificateResponse{}
	certificateDb, err := database.CertificateFromProto(in)
	if err != nil {
		return cert, status.Error(codes.InvalidArgument, err.Error())
	}
	certificateId, err := h.services.Mnemosyne.CreateCertificate(ctx, certificateDb)
	if err != nil {
		return cert, status.Error(codes.Internal, err.Error())
	}
	cert.Id = uint32(certificateId)
	return
}

func (h *Handler) GetCertificates(ctx context.Context, in *certificate.Filter) (certificates *certificate.Certificates, err error) {
	certificateDb, err := h.services.Mnemosyne.GetCertificates(ctx, int(in.GetUserId()))
	if err != nil {
		return certificates, status.Error(codes.Internal, err.Error())
	}
	var slice []*certificate.CertificateResponse
	for _, i := range certificateDb {
		slice = append(slice, i.CertificateToProto())
	}
	return &certificate.Certificates{Certificates: slice}, err
}

func (h *Handler) UpdateCertificate(ctx context.Context, in *certificate.CertificateRequest) (cert *certificate.CertificateResponse, err error) {
	certificateDb, err := database.CertificateFromProto(in)
	if err != nil {
		return cert, status.Error(codes.InvalidArgument, err.Error())
	}
	err = h.services.Mnemosyne.UpdateCertificate(ctx, certificateDb)
	if err != nil {
		return cert, status.Error(codes.Internal, err.Error())
	}
	return
}

func (h *Handler) DeactivateCertificate(ctx context.Context, in *certificate.Id) (certificates *certificate.CertificateResponse, err error) {
	err = h.services.Mnemosyne.DeactivateCertificate(ctx, int(in.GetId()))
	if err != nil {
		return certificates, status.Error(codes.Internal, err.Error())
	}
	return
}

func (h *Handler) ActivateCertificate(ctx context.Context, in *certificate.Id) (certificates *certificate.CertificateResponse, err error) {
	err = h.services.Mnemosyne.ActivateCertificate(ctx, int(in.GetId()))
	if err != nil {
		return certificates, status.Error(codes.Internal, err.Error())
	}
	return
}
