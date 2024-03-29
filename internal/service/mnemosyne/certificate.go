package mnemosyne

import (
	"github.com/GolangUnited/students-dataservice-mnemosyne/models/database"
)
import "context"

func (s *Service) CreateCertificate(ctx context.Context, certificate database.Certificate) (certificateId uint32, err error) {
	return s.reposCertificate.CreateCertificate(ctx, certificate)

}

func (s *Service) GetCertificates(ctx context.Context, userId uint32) (certificates []database.Certificate, err error) {
	return s.reposCertificate.GetCertificates(ctx, userId)

}

func (s *Service) UpdateCertificate(ctx context.Context, certificate database.Certificate) error {
	return s.reposCertificate.UpdateCertificates(ctx, certificate)

}
func (s *Service) DeactivateCertificate(ctx context.Context, certificateId uint32) error {
	return s.reposCertificate.DeactivateCertificate(ctx, certificateId)

}

func (s *Service) ActivateCertificate(ctx context.Context, certificateId uint32) error {
	return s.reposCertificate.ActivateCertificate(ctx, certificateId)

}
