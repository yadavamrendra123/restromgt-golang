package types

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type DateOnly struct {
	time.Time
}

func (d *DateOnly) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	parsed, err := time.Parse("2006-01-02", str)
	if err != nil {
		return err
	}
	d.Time = parsed
	return nil
}

func (d DateOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format("2006-01-02"))
}

func (d DateOnly) Value() (driver.Value, error) {
	return d.Time.Format("2006-01-02"), nil
}

func (d *DateOnly) Scan(value interface{}) error {
	if value == nil {
		*d = DateOnly{time.Time{}}
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		*d = DateOnly{v}
	case []byte:
		t, err := time.Parse("2006-01-02", string(v))
		if err != nil {
			return err
		}
		*d = DateOnly{t}
	default:
		return errors.New("invalid type for DateOnly")
	}
	return nil
}
