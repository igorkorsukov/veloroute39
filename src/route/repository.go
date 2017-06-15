package route

import (
	"context"
	"infra/datastore"
)

type RouteRepository interface {
	List() ([]Route, error)
	Find(uid string) (Route, error)
	Add(r Route) error
	Delete(uid string) error
}

type repository struct {
	ds     datastore.DataStore
	bucket string
}

func NewRepository(ds datastore.DataStore) RouteRepository {
	r := &repository{}
	r.ds = ds
	if r.ds == nil {
		r.ds = datastore.NewMemDS()
	}
	r.bucket = "Route"
	return r
}

func (r *repository) List() ([]Route, error) {

	ctx := context.Background()

	var rts []Route
	err := r.ds.GetAll(ctx, r.bucket, &rts)
	if err != nil {
		return nil, err
	}

	return rts, nil
}

func (r *repository) Find(uid string) (Route, error) {
	rt := Route{}
	err := r.ds.Get(context.Background(), r.bucket, uid, &rt)
	return rt, err
}

func (r *repository) Add(rt Route) error {
	return r.ds.Put(context.Background(), r.bucket, rt.UID, &rt)
}

func (r *repository) Delete(uid string) error {
	return r.ds.Delete(context.Background(), r.bucket, uid)
}
