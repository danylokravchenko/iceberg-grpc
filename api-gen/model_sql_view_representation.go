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

// checks if the SQLViewRepresentation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SQLViewRepresentation{}

// SQLViewRepresentation struct for SQLViewRepresentation
type SQLViewRepresentation struct {
	Type string `json:"type"`
	Sql string `json:"sql"`
	Dialect string `json:"dialect"`
}

type _SQLViewRepresentation SQLViewRepresentation

// NewSQLViewRepresentation instantiates a new SQLViewRepresentation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSQLViewRepresentation(type_ string, sql string, dialect string) *SQLViewRepresentation {
	this := SQLViewRepresentation{}
	this.Type = type_
	this.Sql = sql
	this.Dialect = dialect
	return &this
}

// NewSQLViewRepresentationWithDefaults instantiates a new SQLViewRepresentation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSQLViewRepresentationWithDefaults() *SQLViewRepresentation {
	this := SQLViewRepresentation{}
	return &this
}

// GetType returns the Type field value
func (o *SQLViewRepresentation) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SQLViewRepresentation) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SQLViewRepresentation) SetType(v string) {
	o.Type = v
}

// GetSql returns the Sql field value
func (o *SQLViewRepresentation) GetSql() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Sql
}

// GetSqlOk returns a tuple with the Sql field value
// and a boolean to check if the value has been set.
func (o *SQLViewRepresentation) GetSqlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sql, true
}

// SetSql sets field value
func (o *SQLViewRepresentation) SetSql(v string) {
	o.Sql = v
}

// GetDialect returns the Dialect field value
func (o *SQLViewRepresentation) GetDialect() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Dialect
}

// GetDialectOk returns a tuple with the Dialect field value
// and a boolean to check if the value has been set.
func (o *SQLViewRepresentation) GetDialectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dialect, true
}

// SetDialect sets field value
func (o *SQLViewRepresentation) SetDialect(v string) {
	o.Dialect = v
}

func (o SQLViewRepresentation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SQLViewRepresentation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["sql"] = o.Sql
	toSerialize["dialect"] = o.Dialect
	return toSerialize, nil
}

func (o *SQLViewRepresentation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"sql",
		"dialect",
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

	varSQLViewRepresentation := _SQLViewRepresentation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSQLViewRepresentation)

	if err != nil {
		return err
	}

	*o = SQLViewRepresentation(varSQLViewRepresentation)

	return err
}

type NullableSQLViewRepresentation struct {
	value *SQLViewRepresentation
	isSet bool
}

func (v NullableSQLViewRepresentation) Get() *SQLViewRepresentation {
	return v.value
}

func (v *NullableSQLViewRepresentation) Set(val *SQLViewRepresentation) {
	v.value = val
	v.isSet = true
}

func (v NullableSQLViewRepresentation) IsSet() bool {
	return v.isSet
}

func (v *NullableSQLViewRepresentation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSQLViewRepresentation(val *SQLViewRepresentation) *NullableSQLViewRepresentation {
	return &NullableSQLViewRepresentation{value: val, isSet: true}
}

func (v NullableSQLViewRepresentation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSQLViewRepresentation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


