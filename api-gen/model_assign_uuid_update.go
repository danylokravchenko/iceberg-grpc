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

// checks if the AssignUUIDUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssignUUIDUpdate{}

// AssignUUIDUpdate Assigning a UUID to a table/view should only be done when creating the table/view. It is not safe to re-assign the UUID if a table/view already has a UUID assigned
type AssignUUIDUpdate struct {
	BaseUpdate
	Action *string `json:"action,omitempty"`
	Uuid string `json:"uuid"`
}

type _AssignUUIDUpdate AssignUUIDUpdate

// NewAssignUUIDUpdate instantiates a new AssignUUIDUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssignUUIDUpdate(uuid string) *AssignUUIDUpdate {
	this := AssignUUIDUpdate{}
	return &this
}

// NewAssignUUIDUpdateWithDefaults instantiates a new AssignUUIDUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssignUUIDUpdateWithDefaults() *AssignUUIDUpdate {
	this := AssignUUIDUpdate{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *AssignUUIDUpdate) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssignUUIDUpdate) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *AssignUUIDUpdate) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *AssignUUIDUpdate) SetAction(v string) {
	o.Action = &v
}

// GetUuid returns the Uuid field value
func (o *AssignUUIDUpdate) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *AssignUUIDUpdate) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *AssignUUIDUpdate) SetUuid(v string) {
	o.Uuid = v
}

func (o AssignUUIDUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssignUUIDUpdate) ToMap() (map[string]interface{}, error) {
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
	toSerialize["uuid"] = o.Uuid
	return toSerialize, nil
}

func (o *AssignUUIDUpdate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uuid",
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

	varAssignUUIDUpdate := _AssignUUIDUpdate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAssignUUIDUpdate)

	if err != nil {
		return err
	}

	*o = AssignUUIDUpdate(varAssignUUIDUpdate)

	return err
}

type NullableAssignUUIDUpdate struct {
	value *AssignUUIDUpdate
	isSet bool
}

func (v NullableAssignUUIDUpdate) Get() *AssignUUIDUpdate {
	return v.value
}

func (v *NullableAssignUUIDUpdate) Set(val *AssignUUIDUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableAssignUUIDUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableAssignUUIDUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssignUUIDUpdate(val *AssignUUIDUpdate) *NullableAssignUUIDUpdate {
	return &NullableAssignUUIDUpdate{value: val, isSet: true}
}

func (v NullableAssignUUIDUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssignUUIDUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


