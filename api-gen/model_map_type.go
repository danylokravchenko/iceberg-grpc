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

// checks if the MapType type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MapType{}

// MapType struct for MapType
type MapType struct {
	Type string `json:"type"`
	KeyId int32 `json:"key-id"`
	Key Type `json:"key"`
	ValueId int32 `json:"value-id"`
	Value Type `json:"value"`
	ValueRequired bool `json:"value-required"`
}

type _MapType MapType

// NewMapType instantiates a new MapType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMapType(type_ string, keyId int32, key Type, valueId int32, value Type, valueRequired bool) *MapType {
	this := MapType{}
	this.Type = type_
	this.KeyId = keyId
	this.Key = key
	this.ValueId = valueId
	this.Value = value
	this.ValueRequired = valueRequired
	return &this
}

// NewMapTypeWithDefaults instantiates a new MapType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMapTypeWithDefaults() *MapType {
	this := MapType{}
	return &this
}

// GetType returns the Type field value
func (o *MapType) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *MapType) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *MapType) SetType(v string) {
	o.Type = v
}

// GetKeyId returns the KeyId field value
func (o *MapType) GetKeyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.KeyId
}

// GetKeyIdOk returns a tuple with the KeyId field value
// and a boolean to check if the value has been set.
func (o *MapType) GetKeyIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.KeyId, true
}

// SetKeyId sets field value
func (o *MapType) SetKeyId(v int32) {
	o.KeyId = v
}

// GetKey returns the Key field value
func (o *MapType) GetKey() Type {
	if o == nil {
		var ret Type
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *MapType) GetKeyOk() (*Type, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *MapType) SetKey(v Type) {
	o.Key = v
}

// GetValueId returns the ValueId field value
func (o *MapType) GetValueId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ValueId
}

// GetValueIdOk returns a tuple with the ValueId field value
// and a boolean to check if the value has been set.
func (o *MapType) GetValueIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueId, true
}

// SetValueId sets field value
func (o *MapType) SetValueId(v int32) {
	o.ValueId = v
}

// GetValue returns the Value field value
func (o *MapType) GetValue() Type {
	if o == nil {
		var ret Type
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *MapType) GetValueOk() (*Type, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *MapType) SetValue(v Type) {
	o.Value = v
}

// GetValueRequired returns the ValueRequired field value
func (o *MapType) GetValueRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ValueRequired
}

// GetValueRequiredOk returns a tuple with the ValueRequired field value
// and a boolean to check if the value has been set.
func (o *MapType) GetValueRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValueRequired, true
}

// SetValueRequired sets field value
func (o *MapType) SetValueRequired(v bool) {
	o.ValueRequired = v
}

func (o MapType) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MapType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["key-id"] = o.KeyId
	toSerialize["key"] = o.Key
	toSerialize["value-id"] = o.ValueId
	toSerialize["value"] = o.Value
	toSerialize["value-required"] = o.ValueRequired
	return toSerialize, nil
}

func (o *MapType) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"key-id",
		"key",
		"value-id",
		"value",
		"value-required",
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

	varMapType := _MapType{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMapType)

	if err != nil {
		return err
	}

	*o = MapType(varMapType)

	return err
}

type NullableMapType struct {
	value *MapType
	isSet bool
}

func (v NullableMapType) Get() *MapType {
	return v.value
}

func (v *NullableMapType) Set(val *MapType) {
	v.value = val
	v.isSet = true
}

func (v NullableMapType) IsSet() bool {
	return v.isSet
}

func (v *NullableMapType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMapType(val *MapType) *NullableMapType {
	return &NullableMapType{value: val, isSet: true}
}

func (v NullableMapType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMapType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


