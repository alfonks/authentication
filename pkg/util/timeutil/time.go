package timeutil

import (
	"time"
)

func FindDifferenceFromUTC() time.Duration {
	now := time.Now()
	utcNow := now.UTC()
	return utcNow.Sub(now)
}
