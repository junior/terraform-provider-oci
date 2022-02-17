// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// UpdateMaintenanceRunDetails Describes the modification parameters for the maintenance run.
type UpdateMaintenanceRunDetails struct {

	// If `FALSE`, skips the maintenance run.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// The scheduled date and time of the maintenance run to update.
	TimeScheduled *common.SDKTime `mandatory:"false" json:"timeScheduled"`

	// If set to `TRUE`, starts patching immediately.
	IsPatchNowEnabled *bool `mandatory:"false" json:"isPatchNowEnabled"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the patch to be applied in the maintenance run.
	PatchId *string `mandatory:"false" json:"patchId"`

	// Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.
	// *IMPORTANT*: Non-rolling infrastructure patching involves system down time. See Oracle-Managed Infrastructure Maintenance Updates (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle) for more information.
	PatchingMode UpdateMaintenanceRunDetailsPatchingModeEnum `mandatory:"false" json:"patchingMode,omitempty"`
}

func (m UpdateMaintenanceRunDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateMaintenanceRunDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateMaintenanceRunDetailsPatchingModeEnum(string(m.PatchingMode)); !ok && m.PatchingMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PatchingMode: %s. Supported values are: %s.", m.PatchingMode, strings.Join(GetUpdateMaintenanceRunDetailsPatchingModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateMaintenanceRunDetailsPatchingModeEnum Enum with underlying type: string
type UpdateMaintenanceRunDetailsPatchingModeEnum string

// Set of constants representing the allowable values for UpdateMaintenanceRunDetailsPatchingModeEnum
const (
	UpdateMaintenanceRunDetailsPatchingModeRolling    UpdateMaintenanceRunDetailsPatchingModeEnum = "ROLLING"
	UpdateMaintenanceRunDetailsPatchingModeNonrolling UpdateMaintenanceRunDetailsPatchingModeEnum = "NONROLLING"
)

var mappingUpdateMaintenanceRunDetailsPatchingModeEnum = map[string]UpdateMaintenanceRunDetailsPatchingModeEnum{
	"ROLLING":    UpdateMaintenanceRunDetailsPatchingModeRolling,
	"NONROLLING": UpdateMaintenanceRunDetailsPatchingModeNonrolling,
}

// GetUpdateMaintenanceRunDetailsPatchingModeEnumValues Enumerates the set of values for UpdateMaintenanceRunDetailsPatchingModeEnum
func GetUpdateMaintenanceRunDetailsPatchingModeEnumValues() []UpdateMaintenanceRunDetailsPatchingModeEnum {
	values := make([]UpdateMaintenanceRunDetailsPatchingModeEnum, 0)
	for _, v := range mappingUpdateMaintenanceRunDetailsPatchingModeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateMaintenanceRunDetailsPatchingModeEnumStringValues Enumerates the set of values in String for UpdateMaintenanceRunDetailsPatchingModeEnum
func GetUpdateMaintenanceRunDetailsPatchingModeEnumStringValues() []string {
	return []string{
		"ROLLING",
		"NONROLLING",
	}
}

// GetMappingUpdateMaintenanceRunDetailsPatchingModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateMaintenanceRunDetailsPatchingModeEnum(val string) (UpdateMaintenanceRunDetailsPatchingModeEnum, bool) {
	mappingUpdateMaintenanceRunDetailsPatchingModeEnumIgnoreCase := make(map[string]UpdateMaintenanceRunDetailsPatchingModeEnum)
	for k, v := range mappingUpdateMaintenanceRunDetailsPatchingModeEnum {
		mappingUpdateMaintenanceRunDetailsPatchingModeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUpdateMaintenanceRunDetailsPatchingModeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
