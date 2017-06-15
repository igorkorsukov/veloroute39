package datastore

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type memDS struct {
	data map[string]map[string][]byte
}

func NewMemDS() DataStore {
	s := &memDS{}
	s.data = make(map[string]map[string][]byte)
	return s
}

func (s *memDS) Open() error {
	return nil
}

func (s *memDS) Close() error {
	return nil
}

func (s *memDS) Put(ctx context.Context, bucket, uid string, src interface{}) error {

	b, ok := s.data[bucket]
	if !ok {
		b = make(map[string][]byte)
	}

	val, err := json.Marshal(src)
	if err != nil {
		return err
	}

	b[uid] = val
	s.data[bucket] = b

	return nil
}

func (s *memDS) Get(ctx context.Context, bucket, uid string, dst interface{}) error {

	b, ok := s.data[bucket]
	if !ok {
		return fmt.Errorf(DSErrNotFoundBucket, bucket)
	}

	val, ok := b[uid]
	if !ok {
		return fmt.Errorf(DSErrNotFoundValue, uid)
	}

	return json.Unmarshal(val, dst)
}

func (s *memDS) GetAll(ctx context.Context, bucket string, dst interface{}) error {

	buck, ok := s.data[bucket]
	if !ok {
		return fmt.Errorf(DSErrNotFoundBucket, bucket)
	}

	v := reflect.ValueOf(dst)
	if v.Kind() != reflect.Ptr {
		return errors.New("dst must be pointer")
	}

	v = v.Elem()
	if v.Kind() != reflect.Slice {
		return errors.New("dst must be slice")
	}

	for _, data := range buck {

		ptr := reflect.New(v.Type().Elem()).Interface()
		err := json.Unmarshal(data, ptr)
		if err != nil {
			return err
		}

		pv := reflect.ValueOf(ptr).Elem()
		v.Set(reflect.Append(v, pv))
	}

	return nil
}

func (s *memDS) Keys(ctx context.Context, bucket string) ([]string, error) {

	b, ok := s.data[bucket]
	if !ok {
		return nil, fmt.Errorf(DSErrNotFoundBucket, bucket)
	}

	var keys []string
	for k := range b {
		keys = append(keys, k)
	}

	return keys, nil
}

func (s *memDS) Delete(ctx context.Context, bucket, uid string) error {

	b, ok := s.data[bucket]
	if !ok {
		return nil
	}

	delete(b, uid)

	return nil
}
