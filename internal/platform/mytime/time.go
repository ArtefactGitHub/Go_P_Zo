package mytime

import (
	"database/sql"
	"time"
)

var (
	ToTimeLayout = "2006-01-02 15:04"
)

func ToTime(str string) time.Time {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	t, _ := time.ParseInLocation(ToTimeLayout, str, jst)
	return t
}

func ToNullTime(str string) sql.NullTime {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	t, _ := time.ParseInLocation(ToTimeLayout, str, jst)
	return sql.NullTime{Time: t, Valid: true}
}
