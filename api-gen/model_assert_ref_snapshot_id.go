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

// checks if the AssertRefSnapshotId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssertRefSnapshotId{}

// AssertRefSnapshotId The table branch or tag identified by the requirement's `ref` must reference the requirement's `snapshot-id`; if `snapshot-id` is `null` or missing, the ref must not already exist
type AssertRefSnapshotId struct {
	TableRequirement
	Type *string `json:"type,omitempty"`
	Ref string `json:"ref"`
	SnapshotId int64 `json:"snapshot-id"`
}

type _AssertRefSnapshotId AssertRefSnapshotId

// NewAssertRefSnapshotId instantiates a new AssertRefSnapshotId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssertRefSnapshotId(ref string, snapshotId int64) *AssertRefSnapshotId {
	this := AssertRefSnapshotId{}
	return &this
}

// NewAssertRefSnapshotIdWithDefaults instantiates a new AssertRefSnapshotId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssertRefSnapshotIdWithDefaults() *AssertRefSnapshotId {
	this := AssertRefSnapshotId{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AssertRefSnapshotId) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssertRefSnapshotId) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AssertRefSnapshotId) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AssertRefSnapshotId) SetType(v string) {
	o.Type = &v
}

// GetRef returns the Ref field value
func (o *AssertRefSnapshotId) GetRef() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Ref
}

// GetRefOk returns a tuple with the Ref field value
// and a boolean to check if the value has been set.
func (o *AssertRefSnapshotId) GetRefOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ref, true
}

// SetRef sets field value
func (o *AssertRefSnapshotId) SetRef(v string) {
	o.Ref = v
}

// GetSnapshotId returns the SnapshotId field value
func (o *AssertRefSnapshotId) GetSnapshotId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value
// and a boolean to check if the value has been set.
func (o *AssertRefSnapshotId) GetSnapshotIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SnapshotId, true
}

// SetSnapshotId sets field value
func (o *AssertRefSnapshotId) SetSnapshotId(v int64) {
	o.SnapshotId = v
}

func (o AssertRefSnapshotId) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssertRefSnapshotId) ToMap() (map[string]interface{}, error) {
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
	toSerialize["ref"] = o.Ref
	toSerialize["snapshot-id"] = o.SnapshotId
	return toSerialize, nil
}

func (o *AssertRefSnapshotId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"ref",
		"snapshot-id",
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

	varAssertRefSnapshotId := _AssertRefSnapshotId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAssertRefSnapshotId)

	if err != nil {
		return err
	}

	*o = AssertRefSnapshotId(varAssertRefSnapshotId)

	return err
}

type NullableAssertRefSnapshotId struct {
	value *AssertRefSnapshotId
	isSet bool
}

func (v NullableAssertRefSnapshotId) Get() *AssertRefSnapshotId {
	return v.value
}

func (v *NullableAssertRefSnapshotId) Set(val *AssertRefSnapshotId) {
	v.value = val
	v.isSet = true
}

func (v NullableAssertRefSnapshotId) IsSet() bool {
	return v.isSet
}

func (v *NullableAssertRefSnapshotId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssertRefSnapshotId(val *AssertRefSnapshotId) *NullableAssertRefSnapshotId {
	return &NullableAssertRefSnapshotId{value: val, isSet: true}
}

func (v NullableAssertRefSnapshotId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssertRefSnapshotId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


