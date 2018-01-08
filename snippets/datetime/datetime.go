package datetime

import "time"

func ToUTCTime(unixMS int64) time.Time {
	return time.Unix(0, unixMS*int64(time.Millisecond)).UTC()
}
