package hook

import "github.com/pocketbase/pocketbase/core"

func (h *Hook) otpRequest(e *core.RecordCreateOTPRequestEvent) error {
	return h.service.Authorization().OtpRequest(e)
}

func (h *Hook) passwordResetRequest(e *core.RecordRequestPasswordResetRequestEvent) error {
	return h.service.Authorization().ResetPasswordRequest(e)
}

func (h *Hook) checkPhoneNumber(e *core.RecordRequestEvent) error {
	return h.service.Users().CheckPhoneNumber(e)
}

func (h *Hook) CheckBirthDate(e *core.RecordRequestEvent) error{
	return h.service.Users().CheckBirthDate(e)
}