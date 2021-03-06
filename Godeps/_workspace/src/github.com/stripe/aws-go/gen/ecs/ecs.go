// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package ecs provides a client for Amazon EC2 Container Service.
package ecs

import (
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/gen/endpoints"
)

import (
	"encoding/xml"
	"io"
)

// ECS is a client for Amazon EC2 Container Service.
type ECS struct {
	client *aws.QueryClient
}

// New returns a new ECS client.
func New(creds aws.CredentialsProvider, region string, client *http.Client) *ECS {
	if client == nil {
		client = http.DefaultClient
	}

	endpoint, service, region := endpoints.Lookup("ecs", region)

	return &ECS{
		client: &aws.QueryClient{
			Context: aws.Context{
				Credentials: creds,
				Service:     service,
				Region:      region,
			},
			Client:     client,
			Endpoint:   endpoint,
			APIVersion: "2014-11-13",
		},
	}
}

// CreateCluster creates a new Amazon ECS cluster. By default, your account
// will receive a default cluster when you launch your first container
// instance. However, you can create your own cluster with a unique name
// with the CreateCluster action. During the preview, each account is
// limited to two clusters.
func (c *ECS) CreateCluster(req *CreateClusterRequest) (resp *CreateClusterResult, err error) {
	resp = &CreateClusterResult{}
	err = c.client.Do("CreateCluster", "POST", "/", req, resp)
	return
}

// DeleteCluster deletes the specified cluster. You must deregister all
// container instances from this cluster before you may delete it. You can
// list the container instances in a cluster with ListContainerInstances
// and deregister them with DeregisterContainerInstance
func (c *ECS) DeleteCluster(req *DeleteClusterRequest) (resp *DeleteClusterResult, err error) {
	resp = &DeleteClusterResult{}
	err = c.client.Do("DeleteCluster", "POST", "/", req, resp)
	return
}

// DeregisterContainerInstance deregisters an Amazon ECS container instance
// from the specified cluster. This instance will no longer be available to
// run tasks.
func (c *ECS) DeregisterContainerInstance(req *DeregisterContainerInstanceRequest) (resp *DeregisterContainerInstanceResult, err error) {
	resp = &DeregisterContainerInstanceResult{}
	err = c.client.Do("DeregisterContainerInstance", "POST", "/", req, resp)
	return
}

// DeregisterTaskDefinition deregisters the specified task definition. You
// will no longer be able to run tasks from this definition after
// deregistration.
func (c *ECS) DeregisterTaskDefinition(req *DeregisterTaskDefinitionRequest) (resp *DeregisterTaskDefinitionResult, err error) {
	resp = &DeregisterTaskDefinitionResult{}
	err = c.client.Do("DeregisterTaskDefinition", "POST", "/", req, resp)
	return
}

// DescribeClusters is undocumented.
func (c *ECS) DescribeClusters(req *DescribeClustersRequest) (resp *DescribeClustersResult, err error) {
	resp = &DescribeClustersResult{}
	err = c.client.Do("DescribeClusters", "POST", "/", req, resp)
	return
}

// DescribeContainerInstances describes Amazon EC2 Container Service
// container instances. Returns metadata about registered and remaining
// resources on each container instance requested.
func (c *ECS) DescribeContainerInstances(req *DescribeContainerInstancesRequest) (resp *DescribeContainerInstancesResult, err error) {
	resp = &DescribeContainerInstancesResult{}
	err = c.client.Do("DescribeContainerInstances", "POST", "/", req, resp)
	return
}

// DescribeTaskDefinition is undocumented.
func (c *ECS) DescribeTaskDefinition(req *DescribeTaskDefinitionRequest) (resp *DescribeTaskDefinitionResult, err error) {
	resp = &DescribeTaskDefinitionResult{}
	err = c.client.Do("DescribeTaskDefinition", "POST", "/", req, resp)
	return
}

// DescribeTasks is undocumented.
func (c *ECS) DescribeTasks(req *DescribeTasksRequest) (resp *DescribeTasksResult, err error) {
	resp = &DescribeTasksResult{}
	err = c.client.Do("DescribeTasks", "POST", "/", req, resp)
	return
}

