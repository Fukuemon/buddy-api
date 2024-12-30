package db

import (
	"context"

	"gorm.io/gorm"
)

type TransactionHandler func(tx *gorm.DB) error

func WithTransaction(ctx context.Context, handler TransactionHandler) error {
	db := GetDB()
	tx := db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	// ハンドラーを実行
	if err := handler(tx); err != nil {
		tx.Rollback()
		return err
	}

	// コミット
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
