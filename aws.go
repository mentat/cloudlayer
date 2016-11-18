package cloudlayer

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type AWSSessionCache struct {
	cache map[string]*session.Session
	mutex sync.RWMutex
}

func (this *AWSSessionCache) GetSession(region, apiId, apiKey string) (*session.Session, error) {
	if this.cache == nil {
		this.mutex.Lock()
		this.cache = make(map[string]*session.Session)
		this.mutex.Unlock()
	}
	this.mutex.RLock()
	if sess, ok := this.cache[region]; ok {
		this.mutex.RUnlock()
		return sess, nil
	} else {
		this.mutex.RUnlock()
		sess, err := session.NewSession(&aws.Config{
			Region:      aws.String(region),
			Credentials: credentials.NewStaticCredentials(apiId, apiKey, ""),
		})
		if err != nil {
			logger.Errorf("Could not create session: %s", err)
			return nil, err
		}
		this.mutex.Lock()
		defer this.mutex.Unlock()
		this.cache[region] = sess
		return sess, nil
	}
}

var awsSessionCache AWSSessionCache
var stateMap map[string]string = map[string]string{
	"pending":       "PENDING",
	"running":       "RUNNING",
	"shutting-down": "STOPPED",
	"terminated":    "STOPPED",
	"stopping":      "STOPPED",
	"stopped":       "STOPPED",
}

var stateOpMap map[string]string = map[string]string{
	"pending":       "RUNNING",
	"running":       "DONE",
	"shutting-down": "RUNNING",
	"terminated":    "DONE",
	"stopping":      "RUNNING",
	"stopped":       "DONE",
}

type AWSLayer struct {
	apiKey string
	apiId  string
}

type instanceIdentifier struct {
	ID     string
	Type   string
	Region string
	Zone   string
}

func (this AWSLayer) mapStateValues(input string) string {
	return stateMap[input]
}

func (this AWSLayer) mapStateOpValues(input string) string {
	return stateOpMap[input]
}

func (this AWSLayer) decodeId(input string) *instanceIdentifier {
	parts := strings.Split(input, ":")
	ident := &instanceIdentifier{
		ID:     parts[2],
		Type:   parts[0],
		Zone:   parts[1],
		Region: parts[1][0 : len(parts[1])-1],
	}
	return ident
}

func (this *AWSLayer) SimpleAuthorize(apiId, apiKey string) error {
	this.apiId = apiId
	this.apiKey = apiKey
	return nil
}

func (this *AWSLayer) DetailedAuthorize(authDetails map[string]string) error {

	return nil
}

func (this *AWSLayer) CreateInstance(details InstanceDetails) (*Instance, error) {

	sess, err := awsSessionCache.GetSession(details.Region, this.apiId, this.apiKey)
	if err != nil {
		return nil, err
	}

	svc := ec2.New(sess)

	params := &ec2.RunInstancesInput{
		ImageId:  aws.String("ami-3c3b632b"), // Required
		MaxCount: aws.Int64(1),               // Required
		MinCount: aws.Int64(1),               // Required
		//AdditionalInfo: aws.String("String"),
		/*BlockDeviceMappings: []*ec2.BlockDeviceMapping{
			{ // Required
				DeviceName: aws.String("String"),
				Ebs: &ec2.EbsBlockDevice{
					DeleteOnTermination: aws.Bool(true),
					Encrypted:           aws.Bool(true),
					Iops:                aws.Int64(1),
					SnapshotId:          aws.String("String"),
					VolumeSize:          aws.Int64(1),
					VolumeType:          aws.String("VolumeType"),
				},
				NoDevice:    aws.String("String"),
				VirtualName: aws.String("String"),
			},
			// More values...
		},*/
		//ClientToken:           aws.String("String"),
		//DisableApiTermination: aws.Bool(false),
		//DryRun:                aws.Bool(true),
		//EbsOptimized: aws.Bool(true),
		/*IamInstanceProfile: &ec2.IamInstanceProfileSpecification{
			Arn:  aws.String("String"),
			Name: aws.String("String"),
		},*/
		//InstanceInitiatedShutdownBehavior: aws.String("ShutdownBehavior"),
		InstanceType: aws.String(details.InstanceType),
		//KernelId:                          aws.String("String"),
		//KeyName:                           aws.String("String"),
		Monitoring: &ec2.RunInstancesMonitoringEnabled{
			Enabled: aws.Bool(true), // Required
		},
		/*NetworkInterfaces: []*ec2.InstanceNetworkInterfaceSpecification{
			{ // Required
				AssociatePublicIpAddress: aws.Bool(true),
				DeleteOnTermination:      aws.Bool(true),
				Description:              aws.String("String"),
				DeviceIndex:              aws.Int64(1),
				Groups: []*string{
					aws.String("String"), // Required
					// More values...
				},
				NetworkInterfaceId: aws.String("String"),
				PrivateIpAddress:   aws.String("String"),
				PrivateIpAddresses: []*ec2.PrivateIpAddressSpecification{
					{ // Required
						PrivateIpAddress: aws.String("String"), // Required
						Primary:          aws.Bool(true),
					},
					// More values...
				},
				SecondaryPrivateIpAddressCount: aws.Int64(1),
				SubnetId:                       aws.String("String"),
			},
			// More values...
		},*/
		/*Placement: &ec2.Placement{
			Affinity:         aws.String("String"),
			AvailabilityZone: aws.String("String"),
			GroupName:        aws.String("String"),
			HostId:           aws.String("String"),
			Tenancy:          aws.String("Tenancy"),
		},*/
		//PrivateIpAddress: aws.String("String"),
		//RamdiskId:        aws.String("String"),
		/*SecurityGroupIds: []*string{
			aws.String("String"), // Required
			// More values...
		},
		SecurityGroups: []*string{
			aws.String("String"), // Required
			// More values...
		},*/
		//SubnetId: aws.String("String"),
		//UserData: aws.String("String"),
	}

	resp, err := svc.RunInstances(params)

	if err != nil {
		return nil, err
	}

	newId := fmt.Sprintf("i:%s:%s",
		SafeString(resp.Instances[0].Placement.AvailabilityZone),
		SafeString(resp.Instances[0].InstanceId))

	inst := &Instance{
		ID:     newId,
		Status: this.mapStateValues(SafeString(resp.Instances[0].State.Name)),
		CurrentOperation: Operation{
			ID:        newId,
			StartTime: resp.Instances[0].LaunchTime,
		},
		Details: InstanceDetails{
			PublicIP:     SafeString(resp.Instances[0].PublicIpAddress),
			PrivateIP:    SafeString(resp.Instances[0].PrivateIpAddress),
			LaunchTime:   resp.Instances[0].LaunchTime,
			InstanceType: SafeString(resp.Instances[0].InstanceType),
		},
	}

	return inst, nil
}

