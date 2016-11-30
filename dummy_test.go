package cloudlayer

import "testing"

func TestDummyAuthorize(t *testing.T) {
	_, err := NewCloudLayer("dummy")
	if err != nil {
		t.Fatalf("Could not create Dummy layer: %s", err)
	}

}
