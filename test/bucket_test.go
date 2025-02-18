package main

import (
    "testing"
    "go-my-kv-storage/core/vault"
)

func TestBucketPutAndGet(t *testing.T) {
    bucket := vault.NewBucket()

    key := "testKey"
    value := "testValue"

    err := bucket.Put(key, value)
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

    got, err := bucket.Get(key)
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

    if got != value {
        t.Fatalf("expected %v, got %v", value, got)
    }
}

func TestBucketDelete(t *testing.T) {
    bucket := vault.NewBucket()

    key := "testKey"
    value := "testValue"

    err := bucket.Put(key, value)
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

    err = bucket.Delete(key)
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

    _, err = bucket.Get(key)
    if err == nil {
        t.Fatalf("expected error, got nil")
    }
}

func TestBucketClear(t *testing.T) {
	bucket := vault.NewBucket()

    key1 := "testKey1"
    value1 := "testValue1"
    key2 := "testKey2"
    value2 := "testValue2"

    bucket.Put(key1, value1)
    bucket.Put(key2, value2)

    err := bucket.Clear()
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

    _, err = bucket.Get(key1)
    if err == nil {
        t.Fatalf("expected error, got nil")
    }

    _, err = bucket.Get(key2)
    if err == nil {
        t.Fatalf("expected error, got nil")
    }
}