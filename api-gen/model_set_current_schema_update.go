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

// checks if the SetCurrentSchemaUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetCurrentSchemaUpdate{}

// SetCurrentSchemaUpdate struct for SetCurrentSchemaUpdate
type SetCurrentSchemaUpdate struct {
	BaseUpdate
	Action *string `json:"action,omitempty"`
	// Schema ID to set as current, or -1 to set last added schema
	SchemaId int32 `json:"schema-id"`
}

type _SetCurrentSchemaUpdate SetCurrentSchemaUpdate

// NewSetCurrentSchemaUpdate instantiates a new SetCurrentSchemaUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetCurrentSchemaUpdate(schemaId int32) *SetCurrentSchemaUpdate {
	this := SetCurrentSchemaUpdate{}
	return &this
}

// NewSetCurrentSchemaUpdateWithDefaults instantiates a new SetCurrentSchemaUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetCurrentSchemaUpdateWithDefaults() *SetCurrentSchemaUpdate {
	this := SetCurrentSchemaUpdate{}
	return &this
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *SetCurrentSchemaUpdate) GetAction() string {
	if o == nil || IsNil(o.Action) {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetCurrentSchemaUpdate) GetActionOk() (*string, bool) {
	if o == nil || IsNil(o.Action) {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *SetCurrentSchemaUpdate) HasAction() bool {
	if o != nil && !IsNil(o.Action) {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *SetCurrentSchemaUpdate) SetAction(v string) {
	o.Action = &v
}

// GetSchemaId returns the SchemaId field value
func (o *SetCurrentSchemaUpdate) GetSchemaId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SchemaId
}

// GetSchemaIdOk returns a tuple with the SchemaId field value
// and a boolean to check if the value has been set.
func (o *SetCurrentSchemaUpdate) GetSchemaIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SchemaId, true
}

// SetSchemaId sets field value
func (o *SetCurrentSchemaUpdate) SetSchemaId(v int32) {
	o.SchemaId = v
}

func (o SetCurrentSchemaUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetCurrentSchemaUpdate) ToMap() (map[string]interface{}, error) {
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
	toSerialize["schema-id"] = o.SchemaId
	return toSerialize, nil
}

func (o *SetCurrentSchemaUpdate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"schema-id",
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

	varSetCurrentSchemaUpdate := _SetCurrentSchemaUpdate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetCurrentSchemaUpdate)

	if err != nil {
		return err
	}

	*o = SetCurrentSchemaUpdate(varSetCurrentSchemaUpdate)

	return err
}

type NullableSetCurrentSchemaUpdate struct {
	value *SetCurrentSchemaUpdate
	isSet bool
}

func (v NullableSetCurrentSchemaUpdate) Get() *SetCurrentSchemaUpdate {
	return v.value
}

func (v *NullableSetCurrentSchemaUpdate) Set(val *SetCurrentSchemaUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSetCurrentSchemaUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSetCurrentSchemaUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetCurrentSchemaUpdate(val *SetCurrentSchemaUpdate) *NullableSetCurrentSchemaUpdate {
	return &NullableSetCurrentSchemaUpdate{value: val, isSet: true}
}

func (v NullableSetCurrentSchemaUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetCurrentSchemaUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


