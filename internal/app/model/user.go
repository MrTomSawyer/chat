package model

import (
	"github.com/MrTomSawyer/chat/internal/app/apperrors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	DefaultFields
	Name         string `json:"name"`
	PasswordHash string `json:"password_hash"`
}

func (u *User) SetPassword(password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return apperrors.ErrFailedtoHashPasswd
	}
	u.PasswordHash = string(hash)
	return nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}
