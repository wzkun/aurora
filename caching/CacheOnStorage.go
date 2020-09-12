package caching

// CacheOnStorage struct
type CacheOnStorage struct {
	cache Storage
	store Storage
}

// NewCacheOnStorage create CacheOnStorage
func NewCacheOnStorage(cache, store Storage) *CacheOnStorage {
	o := new(CacheOnStorage)
	o.cache = cache
	o.store = store
	return o
}

// Get function
func (m *CacheOnStorage) Get(k Key) (Item, error) {
	item, err := m.cache.Get(k)
	if err != nil {
		return nil, err
	}

	if item != nil {
		return item, nil
	}

	item, err = m.store.Get(k)
	if err != nil {
		return nil, err
	}

	err = m.cache.Set(item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

// GetMulti function
func (m *CacheOnStorage) GetMulti(ks ...Key) ([]Item, error) {
	results, err := m.cache.GetMulti(ks...)
	if err != nil {
		return nil, err
	}

	indices := make([]int, 0, len(ks))
	uncachedKs := make([]Key, 0, len(ks))
	for i, item := range results {
		if item == nil {
			indices = append(indices, i)
			uncachedKs = append(uncachedKs, ks[i])
		}
	}

	if len(uncachedKs) <= 0 {
		return results, nil
	}

	uncachedItems, err := m.store.GetMulti(uncachedKs...)
	if err != nil {
		return nil, err
	}

	foundItems := make([]Item, 0, len(uncachedItems))
	for _, item := range uncachedItems {
		if item != nil {
			foundItems = append(foundItems, item)
		}
	}

	if _, err := m.cache.SetMulti(foundItems...); err != nil {
		return nil, err
	}

	for i, index := range indices {
		results[index] = uncachedItems[i]
	}

	return results, nil
}

// Set function
func (m *CacheOnStorage) Set(item Item) error {
	if err := m.store.Set(item); err != nil {
		return err
	}

	return m.cache.Set(item)
}

// SetMulti function
func (m *CacheOnStorage) SetMulti(items ...Item) ([]Item, error) {
	if statuses, err := m.store.SetMulti(items...); err != nil {
		return statuses, err
	}

	return m.cache.SetMulti(items...)
}

// SetIfAbsent function
func (m *CacheOnStorage) SetIfAbsent(item Item) error {
	if err := m.store.SetIfAbsent(item); err != nil {
		return err
	}

	return m.cache.SetIfAbsent(item)
}

// SetMultiIfAbsent function
func (m *CacheOnStorage) SetMultiIfAbsent(items ...Item) ([]Item, error) {
	if statuses, err := m.store.SetMultiIfAbsent(items...); err != nil {
		return statuses, err
	}

	return m.cache.SetMultiIfAbsent(items...)
}

// Update function
func (m *CacheOnStorage) Update(item Item) error {
	err := m.store.Update(item)
	if err != nil {
		return err
	}

	return m.cache.Update(item)
}

// UpdateMulti function
func (m *CacheOnStorage) UpdateMulti(items ...Item) ([]Item, error) {
	if statuses, err := m.store.UpdateMulti(items...); err != nil {
		return statuses, err
	}

	return m.cache.UpdateMulti(items...)
}

// Delete function
func (m *CacheOnStorage) Delete(k Key) error {
	if err := m.store.Delete(k); err != nil {
		return err
	}

	return m.cache.Delete(k)
}

// DeleteMulti function
func (m *CacheOnStorage) DeleteMulti(ks ...Key) ([]Key, error) {
	if statuses, err := m.store.DeleteMulti(ks...); err != nil {
		m.cache.DeleteMulti(statuses...)
		return statuses, err
	}

	return m.cache.DeleteMulti(ks...)
}
