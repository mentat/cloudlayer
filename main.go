package cloudlayer

type CloudLayer interface {
	// Authorize this cloud layer with just an API key.
	SimpleAuthorize(apiKey string) error

	// Authorize this cloud layer with a dictionary of values.
	DetailedAuthorize(authDetails map[string]string) error

	// Create a new instance in this cloud layer.
	CreateInstance(details InstanceDetails) (*Instance, error)
	// Remove an instance from this cloudl layer.
	RemoveInstance(instanceId string) (*Operation, error)

	// Check the status of a long running operation.
	CheckOperationStatus(operationId string) (*Operation, error)

	// Create a new data storage volume.
	CreateVolume(details VolumeDetails) (*Operation, error)
	// Remove a data storage volume.
	RemoveVolume(volumeId string) (*Operation, error)

	// Create a volume snapshot
	CreateSnapshot(volumnId string) (*Operation, error)
	// Remove a volume snapshot
	RemoveSnapshot(volumnId string) (*Operation, error)
	// List current snapshots for the current account
	ListSnapshots() ([]SnapshotDetails, error)
}

func GetCloudLayer(cloudName string) (CloudLayer, error) {
	switch cloudName {
	case "openstack":
		return &OpenStackLayer{}, nil
	}
	return nil, nil
}
