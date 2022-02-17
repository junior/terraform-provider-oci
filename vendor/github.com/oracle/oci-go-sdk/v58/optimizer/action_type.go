// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// Use the Cloud Advisor API to find potential inefficiencies in your tenancy and address them.
// Cloud Advisor can help you save money, improve performance, strengthen system resilience, and improve security.
// For more information, see Cloud Advisor (https://docs.cloud.oracle.com/Content/CloudAdvisor/Concepts/cloudadvisoroverview.htm).
//

package optimizer

import (
	"strings"
)

// ActionTypeEnum Enum with underlying type: string
type ActionTypeEnum string

// Set of constants representing the allowable values for ActionTypeEnum
const (
	ActionTypeKbArticle ActionTypeEnum = "KB_ARTICLE"
)

var mappingActionTypeEnum = map[string]ActionTypeEnum{
	"KB_ARTICLE": ActionTypeKbArticle,
}

// GetActionTypeEnumValues Enumerates the set of values for ActionTypeEnum
func GetActionTypeEnumValues() []ActionTypeEnum {
	values := make([]ActionTypeEnum, 0)
	for _, v := range mappingActionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetActionTypeEnumStringValues Enumerates the set of values in String for ActionTypeEnum
func GetActionTypeEnumStringValues() []string {
	return []string{
		"KB_ARTICLE",
	}
}

// GetMappingActionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingActionTypeEnum(val string) (ActionTypeEnum, bool) {
	mappingActionTypeEnumIgnoreCase := make(map[string]ActionTypeEnum)
	for k, v := range mappingActionTypeEnum {
		mappingActionTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingActionTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
