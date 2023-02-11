package persistence

import (
	"context"
	"errors"
)

var ErrKeyNotFound = errors.New("key was not found")

type Entry struct {
	Key   string
	Value []byte
}

type Persistence interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Exists(ctx context.Context, key string) (bool, error)
	Set(ctx context.Context, key string, value []byte) error
	Delete(ctx context.Context, key string) error
	List(ctx context.Context, prefix string) ([]Entry, error)
	IsLeader(ctx context.Context) (bool, error)
}
