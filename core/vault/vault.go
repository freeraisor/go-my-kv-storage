package vault

type IBucket interface {
	Put(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
	Clear() error
	Size() int
}

type IVaultCore interface {
	GetBucket(id int) (IBucket, error)
	DeleteBucket(id int) error
	Size() int
}

type IVault interface {
	IVaultCore
	Open() error
	Close() error
	Put(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
	Clear() error
	IsOpen() bool
}