// DiscoverPollEndpoint this action is only used by the Amazon EC2
// Container Service agent, and it is not intended for use outside of the
// agent. Returns an endpoint for the Amazon EC2 Container Service agent to
// poll for updates.
func (c *ECS) DiscoverPollEndpoint(req *DiscoverPollEndpointRequest) (resp *DiscoverPollEndpointResult, err error) {
	resp = &DiscoverPollEndpointResult{}
	err = c.client.Do("DiscoverPollEndpoint", "POST", "/", req, resp)
	return
}

// ListClusters is undocumented.
func (c *ECS) ListClusters(req *ListClustersRequest) (resp *ListClustersResult, err error) {
	resp = &ListClustersResult{}
	err = c.client.Do("ListClusters", "POST", "/", req, resp)
	return
}

// ListContainerInstances returns a list of container instances in a
// specified cluster.
func (c *ECS) ListContainerInstances(req *ListContainerInstancesRequest) (resp *ListContainerInstancesResult, err error) {
	resp = &ListContainerInstancesResult{}
	err = c.client.Do("ListContainerInstances", "POST", "/", req, resp)
	return
}

// ListTaskDefinitions returns a list of task definitions that are
// registered to your account. You can filter the results by family name
// with the familyPrefix parameter.
func (c *ECS) ListTaskDefinitions(req *ListTaskDefinitionsRequest) (resp *ListTaskDefinitionsResult, err error) {
	resp = &ListTaskDefinitionsResult{}
	err = c.client.Do("ListTaskDefinitions", "POST", "/", req, resp)
	return
}

// ListTasks returns a list of tasks for a specified cluster. You can
// filter the results by family name or by a particular container instance
// with the family and containerInstance parameters.
func (c *ECS) ListTasks(req *ListTasksRequest) (resp *ListTasksResult, err error) {
	resp = &ListTasksResult{}
	err = c.client.Do("ListTasks", "POST", "/", req, resp)
	return
}

// RegisterContainerInstance this action is only used by the Amazon EC2
// Container Service agent, and it is not intended for use outside of the
// agent. Registers an Amazon EC2 instance into the specified cluster. This
// instance will become available to place containers on.
func (c *ECS) RegisterContainerInstance(req *RegisterContainerInstanceRequest) (resp *RegisterContainerInstanceResult, err error) {
	resp = &RegisterContainerInstanceResult{}
	err = c.client.Do("RegisterContainerInstance", "POST", "/", req, resp)
	return
}

// RegisterTaskDefinition registers a new task definition from the supplied
// family and containerDefinitions
func (c *ECS) RegisterTaskDefinition(req *RegisterTaskDefinitionRequest) (resp *RegisterTaskDefinitionResult, err error) {
	resp = &RegisterTaskDefinitionResult{}
	err = c.client.Do("RegisterTaskDefinition", "POST", "/", req, resp)
	return
}

// RunTask start a task using random placement and the default Amazon ECS
// scheduler. If you want to use your own scheduler or place a task on a
// specific container instance, use StartTask instead.
func (c *ECS) RunTask(req *RunTaskRequest) (resp *RunTaskResult, err error) {
	resp = &RunTaskResult{}
	err = c.client.Do("RunTask", "POST", "/", req, resp)
	return
}

// StartTask starts a new task from the specified task definition on the
// specified container instance or instances. If you want to use the
// default Amazon ECS scheduler to place your task, use RunTask instead.
func (c *ECS) StartTask(req *StartTaskRequest) (resp *StartTaskResult, err error) {
	resp = &StartTaskResult{}
	err = c.client.Do("StartTask", "POST", "/", req, resp)
	return
}

// StopTask is undocumented.
func (c *ECS) StopTask(req *StopTaskRequest) (resp *StopTaskResult, err error) {
	resp = &StopTaskResult{}
	err = c.client.Do("StopTask", "POST", "/", req, resp)
	return
}

