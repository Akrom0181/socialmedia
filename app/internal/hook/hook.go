package hook

import (
	"log/slog"

	"github.com/pocketbase/pocketbase"
	"gitlab.saidoff.uz/company/muslim-administration/reading/back/internal/model"
	"gitlab.saidoff.uz/company/muslim-administration/reading/back/internal/service"
)

type Hook struct {
	logger  *slog.Logger
	service service.I
}

func (h *Hook) Register(app *pocketbase.PocketBase) {
	//app.OnRecordRequestOTPRequest(model.UsersCollection).BindFunc(recordRequestOTPRequestEventWrapper(h.otpRequest))
	//app.OnRecordRequestPasswordResetRequest(model.UsersCollection).BindFunc(recordRequestPasswordResetRequestEventWrapper(h.passwordResetRequest))
	//
	//app.OnMailerRecordPasswordResetSend().BindFunc(mailerRecordPasswordResetSendEventWrapper)
	//app.OnMailerRecordOTPSend().BindFunc(mailerRecordOTPSendEventWrapper)
	app.OnRecordCreateRequest(model.UsersCollection).BindFunc(recordRequestCheckPhoneNumberWrapper(h.checkPhoneNumber))
	app.OnRecordUpdateRequest(model.UsersCollection).BindFunc(recordRequestCheckPhoneNumberWrapper(h.checkPhoneNumber))
	app.OnRecordCreateRequest(model.UsersCollection).BindFunc(recordRequestCheckBirthDateWrapper(h.CheckBirthDate))
	app.OnRecordUpdateRequest(model.UsersCollection).BindFunc(recordRequestCheckBirthDateWrapper(h.CheckBirthDate))
}

func New(logger *slog.Logger, service service.I) *Hook {
	return &Hook{
		logger:  logger,
		service: service,
	}
}
