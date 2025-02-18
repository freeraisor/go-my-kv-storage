package main

import (
	"testing"
	"go-my-kv-storage/core/vault"
)

func TestBasicVaultCore_GetBucket(t *testing.T) {
	v := vault.NewBasicVaultCore(8)
	bucketID := 3

	bucket, err := v.GetBucket(bucketID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if bucket == nil {
		t.Fatalf("expected a bucket, got nil")
	}
}

func TestBasicVaultCore_DeleteBucket(t *testing.T) {
	v := vault.NewBasicVaultCore(8)
	bucketID := 3

	err := v.DeleteBucket(bucketID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if v.Size() != 7 {
		t.Fatalf("expected size 7, got %v", v.Size())
	}
}

func TestBasicVaultCore_DeleteBucket_InvalidID(t *testing.T) {
	v := vault.NewBasicVaultCore(8)
	invalidID := 10

	err := v.DeleteBucket(invalidID)
	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
}

func TestBasicVaultCore_Size(t *testing.T) {
	v := vault.NewBasicVaultCore(8)

	if v.Size() != 8 {
		t.Fatalf("expected size 8, got %v", v.Size())
	}
}