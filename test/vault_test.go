package main

import (
	"go-my-kv-storage/core/vault"
	"go-my-kv-storage/mocks"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestBasicVault_Open(t *testing.T) {
	vault := vault.NewBasicVault(3)
	err := vault.Open()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !vault.IsOpen() {
		t.Fatalf("expected vault to be open")
	}
}

func TestBasicVault_Close(t *testing.T) {
	vault := vault.NewBasicVault(3)
	vault.Open()
	err := vault.Close()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if vault.IsOpen() {
		t.Fatalf("expected vault to be closed")
	}
}

func TestBasicVault_PutAndGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockVaultCore := mocks.NewMockIVaultCore(ctrl)
	mockBucket := mocks.NewMockIBucket(ctrl)

	bucketCount := 3
	v := vault.NewBasicVaultWithDependencies(mockVaultCore, bucketCount, *vault.NewRandHelper())
	if err := v.Open(); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	key := "testKey"
	value := "testValue"

	mockVaultCore.EXPECT().GetBucket(gomock.Any()).Return(mockBucket, nil).AnyTimes()
	mockVaultCore.EXPECT().Size().Return(bucketCount).AnyTimes()
	mockBucket.EXPECT().Put(key, value).Return(nil)
	mockBucket.EXPECT().Get(key).Return(value, nil)

	err := v.Put(key, value)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	got, err := v.Get(key)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if got != value {
		t.Fatalf("expected %v, got %v", value, got)
	}
}

func TestBasicVault_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockVaultCore := mocks.NewMockIVaultCore(ctrl)
	mockBucket := mocks.NewMockIBucket(ctrl)

	bucketCount := 3
	v := vault.NewBasicVaultWithDependencies(mockVaultCore, bucketCount, *vault.NewRandHelper())
	if err := v.Open(); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	key := "testKey"
	value := "testValue"

	mockVaultCore.EXPECT().GetBucket(gomock.Any()).Return(mockBucket, nil).AnyTimes()
	mockVaultCore.EXPECT().Size().Return(bucketCount).AnyTimes()
	mockBucket.EXPECT().Put(key, value).Return(nil)
	mockBucket.EXPECT().Delete(key).Return(nil)
	mockBucket.EXPECT().Get(key).Return("", vault.ErrKeyNotFound).MinTimes(1).MaxTimes(bucketCount)

	v.Put(key, value)
	err := v.Delete(key)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = v.Get(key)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestBasicVault_Clear(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockVaultCore := mocks.NewMockIVaultCore(ctrl)
	mockBucket := mocks.NewMockIBucket(ctrl)

	bucketCount := 3
	v := vault.NewBasicVaultWithDependencies(mockVaultCore, bucketCount, *vault.NewRandHelper())
	if err := v.Open(); err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	key := "testKey"
	value := "testValue"

	mockVaultCore.EXPECT().GetBucket(gomock.Any()).Return(mockBucket, nil).AnyTimes()
	mockVaultCore.EXPECT().Size().Return(bucketCount).AnyTimes()
	mockBucket.EXPECT().Put(key, value).Return(nil)
	mockBucket.EXPECT().Clear().Return(nil).MinTimes(1).MaxTimes(bucketCount)
	mockBucket.EXPECT().Get(key).Return("", vault.ErrKeyNotFound).MinTimes(1).MaxTimes(bucketCount)

	v.Put(key, value)
	err := v.Clear()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	_, err = v.Get(key)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}
