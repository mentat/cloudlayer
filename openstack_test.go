// +build openstack

package cloudlayer

import "testing"

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
		Region:       "RegionOne",
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
}
