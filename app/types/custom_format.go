package types

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type CustomFormat struct {
	Value string
}

// UnmarshalJSON parses the JSON-encoded data and stores the result.
func (cf *CustomFormat) UnmarshalJSON(data []byte) error {
	var obj map[string]string
	if err := json.Unmarshal(data, &obj); err != nil {
		return err
	}

	str, ok := obj["value"]
	if !ok {
		return errors.New("missing value field")
	}

	parts := strings.Split(str, "_$$$_")
	if len(parts) != 2 || parts[0] != "ALICE" {
		return errors.New("invalid format")
	}

	cf.Value = parts[1]
	return nil
}

// MarshalJSON returns the JSON encoding of CustomFormat.
func (cf CustomFormat) MarshalJSON() ([]byte, error) {
	str := fmt.Sprintf("ALICE_$$$_%s", cf.Value)
	return json.Marshal(map[string]string{"value": str})
}

// ToString returns the string representation of CustomFormat.
func (cf CustomFormat) ToString() string {
	return cf.Value
}

// FromString parses a string into CustomFormat.
func (cf *CustomFormat) FromString(str string) error {
	cf.Value = str
	return nil
}
