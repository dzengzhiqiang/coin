package utils

import (
	"fmt"
	"time"
)

type MilliSecond int64

func (m MilliSecond) String() string {
	return fmt.Sprintf("%d", m)
}

func UnixMilliSecond() MilliSecond {
	now := time.Now()
	return MilliSecond(now.UnixNano() / 1000000)
}