// SubmitContainerStateChange this action is only used by the Amazon EC2
// Container Service agent, and it is not intended for use outside of the
// agent. Sent to acknowledge that a container changed states.
func (c *ECS) SubmitContainerStateChange(req *SubmitContainerStateChangeRequest) (resp *SubmitContainerStateChangeResult, err error) {
	resp = &SubmitContainerStateChangeResult{}
	err = c.client.Do("SubmitContainerStateChange", "POST", "/", req, resp)
	return
}

// SubmitTaskStateChange this action is only used by the Amazon EC2
// Container Service agent, and it is not intended for use outside of the
// agent. Sent to acknowledge that a task changed states.
func (c *ECS) SubmitTaskStateChange(req *SubmitTaskStateChangeRequest) (resp *SubmitTaskStateChangeResult, err error) {
	resp = &SubmitTaskStateChangeResult{}
	err = c.client.Do("SubmitTaskStateChange", "POST", "/", req, resp)
	return
}

// Cluster is undocumented.
type Cluster struct {
	ClusterARN  aws.StringValue `query:"clusterArn" xml:"clusterArn"`
	ClusterName aws.StringValue `query:"clusterName" xml:"clusterName"`
	Status      aws.StringValue `query:"status" xml:"status"`
}

// Container is undocumented.
type Container struct {
	ContainerARN    aws.StringValue  `query:"containerArn" xml:"containerArn"`
	ExitCode        aws.IntegerValue `query:"exitCode" xml:"exitCode"`
	LastStatus      aws.StringValue  `query:"lastStatus" xml:"lastStatus"`
	Name            aws.StringValue  `query:"name" xml:"name"`
	NetworkBindings []NetworkBinding `query:"networkBindings.member" xml:"networkBindings>member"`
	Reason          aws.StringValue  `query:"reason" xml:"reason"`
	TaskARN         aws.StringValue  `query:"taskArn" xml:"taskArn"`
}

// ContainerDefinition is undocumented.
type ContainerDefinition struct {
	Command      []string         `query:"command.member" xml:"command>member"`
	CPU          aws.IntegerValue `query:"cpu" xml:"cpu"`
	EntryPoint   []string         `query:"entryPoint.member" xml:"entryPoint>member"`
	Environment  []KeyValuePair   `query:"environment.member" xml:"environment>member"`
	Essential    aws.BooleanValue `query:"essential" xml:"essential"`
	Image        aws.StringValue  `query:"image" xml:"image"`
	Links        []string         `query:"links.member" xml:"links>member"`
	Memory       aws.IntegerValue `query:"memory" xml:"memory"`
	Name         aws.StringValue  `query:"name" xml:"name"`
	PortMappings []PortMapping    `query:"portMappings.member" xml:"portMappings>member"`
}

// ContainerInstance is undocumented.
type ContainerInstance struct {
	AgentConnected       aws.BooleanValue `query:"agentConnected" xml:"agentConnected"`
	ContainerInstanceARN aws.StringValue  `query:"containerInstanceArn" xml:"containerInstanceArn"`
	EC2InstanceID        aws.StringValue  `query:"ec2InstanceId" xml:"ec2InstanceId"`
	RegisteredResources  []Resource       `query:"registeredResources.member" xml:"registeredResources>member"`
	RemainingResources   []Resource       `query:"remainingResources.member" xml:"remainingResources>member"`
	Status               aws.StringValue  `query:"status" xml:"status"`
}

// ContainerOverride is undocumented.
type ContainerOverride struct {
	Command []string        `query:"command.member" xml:"command>member"`
	Name    aws.StringValue `query:"name" xml:"name"`
}

// CreateClusterRequest is undocumented.
type CreateClusterRequest struct {
	ClusterName aws.StringValue `query:"clusterName" xml:"clusterName"`
}

// CreateClusterResponse is undocumented.
type CreateClusterResponse struct {
	Cluster *Cluster `query:"cluster" xml:"CreateClusterResult>cluster"`
}

// DeleteClusterRequest is undocumented.
type DeleteClusterRequest struct {
	Cluster aws.StringValue `query:"cluster" xml:"cluster"`
}

