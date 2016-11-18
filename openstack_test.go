package cloudlayer

import "testing"

func _TestOpenStackAuthorize(t *testing.T) {

	layer, err := NewCloudLayer("openstack")
	if err != nil {
		t.Fatalf("Could not create Openstack layer.")
	}
	err = layer.DetailedAuthorize(map[string]string{
		"identityEndpoint": "https://...",
		"password":         "...",
		"tenantId":         "...",
	})
	if err != nil {
		t.Fatalf("Could not authorize: %s", err)
	}
}
