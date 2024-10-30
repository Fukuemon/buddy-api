package repository

import (
	"api-buddy/domain/common"
	errorDomain "api-buddy/domain/error"
	userDomain "api-buddy/domain/user"
	"api-buddy/infrastructure/mysql/db"
	"context"

	"github.com/Fukuemon/go-pkg/query"

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

func (r *UserRepository) FindByFacilityID(ctx context.Context, facility_id string, filters []query.Filter, sort query.SortOption) ([]*userDomain.User, error) {
	// Queryオブジェクトを初期化
	q := query.NewQuery()

	// Queryオブジェクトにフィルターを適用
	for _, filter := range filters {
		filter.Apply(q)
	}

	dbQuery := r.db

	// リレーションテーブルに基づいたフィルタリングのQueryを適用
	for _, mapping := range userDomain.UserRelationMappings {
		if value, exists := q.Filters[mapping.FilterField]; exists {
			dbQuery = dbQuery.Joins("JOIN "+mapping.TableName+" ON "+mapping.JoinKey).
				Where(mapping.FilterField+" = ?", value)
		}
	}

	// 残りのフィルタリング（`users` テーブルに対するフィルター）を適用
	for key, value := range q.Filters {
		if _, isRelationField := userDomain.UserRelationMappings[key]; !isRelationField {
			dbQuery = dbQuery.Where(key, value)
		}
	}

	// ソートオプションの適用
	if sort.Field != "" {
		// ソートフィールドがリレーションに含まれるかを確認
		if mapping, exists := userDomain.UserRelationMappings[sort.Field]; exists {
			// リレーションテーブルのフィールドをソートに使用
			dbQuery = dbQuery.Joins("JOIN " + mapping.TableName + " ON " + mapping.JoinKey).
				Order(mapping.FilterField + " " + sort.Order)
		} else {
			// ユーザーテーブルのフィールドをソートに使用
			dbQuery = dbQuery.Order(sort.Field + " " + sort.Order)
		}
	} else {
		dbQuery = dbQuery.Order(common.UpdatedAt + " " + query.DESC) // デフォルトのソート条件
	}

	// クエリ実行
	var users []*userDomain.User
	err := dbQuery.Preload("Facility").Preload("Position").Preload("Team").Preload("Department").Preload("Policies").Preload("Area").
		Where("users.facility_id = ?", facility_id).Find(&users).Error
	if err != nil {
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}

	return users, nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*userDomain.User, error) {
	var user userDomain.User
	err := r.db.Preload("Facility").Preload("Position").Preload("Team").Preload("Department").Preload("Policies").Preload("Area").Where("id = ?", id).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorDomain.WrapError(errorDomain.NotFoundErr, err)
		}
		return nil, errorDomain.WrapError(errorDomain.GeneralDBError, err)
	}
	return &user, nil
}
