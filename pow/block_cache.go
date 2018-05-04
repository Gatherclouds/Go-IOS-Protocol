package pow

const (
	MaxCacheDepth = 6
)

type CacheStatus int

const (
	Extend     CacheStatus = iota
	Fork
	NotFound
	ErrorBlock
)
