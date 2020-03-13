package common

import (
	"strconv"
	"time"
)

func GetTimeStamp() string {
	timeInt := time.Now().UnixNano() / 1e6
	return strconv.FormatInt(timeInt, 10)
}
