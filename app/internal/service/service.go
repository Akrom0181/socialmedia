package service

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
	"gitlab.saidoff.uz/company/muslim-administration/reading/back/internal/model"
)

type AuthorizationI interface {
	OtpRequest(e *core.RecordCreateOTPRequestEvent) error
	ResetPasswordRequest(e *core.RecordRequestPasswordResetRequestEvent) error
	ResetPasswordOTPConfirm(req *model.PasswordResetOTPConfirmRequest) (string, error)
}

type UsersI interface {
	CheckPhoneNumber(e *core.RecordRequestEvent) error
	CheckBirthDate(e *core.RecordRequestEvent) error
}

type I interface {
	Authorization() AuthorizationI
	Users() UsersI
}

type service struct {
	AuthorizationI
	UsersI
}

func (s *service) Authorization() AuthorizationI {
	return s.AuthorizationI
}

func (s *service) Users() UsersI {
	return s.UsersI
}

func NewService(db dbx.Builder) I {
	return &service{
		AuthorizationI: NewAuthorizationS(db),
		UsersI:         NewUsersS(db),
	}
}
