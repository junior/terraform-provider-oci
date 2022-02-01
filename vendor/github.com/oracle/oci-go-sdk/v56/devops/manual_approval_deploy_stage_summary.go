// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ManualApprovalDeployStageSummary Specifies the manual approval stage.
type ManualApprovalDeployStageSummary struct {

	// Unique identifier that is immutable on creation.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of a project.
	ProjectId *string `mandatory:"true" json:"projectId"`

	// The OCID of a pipeline.
	DeployPipelineId *string `mandatory:"true" json:"deployPipelineId"`

	// The OCID of a compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	ApprovalPolicy ApprovalPolicy `mandatory:"true" json:"approvalPolicy"`

	// Optional description about the deployment stage.
	Description *string `mandatory:"false" json:"description"`

	// Deployment stage display name, which can be renamed and is not necessarily unique. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Time the deployment stage was created. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time the deployment stage was updated. Format defined by RFC3339 (https://datatracker.ietf.org/doc/html/rfc3339).
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	DeployStagePredecessorCollection *DeployStagePredecessorCollection `mandatory:"false" json:"deployStagePredecessorCollection"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace": {"bar-key": "value"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces. See Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`

	// The current state of the deployment stage.
	LifecycleState DeployStageLifecycleStateEnum `mandatory:"false" json:"lifecycleState,omitempty"`
}

//GetId returns Id
func (m ManualApprovalDeployStageSummary) GetId() *string {
	return m.Id
}

//GetDescription returns Description
func (m ManualApprovalDeployStageSummary) GetDescription() *string {
	return m.Description
}

//GetDisplayName returns DisplayName
func (m ManualApprovalDeployStageSummary) GetDisplayName() *string {
	return m.DisplayName
}

//GetProjectId returns ProjectId
func (m ManualApprovalDeployStageSummary) GetProjectId() *string {
	return m.ProjectId
}

//GetDeployPipelineId returns DeployPipelineId
func (m ManualApprovalDeployStageSummary) GetDeployPipelineId() *string {
	return m.DeployPipelineId
}

//GetCompartmentId returns CompartmentId
func (m ManualApprovalDeployStageSummary) GetCompartmentId() *string {
	return m.CompartmentId
}

//GetTimeCreated returns TimeCreated
func (m ManualApprovalDeployStageSummary) GetTimeCreated() *common.SDKTime {
	return m.TimeCreated
}

//GetTimeUpdated returns TimeUpdated
func (m ManualApprovalDeployStageSummary) GetTimeUpdated() *common.SDKTime {
	return m.TimeUpdated
}

//GetLifecycleState returns LifecycleState
func (m ManualApprovalDeployStageSummary) GetLifecycleState() DeployStageLifecycleStateEnum {
	return m.LifecycleState
}

//GetLifecycleDetails returns LifecycleDetails
func (m ManualApprovalDeployStageSummary) GetLifecycleDetails() *string {
	return m.LifecycleDetails
}

//GetDeployStagePredecessorCollection returns DeployStagePredecessorCollection
func (m ManualApprovalDeployStageSummary) GetDeployStagePredecessorCollection() *DeployStagePredecessorCollection {
	return m.DeployStagePredecessorCollection
}

//GetFreeformTags returns FreeformTags
func (m ManualApprovalDeployStageSummary) GetFreeformTags() map[string]string {
	return m.FreeformTags
}

//GetDefinedTags returns DefinedTags
func (m ManualApprovalDeployStageSummary) GetDefinedTags() map[string]map[string]interface{} {
	return m.DefinedTags
}

//GetSystemTags returns SystemTags
func (m ManualApprovalDeployStageSummary) GetSystemTags() map[string]map[string]interface{} {
	return m.SystemTags
}

func (m ManualApprovalDeployStageSummary) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ManualApprovalDeployStageSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeManualApprovalDeployStageSummary ManualApprovalDeployStageSummary
	s := struct {
		DiscriminatorParam string `json:"deployStageType"`
		MarshalTypeManualApprovalDeployStageSummary
	}{
		"MANUAL_APPROVAL",
		(MarshalTypeManualApprovalDeployStageSummary)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ManualApprovalDeployStageSummary) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description                      *string                           `json:"description"`
		DisplayName                      *string                           `json:"displayName"`
		TimeCreated                      *common.SDKTime                   `json:"timeCreated"`
		TimeUpdated                      *common.SDKTime                   `json:"timeUpdated"`
		LifecycleState                   DeployStageLifecycleStateEnum     `json:"lifecycleState"`
		LifecycleDetails                 *string                           `json:"lifecycleDetails"`
		DeployStagePredecessorCollection *DeployStagePredecessorCollection `json:"deployStagePredecessorCollection"`
		FreeformTags                     map[string]string                 `json:"freeformTags"`
		DefinedTags                      map[string]map[string]interface{} `json:"definedTags"`
		SystemTags                       map[string]map[string]interface{} `json:"systemTags"`
		Id                               *string                           `json:"id"`
		ProjectId                        *string                           `json:"projectId"`
		DeployPipelineId                 *string                           `json:"deployPipelineId"`
		CompartmentId                    *string                           `json:"compartmentId"`
		ApprovalPolicy                   approvalpolicy                    `json:"approvalPolicy"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.DisplayName = model.DisplayName

	m.TimeCreated = model.TimeCreated

	m.TimeUpdated = model.TimeUpdated

	m.LifecycleState = model.LifecycleState

	m.LifecycleDetails = model.LifecycleDetails

	m.DeployStagePredecessorCollection = model.DeployStagePredecessorCollection

	m.FreeformTags = model.FreeformTags

	m.DefinedTags = model.DefinedTags

	m.SystemTags = model.SystemTags

	m.Id = model.Id

	m.ProjectId = model.ProjectId

	m.DeployPipelineId = model.DeployPipelineId

	m.CompartmentId = model.CompartmentId

	nn, e = model.ApprovalPolicy.UnmarshalPolymorphicJSON(model.ApprovalPolicy.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.ApprovalPolicy = nn.(ApprovalPolicy)
	} else {
		m.ApprovalPolicy = nil
	}

	return
}
