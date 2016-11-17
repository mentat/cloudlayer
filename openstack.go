package cloudlayer

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
)

type OpenStackLayer struct {
	client *gophercloud.ProviderClient
}

func (this *OpenStackLayer) SimpleAuthorize(apiKey string) error {
	return fmt.Errorf("Cannot simple authorize on OpenStack")
}

func (this *OpenStackLayer) DetailedAuthorize(authDetails map[string]string) error {
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: authDetails["identityEndpoint"],
		Username:         authDetails["userName"],
		Password:         authDetails["password"],
		TenantID:         authDetails["tenantId"],
	}

	provider, err := openstack.AuthenticatedClient(opts)

	if err != nil {
		return err
	}

	this.client = provider

	return nil
}

func (this *OpenStackLayer) CreateInstance(details InstanceDetails) (*Instance, error) {
	return nil, nil
}

func (this *OpenStackLayer) RemoveInstance(instanceId string) (*Operation, error) {
	return nil, nil
}

func (this *OpenStackLayer) CheckOperationStatus(operationId string) (*Operation, error) {
	return nil, nil
}

func (this *OpenStackLayer) CreateVolume(details VolumeDetails) (*Operation, error) {
	return nil, nil
}

func (this *OpenStackLayer) RemoveVolume(volumeId string) (*Operation, error) {
	return nil, nil
}

// Create a volume snapshot
func (this *OpenStackLayer) CreateSnapshot(volumnId string) (*Operation, error) {
	return nil, nil
}

// Remove a volume snapshot
func (this *OpenStackLayer) RemoveSnapshot(volumnId string) (*Operation, error) {
	return nil, nil
}

// List current snapshots for the current account
func (this *OpenStackLayer) ListSnapshots() ([]SnapshotDetails, error) {
	return nil, nil
}
