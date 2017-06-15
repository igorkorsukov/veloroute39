package datastore

import (
	"context"
)

var (
	DSErrNotFoundBucket string = "not found bucket: %s"
	DSErrNotFoundValue  string = "not found value: %s"
)

type DataStore interface {
	Open() error
	Close() error

	Put(ctx context.Context, bucket, uid string, src interface{}) error
	Get(ctx context.Context, bucket, uid string, dst interface{}) error
	GetAll(ctx context.Context, bucket string, dst interface{}) error
	Delete(ctx context.Context, bucket, uid string) error
}
