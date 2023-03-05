package gothonMaps

//the Keys() method returns the keys associated to a map.
func Keys[K comparable, V any](m map[K]V) []K {
	var keys []K
	for k, _ := range m {
		keys = append(keys, k)
	}

	return keys
}

//the Values() method returns the values associated to a map.
func Values[K comparable, V any](m map[K]V) []V {
	var values []V
	for _, v := range m {
		values = append(values, v)
	}

	return values
}

//the Delete() method return the deleted key, value from a given map.
func Delete[K comparable, V any](m map[K]V, key K) (delKey K, delValue V) {
	value, ok := m[key]
	if !ok {
		return
	}
	delete(m, key)

	return key, value
}

//the SetDefault()  method returns the value of the item with the specified key.
//If the key does not exist, insert the key, with the specified value.
func SetDefault[K comparable, V any](m map[K]V, key K, value V) V {
	val, ok := m[key]
	if !ok {
		m[key] = value
		val = value
	}

	return val
}

func Items[K comparable, V any](m map[K]V) []Pair[K, V] {
	var items []Pair[K, V]
	for key, value := range m {
		items = append(items, Pair[K, V]{key, value})
	}

	return items
}