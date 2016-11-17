// +build openstack

package cloudlayer

import (
	"testing"
	"time"
)

func TestOpenStackCreate(t *testing.T) {

	layer, err := NewCloudLayer("openstack")
	if err != nil {
		t.Fatalf("Could not create Openstack layer.")
	}
	// This is setup for little_mimic
	err = layer.DetailedAuthorize(map[string]string{
		"identityEndpoint": "http://localhost:8080/v3",
		"userName":         "testuser",
		"password":         "testpass",
		"tenantId":         "352a891c064648f38983f165e5aeb28d",
		"domainName":       "default",
	})
	if err != nil {
		t.Fatalf("Could not authorize: %s", err)
	}

	details := InstanceDetails{
		Hostname:     "polka1",
		InstanceType: "m1.small",
		BaseImage:    "trusty-server-cloudimg-amd64-disk1.img",
		Networks: []NetworkDetails{
			{
				ID: "edd948da-6620-4330-a319-ed89547dcc0a",
			},
			{
				ID: "103143e6-867a-4673-8eeb-f2797608f862",
			},
		},
	}
	inst, err := layer.CreateInstance(details)
	if err != nil {
		t.Fatalf("Could not create an openstack instance: %s", err)
	}
	if inst.Details.Hostname != "polka1" {
		t.Fatalf("Not enough polka")
	}
	var instGet *Instance
	for i := 0; i < 100; i++ {
		instGet, err = layer.GetInstance(inst.ID)
		if err != nil {
			t.Fatalf("Could not get instance %s: %s", inst.ID, err)
		}
		logger.Debugf("test: instGet is %s", instGet)
		logger.Debugf("test: instGet.status is %s", instGet.Status)
		logger.Debugf("test: instGet.details is %#v", instGet.Details)
		if instGet.Status == "ACTIVE" {
			verifyInstanceDetails(t, *instGet)
			break
		}
		time.Sleep(2 * time.Second)
	}
	if instGet.Status != "ACTIVE" {
		t.Fatalf("Instance %s never went active, stuck in %s", inst.ID, inst.Status)
	}
	_, err = layer.RemoveInstance(inst.ID)
	if err != nil {
		t.Fatalf("Failed to delete instance %s: %s", inst.ID, err)
	}
}

// verifyInstanceDetails - check an active nova instance for things
func verifyInstanceDetails(t *testing.T, inst Instance) error {

	if len(inst.Details.Hostname) == 0 {
		t.Fatalf("Hostname is empty")
	}
	if len(inst.Details.InstanceType) == 0 {
		t.Fatalf("InstanceType is empty")
	}
	if len(inst.Details.BaseImage) == 0 {
		t.Fatalf("BaseImage is empty")
	}
	if len(inst.Details.PublicIP) == 0 {
		t.Fatalf("PublicIP is empty")
	}
	if len(inst.Details.PrivateIP) == 0 {
		t.Fatalf("PrivateIP is empty")
	}
	return nil
}
