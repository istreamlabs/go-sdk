/*
 * iStreamPlanet Channels API
 *
 * API version: 0.0.0
 * Contact: support@istreamplanet.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package isp

import (
	"encoding/json"
)

// checks if the GenericSignalResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericSignalResult{}

// GenericSignalResult struct for GenericSignalResult
type GenericSignalResult struct {
	// Error details from the signaling subsystem
	ErrorMessage string `json:"error_message" doc:"Error details from the signaling subsystem"`
	// Result of signal; 'accepted' means no error
	ResultCode string `json:"result_code" enum:"accepted,not_found,already_exists,invalid_argument,internal_error" doc:"Result of signal; 'accepted' means no error"`
}

// NewGenericSignalResult instantiates a new GenericSignalResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericSignalResult(errorMessage string, resultCode string) *GenericSignalResult {
	this := GenericSignalResult{}
	this.ErrorMessage = errorMessage
	this.ResultCode = resultCode
	return &this
}

// NewGenericSignalResultWithDefaults instantiates a new GenericSignalResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericSignalResultWithDefaults() *GenericSignalResult {
	this := GenericSignalResult{}
	return &this
}

// GetErrorMessage returns the ErrorMessage field value
func (o *GenericSignalResult) GetErrorMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ErrorMessage
}

// GetErrorMessageOk returns a tuple with the ErrorMessage field value
// and a boolean to check if the value has been set.
func (o *GenericSignalResult) GetErrorMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ErrorMessage, true
}

// SetErrorMessage sets field value
func (o *GenericSignalResult) SetErrorMessage(v string) {
	o.ErrorMessage = v
}

// GetResultCode returns the ResultCode field value
func (o *GenericSignalResult) GetResultCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value
// and a boolean to check if the value has been set.
func (o *GenericSignalResult) GetResultCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ResultCode, true
}

// SetResultCode sets field value
func (o *GenericSignalResult) SetResultCode(v string) {
	o.ResultCode = v
}

func (o GenericSignalResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericSignalResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["error_message"] = o.ErrorMessage
	toSerialize["result_code"] = o.ResultCode
	return toSerialize, nil
}

type NullableGenericSignalResult struct {
	value *GenericSignalResult
	isSet bool
}

func (v NullableGenericSignalResult) Get() *GenericSignalResult {
	return v.value
}

func (v *NullableGenericSignalResult) Set(val *GenericSignalResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericSignalResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericSignalResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericSignalResult(val *GenericSignalResult) *NullableGenericSignalResult {
	return &NullableGenericSignalResult{value: val, isSet: true}
}

func (v NullableGenericSignalResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericSignalResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


