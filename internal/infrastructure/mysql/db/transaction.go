package db

import (
	"context"

	"gorm.io/gorm"
)

// TransactionHandler はトランザクション内での処理を定義する関数
type TransactionHandler func(tx *gorm.DB) error

// WithTransaction はトランザクションを開始し、指定されたハンドラーを実行します。
func WithTransaction(ctx context.Context, handler TransactionHandler) error {
	db := GetDB()
	tx := db.Begin() // トランザクション開始
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
