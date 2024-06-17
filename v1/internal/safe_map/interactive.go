package safe_map

func (sm *SafeMap[T]) Has(key string) bool {
	sm.mu.RLock()
	_, ok := sm.m[key]
	sm.mu.RUnlock()
	return ok
}

func (sm *SafeMap[T]) Len() int {
	sm.mu.RLock()
	l := len(sm.m)
	sm.mu.RUnlock()
	return l
}

func (sm *SafeMap[T]) Keys() []string {
	sm.mu.RLock()
	keys := make([]string, 0, sm.Len())
	for k := range sm.m {
		keys = append(keys, k)
	}
	sm.mu.RUnlock()
	return keys
}

func (sm *SafeMap[T]) Values() []T {
	sm.mu.RLock()
	values := make([]T, 0, sm.Len())
	for _, v := range sm.m {
		values = append(values, v)
	}
	sm.mu.RUnlock()
	return values
}

func (sm *SafeMap[T]) Copy() map[string]T {
	sm.mu.RLock()
	m := make(map[string]T, sm.Len())
	for k, v := range sm.m {
		m[k] = v
	}
	sm.mu.RUnlock()
	return m
}

func (sm *SafeMap[T]) Uion(m *SafeMap[T]) map[string]T {
	sm.mu.Lock()

	data := sm.Copy()

	for k, v := range m.Copy() {
		if _, ok := data[k]; !ok {
			data[k] = v
		}
	}

	sm.mu.Unlock()

	return data
}