// DeleteClusterResponse is undocumented.
type DeleteClusterResponse struct {
	Cluster *Cluster `query:"cluster" xml:"DeleteClusterResult>cluster"`
}

// DeregisterContainerInstanceRequest is undocumented.
type DeregisterContainerInstanceRequest struct {
	Cluster           aws.StringValue  `query:"cluster" xml:"cluster"`
	ContainerInstance aws.StringValue  `query:"containerInstance" xml:"containerInstance"`
	Force             aws.BooleanValue `query:"force" xml:"force"`
}

// DeregisterContainerInstanceResponse is undocumented.
type DeregisterContainerInstanceResponse struct {
	ContainerInstance *ContainerInstance `query:"containerInstance" xml:"DeregisterContainerInstanceResult>containerInstance"`
}

// DeregisterTaskDefinitionRequest is undocumented.
type DeregisterTaskDefinitionRequest struct {
	TaskDefinition aws.StringValue `query:"taskDefinition" xml:"taskDefinition"`
}

// DeregisterTaskDefinitionResponse is undocumented.
type DeregisterTaskDefinitionResponse struct {
	TaskDefinition *TaskDefinition `query:"taskDefinition" xml:"DeregisterTaskDefinitionResult>taskDefinition"`
}

// DescribeClustersRequest is undocumented.
type DescribeClustersRequest struct {
	Clusters []string `query:"clusters.member" xml:"clusters>member"`
}

// DescribeClustersResponse is undocumented.
type DescribeClustersResponse struct {
	Clusters []Cluster `query:"clusters.member" xml:"DescribeClustersResult>clusters>member"`
	Failures []Failure `query:"failures.member" xml:"DescribeClustersResult>failures>member"`
}

// DescribeContainerInstancesRequest is undocumented.
type DescribeContainerInstancesRequest struct {
	Cluster            aws.StringValue `query:"cluster" xml:"cluster"`
	ContainerInstances []string        `query:"containerInstances.member" xml:"containerInstances>member"`
}

// DescribeContainerInstancesResponse is undocumented.
type DescribeContainerInstancesResponse struct {
	ContainerInstances []ContainerInstance `query:"containerInstances.member" xml:"DescribeContainerInstancesResult>containerInstances>member"`
	Failures           []Failure           `query:"failures.member" xml:"DescribeContainerInstancesResult>failures>member"`
}

// DescribeTaskDefinitionRequest is undocumented.
type DescribeTaskDefinitionRequest struct {
	TaskDefinition aws.StringValue `query:"taskDefinition" xml:"taskDefinition"`
}

// DescribeTaskDefinitionResponse is undocumented.
type DescribeTaskDefinitionResponse struct {
	TaskDefinition *TaskDefinition `query:"taskDefinition" xml:"DescribeTaskDefinitionResult>taskDefinition"`
}

// DescribeTasksRequest is undocumented.
type DescribeTasksRequest struct {
	Cluster aws.StringValue `query:"cluster" xml:"cluster"`
	Tasks   []string        `query:"tasks.member" xml:"tasks>member"`
}

// DescribeTasksResponse is undocumented.
type DescribeTasksResponse struct {
	Failures []Failure `query:"failures.member" xml:"DescribeTasksResult>failures>member"`
	Tasks    []Task    `query:"tasks.member" xml:"DescribeTasksResult>tasks>member"`
}

// DiscoverPollEndpointRequest is undocumented.
type DiscoverPollEndpointRequest struct {
	ContainerInstance aws.StringValue `query:"containerInstance" xml:"containerInstance"`
}

// DiscoverPollEndpointResponse is undocumented.
type DiscoverPollEndpointResponse struct {
	Endpoint aws.StringValue `query:"endpoint" xml:"DiscoverPollEndpointResult>endpoint"`
}

// Failure is undocumented.
type Failure struct {
	ARN    aws.StringValue `query:"arn" xml:"arn"`
	Reason aws.StringValue `query:"reason" xml:"reason"`
}

