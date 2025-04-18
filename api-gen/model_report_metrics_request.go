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


// ReportMetricsRequest struct for ReportMetricsRequest
type ReportMetricsRequest struct {
	CommitReport *CommitReport
	ScanReport *ScanReport
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReportMetricsRequest) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into CommitReport
	err = json.Unmarshal(data, &dst.CommitReport);
	if err == nil {
		jsonCommitReport, _ := json.Marshal(dst.CommitReport)
		if string(jsonCommitReport) == "{}" { // empty struct
			dst.CommitReport = nil
		} else {
			return nil // data stored in dst.CommitReport, return on the first match
		}
	} else {
		dst.CommitReport = nil
	}

	// try to unmarshal JSON data into ScanReport
	err = json.Unmarshal(data, &dst.ScanReport);
	if err == nil {
		jsonScanReport, _ := json.Marshal(dst.ScanReport)
		if string(jsonScanReport) == "{}" { // empty struct
			dst.ScanReport = nil
		} else {
			return nil // data stored in dst.ScanReport, return on the first match
		}
	} else {
		dst.ScanReport = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ReportMetricsRequest)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ReportMetricsRequest) MarshalJSON() ([]byte, error) {
	if src.CommitReport != nil {
		return json.Marshal(&src.CommitReport)
	}

	if src.ScanReport != nil {
		return json.Marshal(&src.ScanReport)
	}

	return nil, nil // no data in anyOf schemas
}


type NullableReportMetricsRequest struct {
	value *ReportMetricsRequest
	isSet bool
}

func (v NullableReportMetricsRequest) Get() *ReportMetricsRequest {
	return v.value
}

func (v *NullableReportMetricsRequest) Set(val *ReportMetricsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableReportMetricsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableReportMetricsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportMetricsRequest(val *ReportMetricsRequest) *NullableReportMetricsRequest {
	return &NullableReportMetricsRequest{value: val, isSet: true}
}

func (v NullableReportMetricsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportMetricsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


