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

// checks if the MetadataLogInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetadataLogInner{}

// MetadataLogInner struct for MetadataLogInner
type MetadataLogInner struct {
	MetadataFile string `json:"metadata-file"`
	TimestampMs int64 `json:"timestamp-ms"`
}

type _MetadataLogInner MetadataLogInner

// NewMetadataLogInner instantiates a new MetadataLogInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetadataLogInner(metadataFile string, timestampMs int64) *MetadataLogInner {
	this := MetadataLogInner{}
	this.MetadataFile = metadataFile
	this.TimestampMs = timestampMs
	return &this
}

// NewMetadataLogInnerWithDefaults instantiates a new MetadataLogInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetadataLogInnerWithDefaults() *MetadataLogInner {
	this := MetadataLogInner{}
	return &this
}

// GetMetadataFile returns the MetadataFile field value
func (o *MetadataLogInner) GetMetadataFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MetadataFile
}

// GetMetadataFileOk returns a tuple with the MetadataFile field value
// and a boolean to check if the value has been set.
func (o *MetadataLogInner) GetMetadataFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MetadataFile, true
}

// SetMetadataFile sets field value
func (o *MetadataLogInner) SetMetadataFile(v string) {
	o.MetadataFile = v
}

// GetTimestampMs returns the TimestampMs field value
func (o *MetadataLogInner) GetTimestampMs() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TimestampMs
}

// GetTimestampMsOk returns a tuple with the TimestampMs field value
// and a boolean to check if the value has been set.
func (o *MetadataLogInner) GetTimestampMsOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TimestampMs, true
}

// SetTimestampMs sets field value
func (o *MetadataLogInner) SetTimestampMs(v int64) {
	o.TimestampMs = v
}

func (o MetadataLogInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetadataLogInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["metadata-file"] = o.MetadataFile
	toSerialize["timestamp-ms"] = o.TimestampMs
	return toSerialize, nil
}

func (o *MetadataLogInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"metadata-file",
		"timestamp-ms",
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

	varMetadataLogInner := _MetadataLogInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMetadataLogInner)

	if err != nil {
		return err
	}

	*o = MetadataLogInner(varMetadataLogInner)

	return err
}

type NullableMetadataLogInner struct {
	value *MetadataLogInner
	isSet bool
}

func (v NullableMetadataLogInner) Get() *MetadataLogInner {
	return v.value
}

func (v *NullableMetadataLogInner) Set(val *MetadataLogInner) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataLogInner) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataLogInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataLogInner(val *MetadataLogInner) *NullableMetadataLogInner {
	return &NullableMetadataLogInner{value: val, isSet: true}
}

func (v NullableMetadataLogInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataLogInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