// KeyValuePair is undocumented.
type KeyValuePair struct {
	Name  aws.StringValue `query:"name" xml:"name"`
	Value aws.StringValue `query:"value" xml:"value"`
}

// ListClustersRequest is undocumented.
type ListClustersRequest struct {
	MaxResults aws.IntegerValue `query:"maxResults" xml:"maxResults"`
	NextToken  aws.StringValue  `query:"nextToken" xml:"nextToken"`
}

// ListClustersResponse is undocumented.
type ListClustersResponse struct {
	ClusterARNs []string        `query:"clusterArns.member" xml:"ListClustersResult>clusterArns>member"`
	NextToken   aws.StringValue `query:"nextToken" xml:"ListClustersResult>nextToken"`
}

// ListContainerInstancesRequest is undocumented.
type ListContainerInstancesRequest struct {
	Cluster    aws.StringValue  `query:"cluster" xml:"cluster"`
	MaxResults aws.IntegerValue `query:"maxResults" xml:"maxResults"`
	NextToken  aws.StringValue  `query:"nextToken" xml:"nextToken"`
}

// ListContainerInstancesResponse is undocumented.
type ListContainerInstancesResponse struct {
	ContainerInstanceARNs []string        `query:"containerInstanceArns.member" xml:"ListContainerInstancesResult>containerInstanceArns>member"`
	NextToken             aws.StringValue `query:"nextToken" xml:"ListContainerInstancesResult>nextToken"`
}

// ListTaskDefinitionsRequest is undocumented.
type ListTaskDefinitionsRequest struct {
	FamilyPrefix aws.StringValue  `query:"familyPrefix" xml:"familyPrefix"`
	MaxResults   aws.IntegerValue `query:"maxResults" xml:"maxResults"`
	NextToken    aws.StringValue  `query:"nextToken" xml:"nextToken"`
}

// ListTaskDefinitionsResponse is undocumented.
type ListTaskDefinitionsResponse struct {
	NextToken          aws.StringValue `query:"nextToken" xml:"ListTaskDefinitionsResult>nextToken"`
	TaskDefinitionARNs []string        `query:"taskDefinitionArns.member" xml:"ListTaskDefinitionsResult>taskDefinitionArns>member"`
}

// ListTasksRequest is undocumented.
type ListTasksRequest struct {
	Cluster           aws.StringValue  `query:"cluster" xml:"cluster"`
	ContainerInstance aws.StringValue  `query:"containerInstance" xml:"containerInstance"`
	Family            aws.StringValue  `query:"family" xml:"family"`
	MaxResults        aws.IntegerValue `query:"maxResults" xml:"maxResults"`
	NextToken         aws.StringValue  `query:"nextToken" xml:"nextToken"`
}

// ListTasksResponse is undocumented.
type ListTasksResponse struct {
	NextToken aws.StringValue `query:"nextToken" xml:"ListTasksResult>nextToken"`
	TaskARNs  []string        `query:"taskArns.member" xml:"ListTasksResult>taskArns>member"`
}

// NetworkBinding is undocumented.
type NetworkBinding struct {
	BindIP        aws.StringValue  `query:"bindIP" xml:"bindIP"`
	ContainerPort aws.IntegerValue `query:"containerPort" xml:"containerPort"`
	HostPort      aws.IntegerValue `query:"hostPort" xml:"hostPort"`
}

// PortMapping is undocumented.
type PortMapping struct {
	ContainerPort aws.IntegerValue `query:"containerPort" xml:"containerPort"`
	HostPort      aws.IntegerValue `query:"hostPort" xml:"hostPort"`
}

// RegisterContainerInstanceRequest is undocumented.
type RegisterContainerInstanceRequest struct {
	Cluster                           aws.StringValue `query:"cluster" xml:"cluster"`
	InstanceIdentityDocument          aws.StringValue `query:"instanceIdentityDocument" xml:"instanceIdentityDocument"`
	InstanceIdentityDocumentSignature aws.StringValue `query:"instanceIdentityDocumentSignature" xml:"instanceIdentityDocumentSignature"`
	TotalResources                    []Resource      `query:"totalResources.member" xml:"totalResources>member"`
}

