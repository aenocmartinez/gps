package util

import "time"

const DDMMYYYYhhmmss = "2006-01-02 15:04:05"
const YYYYmmdd = "2006-01-02"

func CurrentDateFormatTimestamp() string {
	return time.Now().Format(DDMMYYYYhhmmss)
}
