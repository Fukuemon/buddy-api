package repository

import (
	errorDomain "api-buddy/domain/error"
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
		return errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*userDomain.User, error) {
	var user userDomain.User
	err := r.db.Preload("Facility").Preload("Position").Preload("Team").Preload("Department").Preload("Policies").Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorDomain.WrapError(errorDomain.NotFoundErr, err)
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return &user, nil
}
