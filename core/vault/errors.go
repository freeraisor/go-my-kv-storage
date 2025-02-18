package vault

import "errors"

var (
	ErrBucketNotFound  = errors.New("bucket not found")
	ErrBucketExists    = errors.New("bucket already exists")
	ErrBucketNameEmpty = errors.New("bucket name cannot be empty")

	ErrKeyNotFound  = errors.New("key not found")
	ErrKeyExists    = errors.New("key already exists")
	ErrInvalidKey   = errors.New("invalid key")
	ErrInvalidValue = errors.New("invalid value")

	ErrVaultNotOpen       = errors.New("vault is not open")
	ErrVaultCorrupted     = errors.New("vault is corrupted")
	ErrVaultAlreadyOpened = errors.New("vault is already open")
	ErrVaultAlreadyClosed = errors.New("vault is already closed")
)
