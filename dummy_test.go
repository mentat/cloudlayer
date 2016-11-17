package cloudlayer

import "testing"

func TestDummyAuthorize(t *testing.T) {
	_, err := NewCloudLayer("dummy")
	if err != nil {
		t.Fatalf("Could not create Dummy layer: %s", err)
	}

}

func TestDummyCreate(t *testing.T) {
	layer, err := NewCloudLayer("dummy")
	if err != nil {
		t.Fatalf("Could not create AWS layer: %s", err)
	}
	inst, err := layer.CreateInstance(InstanceDetails{
		BaseImage: "consul",
		ExposedPorts: []PortDetails{
			PortDetails{
				InstancePort: 8400,
				HostPort:     8400,
				Protocol:     "tcp",
			},
		},
	})
	if err != nil {
		t.Fatalf("Could not create docker container: %s", err)
	}

	_, err = layer.RemoveInstance(inst.ID)
	if err != nil {
		t.Fatalf("Could not remove docker container: %s", err)
	}

}