func (this *AWSLayer) GetInstance(instanceId string) (*Instance, error) {

	ident := this.decodeId(instanceId)

	sess, err := awsSessionCache.GetSession(ident.Region, this.apiId, this.apiKey)
	if err != nil {
		return nil, err
	}

	svc := ec2.New(sess)

	params := &ec2.DescribeInstancesInput{
		//DryRun: aws.Bool(true),
		/*Filters: []*ec2.Filter{
			{ // Required
				Name: aws.String("String"),
				Values: []*string{
					aws.String("String"), // Required
					// More values...
				},
			},
			// More values...
		},*/
		InstanceIds: []*string{
			aws.String(ident.ID), // Required
			// More values...
		},
		//MaxResults: aws.Int64(1),
		//NextToken:  aws.String("String"),
	}
	resp, err := svc.DescribeInstances(params)

	if err != nil {
		return nil, err
	}

	rawInstance := resp.Reservations[0].Instances[0]

	inst := &Instance{
		ID:     instanceId,
		Status: this.mapStateValues(SafeString(rawInstance.State.Name)),
		CurrentOperation: Operation{
			ID:        instanceId,
			StartTime: rawInstance.LaunchTime,
			Status:    this.mapStateOpValues(SafeString(rawInstance.State.Name)),
		},
		Details: InstanceDetails{
			PublicIP:     SafeString(rawInstance.PublicIpAddress),
			PrivateIP:    SafeString(rawInstance.PrivateIpAddress),
			LaunchTime:   rawInstance.LaunchTime,
			InstanceType: SafeString(rawInstance.InstanceType),
		},
	}

	return inst, nil
}

func (this *AWSLayer) RemoveInstance(instanceId string) (*Operation, error) {

	ident := this.decodeId(instanceId)

	sess, err := awsSessionCache.GetSession(ident.Region, this.apiId, this.apiKey)
	if err != nil {
		return nil, err
	}

	svc := ec2.New(sess)
	now := time.Now()

	params := &ec2.StopInstancesInput{
		InstanceIds: []*string{ // Required
			aws.String(ident.ID), // Required
			// More values...
		},
		//DryRun: aws.Bool(true),
		Force: aws.Bool(true),
	}

	_, err = svc.StopInstances(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		return nil, err
	}

	op := &Operation{
		ID:        instanceId,
		Name:      "Shutting down instance.",
		Status:    "PENDING",
		StartTime: &now,
	}

	return op, nil
}

func (this *AWSLayer) CheckOperationStatus(operationId string) (*Operation, error) {
	// AWS doesn't have a real operation, but we can kind of fake it
	// by just getting the instance.
	ident := this.decodeId(operationId)

	sess, err := awsSessionCache.GetSession(ident.Region, this.apiId, this.apiKey)
	if err != nil {
		return nil, err
	}

	svc := ec2.New(sess)

	params := &ec2.DescribeInstancesInput{
		InstanceIds: []*string{
			aws.String(ident.ID), // Required
		},
	}
	resp, err := svc.DescribeInstances(params)

	if err != nil {
		return nil, err
	}

	rawInstance := resp.Reservations[0].Instances[0]

	op := &Operation{
		ID:        operationId,
		StartTime: rawInstance.LaunchTime,
		Status:    this.mapStateOpValues(SafeString(rawInstance.State.Name)),
	}

	return op, nil
}

func (this *AWSLayer) CreateVolume(details VolumeDetails) (*Operation, error) {
	return nil, nil
}

func (this *AWSLayer) RemoveVolume(volumeId string) (*Operation, error) {
	return nil, nil
}

// Create a volume snapshot
func (this *AWSLayer) CreateSnapshot(volumnId string) (*Operation, error) {
	return nil, nil
}

// Remove a volume snapshot
func (this *AWSLayer) RemoveSnapshot(volumnId string) (*Operation, error) {
	return nil, nil
}

// List current snapshots for the current account
func (this *AWSLayer) ListSnapshots() ([]SnapshotDetails, error) {
	return nil, nil
}
