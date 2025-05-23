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

// checks if the CounterResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CounterResult{}

// CounterResult struct for CounterResult
type CounterResult struct {
	Unit string `json:"unit"`
	Value int64 `json:"value"`
}

type _CounterResult CounterResult

// NewCounterResult instantiates a new CounterResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCounterResult(unit string, value int64) *CounterResult {
	this := CounterResult{}
	this.Unit = unit
	this.Value = value
	return &this
}

// NewCounterResultWithDefaults instantiates a new CounterResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCounterResultWithDefaults() *CounterResult {
	this := CounterResult{}
	return &this
}

// GetUnit returns the Unit field value
func (o *CounterResult) GetUnit() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Unit
}

// GetUnitOk returns a tuple with the Unit field value
// and a boolean to check if the value has been set.
func (o *CounterResult) GetUnitOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Unit, true
}

// SetUnit sets field value
func (o *CounterResult) SetUnit(v string) {
	o.Unit = v
}

// GetValue returns the Value field value
func (o *CounterResult) GetValue() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *CounterResult) GetValueOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *CounterResult) SetValue(v int64) {
	o.Value = v
}

func (o CounterResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CounterResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["unit"] = o.Unit
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *CounterResult) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"unit",
		"value",
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

	varCounterResult := _CounterResult{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCounterResult)

	if err != nil {
		return err
	}

	*o = CounterResult(varCounterResult)

	return err
}

type NullableCounterResult struct {
	value *CounterResult
	isSet bool
}

func (v NullableCounterResult) Get() *CounterResult {
	return v.value
}

func (v *NullableCounterResult) Set(val *CounterResult) {
	v.value = val
	v.isSet = true
}

func (v NullableCounterResult) IsSet() bool {
	return v.isSet
}

func (v *NullableCounterResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCounterResult(val *CounterResult) *NullableCounterResult {
	return &NullableCounterResult{value: val, isSet: true}
}

func (v NullableCounterResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCounterResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


