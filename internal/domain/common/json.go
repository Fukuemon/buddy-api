package common

import (
	errorDomain "api-buddy/domain/error"
	"database/sql/driver"
	"encoding/json"
)

type JSONSlice[T any] []T

func (j *JSONSlice[T]) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errorDomain.NewError("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, j)
}

// Value implements the driver.Valuer interface
func (j JSONSlice[T]) Value() (driver.Value, error) {
	if len(j) == 0 {
		return "[]", nil
	}
	return json.Marshal(j)
}
