package user

import (
	"context"

	"github.com/MrTomSawyer/chat/internal/app/model"
)

type UserStorageManager interface {
	Create(ctx context.Context, user *model.User) error
}

type Service struct {
	r UserStorageManager
}

func NewUserService(r UserStorageManager) *Service {
	return &Service{
		r: r,
	}
}

func (s *Service) Create(ctx context.Context, user *model.User) error {
	return s.r.Create(ctx, user)
}
