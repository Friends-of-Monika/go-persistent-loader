package types

import (
	"time"
)

func init() {
	class := DatetimeTimedelta(0)
	registerType("datetime", "timedelta", &class)
}

type DatetimeTimedelta time.Duration

func (d DatetimeTimedelta) Call(args ...interface{}) (interface{}, error) {
	return DatetimeTimedelta(args[0].(int)*8.64e+13 + args[1].(int)*1e+9 + args[2].(int)*1000), nil
}
