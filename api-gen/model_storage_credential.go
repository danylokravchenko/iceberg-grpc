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

// checks if the StorageCredential type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StorageCredential{}

// StorageCredential struct for StorageCredential
type StorageCredential struct {
	// Indicates a storage location prefix where the credential is relevant. Clients should choose the most specific prefix (by selecting the longest prefix) if several credentials of the same type are available.
	Prefix string `json:"prefix"`
	Config map[string]string `json:"config"`
}

type _StorageCredential StorageCredential

// NewStorageCredential instantiates a new StorageCredential object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageCredential(prefix string, config map[string]string) *StorageCredential {
	this := StorageCredential{}
	this.Prefix = prefix
	this.Config = config
	return &this
}

// NewStorageCredentialWithDefaults instantiates a new StorageCredential object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageCredentialWithDefaults() *StorageCredential {
	this := StorageCredential{}
	return &this
}

// GetPrefix returns the Prefix field value
func (o *StorageCredential) GetPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value
// and a boolean to check if the value has been set.
func (o *StorageCredential) GetPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Prefix, true
}

// SetPrefix sets field value
func (o *StorageCredential) SetPrefix(v string) {
	o.Prefix = v
}

// GetConfig returns the Config field value
func (o *StorageCredential) GetConfig() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *StorageCredential) GetConfigOk() (map[string]string, bool) {
	if o == nil {
		return map[string]string{}, false
	}
	return o.Config, true
}

// SetConfig sets field value
func (o *StorageCredential) SetConfig(v map[string]string) {
	o.Config = v
}

func (o StorageCredential) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StorageCredential) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["prefix"] = o.Prefix
	toSerialize["config"] = o.Config
	return toSerialize, nil
}

func (o *StorageCredential) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"prefix",
		"config",
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

	varStorageCredential := _StorageCredential{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varStorageCredential)

	if err != nil {
		return err
	}

	*o = StorageCredential(varStorageCredential)

	return err
}

type NullableStorageCredential struct {
	value *StorageCredential
	isSet bool
}

func (v NullableStorageCredential) Get() *StorageCredential {
	return v.value
}

func (v *NullableStorageCredential) Set(val *StorageCredential) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageCredential) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageCredential) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageCredential(val *StorageCredential) *NullableStorageCredential {
	return &NullableStorageCredential{value: val, isSet: true}
}

func (v NullableStorageCredential) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageCredential) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


