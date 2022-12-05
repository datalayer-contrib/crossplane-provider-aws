/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

type ContainerProviderType string

const (
	ContainerProviderType_EKS ContainerProviderType = "EKS"
)

type EndpointState string

const (
	EndpointState_CREATING               EndpointState = "CREATING"
	EndpointState_ACTIVE                 EndpointState = "ACTIVE"
	EndpointState_TERMINATING            EndpointState = "TERMINATING"
	EndpointState_TERMINATED             EndpointState = "TERMINATED"
	EndpointState_TERMINATED_WITH_ERRORS EndpointState = "TERMINATED_WITH_ERRORS"
)

type FailureReason string

const (
	FailureReason_INTERNAL_ERROR      FailureReason = "INTERNAL_ERROR"
	FailureReason_USER_ERROR          FailureReason = "USER_ERROR"
	FailureReason_VALIDATION_ERROR    FailureReason = "VALIDATION_ERROR"
	FailureReason_CLUSTER_UNAVAILABLE FailureReason = "CLUSTER_UNAVAILABLE"
)

type JobRunState string

const (
	JobRunState_PENDING        JobRunState = "PENDING"
	JobRunState_SUBMITTED      JobRunState = "SUBMITTED"
	JobRunState_RUNNING        JobRunState = "RUNNING"
	JobRunState_FAILED         JobRunState = "FAILED"
	JobRunState_CANCELLED      JobRunState = "CANCELLED"
	JobRunState_CANCEL_PENDING JobRunState = "CANCEL_PENDING"
	JobRunState_COMPLETED      JobRunState = "COMPLETED"
)

type PersistentAppUI string

const (
	PersistentAppUI_ENABLED  PersistentAppUI = "ENABLED"
	PersistentAppUI_DISABLED PersistentAppUI = "DISABLED"
)

type VirtualClusterState string

const (
	VirtualClusterState_RUNNING     VirtualClusterState = "RUNNING"
	VirtualClusterState_TERMINATING VirtualClusterState = "TERMINATING"
	VirtualClusterState_TERMINATED  VirtualClusterState = "TERMINATED"
	VirtualClusterState_ARRESTED    VirtualClusterState = "ARRESTED"
)
