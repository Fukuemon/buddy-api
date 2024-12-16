package common

import (
	"database/sql/driver"
	"time"

	errorDomain "api-buddy/domain/error"
	"encoding/json"
)

// 日付 (yyyy/mm/dd) のカスタム型
type Date struct {
	time.Time
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Format("2006-01-02"))
}

// Date型のカスタムフォーマット
func (d *Date) UnmarshalJSON(data []byte) error {
	var dateStr string
	if err := json.Unmarshal(data, &dateStr); err != nil {
		return err
	}
	parsed, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return err
	}
	d.Time = parsed
	return nil
}

// Date型: データベースへの保存
func (d Date) Value() (driver.Value, error) {
	return d.Format("2006-01-02"), nil
}

// Date型: データベースからの読み取り
func (d *Date) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		d.Time = v
		return nil
	case string:
		parsed, err := time.Parse("2006-01-02", v)
		if err != nil {
			return err
		}
		d.Time = parsed
		return nil
	default:
		return errorDomain.NewError("Date型に変換できません")
	}
}

// 時刻 (hh:mm) のカスタム型
type Time struct {
	time.Time
}

// Time型: JSONへの出力
func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Format("15:04")) // hh:mm形式
}

// Time型: JSONからの入力
func (t *Time) UnmarshalJSON(data []byte) error {
	var timeStr string
	if err := json.Unmarshal(data, &timeStr); err != nil {
		return err
	}
	parsed, err := time.Parse("15:04", timeStr)
	if err != nil {
		return err
	}
	t.Time = parsed
	return nil
}

// Time型: データベースへの保存
func (t Time) Value() (driver.Value, error) {
	return t.Format("15:04:05"), nil
}

// Time型: データベースからの読み取り
func (t *Time) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		t.Time = v
		return nil
	case string:
		parsed, err := time.Parse("15:04:05", v)
		if err != nil {
			return err
		}
		t.Time = parsed
		return nil
	case []uint8:
		strValue := string(v)
		parsed, err := time.Parse("15:04:05", strValue)
		if err != nil {
			return err
		}
		t.Time = parsed
		return nil
	default:
		return errorDomain.NewError("Time型に変換できません")
	}
}
