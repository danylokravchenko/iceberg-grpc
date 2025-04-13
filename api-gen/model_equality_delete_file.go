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

// checks if the EqualityDeleteFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EqualityDeleteFile{}

// EqualityDeleteFile struct for EqualityDeleteFile
type EqualityDeleteFile struct {
	ContentFile
	Content string `json:"content"`
	// List of equality field IDs
	EqualityIds []int32 `json:"equality-ids,omitempty"`
}

type _EqualityDeleteFile EqualityDeleteFile

// NewEqualityDeleteFile instantiates a new EqualityDeleteFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEqualityDeleteFile(content string, filePath string, fileFormat FileFormat, specId int32, partition []PrimitiveTypeValue, fileSizeInBytes int64, recordCount int64) *EqualityDeleteFile {
	this := EqualityDeleteFile{}
	this.Content = content
	this.FilePath = filePath
	this.FileFormat = fileFormat
	this.SpecId = specId
	this.Partition = partition
	this.FileSizeInBytes = fileSizeInBytes
	this.RecordCount = recordCount
	return &this
}

// NewEqualityDeleteFileWithDefaults instantiates a new EqualityDeleteFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEqualityDeleteFileWithDefaults() *EqualityDeleteFile {
	this := EqualityDeleteFile{}
	return &this
}

// GetContent returns the Content field value
func (o *EqualityDeleteFile) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *EqualityDeleteFile) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *EqualityDeleteFile) SetContent(v string) {
	o.Content = v
}

// GetEqualityIds returns the EqualityIds field value if set, zero value otherwise.
func (o *EqualityDeleteFile) GetEqualityIds() []int32 {
	if o == nil || IsNil(o.EqualityIds) {
		var ret []int32
		return ret
	}
	return o.EqualityIds
}

// GetEqualityIdsOk returns a tuple with the EqualityIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EqualityDeleteFile) GetEqualityIdsOk() ([]int32, bool) {
	if o == nil || IsNil(o.EqualityIds) {
		return nil, false
	}
	return o.EqualityIds, true
}

// HasEqualityIds returns a boolean if a field has been set.
func (o *EqualityDeleteFile) HasEqualityIds() bool {
	if o != nil && !IsNil(o.EqualityIds) {
		return true
	}

	return false
}

// SetEqualityIds gets a reference to the given []int32 and assigns it to the EqualityIds field.
func (o *EqualityDeleteFile) SetEqualityIds(v []int32) {
	o.EqualityIds = v
}

func (o EqualityDeleteFile) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EqualityDeleteFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedContentFile, errContentFile := json.Marshal(o.ContentFile)
	if errContentFile != nil {
		return map[string]interface{}{}, errContentFile
	}
	errContentFile = json.Unmarshal([]byte(serializedContentFile), &toSerialize)
	if errContentFile != nil {
		return map[string]interface{}{}, errContentFile
	}
	toSerialize["content"] = o.Content
	if !IsNil(o.EqualityIds) {
		toSerialize["equality-ids"] = o.EqualityIds
	}
	return toSerialize, nil
}

func (o *EqualityDeleteFile) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"content",
		"file-path",
		"file-format",
		"spec-id",
		"partition",
		"file-size-in-bytes",
		"record-count",
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

	varEqualityDeleteFile := _EqualityDeleteFile{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEqualityDeleteFile)

	if err != nil {
		return err
	}

	*o = EqualityDeleteFile(varEqualityDeleteFile)

	return err
}

type NullableEqualityDeleteFile struct {
	value *EqualityDeleteFile
	isSet bool
}

func (v NullableEqualityDeleteFile) Get() *EqualityDeleteFile {
	return v.value
}

func (v *NullableEqualityDeleteFile) Set(val *EqualityDeleteFile) {
	v.value = val
	v.isSet = true
}

func (v NullableEqualityDeleteFile) IsSet() bool {
	return v.isSet
}

func (v *NullableEqualityDeleteFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEqualityDeleteFile(val *EqualityDeleteFile) *NullableEqualityDeleteFile {
	return &NullableEqualityDeleteFile{value: val, isSet: true}
}

func (v NullableEqualityDeleteFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEqualityDeleteFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


