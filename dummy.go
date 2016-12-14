package cloudlayer

import "math/rand"

// DummyLayer - a sample interface for new layer types.
type DummyLayer struct {
	instances map[string]*Instance
}

// NewDummyLayer - Create a new Dummy Layer
func NewDummyLayer() *DummyLayer {
	layer := &DummyLayer{
		instances: make(map[string]*Instance),
	}
	return layer
}

func randStringBytesRmndr(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
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
	inst := &Instance{
		Details:          details,
		CurrentOperation: Operation{},
		Status:           "RUNNING",
		ID:               randStringBytesRmndr(10),
	}
	dummy.instances[inst.ID] = inst
	return inst, nil
}

// RemoveInstance - Remove an instance from this cloudl layer.
func (dummy DummyLayer) RemoveInstance(instanceID string) (*Operation, error) {
	return nil, nil
}

// GetInstance - Get an instance from the layer.
func (dummy DummyLayer) GetInstance(instanceID string) (*Instance, error) {
	return nil, nil
}

// ListInstances - List the instances in this layer.
func (dummy DummyLayer) ListInstances() ([]*Instance, error) {
	ret := make([]*Instance, 0, len(dummy.instances))
	for _, v := range dummy.instances {
		ret = append(ret, v)
	}
	return ret, nil
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
