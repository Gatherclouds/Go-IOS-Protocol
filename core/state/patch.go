package state

// state pool记录状态转移的结构，可以序列化到其他地方存储
type Patch struct {
	m map[Key]Value
}

func (p *Patch) Put(key Key, value Value) {
	p.m[key] = value
}
func (p *Patch) Get(key Key) Value {
	val, ok := p.m[key]
	if !ok {
		return nil
	}
	return val
}
func (p *Patch) Has(key Key) bool {
	_, ok := p.m[key]
	return ok
}

