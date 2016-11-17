package cloudlayer

import "time"

type VolumeDetails struct {
	VolumeSize     int    // The size of the volume in gigabytes
	Zone           string // Availability zode
	Region         string // The region or DC for this instance.
	OriginSnapshot string // The ID of the snapshot to create this volume from.
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
