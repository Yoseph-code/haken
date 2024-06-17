package safe_map

import "sync"

type SafeMap[T interface{}] struct {
	mu sync.RWMutex
	m  map[string]T
}

func NewSafeMap[T interface{}]() *SafeMap[T] {
	return &SafeMap[T]{
		m: make(map[string]T),
	}
}
