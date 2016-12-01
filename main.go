package cloudlayer

import (
	"fmt"

	"github.com/juju/loggo"
)

var logger = loggo.GetLogger("cloudlayer")

type CloudLayer interface {
	// SimpleAuthorize - Authorize this cloud layer with just an API key.
	SimpleAuthorize(apiID, apiKey string) error

	// DetailedAuthorize - Authorize this cloud layer with a dictionary of values.
	DetailedAuthorize(authDetails map[string]string) error

	// CreateInstance - Create a new instance in this cloud layer.
	CreateInstance(details InstanceDetails) (*Instance, error)

	// Remove an instance from this cloudl layer.
	RemoveInstance(instanceID string) (*Operation, error)

	// GetInstance - Get an instance from the layer.
	GetInstance(instanceID string) (*Instance, error)

	// CheckOperationStatus - Check the status of a long running operation.
	CheckOperationStatus(operationID string) (*Operation, error)

	// CreateVolume - Create a new data storage volume.
	CreateVolume(details VolumeDetails) (*Operation, error)

	// RemoveVolume - Remove a data storage volume.
	RemoveVolume(volumeID string) (*Operation, error)

	// CreateSnapshot - Create a volume snapshot
	CreateSnapshot(volumnID string) (*Operation, error)

	// RemoveSnapshot - Remove a volume snapshot
	RemoveSnapshot(volumnID string) (*Operation, error)

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
		return NewDockerLayer(), nil
	}
	return nil, fmt.Errorf("Could not find cloud layer: %s", cloudName)
}
