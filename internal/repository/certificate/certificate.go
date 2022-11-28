package certificate

import (
	"context"
	"time"

	"github.com/NEKETSKY/mnemosyne/models/database"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

type CertificateRepository struct {
	db *pgx.Conn
}

func NewCertificateRepository(db *pgx.Conn) *CertificateRepository {
	return &CertificateRepository{
		db: db,
	}
}

func (u *CertificateRepository) AddCertificate(ctx context.Context, certificate database.Certificate) (certificateId int, err error) {
	row := u.db.QueryRow(
		ctx,
		AddCertificate,
		certificate.UserId,
		certificate.IssueDate,
		certificate.ExpireDate,
	)
	if err != nil {
		return 0, errors.Wrap(err, "AddCertificate query error")
	}
	err = row.Scan(&certificateId)

	if err != nil {
		return 0, errors.Wrap(err, "AddCertificate error while query executing")
	}

	return
}

func (u *CertificateRepository) GetCertificates(ctx context.Context) (certificates []database.Certificate, err error) {
	rows, _ := u.db.Query(ctx, GetAllCertificates)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get certificates from db")
	}
	certificates, err = pgx.CollectRows(rows, pgx.RowToStructByName[database.Certificate])
	if err != nil {
		return nil, errors.Wrap(err, "GetCertificates - unable to collect rows ")
	}
	return certificates, err
}

func (u *CertificateRepository) UpdateCertificatesById(ctx context.Context, certificate database.Certificate) (err error) {
	_, err = u.db.Exec(
		ctx,
		UpdateCertificatesById,
		certificate.UserId,
		certificate.IssueDate,
		certificate.ExpireDate,
	)
	if err != nil {
		return errors.Wrap(err, "UpdateCertificates - unable to execute update statement")
	}
	return err
}

func (u *CertificateRepository) ActivateCertificateById(ctx context.Context, certificateId int) (err error) {
	_, err = u.db.Exec(ctx, ActivateById, certificateId, time.Now())
	if err != nil {
		return errors.Wrapf(err, "unable to set certificate %d as active", certificateId)
	}
	return err
}

func (u *CertificateRepository) DeactivateCertificateById(ctx context.Context, certificateId int) (err error) {
	_, err = u.db.Exec(ctx, DeactivateById, certificateId, time.Now())
	if err != nil {
		return errors.Wrapf(err, "unable to set certificate %d as deleted", certificateId)
	}
	return err
}
