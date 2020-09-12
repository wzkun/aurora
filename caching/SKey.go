package caching

import (
	"bytes"
)

// SKey objects
type SKey struct {
	id        string
	tableName string

	proxy Proxy
}

//NewKey Function
func NewKey(proxy Proxy, tableName, id string) Key {
	p := new(SKey)
	p.id = id
	p.tableName = tableName
	p.proxy = proxy

	return p
}

//NewKeys Function
func NewKeys(proxy Proxy, tableName string, ids ...string) []Key {
	ks := make([]Key, len(ids))
	for i, id := range ids {
		k := NewKey(proxy, tableName, id)
		ks[i] = k
	}
	return ks
}

// UID return the table name
func (k *SKey) UID() string {
	b := bytes.Buffer{}
	b.WriteString(k.tableName)
	b.WriteString(k.id)
	return b.String()
}

// Idx return the table name
func (k *SKey) Idx() string {
	return k.id
}

// TableName return the table name
func (k *SKey) TableName() string {
	return k.tableName
}

// Maker return the table name
func (k *SKey) Maker() Proxy {
	return k.proxy
}

// KeysToStrings Function
func KeysToStrings(ks []Key) []string {
	items := make([]string, 0, len(ks))
	for _, k := range ks {
		if k == nil {
			continue
		}
		items = append(items, k.Idx())
	}
	return items
}
