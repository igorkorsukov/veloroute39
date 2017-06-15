package datastore

import (
	"os"
	"path/filepath"
	"testing"

	"context"

	"github.com/stretchr/testify/assert"
	"fmt"
)

type testDSVal struct {
	Val string
}

func TestGcloudDS_Put(t *testing.T) {

	cr, _ := filepath.Abs("veloroute39-78759ec9aa23.json")
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", cr)

	ds := NewGCloudDS()

	err := ds.Open()
	assert.NoError(t, err)

	ctx := context.Background()

	val_src := testDSVal{Val:"val1"}
	err = ds.Put(ctx, "test_bucket", "key1", &val_src)
	assert.NoError(t, err)

	val_dst := &testDSVal{}
	err = ds.Get(ctx, "test_bucket", "key1", val_dst)
	assert.NoError(t, err)
	fmt.Println("val_dst: ", val_dst)
	assert.Equal(t, "val1", val_dst.Val)
}
