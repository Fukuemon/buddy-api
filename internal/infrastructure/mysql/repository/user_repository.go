package repository

import (
	userDomain "api-buddy/domain/user"
	"api-buddy/infrastructure/mysql/db"
	"context"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() userDomain.UserRepository {
	return &UserRepository{
		db: db.GetDB(),
	}
}

func (r *UserRepository) Create(ctx context.Context, user *userDomain.User) error {
	err := r.db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*userDomain.User, error) {
	var user userDomain.User
	err := r.db.Preload("Facility").Preload("Position").Preload("Team").Preload("Department").Preload("Policies").Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
