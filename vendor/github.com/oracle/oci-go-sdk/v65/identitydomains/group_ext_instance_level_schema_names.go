// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// GroupExtInstanceLevelSchemaNames DBCS instance-level schema-names. Each schema-name is specific to a DB Instance.
// **Added In:** 18.2.4
// **SCIM++ Properties:**
//   - idcsCompositeKey: [dbInstanceId, schemaName]
//   - idcsSearchable: true
//   - multiValued: true
//   - mutability: readOnly
//   - required: false
//   - returned: request
//   - type: complex
type GroupExtInstanceLevelSchemaNames struct {

	// App Id of DBCS App instance
	// **Added In:** 18.2.4
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	DbInstanceId *string `mandatory:"true" json:"dbInstanceId"`

	// The DBCS schema-name granted to this Group for the DB instance that 'dbInstanceId' specifies.
	// **Added In:** 18.2.4
	// **SCIM++ Properties:**
	//  - caseExact: true
	//  - idcsSearchable: true
	//  - multiValued: false
	//  - mutability: readOnly
	//  - required: true
	//  - returned: default
	//  - type: string
	//  - uniqueness: none
	SchemaName *string `mandatory:"true" json:"schemaName"`
}

func (m GroupExtInstanceLevelSchemaNames) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m GroupExtInstanceLevelSchemaNames) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
