package host

var l log.Logger

func Put(pool state.Pool, key state.Key, value state.Value) bool {
	pool.Put(key, value)
	return true
}

