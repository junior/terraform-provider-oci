// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"net/http"
	"strings"
)

// ListLogSetsRequest wrapper for the ListLogSets operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/loganalytics/ListLogSets.go.html to see an example of how to use ListLogSetsRequest.
type ListLogSetsRequest struct {

	// The Logging Analytics namespace used for the request.
	NamespaceName *string `mandatory:"true" contributesTo:"path" name:"namespaceName"`

	// The client request ID for tracing.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use, either ascending (`ASC`) or descending (`DESC`).
	SortOrder ListLogSetsSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// If this filter is present, each of the logsets returned must contain the value of this filter.
	LogSetNameContains []string `contributesTo:"query" name:"logSetNameContains" collectionFormat:"multi"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListLogSetsRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListLogSetsRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	_, err := request.ValidateEnumValue()
	if err != nil {
		return http.Request{}, err
	}
	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListLogSetsRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListLogSetsRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (request ListLogSetsRequest) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingListLogSetsSortOrderEnum(string(request.SortOrder)); !ok && request.SortOrder != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SortOrder: %s. Supported values are: %s.", request.SortOrder, strings.Join(GetListLogSetsSortOrderEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ListLogSetsResponse wrapper for the ListLogSets operation
type ListLogSetsResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of LogSetCollection instances
	LogSetCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. When you contact Oracle about a specific request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response,
	// then additional items may be available on the next page of the list. Include this value as the `page` parameter for the
	// subsequent request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListLogSetsResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListLogSetsResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListLogSetsSortOrderEnum Enum with underlying type: string
type ListLogSetsSortOrderEnum string

// Set of constants representing the allowable values for ListLogSetsSortOrderEnum
const (
	ListLogSetsSortOrderAsc  ListLogSetsSortOrderEnum = "ASC"
	ListLogSetsSortOrderDesc ListLogSetsSortOrderEnum = "DESC"
)

var mappingListLogSetsSortOrderEnum = map[string]ListLogSetsSortOrderEnum{
	"ASC":  ListLogSetsSortOrderAsc,
	"DESC": ListLogSetsSortOrderDesc,
}

// GetListLogSetsSortOrderEnumValues Enumerates the set of values for ListLogSetsSortOrderEnum
func GetListLogSetsSortOrderEnumValues() []ListLogSetsSortOrderEnum {
	values := make([]ListLogSetsSortOrderEnum, 0)
	for _, v := range mappingListLogSetsSortOrderEnum {
		values = append(values, v)
	}
	return values
}

// GetListLogSetsSortOrderEnumStringValues Enumerates the set of values in String for ListLogSetsSortOrderEnum
func GetListLogSetsSortOrderEnumStringValues() []string {
	return []string{
		"ASC",
		"DESC",
	}
}

// GetMappingListLogSetsSortOrderEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingListLogSetsSortOrderEnum(val string) (ListLogSetsSortOrderEnum, bool) {
	mappingListLogSetsSortOrderEnumIgnoreCase := make(map[string]ListLogSetsSortOrderEnum)
	for k, v := range mappingListLogSetsSortOrderEnum {
		mappingListLogSetsSortOrderEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingListLogSetsSortOrderEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
