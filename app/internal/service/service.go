package service

import (
	"context"
	"time"

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

type RedisI interface {
	SetX(ctx context.Context, key string, value interface{}, duration time.Duration) error
	GetX(ctx context.Context, key string) (interface{}, error)
	DelX(ctx context.Context, key string) error
}

type I interface {
	Authorization() AuthorizationI
	Users() UsersI
	Redis() RedisI
}

type service struct {
	AuthorizationI
	UsersI
	RedisI
}

func (s *service) Authorization() AuthorizationI {
	return s.AuthorizationI
}

func (s *service) Users() UsersI {
	return s.UsersI
}

func NewService(db dbx.Builder, redis Store) I {
	return &service{
		AuthorizationI: NewAuthorizationS(db),
		UsersI:         NewUsersS(db),
		RedisI:         &redis,
	}
}

func (s *service) Redis() RedisI {
	return s.RedisI
}
