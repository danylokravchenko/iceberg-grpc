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

// checks if the SetDefaultSortOrderUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetDefaultSortOrderUpdate{}

// SetDefaultSortOrderUpdate struct for SetDefaultSortOrderUpdate
type SetDefaultSortOrderUpdate struct {
	BaseUpdate
	Action *string `json:"action,omitempty"`
	// Sort order ID to set as the default, or -1 to set last added sort order
	SortOrderId int32 `json:"sort-order-id"`
}

type _SetDefaultSortOrderUpdate SetDefaultSortOrderUpdate

// NewSetDefaultSortOrderUpdate instantiates a new SetDefaultSortOrderUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetDefaultSortOrderUpdate(sortOrderId int32) *SetDefaultSortOrderUpdate {
	this := SetDefaultSortOrderUpdate{}
	return &this
}

// NewSetDefaultSortOrderUpdateWithDefaults instantiates a new SetDefaultSortOrderUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetDefaultSortOrderUpdateWithDefaults() *SetDefaultSortOrderUpdate {
	this := SetDefaultSortOrderUpdate{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *SetDefaultSortOrderUpdate) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetDefaultSortOrderUpdate) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *SetDefaultSortOrderUpdate) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *SetDefaultSortOrderUpdate) SetAction(v string) {
	o.Action = &v
}

// GetSortOrderId returns the SortOrderId field value
func (o *SetDefaultSortOrderUpdate) GetSortOrderId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SortOrderId
}

// GetSortOrderIdOk returns a tuple with the SortOrderId field value
// and a boolean to check if the value has been set.
func (o *SetDefaultSortOrderUpdate) GetSortOrderIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SortOrderId, true
}

// SetSortOrderId sets field value
func (o *SetDefaultSortOrderUpdate) SetSortOrderId(v int32) {
	o.SortOrderId = v
}

func (o SetDefaultSortOrderUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetDefaultSortOrderUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedBaseUpdate, errBaseUpdate := json.Marshal(o.BaseUpdate)
	if errBaseUpdate != nil {
		return map[string]interface{}{}, errBaseUpdate
	}
	errBaseUpdate = json.Unmarshal([]byte(serializedBaseUpdate), &toSerialize)
	if errBaseUpdate != nil {
		return map[string]interface{}{}, errBaseUpdate
	}
	if !IsNil(o.Action) {
		toSerialize["action"] = o.Action
	}
	toSerialize["sort-order-id"] = o.SortOrderId
	return toSerialize, nil
}

func (o *SetDefaultSortOrderUpdate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"sort-order-id",
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

	varSetDefaultSortOrderUpdate := _SetDefaultSortOrderUpdate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetDefaultSortOrderUpdate)

	if err != nil {
		return err
	}

	*o = SetDefaultSortOrderUpdate(varSetDefaultSortOrderUpdate)

	return err
}

type NullableSetDefaultSortOrderUpdate struct {
	value *SetDefaultSortOrderUpdate
	isSet bool
}

func (v NullableSetDefaultSortOrderUpdate) Get() *SetDefaultSortOrderUpdate {
	return v.value
}

func (v *NullableSetDefaultSortOrderUpdate) Set(val *SetDefaultSortOrderUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSetDefaultSortOrderUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSetDefaultSortOrderUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetDefaultSortOrderUpdate(val *SetDefaultSortOrderUpdate) *NullableSetDefaultSortOrderUpdate {
	return &NullableSetDefaultSortOrderUpdate{value: val, isSet: true}
}

func (v NullableSetDefaultSortOrderUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetDefaultSortOrderUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