// RegisterContainerInstanceResponse is undocumented.
type RegisterContainerInstanceResponse struct {
	ContainerInstance *ContainerInstance `query:"containerInstance" xml:"RegisterContainerInstanceResult>containerInstance"`
}

// RegisterTaskDefinitionRequest is undocumented.
type RegisterTaskDefinitionRequest struct {
	ContainerDefinitions []ContainerDefinition `query:"containerDefinitions.member" xml:"containerDefinitions>member"`
	Family               aws.StringValue       `query:"family" xml:"family"`
}

// RegisterTaskDefinitionResponse is undocumented.
type RegisterTaskDefinitionResponse struct {
	TaskDefinition *TaskDefinition `query:"taskDefinition" xml:"RegisterTaskDefinitionResult>taskDefinition"`
}

// Resource is undocumented.
type Resource struct {
	DoubleValue    aws.DoubleValue  `query:"doubleValue" xml:"doubleValue"`
	IntegerValue   aws.IntegerValue `query:"integerValue" xml:"integerValue"`
	LongValue      aws.LongValue    `query:"longValue" xml:"longValue"`
	Name           aws.StringValue  `query:"name" xml:"name"`
	StringSetValue []string         `query:"stringSetValue.member" xml:"stringSetValue>member"`
	Type           aws.StringValue  `query:"type" xml:"type"`
}

// RunTaskRequest is undocumented.
type RunTaskRequest struct {
	Cluster        aws.StringValue  `query:"cluster" xml:"cluster"`
	Count          aws.IntegerValue `query:"count" xml:"count"`
	Overrides      *TaskOverride    `query:"overrides" xml:"overrides"`
	TaskDefinition aws.StringValue  `query:"taskDefinition" xml:"taskDefinition"`
}

// RunTaskResponse is undocumented.
type RunTaskResponse struct {
	Failures []Failure `query:"failures.member" xml:"RunTaskResult>failures>member"`
	Tasks    []Task    `query:"tasks.member" xml:"RunTaskResult>tasks>member"`
}

// StartTaskRequest is undocumented.
type StartTaskRequest struct {
	Cluster            aws.StringValue `query:"cluster" xml:"cluster"`
	ContainerInstances []string        `query:"containerInstances.member" xml:"containerInstances>member"`
	Overrides          *TaskOverride   `query:"overrides" xml:"overrides"`
	TaskDefinition     aws.StringValue `query:"taskDefinition" xml:"taskDefinition"`
}

// StartTaskResponse is undocumented.
type StartTaskResponse struct {
	Failures []Failure `query:"failures.member" xml:"StartTaskResult>failures>member"`
	Tasks    []Task    `query:"tasks.member" xml:"StartTaskResult>tasks>member"`
}

// StopTaskRequest is undocumented.
type StopTaskRequest struct {
	Cluster aws.StringValue `query:"cluster" xml:"cluster"`
	Task    aws.StringValue `query:"task" xml:"task"`
}

// StopTaskResponse is undocumented.
type StopTaskResponse struct {
	Task *Task `query:"task" xml:"StopTaskResult>task"`
}

// SubmitContainerStateChangeRequest is undocumented.
type SubmitContainerStateChangeRequest struct {
	Cluster         aws.StringValue  `query:"cluster" xml:"cluster"`
	ContainerName   aws.StringValue  `query:"containerName" xml:"containerName"`
	ExitCode        aws.IntegerValue `query:"exitCode" xml:"exitCode"`
	NetworkBindings []NetworkBinding `query:"networkBindings.member" xml:"networkBindings>member"`
	Reason          aws.StringValue  `query:"reason" xml:"reason"`
	Status          aws.StringValue  `query:"status" xml:"status"`
	Task            aws.StringValue  `query:"task" xml:"task"`
}

// SubmitContainerStateChangeResponse is undocumented.
type SubmitContainerStateChangeResponse struct {
	Acknowledgment aws.StringValue `query:"acknowledgment" xml:"SubmitContainerStateChangeResult>acknowledgment"`
}

