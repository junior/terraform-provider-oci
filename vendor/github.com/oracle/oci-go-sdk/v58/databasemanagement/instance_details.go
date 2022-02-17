// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// InstanceDetails The details of the Oracle Real Application Clusters (Oracle RAC) database instance.
type InstanceDetails struct {

	// The ID of the Oracle RAC database instance.
	Id *int `mandatory:"true" json:"id"`

	// The name of the Oracle RAC database instance.
	Name *string `mandatory:"true" json:"name"`

	// The name of the host of the Oracle RAC database instance.
	HostName *string `mandatory:"true" json:"hostName"`

	// The status of the Oracle RAC database instance.
	Status InstanceDetailsStatusEnum `mandatory:"true" json:"status"`
}

func (m InstanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingInstanceDetailsStatusEnum(string(m.Status)); !ok && m.Status != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Status: %s. Supported values are: %s.", m.Status, strings.Join(GetInstanceDetailsStatusEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InstanceDetailsStatusEnum Enum with underlying type: string
type InstanceDetailsStatusEnum string

// Set of constants representing the allowable values for InstanceDetailsStatusEnum
const (
	InstanceDetailsStatusUp      InstanceDetailsStatusEnum = "UP"
	InstanceDetailsStatusDown    InstanceDetailsStatusEnum = "DOWN"
	InstanceDetailsStatusUnknown InstanceDetailsStatusEnum = "UNKNOWN"
)

var mappingInstanceDetailsStatusEnum = map[string]InstanceDetailsStatusEnum{
	"UP":      InstanceDetailsStatusUp,
	"DOWN":    InstanceDetailsStatusDown,
	"UNKNOWN": InstanceDetailsStatusUnknown,
}

// GetInstanceDetailsStatusEnumValues Enumerates the set of values for InstanceDetailsStatusEnum
func GetInstanceDetailsStatusEnumValues() []InstanceDetailsStatusEnum {
	values := make([]InstanceDetailsStatusEnum, 0)
	for _, v := range mappingInstanceDetailsStatusEnum {
		values = append(values, v)
	}
	return values
}

// GetInstanceDetailsStatusEnumStringValues Enumerates the set of values in String for InstanceDetailsStatusEnum
func GetInstanceDetailsStatusEnumStringValues() []string {
	return []string{
		"UP",
		"DOWN",
		"UNKNOWN",
	}
}

// GetMappingInstanceDetailsStatusEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInstanceDetailsStatusEnum(val string) (InstanceDetailsStatusEnum, bool) {
	mappingInstanceDetailsStatusEnumIgnoreCase := make(map[string]InstanceDetailsStatusEnum)
	for k, v := range mappingInstanceDetailsStatusEnum {
		mappingInstanceDetailsStatusEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingInstanceDetailsStatusEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
