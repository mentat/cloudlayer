package cloudlayer

import "time"

type VolumeDetails struct {
	VolumeSize     int    // The size of the volume in gigabytes
	Zone           string // Availability zode
	Region         string // The region or DC for this instance.
	OriginSnapshot string
}

type SnapshotDetails struct {
	ID           string
	SnapshotSize int
	CreatedAt    *time.Time
}

type InstanceDetails struct {
	MemorySize   int    // Memory size in gigabytes
	CPUCores     int    // The number of virtual CPU codes
	InstanceType string // A instance-type short name
	BaseImage    string // An AMI or Docker Image
	Zone         string // Availability zode
	Region       string // The region or DC for this instance.

	DiskSize       int    // The size of the
	DiskMountPoint string // Where the disk is to mount on the instance
	DiskVolume     string // The EBS, Cinder, etc disk volume identifier
}

type Error struct {
	Code        string
	NumericCode int
	Description string
}

type Operation struct {
	ID            string
	Name          string // The descriptive name of this operation.
	Status        string // Options are PENDING, RUNNING, or DONE
	StatusMessage string

	StartTime  *time.Time
	EndTime    *time.Time
	IsComplete bool
	IsError    bool
	Errors     []Error
}

// An instance is an active VM/container on a cloud provider.
type Instance struct {
	ID               string
	Details          InstanceDetails
	CurrentOperation Operation
	Status           string
}

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
	// List current snapshots
	ListSnapshots() ([]SnapshotDetails, error)
}

func GetCloudLayer(cloudName string) (CloudLayer, error) {
	switch cloudName {
	case "openstack":
		return &OpenStackLayer
	}
	return nil, nil
}