// SubmitTaskStateChangeRequest is undocumented.
type SubmitTaskStateChangeRequest struct {
	Cluster aws.StringValue `query:"cluster" xml:"cluster"`
	Reason  aws.StringValue `query:"reason" xml:"reason"`
	Status  aws.StringValue `query:"status" xml:"status"`
	Task    aws.StringValue `query:"task" xml:"task"`
}

// SubmitTaskStateChangeResponse is undocumented.
type SubmitTaskStateChangeResponse struct {
	Acknowledgment aws.StringValue `query:"acknowledgment" xml:"SubmitTaskStateChangeResult>acknowledgment"`
}

// Task is undocumented.
type Task struct {
	ClusterARN           aws.StringValue `query:"clusterArn" xml:"clusterArn"`
	ContainerInstanceARN aws.StringValue `query:"containerInstanceArn" xml:"containerInstanceArn"`
	Containers           []Container     `query:"containers.member" xml:"containers>member"`
	DesiredStatus        aws.StringValue `query:"desiredStatus" xml:"desiredStatus"`
	LastStatus           aws.StringValue `query:"lastStatus" xml:"lastStatus"`
	Overrides            *TaskOverride   `query:"overrides" xml:"overrides"`
	TaskARN              aws.StringValue `query:"taskArn" xml:"taskArn"`
	TaskDefinitionARN    aws.StringValue `query:"taskDefinitionArn" xml:"taskDefinitionArn"`
}

// TaskDefinition is undocumented.
type TaskDefinition struct {
	ContainerDefinitions []ContainerDefinition `query:"containerDefinitions.member" xml:"containerDefinitions>member"`
	Family               aws.StringValue       `query:"family" xml:"family"`
	Revision             aws.IntegerValue      `query:"revision" xml:"revision"`
	TaskDefinitionARN    aws.StringValue       `query:"taskDefinitionArn" xml:"taskDefinitionArn"`
}

// TaskOverride is undocumented.
type TaskOverride struct {
	ContainerOverrides []ContainerOverride `query:"containerOverrides.member" xml:"containerOverrides>member"`
}

// CreateClusterResult is a wrapper for CreateClusterResponse.
type CreateClusterResult struct {
	Cluster *Cluster `query:"cluster" xml:"CreateClusterResult>cluster"`
}

// DeleteClusterResult is a wrapper for DeleteClusterResponse.
type DeleteClusterResult struct {
	Cluster *Cluster `query:"cluster" xml:"DeleteClusterResult>cluster"`
}

// DeregisterContainerInstanceResult is a wrapper for DeregisterContainerInstanceResponse.
type DeregisterContainerInstanceResult struct {
	ContainerInstance *ContainerInstance `query:"containerInstance" xml:"DeregisterContainerInstanceResult>containerInstance"`
}

// DeregisterTaskDefinitionResult is a wrapper for DeregisterTaskDefinitionResponse.
type DeregisterTaskDefinitionResult struct {
	TaskDefinition *TaskDefinition `query:"taskDefinition" xml:"DeregisterTaskDefinitionResult>taskDefinition"`
}

// DescribeClustersResult is a wrapper for DescribeClustersResponse.
type DescribeClustersResult struct {
	Clusters []Cluster `query:"clusters.member" xml:"DescribeClustersResult>clusters>member"`
	Failures []Failure `query:"failures.member" xml:"DescribeClustersResult>failures>member"`
}

// DescribeContainerInstancesResult is a wrapper for DescribeContainerInstancesResponse.
type DescribeContainerInstancesResult struct {
	ContainerInstances []ContainerInstance `query:"containerInstances.member" xml:"DescribeContainerInstancesResult>containerInstances>member"`
	Failures           []Failure           `query:"failures.member" xml:"DescribeContainerInstancesResult>failures>member"`
}

// DescribeTaskDefinitionResult is a wrapper for DescribeTaskDefinitionResponse.
type DescribeTaskDefinitionResult struct {
	TaskDefinition *TaskDefinition `query:"taskDefinition" xml:"DescribeTaskDefinitionResult>taskDefinition"`
}

