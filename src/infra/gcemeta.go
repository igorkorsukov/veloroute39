package infra

import (
	"cloud.google.com/go/compute/metadata"
)

func IsGCloud() bool {
	id, _ := metadata.InstanceID()
	return id != ""
}
