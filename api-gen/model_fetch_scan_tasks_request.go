/*
Apache Iceberg REST Catalog API

Defines the specification for the first version of the REST Catalog API. Implementations should ideally support both Iceberg table specs v1 and v2, with priority given to v2.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package icebergclient

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the FetchScanTasksRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchScanTasksRequest{}

// FetchScanTasksRequest struct for FetchScanTasksRequest
type FetchScanTasksRequest struct {
	// An opaque string provided by the REST server that represents a unit of work to produce file scan tasks for scan planning. This allows clients to fetch tasks across multiple requests to accommodate large result sets.
	PlanTask string `json:"plan-task"`
}

type _FetchScanTasksRequest FetchScanTasksRequest

// NewFetchScanTasksRequest instantiates a new FetchScanTasksRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchScanTasksRequest(planTask string) *FetchScanTasksRequest {
	this := FetchScanTasksRequest{}
	this.PlanTask = planTask
	return &this
}

// NewFetchScanTasksRequestWithDefaults instantiates a new FetchScanTasksRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchScanTasksRequestWithDefaults() *FetchScanTasksRequest {
	this := FetchScanTasksRequest{}
	return &this
}

// GetPlanTask returns the PlanTask field value
func (o *FetchScanTasksRequest) GetPlanTask() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlanTask
}

// GetPlanTaskOk returns a tuple with the PlanTask field value
// and a boolean to check if the value has been set.
func (o *FetchScanTasksRequest) GetPlanTaskOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlanTask, true
}

// SetPlanTask sets field value
func (o *FetchScanTasksRequest) SetPlanTask(v string) {
	o.PlanTask = v
}

func (o FetchScanTasksRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchScanTasksRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plan-task"] = o.PlanTask
	return toSerialize, nil
}

func (o *FetchScanTasksRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"plan-task",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFetchScanTasksRequest := _FetchScanTasksRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varFetchScanTasksRequest)

	if err != nil {
		return err
	}

	*o = FetchScanTasksRequest(varFetchScanTasksRequest)

	return err
}

type NullableFetchScanTasksRequest struct {
	value *FetchScanTasksRequest
	isSet bool
}

func (v NullableFetchScanTasksRequest) Get() *FetchScanTasksRequest {
	return v.value
}

func (v *NullableFetchScanTasksRequest) Set(val *FetchScanTasksRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchScanTasksRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchScanTasksRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchScanTasksRequest(val *FetchScanTasksRequest) *NullableFetchScanTasksRequest {
	return &NullableFetchScanTasksRequest{value: val, isSet: true}
}

func (v NullableFetchScanTasksRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchScanTasksRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


