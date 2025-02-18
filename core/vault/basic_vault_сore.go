package vault

type BasicVaultCore struct {
	buckets []BasicBucket
}

func NewBasicVaultCore(count int) *BasicVaultCore {
	return &BasicVaultCore{buckets: make([]BasicBucket, count)}
}

func (v *BasicVaultCore) GetBucket(id int) (IBucket, error) {
	if id < 0 || id >= len(v.buckets) {
		return nil, ErrBucketNotFound
	}
	return &v.buckets[id], nil
}

func (v *BasicVaultCore) DeleteBucket(id int) error {
	if id < 0 || id >= len(v.buckets) {
		return ErrBucketNotFound
	}
	v.buckets = append(v.buckets[:id], v.buckets[id+1:]...)
	return nil
}

func (v *BasicVaultCore) Size() int {
	return len(v.buckets)
}
