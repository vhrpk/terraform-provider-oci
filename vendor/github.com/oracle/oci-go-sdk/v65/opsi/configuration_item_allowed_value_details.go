// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConfigurationItemAllowedValueDetails Allowed value details of configuration item, to validate what value can be assigned to a configuration item.
type ConfigurationItemAllowedValueDetails interface {
}

type configurationitemallowedvaluedetails struct {
	JsonData         []byte
	AllowedValueType string `json:"allowedValueType"`
}

// UnmarshalJSON unmarshals json
func (m *configurationitemallowedvaluedetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerconfigurationitemallowedvaluedetails configurationitemallowedvaluedetails
	s := struct {
		Model Unmarshalerconfigurationitemallowedvaluedetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.AllowedValueType = s.Model.AllowedValueType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *configurationitemallowedvaluedetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.AllowedValueType {
	case "FREE_TEXT":
		mm := ConfigurationItemFreeTextAllowedValueDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "PICK":
		mm := ConfigurationItemPickAllowedValueDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "LIMIT":
		mm := ConfigurationItemLimitAllowedValueDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

func (m configurationitemallowedvaluedetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m configurationitemallowedvaluedetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
