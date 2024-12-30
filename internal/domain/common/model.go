package common

import (
	"strings"
	"time"
)

// カスタム共通モデル
type CommonModel struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// 共通モデルのフィールドを初期化(作成日時、更新日時)
func InitializeCommonModel(m *CommonModel) {
	currentTime := time.Now()
	m.CreatedAt = currentTime
	m.UpdatedAt = currentTime
}

func AddPlusToPhoneNumber(phoneNumber string) string {
	if phoneNumber[0] != '+' {
		phoneNumber = "+" + phoneNumber
	}
	return phoneNumber
}

func IsPhoneNumber(phoneNumber string) bool {
	for _, c := range phoneNumber {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func IsEmail(email string) bool {
	// 文字列に@と.が含まれているかチェック
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return false
	}
	return true
}
func StringPointer(s string) *string {
	return &s
}

const (
	UpdatedAt = "updated_at"
	CreatedAt = "created_at"
)
