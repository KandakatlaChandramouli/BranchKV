package objectstore

import (
	"sync"
)

type Object struct {
	Key  string
	Data []byte
}

type ObjectStore struct {
	objects map[string]Object
	mu      sync.RWMutex
}

func NewObjectStore() *ObjectStore {

	return &ObjectStore{
		objects: make(
			map[string]Object,
		),
	}
}

func (o *ObjectStore) Put(
	key string,
	data []byte,
) {

	o.mu.Lock()
	defer o.mu.Unlock()

	o.objects[key] = Object{
		Key:  key,
		Data: data,
	}
}

func (o *ObjectStore) Get(
	key string,
) ([]byte, bool) {

	o.mu.RLock()
	defer o.mu.RUnlock()

	v, ok := o.objects[key]

	return v.Data, ok
}
