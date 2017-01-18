package cloudlayer

import (
	"fmt"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)

type OpenStackLayer struct {
	client *gophercloud.ProviderClient
}

func (this *OpenStackLayer) SimpleAuthorize(apiId, apiKey string) error {
	return fmt.Errorf("Cannot simple authorize on OpenStack")
}

// DetailedAuthorize - Auth with username, password, and tenant id against openstack identity
func (osl *OpenStackLayer) DetailedAuthorize(authDetails map[string]string) error {
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: authDetails["identityEndpoint"],
		Username:         authDetails["userName"],
		Password:         authDetails["password"],
		TenantID:         authDetails["tenantId"],
		DomainName:       authDetails["domainName"],
	}

	provider, err := openstack.AuthenticatedClient(opts)

	if err != nil {
		return err
	}

	osl.client = provider

	return nil
}

// ListInstances - List the instances in this layer.
func (this *OpenStackLayer) ListInstances() ([]*Instance, error) {
	return nil, nil
}

// CreateInstance - Create a nova server
func (osl OpenStackLayer) CreateInstance(details InstanceDetails) (*Instance, error) {
	// Get a service client for nova
	serviceClient, err := openstack.NewComputeV2(osl.client, gophercloud.EndpointOpts{
		Region: details.Region,
	})
	if err != nil {
		logger.Errorf("Error getting a service client for nova: %s", err)
		return nil, err
	}
	logger.Debugf("Service client endpoint is %s", serviceClient.Endpoint)
	// Set the network options
	networks := make([]servers.Network, 0, len(details.Networks))
	for _, network := range details.Networks {
		n := servers.Network{
			UUID:    network.ID,
			Port:    network.Port,
			FixedIP: network.FixedIP,
		}
		networks = append(networks, n)
	}
	logger.Debugf("Networks is %s", networks)
	// Set the server options
	opts := servers.CreateOpts{
		Name:          details.Hostname,
		FlavorName:    details.InstanceType,
		ImageName:     details.BaseImage,
		ServiceClient: serviceClient,
		Networks:      networks,
	}
	// Create the server
	server, err := servers.Create(serviceClient, opts).Extract()
	if err != nil {
		logger.Errorf("Error in nova create: %s", err)
		return nil, err
	}
	logger.Infof("Provisioned nova server %s", server.ID)
	inst := &Instance{
		ID:      server.ID,
		Details: details,
		Status:  "PENDING",
	}
	return inst, nil
}

func (this *OpenStackLayer) GetInstance(instanceId string) (*Instance, error) {
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
