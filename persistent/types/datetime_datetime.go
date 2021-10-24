package types

import (
	"time"
)

func init() {
	registerType("datetime", "datetime", &DatetimeDatetime{})
}

type DatetimeDatetime time.Time

func (d *DatetimeDatetime) Call(args ...interface{}) (interface{}, error) {
	b := []byte(args[0].(string))
	*d = DatetimeDatetime(time.Date(int(b[0])*256+int(b[1]), time.Month(b[2]), int(b[3]), int(b[4]), int(b[5]), int(b[6]), ((int(b[7])<<8)|int(b[8]))|int(b[9]), time.UTC))
	return d, nil
}
