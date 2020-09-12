package caching

// Proxy
type Proxy interface {
	NewFromZinto(msg interface{}) Item
	NewFromZintos(msg interface{}) []Item
}

// Key A generic value can store in the cache.
type Key interface {
	UID() string
	TableName() string
	Idx() string
	Maker() Proxy
}

// Item
type Item interface {
	UID() string
	Idx() string
	ElasticIndex() string
	MarshalToJson() ([]byte, error)
	TableName() string
}

// Storage A generic key value storage interface.  The storage may be persistent
// (e.g., a database) or volatile (e.g., cache).  All Storage implementations
// must be thread safe.
type Storage interface {
	// This retrieves a single value from the storage.
	Get(key Key) (Item, error)

	// This retrieves multiple values from the storage.  The items are returned
	// in the same order as the input keys.
	GetMulti(keys ...Key) ([]Item, error)

	// This stores a single item into the storage.
	// Generally, all "Set" function map to "Update" except "SetIfAbsent" for Create
	SetIfAbsent(item Item) error
	// This stores multiple items into the storage with key->value pair.
	SetMultiIfAbsent(items ...Item) ([]Item, error)

	// Generally, all "Set" function map to "Update" or "Create"
	Set(item Item) error
	// This stores multiple items into the storage with key->value pair.
	SetMulti(items ...Item) ([]Item, error)

	// Generally, all "Update" function map to "Update"
	Update(item Item) error
	// This stores multiple items into the storage with key->value pair.
	UpdateMulti(items ...Item) ([]Item, error)

	// This removes a single item from the storage.
	Delete(key Key) error
	// This removes multiple items from the storage.
	DeleteMulti(keys ...Key) ([]Key, error)
}
