package common

import "time"

// カスタム共通モデル
type CommonModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}

// 共通モデルのフィールドを初期化(作成日時、更新日時)
func InitializeCommonModel(m *CommonModel) {
	currentTime := time.Now()
	m.CreatedAt = currentTime
	m.UpdatedAt = currentTime
}
