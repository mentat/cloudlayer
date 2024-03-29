// +build docker

package cloudlayer

import "testing"

func TestDockerCreate(t *testing.T) {
	layer, err := NewCloudLayer("docker")
	if err != nil {
		t.Fatalf("Could not create AWS layer: %s", err)
	}

	err = layer.DetailedAuthorize(map[string]string{
		"host": "/var/run/docker.sock",
	})
	if err != nil {
		t.Fatalf("Could not authorize: %s", err)
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

	instances, err := layer.ListInstances()
	if err != nil {
		layer.RemoveInstance(inst.ID)
		t.Fatalf("Could not list instances: %s", err)
	}

	if instances[0].ID != inst.ID {
		layer.RemoveInstance(inst.ID)
		t.Fatalf("Instance ID isn't expected: %s", instances[0].ID)
	}

	inst2, err := layer.GetInstance(inst.ID)
	if err != nil {
		layer.RemoveInstance(inst.ID)
		t.Fatalf("Could not inspect docker container: %s", err)
	}

	if inst2.Details.PublicIP == "" {
		layer.RemoveInstance(inst.ID)
		t.Fatalf("Could not get IP from container.")
	}

	_, err = layer.RemoveInstance(inst.ID)
	if err != nil {
		t.Fatalf("Could not remove docker container: %s", err)
	}

}
