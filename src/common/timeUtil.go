package common

import (
	"fmt"
	"time"
)

type DateTime time.Time

func (t DateTime) MarshalJSON() ([]byte, error) {
	//do your serializing here
	stamp := fmt.Sprintf("\"%s\"", t.ToLocal())
	return []byte(stamp), nil
}
func (t DateTime) ToLocal() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

func (t DateTime) ToLocalDate() string {
	return time.Time(t).Format("2006-01-02")
}

func (t DateTime) ToLocalTime() string {
	return time.Time(t).Format("15:04:05")
}
