package safe_map

func (sm *SafeMap[T]) Get(key string) (T, bool) {
	sm.mu.RLock()
	value, ok := sm.m[key]
	sm.mu.RUnlock()
	return value, ok
}

func (sm *SafeMap[T]) Set(key string, value T) {
	sm.mu.Lock()
	sm.m[key] = value
	sm.mu.Unlock()
}

func (sm *SafeMap[T]) Delete(key string) {
	sm.mu.Lock()
	delete(sm.m, key)
	sm.mu.Unlock()
}

func (sm *SafeMap[T]) Clear() {
	sm.mu.Lock()
	sm.m = make(map[string]T)
	sm.mu.Unlock()
}

func (sm *SafeMap[T]) Update(key string, value T) {
	sm.mu.Lock()
	sm.m[key] = value
	sm.mu.Unlock()
}

func (sm *SafeMap[T]) UpdateAll(m map[string]T) {
	sm.mu.Lock()
	for k, v := range m {
		sm.m[k] = v
	}
	sm.mu.Unlock()
}

func (sm *SafeMap[T]) Merge(m map[string]T) {
	sm.mu.Lock()
	for k, v := range m {
		if _, ok := sm.m[k]; !ok {
			sm.m[k] = v
		}
	}
	sm.mu.Unlock()
}
