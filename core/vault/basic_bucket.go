package vault

type BasicBucket struct {
	storage map[string]string
}

func NewBucket() *BasicBucket {
	return &BasicBucket{storage: make(map[string]string)}
}

func (b *BasicBucket) Put(key, value string) error {
	b.storage[key] = value
	return nil
}

func (b *BasicBucket) Get(key string) (string, error) {
	val, ok := b.storage[key]
	if !ok {
		return "", ErrKeyNotFound
	}
	return val, nil
}

func (b *BasicBucket) Delete(key string) error {
	delete(b.storage, key)
	return nil
}

func (b *BasicBucket) Clear() error {
	b.storage = make(map[string]string)
	return nil
}

func (b *BasicBucket) Size() int {
	return len(b.storage)
}

func (b *BasicBucket) MergeBucket(bucket *BasicBucket) error {
	for key, value := range bucket.storage {
		b.storage[key] = value
	}
	return nil
}
