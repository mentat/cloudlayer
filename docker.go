package cloudlayer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"reflect"
	"strings"
)

// NewDockerLayer - create a new docker layer with appropriate defaults.
func NewDockerLayer() *DockerLayer {
	// Default layer to normal Docker socket.
	layer := &DockerLayer{
		dockerAddress: "/var/run/docker.sock",
	}
	return layer
}

// DockerLayer - the interface to the Docker Remote API.
type DockerLayer struct {
	dockerAddress string
}

func (docker DockerLayer) getURL(url string) string {
	if strings.HasPrefix(docker.dockerAddress, "/") {
		return fmt.Sprintf("http://unix.sock%s", url)
	}
	return fmt.Sprintf("%s%s", docker.dockerAddress, url)
}

func (docker DockerLayer) getHTTPClient() *http.Client {

	dialUnix := func(proto, addr string) (conn net.Conn, err error) {
		return net.Dial("unix", docker.dockerAddress)
	}

	if strings.HasPrefix(docker.dockerAddress, "/") {
		return &http.Client{
			Transport: &http.Transport{
				Dial: dialUnix,
			},
		}
	}
	return &http.Client{}
}

func (docker DockerLayer) get(url string, kind interface{}) (interface{}, error) {

	client := docker.getHTTPClient()

	resp, err := client.Get(docker.getURL(url))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("Error from Docker API, status: %d", resp.StatusCode)
	}

	// Create the type that we expect to recieve.
	kindType := reflect.TypeOf(kind)
	if kindType.Kind() == reflect.Array {
		fmt.Printf("ARRAY")
		innerType := kindType.Elem()
		kindObj := reflect.MakeSlice(innerType, 0, 1)
		err = json.NewDecoder(resp.Body).Decode(&kindObj)
		if err != nil {
			return nil, err
		}
		return kindObj, nil
	} else {
		kindObj := reflect.New(kindType)
		err = json.NewDecoder(resp.Body).Decode(&kindObj)
		if err != nil {
			return nil, err
		}
		return kindObj, nil
	}

}

func (docker DockerLayer) delete(url string, kind interface{}) (interface{}, error) {

	client := docker.getHTTPClient()
	req, err := http.NewRequest("DELETE", docker.getURL(url), nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return nil, fmt.Errorf("Error from Docker API, status: %d", resp.StatusCode)
	}

	// Create the type that we expect to recieve.
	kindObj := reflect.New(reflect.TypeOf(kind))

	err = json.NewDecoder(resp.Body).Decode(&kindObj)
	if err != nil {
		return nil, err
	}
	return kindObj, nil
}

func (docker DockerLayer) post(url string, data interface{}, kind interface{}) (interface{}, error) {

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	client := docker.getHTTPClient()

	resp, err := client.Post(docker.getURL(url), "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, fmt.Errorf("Error from Docker API, status: %d: %s", resp.StatusCode, body)
	}

	if kind != nil {
		// Create the type that we expect to recieve.
		kindObj := reflect.New(reflect.TypeOf(kind)).Interface()
		// Decode the JSON to that type.
		err = json.NewDecoder(resp.Body).Decode(&kindObj)

		if err != nil {
			return nil, err
		}
		return kindObj, nil
	}

	return nil, nil
}

// SimpleAuthorize - Authorize this cloud layer with just an API key.
func (docker DockerLayer) SimpleAuthorize(apiID, apiKey string) error {
	return fmt.Errorf("Docker does not support simple authorization.")
}

// DetailedAuthorize - Authorize a Docker Remote API layer.
func (docker *DockerLayer) DetailedAuthorize(authDetails map[string]string) error {
	v, ok := authDetails["host"]

	if !ok {
		return fmt.Errorf("Host not found in authorization details.")
	}
	docker.dockerAddress = v
	return nil
}

