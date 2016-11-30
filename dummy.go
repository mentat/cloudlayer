package cloudlayer

// DummyLayer - a sample interface for new layer types.
type DummyLayer struct {
}

// SimpleAuthorize - Authorize this cloud layer with just an API key.
func (dummy DummyLayer) SimpleAuthorize(apiID, apiKey string) error {
	return nil
}

// DetailedAuthorize - Authorize this cloud layer with a dictionary of values.
func (dummy DummyLayer) DetailedAuthorize(authDetails map[string]string) error {
	return nil
}

// CreateInstance - Create a new instance in this cloud layer.
func (dummy DummyLayer) CreateInstance(details InstanceDetails) (*Instance, error) {
	return nil, nil
}

// RemoveInstance - Remove an instance from this cloudl layer.
func (dummy DummyLayer) RemoveInstance(instanceID string) (*Operation, error) {
	return nil, nil
}

// GetInstance - Get an instance from the layer.
func (dummy DummyLayer) GetInstance(instanceID string) (*Instance, error) {
	return nil, nil
}

// CheckOperationStatus - Check the status of a long running operation.
func (dummy DummyLayer) CheckOperationStatus(operationID string) (*Operation, error) {
	return nil, nil
}

// CreateVolume - Create a new data storage volume.
func (dummy DummyLayer) CreateVolume(details VolumeDetails) (*Operation, error) {
	return nil, nil
}

// RemoveVolume - Remove a data storage volume.
func (dummy DummyLayer) RemoveVolume(volumeID string) (*Operation, error) {
	return nil, nil
}

// CreateSnapshot - Create a volume snapshot
func (dummy DummyLayer) CreateSnapshot(volumnID string) (*Operation, error) {
	return nil, nil
}

// RemoveSnapshot - Remove a volume snapshot
func (dummy DummyLayer) RemoveSnapshot(volumnID string) (*Operation, error) {
	return nil, nil
}

// ListSnapshots - List current snapshots for the current account
func (dummy DummyLayer) ListSnapshots() ([]SnapshotDetails, error) {
	return nil, nil
}
