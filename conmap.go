package conmap

import "sync"

type Map[T comparable, V comparable] struct {
	m     map[T]V
	mutex sync.RWMutex
}

/*
* Creates a new thread-safe conmap.
 */
func New[T comparable, V comparable]() *Map[T, V] {
	return &Map[T, V]{m: make(map[T]V)}
}

/*
* Returns the value for the given key in the map.
* The second return value indicates if the key exists in the map.
* NOTE: You should always check if the key exists before using the value to avoid nil pointer dereference.
* You can see the usage example in the README file.
 */
func (m *Map[T, V]) Load(key T) (value V, ok bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	value, ok = m.m[key]
	return
}

/*
* Sets the value for the given key in the map.
* If the key already exists, the value is updated.
 */
func (m *Map[T, V]) Store(key T, value V) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	m.m[key] = value
}

/*
* Deletes the element with the given key from the map.
 */
func (m *Map[T, V]) Delete(key T) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	delete(m.m, key)

}

/*
* Iterates over the map and calls the function f for each element.
* If f returns false, the iteration stops.
 */
func (m *Map[T, V]) Range(f func(key T, value V) bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	for k, v := range m.m {
		if !f(k, v) {
			break
		}
	}
}

/*
* Returns the number of elements in the map.
 */
func (m *Map[T, V]) Len() int {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	return len(m.m)
}
