package model

import "fmt"

type UpaUser interface {
	Format() string
}

func f() {
	var u UpaUser
	u.Format()
}

type UserInfo struct {
	CardNo, Name, IdCard, Mobile, AuthCode string
	AuthFlag                               bool
}

func (user *UserInfo) Format() string {
	if user.AuthFlag {
		return fmt.Sprintf("%s:%s:%s:%s:%s:%d", user.CardNo, user.Name, user.IdCard, user.Mobile, user.AuthCode, 1)
	}
	return fmt.Sprintf("%s:%s:%s:%s:%s:%d", user.CardNo, user.Name, user.IdCard, user.Mobile, user.AuthCode, 0)
}

type RecognizeUserInfo struct {
	CardNo, Name, IdCard, Mobile, Mode, MerName, AuthCode string
	AuthFlag                                              bool
}

func (user *RecognizeUserInfo) Format() string {
	if user.AuthFlag {
		return fmt.Sprintf("%s:%s:%s:%s:%s:%s:%s:%d", user.CardNo, user.Name, user.IdCard,
			user.Mobile, user.Mode, user.MerName, user.AuthCode, 1)
	}
	return fmt.Sprintf("%s:%s:%s:%s:%s:%s:%s:%d", user.CardNo, user.Name, user.IdCard,
		user.Mobile, user.Mode, user.MerName, user.AuthCode, 0)
}
