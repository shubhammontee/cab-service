package date_utils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05.05Z"
	apiDBLayout   = "2006-01-02 15:04:05"
)

//GetNow ...
func GetNow() time.Time {
	return time.Now().UTC()
}

//GetNowString date in api string formatt
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

//GetNowDBFormat format in which we want to store in database
func GetNowDBFormat() string {
	return GetNow().Format(apiDBLayout)
}
