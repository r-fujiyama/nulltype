package nulltype

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
)

// Int32 represents a int32 that may be null.
type Int32 struct {
	Int32 int32
	Valid bool
}

// Scan implements the Scanner interface.
func (i *Int32) Scan(value interface{}) error {
	if value == nil {
		i.Int32, i.Valid = 0, false
		return nil
	}

	i.Valid = true
	switch data := value.(type) {
	case string:
		i32, err := strconv.ParseInt(data, 10, 32)
		if err != nil {
			return err
		}
		i.Int32 = int32(i32)
		return nil
	case []byte:
		i32, err := strconv.ParseInt(string(data), 10, 32)
		if err != nil {
			return err
		}
		i.Int32 = int32(i32)
		return nil
	case int:
		i.Int32 = int32(data)
		return nil
	case int16:
		i.Int32 = int32(data)
		return nil
	case int32:
		i.Int32 = data
		return nil
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
}

// Value implements the driver Valuer interface.
func (i Int32) Value() (driver.Value, error) {
	if !i.Valid {
		return nil, nil
	}
	return int64(i.Int32), nil
}

// MarshalJSON encode the value to JSON.
func (i Int32) MarshalJSON() ([]byte, error) {
	if !i.Valid {
		return []byte("null"), nil
	}
	return JSONMarshal(i.Int32)
}

// UnmarshalJSON decode data to the value.
func (i *Int32) UnmarshalJSON(data []byte) error {
	var fl *int32
	if err := json.Unmarshal(data, &fl); err != nil {
		return err
	}
	i.Valid = fl != nil
	if fl == nil {
		i.Int32 = 0
	} else {
		i.Int32 = *fl
	}
	return nil
}

// IsZeroOrNull return true if int32 is 0 or Valid is false.
func (i *Int32) IsZeroOrNull() bool {
	return i.Int32 == 0 || !i.Valid
}

// String return string indicated the value.
func (i Int32) String() string {
	if !i.Valid {
		return "<null>"
	}
	return strconv.Itoa(int(i.Int32))
}
