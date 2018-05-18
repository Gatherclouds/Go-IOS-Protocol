package host

var l log.Logger

func Put(pool state.Pool, key state.Key, value state.Value) bool {
	pool.Put(key, value)
	return true
}

func Get(pool state.Pool, key state.Key) (state.Value, error) {
	return pool.Get(key)
}

func Log(s, cid string) {
	l.D("From Lua %v > %v", cid, s)
}

