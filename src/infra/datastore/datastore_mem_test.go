package datastore

import (
	"context"
	"fmt"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemDS_GetAll(t *testing.T) {

	ctx := context.Background()

	ds := &memDS{}
	ds.data = make(map[string]map[string][]byte)

	ds.Put(ctx, "test_bucket", "key1", &testDSVal{Val: "val1"})
	ds.Put(ctx, "test_bucket", "key2", &testDSVal{Val: "val2"})
	ds.Put(ctx, "test_bucket", "key3", &testDSVal{Val: "val3"})

	var vals []testDSVal
	err := ds.GetAll(ctx, "test_bucket", &vals)
	assert.NoError(t, err)
	fmt.Println("vals: ", vals)
	assert.Equal(t, 3, len(vals))
	//	assert.Equal(t, "val1", vals[0].Val)
}
