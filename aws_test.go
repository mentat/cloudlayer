package cloudlayer

import (
	"fmt"
	"testing"
	"time"
)

const AWS_ID = "AKIAIVH5C42CYRM3PG4A"
const AWS_SECRET = "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"

func TestAWSAuthorize(t *testing.T) {
	layer, err := NewCloudLayer("aws")
	if err != nil {
		t.Fatalf("Could not create AWS layer: %s", err)
	}
	err = layer.SimpleAuthorize(AWS_ID, AWS_SECRET)
	if err != nil {
		t.Fatalf("Could not authorize: %s", err)
	}
}

func TestAWSCreateInstance(t *testing.T) {

	if testing.Short() {
		t.Skip("Skipping test in short mode.")
	}

	layer, err := NewCloudLayer("aws")
	if err != nil {
		t.Fatalf("Could not create AWS layer: %s", err)
	}

	err = layer.SimpleAuthorize(AWS_ID, AWS_SECRET)
	if err != nil {
		t.Fatalf("Could not authorize: %s", err)
	}

	details := InstanceDetails{
		InstanceType: "t1.micro",
		Region:       "us-east-1",
	}

	inst, err := layer.CreateInstance(details)
	if err != nil {
		t.Fatalf("Could not create instance: %s", err)
	}
	t.Logf("Instance is: %s", inst.ID)

	// Wait for instance to come online
	for {
		status, err := layer.GetInstance(inst.ID)
		if err != nil {
			t.Errorf("Problem talking to AWS: %s", err)
			fmt.Printf("Problem talking to AWS: %s\n", err)
		} else if status.Status == "RUNNING" {
			break
		}
		//fmt.Printf("%s\n", status.Status)
		time.Sleep(time.Millisecond * 500)
		t.Logf("Waiting for instance to boot...")
	}

	// Remove instance
	_, err = layer.RemoveInstance(inst.ID)
	if err != nil {
		t.Fatalf("Could not remove instance: %s", err)
	}

}

func TestAWSCreateVolume(t *testing.T) {
	layer, err := NewCloudLayer("aws")
	if err != nil {
		t.Fatalf("Could not create AWS layer: %s", err)
	}
	err = layer.SimpleAuthorize(AWS_ID, AWS_SECRET)
	if err != nil {
		t.Fatalf("Could not authorize: %s", err)
	}
}
