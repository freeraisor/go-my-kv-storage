module go-my-kv-storage/test

go 1.24.0

require go-my-kv-storage/core/vault v0.0.0

replace go-my-kv-storage/core/vault => ../core/vault

require (
	github.com/golang/mock v1.6.0
	go-my-kv-storage/mocks v0.0.0
)

replace go-my-kv-storage/mocks => ./mocks
