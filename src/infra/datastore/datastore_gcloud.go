package datastore

import (
	"context"
	"errors"

	"encoding/json"

	"reflect"

	"cloud.google.com/go/datastore"
)

type gcloudDS struct {
	DB *datastore.Client
}

type gcloudVal struct {
	Val string
}

func NewGCloudDS() DataStore {
	return &gcloudDS{}
}

func (ds *gcloudDS) Open() error {

	ctx := context.Background()
	db, err := datastore.NewClient(ctx, "veloroute39")
	if err != nil {
		return err
	}
	ds.DB = db
	return nil
}

func (ds *gcloudDS) Close() error {
	return nil
}

func (ds *gcloudDS) Put(ctx context.Context, bucket, uid string, src interface{}) error {

	data, err := json.Marshal(src)
	if err != nil {
		return err
	}

	k := datastore.NameKey(bucket, uid, nil)
	_, err = ds.DB.Put(ctx, k, &gcloudVal{Val: string(data)})
	if err != nil {
		return err
	}
	return nil
}

func (ds *gcloudDS) Get(ctx context.Context, bucket, uid string, dst interface{}) error {

	val := &gcloudVal{}
	k := datastore.NameKey(bucket, uid, nil)
	err := ds.DB.Get(ctx, k, val)
	if err != nil {
		return err
	}

	return json.Unmarshal([]byte(val.Val), dst)
}

func (ds *gcloudDS) GetAll(ctx context.Context, bucket string, dst interface{}) error {

	rv := reflect.ValueOf(dst)
	if rv.Kind() != reflect.Ptr {
		return errors.New("dst must be pointer")
	}

	rv = rv.Elem()
	if rv.Kind() != reflect.Slice {
		return errors.New("dst must be slice")
	}

	var vals []*gcloudVal
	_, err := ds.DB.GetAll(ctx, datastore.NewQuery(bucket), &vals)
	if err != nil {
		return err
	}

	for _, v := range vals {

		ptr := reflect.New(rv.Type().Elem()).Interface()
		err := json.Unmarshal([]byte(v.Val), ptr)
		if err != nil {
			return err
		}

		pv := reflect.ValueOf(ptr).Elem()
		rv.Set(reflect.Append(rv, pv))
	}
	return nil
}

func (ds *gcloudDS) Delete(ctx context.Context, bucket, uid string) error {

	k := datastore.NameKey(bucket, uid, nil)
	return ds.DB.Delete(ctx, k)
}
