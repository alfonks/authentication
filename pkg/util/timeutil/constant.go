package timeutil

import "time"

const (
	TimeOneSecond  = time.Second
	TimeOneMinutes = time.Minute
	TimeOneHour    = time.Hour
	TimeOneDay     = TimeOneHour * 24
	TimeOneYear    = TimeOneDay * 365
)
