package caching

import (
	"code.aliyun.com/bim_backend/zoogoer/gun/container/concurrent/lru"
)

// LocalLRUStorage struct
type LocalLRUStorage struct {
	cs *lru.Cache
}

var instance *LocalLRUStorage

func init() {
	instance = &LocalLRUStorage{}
	cs, _ := lru.NewCache(10000)

	instance.cs = cs
}

func GetInstance() *LocalLRUStorage {
	return instance
}

// NewLocalLRUStorage create LocalLRUStorage
func NewLocalLRUStorage(capacity int64) (*LocalLRUStorage, error) {
	cs, err := lru.NewCache(capacity)
	if err != nil {
		return nil, err
	}

	o := new(LocalLRUStorage)
	o.cs = cs
	return o, nil
}

// NewLRUOnStorage create LocalLRUStorage
func NewLRUOnStorage(st Storage) (Storage, error) {
	lru := GetInstance()

	store := NewCacheOnStorage(lru, st)
	return store, nil
}

// Get function
func (m *LocalLRUStorage) Get(k Key) (Item, error) {
	v, hit := m.cs.Get(k.UID())

	if hit {
		return v.(Item), nil
	}

	return nil, nil
}

// GetMulti function
func (m *LocalLRUStorage) GetMulti(ks ...Key) ([]Item, error) {
	results := make([]Item, len(ks))
	for i, k := range ks {
		item, _ := m.Get(k)
		results[i] = item
	}
	return results, nil
}

// Set function
func (m *LocalLRUStorage) Set(item Item) error {
	m.cs.Set(item.UID(), item)
	return nil
}

// SetMulti function
func (m *LocalLRUStorage) SetMulti(items ...Item) ([]Item, error) {
	for _, item := range items {
		m.Set(item)
	}
	return items, nil
}

// SetIfAbsent function
func (m *LocalLRUStorage) SetIfAbsent(item Item) error {
	uid := item.UID()
	if _, hit := m.cs.Get(uid); !hit {
		m.cs.Set(uid, item)
	}
	return nil
}

// SetMultiIfAbsent function
func (m *LocalLRUStorage) SetMultiIfAbsent(items ...Item) ([]Item, error) {
	for _, item := range items {
		m.SetIfAbsent(item)
	}
	return items, nil
}

// Update function
func (m *LocalLRUStorage) Update(item Item) error {
	m.cs.Set(item.UID(), item)
	return nil
}

// UpdateMulti function
func (m *LocalLRUStorage) UpdateMulti(items ...Item) ([]Item, error) {
	for _, item := range items {
		m.Update(item)
	}
	return items, nil
}

// Delete function
func (m *LocalLRUStorage) Delete(k Key) error {
	m.cs.Delete(k.UID())
	return nil
}

// DeleteMulti function
func (m *LocalLRUStorage) DeleteMulti(ks ...Key) ([]Key, error) {
	for _, k := range ks {
		m.Delete(k)
	}
	return ks, nil
}
