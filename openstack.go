package cloudlayer

import (
	"fmt"
	"strings"

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
	// TODO(tvoran): set region in authorize call?
	// TODO(tvoran): prefer openstack auth vars from env over static settings

	// These are the default auth options if none are found in the env
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
	serviceClient, err := openstack.NewComputeV2(osl.client, gophercloud.EndpointOpts{})
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
	logger.Debugf("Networks is %#v", networks)
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
	logger.Infof("Created nova server %s", server.ID)
	inst := &Instance{
		ID:      server.ID,
		Details: details,
		Status:  "PENDING",
	}
	return inst, nil
}

// GetInstance - Get a single nova server's details
func (osl OpenStackLayer) GetInstance(instanceID string) (*Instance, error) {
	// TODO(tvoran): how do we get the region for NewComputeV2?
	serviceClient, err := openstack.NewComputeV2(osl.client, gophercloud.EndpointOpts{})
	if err != nil {
		logger.Errorf("Error getting a service client for nova: %s", err)
		return nil, err
	}
	server, err := servers.Get(serviceClient, instanceID).Extract()
	if err != nil {
		logger.Errorf("Error in nova get for instanceID %s: %s", instanceID, err)
		return nil, err
	}

	var publicIP, privateIP string
	if len(server.Addresses) > 0 {
		for k, v := range server.Addresses {
			// This seems kind of gross
			addrData := v.([]interface{})[0].(map[string]interface{})
			switch {
			case strings.Contains(k, "external"), strings.Contains(k, "public"):
				publicIP = addrData["addr"].(string)
			case strings.Contains(k, "private"):
				privateIP = addrData["addr"].(string)
			}
		}
		// If AccessIPv4 came back from nova, override the publicIP
		if len(server.AccessIPv4) > 0 {
			publicIP = server.AccessIPv4
		}
	}

	details := InstanceDetails{
		Hostname:     server.Name,
		InstanceType: server.Flavor["id"].(string),
		BaseImage:    server.Image["id"].(string),
		PublicIP:     publicIP,
		PrivateIP:    privateIP,
	}
	inst := &Instance{
		ID:      server.ID,
		Details: details,
		Status:  server.Status,
	}

	logger.Infof("Retrieved nova server %s", instanceID)

	return inst, nil
}

// RemoveInstance - Calls nova delete
func (osl OpenStackLayer) RemoveInstance(instanceID string) (*Operation, error) {
	serviceClient, err := openstack.NewComputeV2(osl.client, gophercloud.EndpointOpts{})
	if err != nil {
		logger.Errorf("Error getting a service client for nova: %s", err)
		return nil, err
	}
	err = servers.Delete(serviceClient, instanceID).ExtractErr()
	if err != nil {
		logger.Errorf("Error deleting nova server %s: %s", instanceID, err)
	}

	op := &Operation{
		ID:     instanceID,
		Name:   "Delete",
		Status: "PENDING",
	}

	logger.Infof("Deleted nova server %s", instanceID)

	return op, nil
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
