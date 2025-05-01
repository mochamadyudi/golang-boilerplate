package repository

import (
	"context"

	"core.yuyuid.id/internal/app/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func (dt *UserRepositoryImpl) FindByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	err := dt.DB.WithContext(ctx).Where("email = ? AND deleted_at IS NULL ", email).First(&user).Error
	return user, err
}
func (dt *UserRepositoryImpl) Save(ctx context.Context, user *model.User) error {
	if err := dt.DB.WithContext(ctx).Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (dt *UserRepositoryImpl) FindById(ctx context.Context, id uint, modifier func(*gorm.DB) *gorm.DB) (model.User, error) {
	var user model.User
	db := dt.DB.WithContext(ctx)
	if modifier != nil {
		db = modifier(db)
	}
	if err := db.First(&user, id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (dt *UserRepositoryImpl) IsExistEmail(ctx context.Context, email string) error {
	var user model.User
	if err := dt.DB.WithContext(ctx).Where("email = ? AND deleted_at IS NULL ", email).First(&user).Error; err != nil {
		return err
	}
	return nil
}

func (dt *UserRepositoryImpl) IsExistById(ctx context.Context, id uint) error {
	var user model.User
	err := dt.DB.WithContext(ctx).Where("id = ? AND deleted_at IS NULL", id).First(&user).Error
	return err
}
func NewUserRepository(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: db}
}

func (dt *UserRepositoryImpl) FindUserWithRoles(ctx context.Context, userId uint) (model.User, error) {
	var user model.User
	if err := dt.DB.WithContext(ctx).Preload("Roles").First(&user, userId).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (dt *UserRepositoryImpl) FindAll(ctx context.Context, modifier func(*gorm.DB) *gorm.DB) ([]model.User, error) {
	var users []model.User

	db := dt.DB.WithContext(ctx)
	if modifier != nil {
		db = modifier(db)
	}
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
