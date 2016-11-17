package cloudlayer

import "time"

// VolumeDetails -
type VolumeDetails struct {
	ID             string // The cloud layer specific ID for this volume.
	VolumeSize     int    // The size of the volume in gigabytes
	Zone           string // Availability zode
	Region         string // The region or DC for this instance.
	OriginSnapshot string // The ID of the snapshot to create this volume from.
	MountPoint     string // If the volume is mounted, where is it mounted.
}

// SnapshotDetails -
type SnapshotDetails struct {
	ID           string     // The cloud layer specific ID for this snapshot.
	SnapshotSize int        // The size of this snapshot in gigabytes.
	CreatedAt    *time.Time // When this snapshot was created.
}

// PortDetails -
type PortDetails struct {
	InstancePort int    // The port number to open on the instance, i.e. 22, 80, 443
	HostPort     int    // Opional: the port number to forward from the host (used for containers)
	Protocol     string // The protocol, i.e. tcp, udp, icmp
}

// NetworkDetails - Either ID or Port is required for openstack
type NetworkDetails struct {
	ID      string // UUID of the network (TODO(tvoran): support names here too)
	Port    string // Specify a port, get the network associated with it
	FixedIP string // Optional: specific IP address for this network
}

// InstanceDetails -
type InstanceDetails struct {
	Hostname        string
	MemorySize      int    // Memory size in gigabytes
	CPUCores        int    // The number of virtual CPU codes
	InstanceType    string // A instance-type short name or flavor
	BaseImage       string // An AMI or Docker Image to boot from
	Zone            string // Availability zode
	Region          string // The region or DC for this instance.
	Volumes         []VolumeDetails
	LaunchTime      *time.Time
	Variables       map[string]string // Environmental variables
	PublicIP        string
	PrivateIP       string
	SubnetID        string        // The ID of the subnet this instance is in.
	VPCID           string        // The virtual private cloud ID, if any
	ExposedPorts    []PortDetails // Ports to expose on the instance, in the format: <portNum>/<protocol>, i.e. 22/tcp. 53/udp
	LinkedInstances []string
	Networks        []NetworkDetails // Insert NetworkDetails here
}

// Error -
type Error struct {
	Code        string // The string identifier for this error.
	NumericCode int    // A numeric identifier for this error.
	Description string // A description of the error.
}

// Operation -
type Operation struct {
	ID            string
	Name          string // The descriptive name of this operation.
	Status        string // Options are PENDING, RUNNING, or DONE
	StatusMessage string
	StartTime     *time.Time // When this operation started
	EndTime       *time.Time // When this operation ended
	IsComplete    bool       // Is the operation complete.
	IsError       bool       // Was there an error with this operation.
	Errors        []Error    // List of all errors.
}

// Instance - An instance is an active VM/container on a cloud provider.
type Instance struct {
	ID               string          // Unique identifier for this instance on layer.
	Details          InstanceDetails // The details of the instance.
	CurrentOperation Operation       // The current operation.
	Status           string          // Options are PENDING, RUNNING, or STOPPED
}
