package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type JSONB map[string]interface{}

func (j JSONB) Value() (driver.Value, error) {
	value, err := json.Marshal(j)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal JSONB value: %w", err)
	}
	return string(value), nil
}

func (j *JSONB) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to unmarshal JSONB value: %v", value)
	}
	err := json.Unmarshal(bytes, j)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSONB value: %w", err)
	}
	return nil
}
