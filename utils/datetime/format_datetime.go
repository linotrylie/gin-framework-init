package datetime

import (
	"database/sql/driver"
	"equity/utils/strutil"
	"fmt"
	"time"
)

// JSONTime format json time field by myself
type JSONTime struct {
	time.Time
}

// MarshalJSON on JSONTime format Time field with %Y-%m-%d %H:%M:%S
func (t JSONTime) MarshalJSON() ([]byte, error) {
	if (t == JSONTime{}) {
		formatted := fmt.Sprintf("\"%s\"", "")
		return []byte(formatted), nil
	}

	formatted := fmt.Sprintf("\"%s\"", t.Format(DateYYYYMMDDhhmmssLayout))
	if formatted == `"1970-01-01 08:00:00"` {
		return []byte(`""`), nil
	}
	return []byte(formatted), nil
}

// Value insert timestamp into mysql need this function.
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan valueof time.Time
func (t *JSONTime) Scan(v interface{}) error {
	// 时间戳转换
	timeStampUint8, okUint8 := v.([]uint8)
	if okUint8 {
		timeSec, err := strutil.Atoi(string(timeStampUint8))
		if err != nil {
			return err
		}
		value := time.Unix(int64(timeSec), 0)
		*t = JSONTime{Time: value}
		return nil
	}

	timeStamp, okInt64 := v.(int64)
	if okInt64 {
		value := time.Unix(timeStamp, 0)
		*t = JSONTime{Time: value}
		return nil
	}

	// 日期时间转换
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Time: value}
		return nil
	}
	return nil
}
