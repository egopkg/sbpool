package sbpool

import (
	"strings"
	"sync"
)

var (
	stringsBuilderPool sync.Pool
)

// AcquireStringsBuilder returns an empty Request strings.Builder from pool.
//
// The returned strings.Builder instance may be passed to ReleaseStringsBuilder when it is
// no longer needed. This allows strings.Builder recycling, reduces GC pressure
// and usually improves performance.
func AcquireStringsBuilder() *strings.Builder {
	v := stringsBuilderPool.Get()
	if v == nil {
		return &strings.Builder{}
	}
	return v.(*strings.Builder)
}

// ReleaseStringsBuilder return strings.Builder acquired via AcquireStringsBuilder to the pool.
//
// It is forbidden accessing strings.Builder and/or its' members after returning it to response pool.
func ReleaseStringsBuilder(u *strings.Builder) {
	u.Reset()
	stringsBuilderPool.Put(u)
}
