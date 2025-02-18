package vault

import (
	"math/rand"
	"time"
)

type RandHelper struct {
	r *rand.Rand
}

func NewRandHelper() *RandHelper {
	return &RandHelper{r: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func (rh *RandHelper) GetRandomNumber(max int) int {
	return rh.r.Intn(max)
}

type BasicVault struct {
	data         IVaultCore
	bucket_count int
	open         bool
	rand         RandHelper
}

func NewBasicVault(count int) *BasicVault {
    return &BasicVault{
		data:         NewBasicVaultCore(count),
        bucket_count: count,
        open:         false,
        rand:         *NewRandHelper(),
    }
}

func NewBasicVaultWithDependencies(data IVaultCore, count int, rand RandHelper) *BasicVault {
    return &BasicVault{
        data:         data,
        bucket_count: count,
        open:         false,
        rand:         rand,
    }
}

func (v *BasicVault) Open() error {
	if v.open {
		return ErrVaultAlreadyOpened
	}
	v.open = true
	return nil
}

func (v *BasicVault) Close() error {
	if !v.open {
		return ErrVaultAlreadyClosed
	}
	v.open = false
	return nil
}

func (v *BasicVault) GetRandomBucket() (int, error) {
	index := v.rand.GetRandomNumber(v.bucket_count)
	return index, nil
}

func (v *BasicVault) Put(key, value string) error {
	if !v.open {
		return ErrVaultNotOpen
	}
	bucket, nil := v.data.GetBucket(v.rand.GetRandomNumber(v.bucket_count))
	bucket.Put(key, value)
	return nil
}

func (v *BasicVault) Get(key string) (string, error) {
	if !v.open {
		return "", ErrVaultNotOpen
	}
	for bucketId := range v.data.Size() {
		bucket, err := v.data.GetBucket(bucketId)
		if err != nil {
			continue
		}
		if value, err := bucket.Get(key); err == nil {
			return value, nil
		}
	}
	return "", ErrKeyNotFound
}

func (v *BasicVault) Delete(key string) error {
	if !v.open {
		return ErrVaultNotOpen
	}
	for bucketId := range v.data.Size() {
		bucket, err := v.data.GetBucket(bucketId)
		if err != nil {
			continue
		}
		if err := bucket.Delete(key); err == nil {
			return nil
		}
	}
	return ErrKeyNotFound
}

func (v *BasicVault) Clear() error {
	if !v.open {
		return ErrVaultNotOpen
	}
	for bucketId := range v.data.Size() {
		bucket, _ := v.data.GetBucket(bucketId)
		bucket.Clear()
	}
	return nil
}

func (v *BasicVault) IsOpen() bool {
	return v.open
}
