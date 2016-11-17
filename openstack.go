package cloudlayer

import (
	"fmt"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
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
