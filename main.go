package cloudlayer

import (
	"fmt"

	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("cloudlayer")

type CloudLayer interface {
	// SimpleAuthorize - Authorize this cloud layer with just an API key.
	SimpleAuthorize(apiId, apiKey string) error

	// DetailedAuthorize - Authorize this cloud layer with a dictionary of values.
	DetailedAuthorize(authDetails map[string]string) error

	// CreateInstance - Create a new instance in this cloud layer.
	CreateInstance(details InstanceDetails) (*Instance, error)

	// Remove an instance from this cloudl layer.
	RemoveInstance(instanceId string) (*Operation, error)

	// GetInstance - Get an instance from the layer.
	GetInstance(instanceId string) (*Instance, error)

	// CheckOperationStatus - Check the status of a long running operation.
	CheckOperationStatus(operationId string) (*Operation, error)

	// CreateVolume - Create a new data storage volume.
	CreateVolume(details VolumeDetails) (*Operation, error)

	// RemoveVolume - Remove a data storage volume.
	RemoveVolume(volumeId string) (*Operation, error)

	// CreateSnapshot - Create a volume snapshot
	CreateSnapshot(volumnId string) (*Operation, error)

	// RemoveSnapshot - Remove a volume snapshot
	RemoveSnapshot(volumnId string) (*Operation, error)

	// ListSnapshots - List current snapshots for the current account
	ListSnapshots() ([]SnapshotDetails, error)
}

// NewCloudLayer - Create and return a new cloud layer.
func NewCloudLayer(cloudName string) (CloudLayer, error) {
	switch cloudName {
	case "openstack":
		return &OpenStackLayer{}, nil
	case "aws":
		return &AWSLayer{}, nil
	case "dummy":
		return &DummyLayer{}, nil
	case "docker":
		return &DockerLayer{}, nil
	}
	return nil, fmt.Errorf("Could not find cloud layer: %s", cloudName)
}
