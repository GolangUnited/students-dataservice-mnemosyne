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

func (c *CertificateRepository) CreateCertificate(ctx context.Context, certificate database.Certificate) (certificateId int, err error) {
	row := c.db.QueryRow(
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
func (c *CertificateRepository) GetCertificateById(ctx context.Context, certificateId int) (certificate database.Certificate, err error) {

	rows, err := c.db.Query(ctx, GetCertificateById, certificateId)
	if err != nil {
		return database.Certificate{}, errors.Wrap(err, "GetCertificateById query error")
	}
	certificate, err = pgx.CollectOneRow(rows, pgx.RowToStructByName[database.Certificate])
	if err != nil {
		return database.Certificate{}, errors.Wrap(err, "GetCertificateById CollectOneRow error")
	}
	return
}
func (c *CertificateRepository) GetCertificates(ctx context.Context, userId int) (certificates []database.Certificate, err error) {
	rows, _ := c.db.Query(ctx, GetAllCertificates, userId)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get certificates from db")
	}
	certificates, err = pgx.CollectRows(rows, pgx.RowToStructByName[database.Certificate])
	if err != nil {
		return nil, errors.Wrap(err, "GetCertificates - unable to collect rows ")
	}
	return certificates, err
}

func (c *CertificateRepository) UpdateCertificates(ctx context.Context, certificate database.Certificate) (err error) {
	_, err = c.db.Exec(
		ctx,
		UpdateCertificateById,
		certificate.UserId,
		certificate.IssueDate,
		certificate.ExpireDate,
	)
	if err != nil {
		return errors.Wrap(err, "UpdateCertificates - unable to execute update statement")
	}
	return err
}

func (c *CertificateRepository) ActivateCertificate(ctx context.Context, certificateId int) (err error) {
	_, err = c.db.Exec(ctx, ActivateById, certificateId, time.Now())
	if err != nil {
		return errors.Wrapf(err, "unable to set certificate %d as active", certificateId)
	}
	return err
}

func (c *CertificateRepository) DeactivateCertificate(ctx context.Context, certificateId int) (err error) {
	_, err = c.db.Exec(ctx, DeactivateById, certificateId, time.Now())
	if err != nil {
		return errors.Wrapf(err, "unable to set certificate %d as deleted", certificateId)
	}
	return err
}
