// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud AI Services API
//
// OCI AI Service solutions can help Enterprise customers integrate AI into their products immediately by using our proven,
// pre-trained/custom models or containers, and without a need to set up in house team of AI and ML experts.
// This allows enterprises to focus on business drivers and development work rather than AI/ML operations, shortening the time to market.
//

package aianomalydetection

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// DetectAnomaliesDetails Base class for the DetectAnomalies call. It contains the identifier that will
// be used for deciding what type of request this is.
type DetectAnomaliesDetails interface {

	// The OCID of the trained model。
	GetModelId() *string
}

type detectanomaliesdetails struct {
	JsonData    []byte
	ModelId     *string `mandatory:"true" json:"modelId"`
	RequestType string  `json:"requestType"`
}

// UnmarshalJSON unmarshals json
func (m *detectanomaliesdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerdetectanomaliesdetails detectanomaliesdetails
	s := struct {
		Model Unmarshalerdetectanomaliesdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ModelId = s.Model.ModelId
	m.RequestType = s.Model.RequestType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *detectanomaliesdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.RequestType {
	case "INLINE":
		mm := InlineDetectAnomaliesRequest{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "BASE64_ENCODED":
		mm := EmbeddedDetectAnomaliesRequest{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetModelId returns ModelId
func (m detectanomaliesdetails) GetModelId() *string {
	return m.ModelId
}

func (m detectanomaliesdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m detectanomaliesdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DetectAnomaliesDetailsRequestTypeEnum Enum with underlying type: string
type DetectAnomaliesDetailsRequestTypeEnum string

// Set of constants representing the allowable values for DetectAnomaliesDetailsRequestTypeEnum
const (
	DetectAnomaliesDetailsRequestTypeInline        DetectAnomaliesDetailsRequestTypeEnum = "INLINE"
	DetectAnomaliesDetailsRequestTypeBase64Encoded DetectAnomaliesDetailsRequestTypeEnum = "BASE64_ENCODED"
)

var mappingDetectAnomaliesDetailsRequestTypeEnum = map[string]DetectAnomaliesDetailsRequestTypeEnum{
	"INLINE":         DetectAnomaliesDetailsRequestTypeInline,
	"BASE64_ENCODED": DetectAnomaliesDetailsRequestTypeBase64Encoded,
}

// GetDetectAnomaliesDetailsRequestTypeEnumValues Enumerates the set of values for DetectAnomaliesDetailsRequestTypeEnum
func GetDetectAnomaliesDetailsRequestTypeEnumValues() []DetectAnomaliesDetailsRequestTypeEnum {
	values := make([]DetectAnomaliesDetailsRequestTypeEnum, 0)
	for _, v := range mappingDetectAnomaliesDetailsRequestTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetDetectAnomaliesDetailsRequestTypeEnumStringValues Enumerates the set of values in String for DetectAnomaliesDetailsRequestTypeEnum
func GetDetectAnomaliesDetailsRequestTypeEnumStringValues() []string {
	return []string{
		"INLINE",
		"BASE64_ENCODED",
	}
}

// GetMappingDetectAnomaliesDetailsRequestTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDetectAnomaliesDetailsRequestTypeEnum(val string) (DetectAnomaliesDetailsRequestTypeEnum, bool) {
	mappingDetectAnomaliesDetailsRequestTypeEnumIgnoreCase := make(map[string]DetectAnomaliesDetailsRequestTypeEnum)
	for k, v := range mappingDetectAnomaliesDetailsRequestTypeEnum {
		mappingDetectAnomaliesDetailsRequestTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingDetectAnomaliesDetailsRequestTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
