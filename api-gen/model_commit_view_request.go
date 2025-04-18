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

// checks if the CommitViewRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommitViewRequest{}

// CommitViewRequest struct for CommitViewRequest
type CommitViewRequest struct {
	// View identifier to update
	Identifier *TableIdentifier `json:"identifier,omitempty"`
	Requirements []ViewRequirement `json:"requirements,omitempty"`
	Updates []ViewUpdate `json:"updates"`
}

type _CommitViewRequest CommitViewRequest

// NewCommitViewRequest instantiates a new CommitViewRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitViewRequest(updates []ViewUpdate) *CommitViewRequest {
	this := CommitViewRequest{}
	this.Updates = updates
	return &this
}

// NewCommitViewRequestWithDefaults instantiates a new CommitViewRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitViewRequestWithDefaults() *CommitViewRequest {
	this := CommitViewRequest{}
	return &this
}

// GetIdentifier returns the Identifier field value if set, zero value otherwise.
func (o *CommitViewRequest) GetIdentifier() TableIdentifier {
	if o == nil || IsNil(o.Identifier) {
		var ret TableIdentifier
		return ret
	}
	return *o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitViewRequest) GetIdentifierOk() (*TableIdentifier, bool) {
	if o == nil || IsNil(o.Identifier) {
		return nil, false
	}
	return o.Identifier, true
}

// HasIdentifier returns a boolean if a field has been set.
func (o *CommitViewRequest) HasIdentifier() bool {
	if o != nil && !IsNil(o.Identifier) {
		return true
	}

	return false
}

// SetIdentifier gets a reference to the given TableIdentifier and assigns it to the Identifier field.
func (o *CommitViewRequest) SetIdentifier(v TableIdentifier) {
	o.Identifier = &v
}

// GetRequirements returns the Requirements field value if set, zero value otherwise.
func (o *CommitViewRequest) GetRequirements() []ViewRequirement {
	if o == nil || IsNil(o.Requirements) {
		var ret []ViewRequirement
		return ret
	}
	return o.Requirements
}

// GetRequirementsOk returns a tuple with the Requirements field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitViewRequest) GetRequirementsOk() ([]ViewRequirement, bool) {
	if o == nil || IsNil(o.Requirements) {
		return nil, false
	}
	return o.Requirements, true
}

// HasRequirements returns a boolean if a field has been set.
func (o *CommitViewRequest) HasRequirements() bool {
	if o != nil && !IsNil(o.Requirements) {
		return true
	}

	return false
}

// SetRequirements gets a reference to the given []ViewRequirement and assigns it to the Requirements field.
func (o *CommitViewRequest) SetRequirements(v []ViewRequirement) {
	o.Requirements = v
}

// GetUpdates returns the Updates field value
func (o *CommitViewRequest) GetUpdates() []ViewUpdate {
	if o == nil {
		var ret []ViewUpdate
		return ret
	}

	return o.Updates
}

// GetUpdatesOk returns a tuple with the Updates field value
// and a boolean to check if the value has been set.
func (o *CommitViewRequest) GetUpdatesOk() ([]ViewUpdate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Updates, true
}

// SetUpdates sets field value
func (o *CommitViewRequest) SetUpdates(v []ViewUpdate) {
	o.Updates = v
}

func (o CommitViewRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommitViewRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Identifier) {
		toSerialize["identifier"] = o.Identifier
	}
	if !IsNil(o.Requirements) {
		toSerialize["requirements"] = o.Requirements
	}
	toSerialize["updates"] = o.Updates
	return toSerialize, nil
}

func (o *CommitViewRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"updates",
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

	varCommitViewRequest := _CommitViewRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCommitViewRequest)

	if err != nil {
		return err
	}

	*o = CommitViewRequest(varCommitViewRequest)

	return err
}

type NullableCommitViewRequest struct {
	value *CommitViewRequest
	isSet bool
}

func (v NullableCommitViewRequest) Get() *CommitViewRequest {
	return v.value
}

func (v *NullableCommitViewRequest) Set(val *CommitViewRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitViewRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitViewRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitViewRequest(val *CommitViewRequest) *NullableCommitViewRequest {
	return &NullableCommitViewRequest{value: val, isSet: true}
}

func (v NullableCommitViewRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitViewRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


