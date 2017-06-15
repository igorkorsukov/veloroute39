package datastore

import (
	"context"
	"errors"

	"encoding/json"

	"fmt"

	"reflect"

	"github.com/boltdb/bolt"
)

type boltDS struct {
	DB *bolt.DB
}

func NewBoltDS() DataStore {
	return &boltDS{}
}

func (ds *boltDS) Open() error {

	db, err := bolt.Open("data.db", 0600, nil)
	if err != nil {
		return err
	}
	ds.DB = db
	return nil
}

func (ds *boltDS) Close() error {
	return ds.DB.Close()
}

func (ds *boltDS) Put(ctx context.Context, bucket, uid string, src interface{}) error {

	return ds.DB.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return err
		}

		data, err := json.Marshal(src)
		if err != nil {
			return err
		}

		return b.Put([]byte(uid), data)
	})
}

func (ds *boltDS) Get(ctx context.Context, bucket, uid string, dst interface{}) error {

	return ds.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf(DSErrNotFoundBucket, bucket)
		}

		data := b.Get([]byte(uid))
		if data == nil {
			return fmt.Errorf(DSErrNotFoundValue, uid)
		}

		return json.Unmarshal(data, dst)
	})
}

func (ds *boltDS) GetAll(ctx context.Context, bucket string, dst interface{}) error {

	rv := reflect.ValueOf(dst)
	if rv.Kind() != reflect.Ptr {
		return errors.New("dst must be pointer")
	}

	rv = rv.Elem()
	if rv.Kind() != reflect.Slice {
		return errors.New("dst must be slice")
	}

	return ds.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf(DSErrNotFoundBucket, bucket)
		}

		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {

			ptr := reflect.New(rv.Type().Elem()).Interface()
			err := json.Unmarshal(v, ptr)
			if err != nil {
				return err
			}

			pv := reflect.ValueOf(ptr).Elem()
			rv.Set(reflect.Append(rv, pv))
		}

		return nil
	})
}

func (ds *boltDS) Delete(ctx context.Context, bucket, uid string) error {

	ds.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if b == nil {
			return fmt.Errorf(DSErrNotFoundBucket, bucket)
		}

		return b.Delete([]byte(uid))
	})
	return nil
}
