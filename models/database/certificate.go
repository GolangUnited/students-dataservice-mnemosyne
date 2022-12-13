package database

import (
	"github.com/NEKETSKY/mnemosyne/pkg/api/certificate"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Certificate struct {
	Id         uint32    `json:"id" db:"id"`
	UserId     uint32    `json:"user_id" db:"user_id"`
	IssueDate  time.Time `json:"issue_date" db:"issue_date"`
	ExpireDate time.Time `json:"expire_date" db:"expire_date"`
}

func (c *Certificate) CertificateToProto() *certificate.CertificateResponse {
	return &certificate.CertificateResponse{
		Id:         c.Id,
		UserId:     c.UserId,
		IssueDate:  timestamppb.New(c.IssueDate),
		ExpireDate: timestamppb.New(c.ExpireDate),
	}
}

func CertificateFromProto(protoCertificate *certificate.CertificateRequest) (c Certificate, err error) {
	c.Id = protoCertificate.GetId()
	c.UserId = protoCertificate.GetUserId()
	c.IssueDate = protoCertificate.IssueDate.AsTime()
	c.ExpireDate = protoCertificate.ExpireDate.AsTime()
	
	return
}
