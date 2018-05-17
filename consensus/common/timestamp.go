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

// 根据一个Unix时间点（秒为单位）返回一个时间戳
func GetTimestamp(timeSec int64) Timestamp {
	return Timestamp{(timeSec - Epoch) / SlotLength}
}

func (t *Timestamp) AddDay(intervalDay int) {
	t.Slot = t.Slot + int64(intervalDay)*SecondsInDay/SlotLength
}

func (t *Timestamp) AddHour(intervalHour int) {
	t.Slot = t.Slot + int64(intervalHour)*SecondsInHour/SlotLength
}