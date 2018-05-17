package common

import "time"

const (
	SlotLength    = 3 //一个slot设为3秒
	SecondsInHour = 3600
	SecondsInDay  = 24 * 3600
	Epoch         = 1522540800 //设为2018-04-01 00:00:00的Unix秒数
)

type Timestamp struct {
	Slot int64
}

// 返回当前时间对应的时间戳
func GetCurrentTimestamp() Timestamp {
	t := time.Now()
	return GetTimestamp(t.Unix())
}