// CreateInstance - Create a new instance in this cloud layer.
func (docker DockerLayer) CreateInstance(details InstanceDetails) (*Instance, error) {

	// Create ExposedPorts map
	exposedPorts := make(map[string]struct{}, len(details.ExposedPorts))
	for _, v := range details.ExposedPorts {
		exposedPorts[fmt.Sprintf("%d/%s", v.InstancePort, v.Protocol)] = struct{}{}
	}

	// Create host map for ports
	hostPorts := make(dockerPortMap, len(details.ExposedPorts))
	for _, v := range details.ExposedPorts {
		hostPorts[fmt.Sprintf("%d/%s", v.InstancePort, v.Protocol)] = []dockerHostPort{
			dockerHostPort{
				HostPort: fmt.Sprintf("%d", v.HostPort),
			},
		}
	}

	env := make([]string, 0, len(details.Variables))
	for k, v := range details.Variables {
		env = append(env, fmt.Sprintf("%s=%s", k, v))
	}

	requestObj := dockerCreateContainerRequest{
		Hostname: details.Hostname,
		Image:    details.BaseImage,
		HostConfig: dockerHostConfig{
			PortBindings: hostPorts,
		},
		ExposedPorts: exposedPorts,
		Env:          env,
	}

	respRaw, err := docker.post("/containers/create", &requestObj, dockerCreateResponse{})
	if err != nil {
		logger.Errorf("Error creating container: %s", err)
		return nil, err
	}

	// Type assert the interface from the post method to a hard type.
	resp, ok := respRaw.(*dockerCreateResponse)
	if !ok {
		return nil, fmt.Errorf("Could not type assert response.")
	}

	_, err = docker.post(fmt.Sprintf("/containers/%s/start", resp.ID), struct{}{}, nil)

	if err != nil {
		logger.Errorf("Error starting container: %s %s", resp.ID, err)
		return nil, err
	}

	inst := &Instance{
		ID:      resp.ID,
		Details: details,
		CurrentOperation: Operation{
			ID:     resp.ID,
			Status: "PENDING",
		},
		Status: "PENDING",
	}

	return inst, nil
}

// RemoveInstance - Remove/stop an instance from this cloud layer.
func (docker DockerLayer) RemoveInstance(instanceID string) (*Operation, error) {

	_, err := docker.post(fmt.Sprintf("/containers/%s/stop", instanceID), struct{}{}, nil)

	if err != nil {
		logger.Errorf("Error stopping container: %s %s", instanceID, err)
		return nil, err
	}

	op := &Operation{
		ID:     instanceID,
		Status: "PENDING",
	}

	return op, nil
}

// ListInstances - List the instances in this layer.
func (docker DockerLayer) ListInstances() ([]*Instance, error) {
	resp, err := docker.get("/containers/json", []dockerCreateContainerRequest{})

	if err != nil {
		logger.Errorf("Error listing containers: %s", err)
		return nil, err
	}

	respReal, ok := resp.([]dockerCreateContainerRequest)
	if !ok {
		return nil, fmt.Errorf("Cannot convert types correctly.")
	}

	ret := make([]*Instance, len(respReal))
	for i, v := range respReal {
		inst := &Instance{
			ID: v.ID,
		}
		ret[i] = inst
	}

	return ret, nil
}

// GetInstance - Get an instance from the layer.
func (docker DockerLayer) GetInstance(instanceID string) (*Instance, error) {
	resp, err := docker.get(fmt.Sprintf("/containers/%s/json", instanceID), dockerInspectResponse{})

	if err != nil {
		logger.Errorf("Error inspecting container: %s", err)
		return nil, err
	}

	respReal, ok := resp.(dockerInspectResponse)
	if !ok {
		return nil, fmt.Errorf("Cannot convert types correctly.")
	}

	inst := &Instance{
		ID:      respReal.ID,
		Details: InstanceDetails{},
		Status:  "",
	}

	return inst, nil
}

// CheckOperationStatus - Check the status of a long running operation.
func (docker DockerLayer) CheckOperationStatus(operationID string) (*Operation, error) {
	return nil, nil
}

// CreateVolume - Create a new data storage volume.
func (docker DockerLayer) CreateVolume(details VolumeDetails) (*Operation, error) {
	return nil, nil
}

// RemoveVolume - Remove a data storage volume.
func (docker DockerLayer) RemoveVolume(volumeID string) (*Operation, error) {
	return nil, nil
}

// CreateSnapshot - Create a volume snapshot
func (docker DockerLayer) CreateSnapshot(volumnID string) (*Operation, error) {
	return nil, nil
}

// RemoveSnapshot - Remove a volume snapshot
func (docker DockerLayer) RemoveSnapshot(volumnID string) (*Operation, error) {
	return nil, nil
}

// ListSnapshots - List current snapshots for the current account
func (docker DockerLayer) ListSnapshots() ([]SnapshotDetails, error) {
	return nil, nil
}
