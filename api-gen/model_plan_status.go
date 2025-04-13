/*
Apache Iceberg REST Catalog API

Defines the specification for the first version of the REST Catalog API. Implementations should ideally support both Iceberg table specs v1 and v2, with priority given to v2.

API version: 0.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package icebergclient

import (
	"encoding/json"
	"fmt"
)

// PlanStatus Status of a server-side planning operation
type PlanStatus string

// List of PlanStatus
const (
	COMPLETED PlanStatus = "completed"
	SUBMITTED PlanStatus = "submitted"
	CANCELLED PlanStatus = "cancelled"
	FAILED PlanStatus = "failed"
)

// All allowed values of PlanStatus enum
var AllowedPlanStatusEnumValues = []PlanStatus{
	"completed",
	"submitted",
	"cancelled",
	"failed",
}

func (v *PlanStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PlanStatus(value)
	for _, existing := range AllowedPlanStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PlanStatus", value)
}

// NewPlanStatusFromValue returns a pointer to a valid PlanStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPlanStatusFromValue(v string) (*PlanStatus, error) {
	ev := PlanStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PlanStatus: valid values are %v", v, AllowedPlanStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PlanStatus) IsValid() bool {
	for _, existing := range AllowedPlanStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PlanStatus value
func (v PlanStatus) Ptr() *PlanStatus {
	return &v
}

type NullablePlanStatus struct {
	value *PlanStatus
	isSet bool
}

func (v NullablePlanStatus) Get() *PlanStatus {
	return v.value
}

func (v *NullablePlanStatus) Set(val *PlanStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanStatus(val *PlanStatus) *NullablePlanStatus {
	return &NullablePlanStatus{value: val, isSet: true}
}

func (v NullablePlanStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

