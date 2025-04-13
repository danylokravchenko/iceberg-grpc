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

// checks if the SortOrder type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SortOrder{}

// SortOrder struct for SortOrder
type SortOrder struct {
	OrderId int32 `json:"order-id"`
	Fields []SortField `json:"fields"`
}

type _SortOrder SortOrder

// NewSortOrder instantiates a new SortOrder object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSortOrder(orderId int32, fields []SortField) *SortOrder {
	this := SortOrder{}
	this.OrderId = orderId
	this.Fields = fields
	return &this
}

// NewSortOrderWithDefaults instantiates a new SortOrder object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSortOrderWithDefaults() *SortOrder {
	this := SortOrder{}
	return &this
}

// GetOrderId returns the OrderId field value
func (o *SortOrder) GetOrderId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.OrderId
}

// GetOrderIdOk returns a tuple with the OrderId field value
// and a boolean to check if the value has been set.
func (o *SortOrder) GetOrderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrderId, true
}

// SetOrderId sets field value
func (o *SortOrder) SetOrderId(v int32) {
	o.OrderId = v
}

// GetFields returns the Fields field value
func (o *SortOrder) GetFields() []SortField {
	if o == nil {
		var ret []SortField
		return ret
	}

	return o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value
// and a boolean to check if the value has been set.
func (o *SortOrder) GetFieldsOk() ([]SortField, bool) {
	if o == nil {
		return nil, false
	}
	return o.Fields, true
}

// SetFields sets field value
func (o *SortOrder) SetFields(v []SortField) {
	o.Fields = v
}

func (o SortOrder) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SortOrder) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["order-id"] = o.OrderId
	toSerialize["fields"] = o.Fields
	return toSerialize, nil
}

func (o *SortOrder) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"order-id",
		"fields",
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

	varSortOrder := _SortOrder{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSortOrder)

	if err != nil {
		return err
	}

	*o = SortOrder(varSortOrder)

	return err
}

type NullableSortOrder struct {
	value *SortOrder
	isSet bool
}

func (v NullableSortOrder) Get() *SortOrder {
	return v.value
}

func (v *NullableSortOrder) Set(val *SortOrder) {
	v.value = val
	v.isSet = true
}

func (v NullableSortOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableSortOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSortOrder(val *SortOrder) *NullableSortOrder {
	return &NullableSortOrder{value: val, isSet: true}
}

func (v NullableSortOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSortOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


