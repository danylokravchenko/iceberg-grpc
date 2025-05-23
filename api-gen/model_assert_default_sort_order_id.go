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

// checks if the AssertDefaultSortOrderId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssertDefaultSortOrderId{}

// AssertDefaultSortOrderId The table's default sort order id must match the requirement's `default-sort-order-id`
type AssertDefaultSortOrderId struct {
	TableRequirement
	Type *string `json:"type,omitempty"`
	DefaultSortOrderId int32 `json:"default-sort-order-id"`
}

type _AssertDefaultSortOrderId AssertDefaultSortOrderId

// NewAssertDefaultSortOrderId instantiates a new AssertDefaultSortOrderId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssertDefaultSortOrderId(defaultSortOrderId int32) *AssertDefaultSortOrderId {
	this := AssertDefaultSortOrderId{}
	return &this
}

// NewAssertDefaultSortOrderIdWithDefaults instantiates a new AssertDefaultSortOrderId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssertDefaultSortOrderIdWithDefaults() *AssertDefaultSortOrderId {
	this := AssertDefaultSortOrderId{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AssertDefaultSortOrderId) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssertDefaultSortOrderId) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AssertDefaultSortOrderId) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AssertDefaultSortOrderId) SetType(v string) {
	o.Type = &v
}

// GetDefaultSortOrderId returns the DefaultSortOrderId field value
func (o *AssertDefaultSortOrderId) GetDefaultSortOrderId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DefaultSortOrderId
}

// GetDefaultSortOrderIdOk returns a tuple with the DefaultSortOrderId field value
// and a boolean to check if the value has been set.
func (o *AssertDefaultSortOrderId) GetDefaultSortOrderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultSortOrderId, true
}

// SetDefaultSortOrderId sets field value
func (o *AssertDefaultSortOrderId) SetDefaultSortOrderId(v int32) {
	o.DefaultSortOrderId = v
}

func (o AssertDefaultSortOrderId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssertDefaultSortOrderId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedTableRequirement, errTableRequirement := json.Marshal(o.TableRequirement)
	if errTableRequirement != nil {
		return map[string]interface{}{}, errTableRequirement
	}
	errTableRequirement = json.Unmarshal([]byte(serializedTableRequirement), &toSerialize)
	if errTableRequirement != nil {
		return map[string]interface{}{}, errTableRequirement
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	toSerialize["default-sort-order-id"] = o.DefaultSortOrderId
	return toSerialize, nil
}

func (o *AssertDefaultSortOrderId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"default-sort-order-id",
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

	varAssertDefaultSortOrderId := _AssertDefaultSortOrderId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAssertDefaultSortOrderId)

	if err != nil {
		return err
	}

	*o = AssertDefaultSortOrderId(varAssertDefaultSortOrderId)

	return err
}

type NullableAssertDefaultSortOrderId struct {
	value *AssertDefaultSortOrderId
	isSet bool
}

func (v NullableAssertDefaultSortOrderId) Get() *AssertDefaultSortOrderId {
	return v.value
}

func (v *NullableAssertDefaultSortOrderId) Set(val *AssertDefaultSortOrderId) {
	v.value = val
	v.isSet = true
}

func (v NullableAssertDefaultSortOrderId) IsSet() bool {
	return v.isSet
}

func (v *NullableAssertDefaultSortOrderId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssertDefaultSortOrderId(val *AssertDefaultSortOrderId) *NullableAssertDefaultSortOrderId {
	return &NullableAssertDefaultSortOrderId{value: val, isSet: true}
}

func (v NullableAssertDefaultSortOrderId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssertDefaultSortOrderId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


