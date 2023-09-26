// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Domains API
//
// Use the Identity Domains API to manage resources within an identity domain, for example, users, dynamic resource groups, groups, and identity providers. For information about managing resources within identity domains, see Identity and Access Management (with identity domains) (https://docs.oracle.com/iaas/Content/Identity/home.htm). This REST API is SCIM compliant.
// Use the table of contents and search tool to explore the Identity Domains API.
//

package identitydomains

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AuthenticationFactorSettingsEndpointRestrictions Settings that describe the set of restrictions that the system should apply to devices and trusted endpoints of a user
// **SCIM++ Properties:**
//   - idcsSearchable: false
//   - multiValued: false
//   - mutability: readWrite
//   - required: true
//   - returned: default
//   - type: complex
//   - uniqueness: none
type AuthenticationFactorSettingsEndpointRestrictions struct {

	// Maximum number of enrolled devices per user
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 20
	//  - idcsMinValue: 1
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MaxEnrolledDevices *int `mandatory:"true" json:"maxEnrolledDevices"`

	// Max number of trusted endpoints per user
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 20
	//  - idcsMinValue: 1
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MaxTrustedEndpoints *int `mandatory:"true" json:"maxTrustedEndpoints"`

	// Maximum number of days until an endpoint can be trusted
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 180
	//  - idcsMinValue: 1
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MaxEndpointTrustDurationInDays *int `mandatory:"true" json:"maxEndpointTrustDurationInDays"`

	// Specify if trusted endpoints are enabled
	// **SCIM++ Properties:**
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: boolean
	//  - uniqueness: none
	TrustedEndpointsEnabled *bool `mandatory:"true" json:"trustedEndpointsEnabled"`

	// An integer that represents the maximum number of failed MFA logins before an account is locked
	// **SCIM++ Properties:**
	//  - idcsMaxValue: 20
	//  - idcsMinValue: 5
	//  - idcsSearchable: false
	//  - multiValued: false
	//  - mutability: readWrite
	//  - required: true
	//  - returned: default
	//  - type: integer
	//  - uniqueness: none
	MaxIncorrectAttempts *int `mandatory:"true" json:"maxIncorrectAttempts"`
}

func (m AuthenticationFactorSettingsEndpointRestrictions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AuthenticationFactorSettingsEndpointRestrictions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
