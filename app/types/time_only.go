package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type TimeOnly struct {
	time.Time
}

func (t *TimeOnly) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	parsed, err := time.Parse("15:04:05", str)
	if err != nil {
		return err
	}
	t.Time = parsed
	return nil
}

func (t TimeOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time.Format("15:04:05"))
}

func (t TimeOnly) Value() (driver.Value, error) {
	return t.Time.Format("15:04:05"), nil
}

func (t *TimeOnly) Scan(value interface{}) error {
	if value == nil {
		*t = TimeOnly{time.Time{}}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*t = TimeOnly{v}
	case []byte:
		parsedTime, err := time.Parse("15:04:05", string(v))
		if err != nil {
			return err
		}
		*t = TimeOnly{parsedTime}
	default:
		return errors.New("invalid type for TimeOnly")
	}
	return nil
}
