package service

import (
	"context"

	"core.yuyuid.id/internal/app/model"
	"core.yuyuid.id/internal/app/repository"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	Repo repository.UserRepositoryImpl
}

func (u *UserServiceImpl) FindByEmail(param any, s string) any {
	panic("unimplemented")
}

func (s *UserServiceImpl) SaveUser(ctx context.Context, user *model.User) error {
	return s.Repo.Save(ctx, user)
}

func (s *UserServiceImpl) FindByID(ctx context.Context, id uint, modifier func(*gorm.DB) *gorm.DB) (model.User, error) {
	return s.Repo.FindById(ctx, id, modifier)
}

func (s *UserServiceImpl) IsExist(ctx context.Context, id uint) error {
	return s.Repo.IsExistById(ctx, id)
}
func NewUserService(repo repository.UserRepositoryImpl) *UserServiceImpl {
	return &UserServiceImpl{
		Repo: repo,
	}
}
