package service

import (
	"regexp"
	"time"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/core"
)

type UsersS struct {
	db dbx.Builder
}

func NewUsersS(db dbx.Builder) *UsersS {
	return &UsersS{db: db}
}

func (u *UsersS) CheckPhoneNumber(e *core.RecordRequestEvent) error {
	phone := e.Record.GetString("phone_number")
	pattern := `^\+998[0-9]{9}$`

	ok, _ := regexp.MatchString(pattern, phone)

	if phone == "" {
		return nil
	}

	if !ok {
		return e.JSON(400, map[string]interface{}{
			"message": "invalid phone number",
			"status":  false,
		})
	}

	return nil
}

func (u *UsersS) CheckBirthDate(e *core.RecordRequestEvent) error {
	birthday := e.Record.GetDateTime("birth_day").Time().Format("2006-01-02")
	layout := "2006-01-02"

	now := time.Now().Format(layout)

	birthdate, err1 := time.Parse(layout, birthday)
	nowday, err2 := time.Parse(layout, now)
	if err2 != nil || err1 != nil {
		return e.JSON(400, map[string]interface{}{
			"message": "invalid date format",
			"status":  false,
		})
	}
	if nowday.Before(birthdate) {
		return e.JSON(400, map[string]interface{}{
			"message": "invalid birthday",
			"status":  false,
		})
	}

	return nil
}
