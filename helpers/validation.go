package helpers

import "time"

func DefaultIfEmpty[T any](value, defaultValue T) T {
	if str, ok := any(value).(string); ok && str == "" {
		return defaultValue
	}

	if num, ok := any(value).(int); ok && num == 0 {
		return defaultValue
	}

	if num64, ok := any(value).(int64); ok && num64 == 0 {
		return defaultValue
	}

	if numFloat, ok := any(value).(float64); ok && numFloat == 0 {
		return defaultValue
	}

	if numFloat32, ok := any(value).(float32); ok && numFloat32 == 0 {
		return defaultValue
	}

	if timeVal, ok := any(value).(time.Time); ok && timeVal.IsZero() {
		return defaultValue
	}

	return value
}
