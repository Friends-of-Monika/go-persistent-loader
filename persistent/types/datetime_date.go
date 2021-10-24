package types

import (
	"time"
)

func init() {
	registerType("datetime", "date", &DatetimeDate{})
}

type DatetimeDate time.Time

func (d DatetimeDate) Call(args ...interface{}) (interface{}, error) {
	b := []byte(args[0].(string))
	return DatetimeDate(time.Date(int(b[0])*256+int(b[1]), time.Month(b[2]), int(b[3]), 0, 0, 0, 0, time.UTC)), nil
}
