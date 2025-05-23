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

// checks if the CatalogConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CatalogConfig{}

// CatalogConfig Server-provided configuration for the catalog.
type CatalogConfig struct {
	// Properties that should be used to override client configuration; applied after defaults and client configuration.
	Overrides map[string]string `json:"overrides"`
	// Properties that should be used as default configuration; applied before client configuration.
	Defaults map[string]string `json:"defaults"`
	// A list of endpoints that the server supports. The format of each endpoint must be \"<HTTP verb> <resource path from OpenAPI REST spec>\". The HTTP verb and the resource path must be separated by a space character.
	Endpoints []string `json:"endpoints,omitempty"`
}

type _CatalogConfig CatalogConfig

// NewCatalogConfig instantiates a new CatalogConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCatalogConfig(overrides map[string]string, defaults map[string]string) *CatalogConfig {
	this := CatalogConfig{}
	this.Overrides = overrides
	this.Defaults = defaults
	return &this
}

// NewCatalogConfigWithDefaults instantiates a new CatalogConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCatalogConfigWithDefaults() *CatalogConfig {
	this := CatalogConfig{}
	return &this
}

// GetOverrides returns the Overrides field value
func (o *CatalogConfig) GetOverrides() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Overrides
}

// GetOverridesOk returns a tuple with the Overrides field value
// and a boolean to check if the value has been set.
func (o *CatalogConfig) GetOverridesOk() (map[string]string, bool) {
	if o == nil {
		return map[string]string{}, false
	}
	return o.Overrides, true
}

// SetOverrides sets field value
func (o *CatalogConfig) SetOverrides(v map[string]string) {
	o.Overrides = v
}

// GetDefaults returns the Defaults field value
func (o *CatalogConfig) GetDefaults() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Defaults
}

// GetDefaultsOk returns a tuple with the Defaults field value
// and a boolean to check if the value has been set.
func (o *CatalogConfig) GetDefaultsOk() (map[string]string, bool) {
	if o == nil {
		return map[string]string{}, false
	}
	return o.Defaults, true
}

// SetDefaults sets field value
func (o *CatalogConfig) SetDefaults(v map[string]string) {
	o.Defaults = v
}

// GetEndpoints returns the Endpoints field value if set, zero value otherwise.
func (o *CatalogConfig) GetEndpoints() []string {
	if o == nil || IsNil(o.Endpoints) {
		var ret []string
		return ret
	}
	return o.Endpoints
}

// GetEndpointsOk returns a tuple with the Endpoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CatalogConfig) GetEndpointsOk() ([]string, bool) {
	if o == nil || IsNil(o.Endpoints) {
		return nil, false
	}
	return o.Endpoints, true
}

// HasEndpoints returns a boolean if a field has been set.
func (o *CatalogConfig) HasEndpoints() bool {
	if o != nil && !IsNil(o.Endpoints) {
		return true
	}

	return false
}

// SetEndpoints gets a reference to the given []string and assigns it to the Endpoints field.
func (o *CatalogConfig) SetEndpoints(v []string) {
	o.Endpoints = v
}

func (o CatalogConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CatalogConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["overrides"] = o.Overrides
	toSerialize["defaults"] = o.Defaults
	if !IsNil(o.Endpoints) {
		toSerialize["endpoints"] = o.Endpoints
	}
	return toSerialize, nil
}

func (o *CatalogConfig) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"overrides",
		"defaults",
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

	varCatalogConfig := _CatalogConfig{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCatalogConfig)

	if err != nil {
		return err
	}

	*o = CatalogConfig(varCatalogConfig)

	return err
}

type NullableCatalogConfig struct {
	value *CatalogConfig
	isSet bool
}

func (v NullableCatalogConfig) Get() *CatalogConfig {
	return v.value
}

func (v *NullableCatalogConfig) Set(val *CatalogConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableCatalogConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableCatalogConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCatalogConfig(val *CatalogConfig) *NullableCatalogConfig {
	return &NullableCatalogConfig{value: val, isSet: true}
}

func (v NullableCatalogConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCatalogConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


