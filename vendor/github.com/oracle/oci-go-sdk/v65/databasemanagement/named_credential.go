// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// NamedCredential The details of a named credential.
type NamedCredential struct {

	// The name of the named credential.
	Name *string `mandatory:"true" json:"name"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the named credential.
	Id *string `mandatory:"true" json:"id"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The current lifecycle state of the named credential.
	LifecycleState LifecycleStatesEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the named credential was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The information specified by the user about the named credential.
	Description *string `mandatory:"false" json:"description"`

	// The details of the lifecycle state.
	LifecycleDetails *string `mandatory:"false" json:"lifecycleDetails"`

	// The scope of the named credential.
	Scope NamedCredentialScopeEnum `mandatory:"false" json:"scope,omitempty"`

	// The type of resource associated with the named credential.
	Type ResourceTypeEnum `mandatory:"false" json:"type,omitempty"`

	Content NamedCredentialContent `mandatory:"false" json:"content"`

	// The date and time the named credential was last updated.
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the resource that
	// is associated to the named credential.
	AssociatedResource *string `mandatory:"false" json:"associatedResource"`
}

func (m NamedCredential) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NamedCredential) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLifecycleStatesEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStatesEnumStringValues(), ",")))
	}

	if _, ok := GetMappingNamedCredentialScopeEnum(string(m.Scope)); !ok && m.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", m.Scope, strings.Join(GetNamedCredentialScopeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingResourceTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetResourceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *NamedCredential) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Description        *string                  `json:"description"`
		LifecycleDetails   *string                  `json:"lifecycleDetails"`
		Scope              NamedCredentialScopeEnum `json:"scope"`
		Type               ResourceTypeEnum         `json:"type"`
		Content            namedcredentialcontent   `json:"content"`
		TimeUpdated        *common.SDKTime          `json:"timeUpdated"`
		AssociatedResource *string                  `json:"associatedResource"`
		Name               *string                  `json:"name"`
		Id                 *string                  `json:"id"`
		CompartmentId      *string                  `json:"compartmentId"`
		LifecycleState     LifecycleStatesEnum      `json:"lifecycleState"`
		TimeCreated        *common.SDKTime          `json:"timeCreated"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	m.Description = model.Description

	m.LifecycleDetails = model.LifecycleDetails

	m.Scope = model.Scope

	m.Type = model.Type

	nn, e = model.Content.UnmarshalPolymorphicJSON(model.Content.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Content = nn.(NamedCredentialContent)
	} else {
		m.Content = nil
	}

	m.TimeUpdated = model.TimeUpdated

	m.AssociatedResource = model.AssociatedResource

	m.Name = model.Name

	m.Id = model.Id

	m.CompartmentId = model.CompartmentId

	m.LifecycleState = model.LifecycleState

	m.TimeCreated = model.TimeCreated

	return
}