// DescribeTasksResult is a wrapper for DescribeTasksResponse.
type DescribeTasksResult struct {
	Failures []Failure `query:"failures.member" xml:"DescribeTasksResult>failures>member"`
	Tasks    []Task    `query:"tasks.member" xml:"DescribeTasksResult>tasks>member"`
}

// DiscoverPollEndpointResult is a wrapper for DiscoverPollEndpointResponse.
type DiscoverPollEndpointResult struct {
	Endpoint aws.StringValue `query:"endpoint" xml:"DiscoverPollEndpointResult>endpoint"`
}

// ListClustersResult is a wrapper for ListClustersResponse.
type ListClustersResult struct {
	ClusterARNs []string        `query:"clusterArns.member" xml:"ListClustersResult>clusterArns>member"`
	NextToken   aws.StringValue `query:"nextToken" xml:"ListClustersResult>nextToken"`
}

// ListContainerInstancesResult is a wrapper for ListContainerInstancesResponse.
type ListContainerInstancesResult struct {
	ContainerInstanceARNs []string        `query:"containerInstanceArns.member" xml:"ListContainerInstancesResult>containerInstanceArns>member"`
	NextToken             aws.StringValue `query:"nextToken" xml:"ListContainerInstancesResult>nextToken"`
}

// ListTaskDefinitionsResult is a wrapper for ListTaskDefinitionsResponse.
type ListTaskDefinitionsResult struct {
	NextToken          aws.StringValue `query:"nextToken" xml:"ListTaskDefinitionsResult>nextToken"`
	TaskDefinitionARNs []string        `query:"taskDefinitionArns.member" xml:"ListTaskDefinitionsResult>taskDefinitionArns>member"`
}

// ListTasksResult is a wrapper for ListTasksResponse.
type ListTasksResult struct {
	NextToken aws.StringValue `query:"nextToken" xml:"ListTasksResult>nextToken"`
	TaskARNs  []string        `query:"taskArns.member" xml:"ListTasksResult>taskArns>member"`
}

// RegisterContainerInstanceResult is a wrapper for RegisterContainerInstanceResponse.
type RegisterContainerInstanceResult struct {
	ContainerInstance *ContainerInstance `query:"containerInstance" xml:"RegisterContainerInstanceResult>containerInstance"`
}

// RegisterTaskDefinitionResult is a wrapper for RegisterTaskDefinitionResponse.
type RegisterTaskDefinitionResult struct {
	TaskDefinition *TaskDefinition `query:"taskDefinition" xml:"RegisterTaskDefinitionResult>taskDefinition"`
}

// RunTaskResult is a wrapper for RunTaskResponse.
type RunTaskResult struct {
	Failures []Failure `query:"failures.member" xml:"RunTaskResult>failures>member"`
	Tasks    []Task    `query:"tasks.member" xml:"RunTaskResult>tasks>member"`
}

// StartTaskResult is a wrapper for StartTaskResponse.
type StartTaskResult struct {
	Failures []Failure `query:"failures.member" xml:"StartTaskResult>failures>member"`
	Tasks    []Task    `query:"tasks.member" xml:"StartTaskResult>tasks>member"`
}

// StopTaskResult is a wrapper for StopTaskResponse.
type StopTaskResult struct {
	Task *Task `query:"task" xml:"StopTaskResult>task"`
}

// SubmitContainerStateChangeResult is a wrapper for SubmitContainerStateChangeResponse.
type SubmitContainerStateChangeResult struct {
	Acknowledgment aws.StringValue `query:"acknowledgment" xml:"SubmitContainerStateChangeResult>acknowledgment"`
}

// SubmitTaskStateChangeResult is a wrapper for SubmitTaskStateChangeResponse.
type SubmitTaskStateChangeResult struct {
	Acknowledgment aws.StringValue `query:"acknowledgment" xml:"SubmitTaskStateChangeResult>acknowledgment"`
}

// avoid errors if the packages aren't referenced
var _ time.Time

var _ xml.Decoder
var _ = io.EOF
