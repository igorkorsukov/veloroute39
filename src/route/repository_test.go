package route

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository_List(t *testing.T) {

	r := NewRepository(nil)

	r.Add(Route{UID: "1", Name: "r1"})
	r.Add(Route{UID: "2", Name: "r2"})
	r.Add(Route{UID: "3", Name: "r3"})

	rts, err := r.List()

	assert.NoError(t, err)
	assert.Equal(t, 3, len(rts))
	assert.True(t, indexOfRoute(rts, "1") > -1)
	assert.True(t, indexOfRoute(rts, "2") > -1)
	assert.True(t, indexOfRoute(rts, "3") > -1)
}

func indexOfRoute(rts []Route, uid string) int {
	for i, r := range rts {
		if r.UID == uid {
			return i
		}
	}
	return -1
}
