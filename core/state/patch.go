package state

// state pool记录状态转移的结构，可以序列化到其他地方存储
type Patch struct {
	m map[Key]Value
}

