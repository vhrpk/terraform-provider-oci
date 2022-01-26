// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

package devops

import (
	"github.com/oracle/oci-go-sdk/v56/common"
	"net/http"
)

// ListRepositoriesRequest wrapper for the ListRepositories operation
//
// See also
//
// Click https://docs.cloud.oracle.com/en-us/iaas/tools/go-sdk-examples/latest/devops/ListRepositories.go.html to see an example of how to use ListRepositoriesRequest.
type ListRepositoriesRequest struct {

	// The OCID of the compartment in which to list resources.
	CompartmentId *string `mandatory:"false" contributesTo:"query" name:"compartmentId"`

	// unique project identifier
	ProjectId *string `mandatory:"false" contributesTo:"query" name:"projectId"`

	// Unique repository identifier.
	RepositoryId *string `mandatory:"false" contributesTo:"query" name:"repositoryId"`

	// A filter to return only resources whose lifecycle state matches the given lifecycle state.
	LifecycleState RepositoryLifecycleStateEnum `mandatory:"false" contributesTo:"query" name:"lifecycleState" omitEmpty:"true"`

	// A filter to return only resources that match the entire name given.
	Name *string `mandatory:"false" contributesTo:"query" name:"name"`

	// The maximum number of items to return.
	Limit *int `mandatory:"false" contributesTo:"query" name:"limit"`

	// The page token representing the page at which to start retrieving results. This is usually retrieved from a previous list call.
	Page *string `mandatory:"false" contributesTo:"query" name:"page"`

	// The sort order to use. Use either ascending or descending.
	SortOrder ListRepositoriesSortOrderEnum `mandatory:"false" contributesTo:"query" name:"sortOrder" omitEmpty:"true"`

	// The field to sort by. Only one sort order may be provided. Default order for time created is descending. Default order for name is ascending. If no value is specified time created is default.
	SortBy ListRepositoriesSortByEnum `mandatory:"false" contributesTo:"query" name:"sortBy" omitEmpty:"true"`

	// Unique Oracle-assigned identifier for the request.  If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `mandatory:"false" contributesTo:"header" name:"opc-request-id"`

	// Metadata about the request. This information will not be transmitted to the service, but
	// represents information that the SDK will consume to drive retry behavior.
	RequestMetadata common.RequestMetadata
}

func (request ListRepositoriesRequest) String() string {
	return common.PointerString(request)
}

// HTTPRequest implements the OCIRequest interface
func (request ListRepositoriesRequest) HTTPRequest(method, path string, binaryRequestBody *common.OCIReadSeekCloser, extraHeaders map[string]string) (http.Request, error) {

	return common.MakeDefaultHTTPRequestWithTaggedStructAndExtraHeaders(method, path, request, extraHeaders)
}

// BinaryRequestBody implements the OCIRequest interface
func (request ListRepositoriesRequest) BinaryRequestBody() (*common.OCIReadSeekCloser, bool) {

	return nil, false

}

// RetryPolicy implements the OCIRetryableRequest interface. This retrieves the specified retry policy.
func (request ListRepositoriesRequest) RetryPolicy() *common.RetryPolicy {
	return request.RequestMetadata.RetryPolicy
}

// ListRepositoriesResponse wrapper for the ListRepositories operation
type ListRepositoriesResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// A list of RepositoryCollection instances
	RepositoryCollection `presentIn:"body"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about a particular request, provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`

	// For pagination of a list of items. When paging through a list, if this header appears in the response, then a partial list might have been returned. Include this value as the `page` parameter for the subsequent GET request to get the next batch of items.
	OpcNextPage *string `presentIn:"header" name:"opc-next-page"`
}

func (response ListRepositoriesResponse) String() string {
	return common.PointerString(response)
}

// HTTPResponse implements the OCIResponse interface
func (response ListRepositoriesResponse) HTTPResponse() *http.Response {
	return response.RawResponse
}

// ListRepositoriesSortOrderEnum Enum with underlying type: string
type ListRepositoriesSortOrderEnum string

// Set of constants representing the allowable values for ListRepositoriesSortOrderEnum
const (
	ListRepositoriesSortOrderAsc  ListRepositoriesSortOrderEnum = "ASC"
	ListRepositoriesSortOrderDesc ListRepositoriesSortOrderEnum = "DESC"
)

var mappingListRepositoriesSortOrder = map[string]ListRepositoriesSortOrderEnum{
	"ASC":  ListRepositoriesSortOrderAsc,
	"DESC": ListRepositoriesSortOrderDesc,
}

// GetListRepositoriesSortOrderEnumValues Enumerates the set of values for ListRepositoriesSortOrderEnum
func GetListRepositoriesSortOrderEnumValues() []ListRepositoriesSortOrderEnum {
	values := make([]ListRepositoriesSortOrderEnum, 0)
	for _, v := range mappingListRepositoriesSortOrder {
		values = append(values, v)
	}
	return values
}

// ListRepositoriesSortByEnum Enum with underlying type: string
type ListRepositoriesSortByEnum string

// Set of constants representing the allowable values for ListRepositoriesSortByEnum
const (
	ListRepositoriesSortByTimecreated ListRepositoriesSortByEnum = "timeCreated"
	ListRepositoriesSortByName        ListRepositoriesSortByEnum = "name"
)

var mappingListRepositoriesSortBy = map[string]ListRepositoriesSortByEnum{
	"timeCreated": ListRepositoriesSortByTimecreated,
	"name":        ListRepositoriesSortByName,
}

// GetListRepositoriesSortByEnumValues Enumerates the set of values for ListRepositoriesSortByEnum
func GetListRepositoriesSortByEnumValues() []ListRepositoriesSortByEnum {
	values := make([]ListRepositoriesSortByEnum, 0)
	for _, v := range mappingListRepositoriesSortBy {
		values = append(values, v)
	}
	return values
}
