/*
 * WBD Aventus Channels API
 *
 * API version: 0.0.0
 * Contact: live-control-plane-devs@wbd.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package isp

import (
	"encoding/json"
)

// checks if the RetryTaskResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RetryTaskResponse{}

// RetryTaskResponse struct for RetryTaskResponse
type RetryTaskResponse struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string                 `json:"$schema,omitempty" format:"uri" doc:"An optional URL to a JSON Schema document describing this resource"`
	Error  CancelTaskResponseError `json:"error"`
}

// NewRetryTaskResponse instantiates a new RetryTaskResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRetryTaskResponse(error_ CancelTaskResponseError) *RetryTaskResponse {
	this := RetryTaskResponse{}
	this.Error = error_
	return &this
}

// NewRetryTaskResponseWithDefaults instantiates a new RetryTaskResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRetryTaskResponseWithDefaults() *RetryTaskResponse {
	this := RetryTaskResponse{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *RetryTaskResponse) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RetryTaskResponse) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *RetryTaskResponse) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *RetryTaskResponse) SetSchema(v string) {
	o.Schema = &v
}

// GetError returns the Error field value
func (o *RetryTaskResponse) GetError() CancelTaskResponseError {
	if o == nil {
		var ret CancelTaskResponseError
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *RetryTaskResponse) GetErrorOk() (*CancelTaskResponseError, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *RetryTaskResponse) SetError(v CancelTaskResponseError) {
	o.Error = v
}

func (o RetryTaskResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RetryTaskResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["error"] = o.Error
	return toSerialize, nil
}

type NullableRetryTaskResponse struct {
	value *RetryTaskResponse
	isSet bool
}

func (v NullableRetryTaskResponse) Get() *RetryTaskResponse {
	return v.value
}

func (v *NullableRetryTaskResponse) Set(val *RetryTaskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRetryTaskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRetryTaskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetryTaskResponse(val *RetryTaskResponse) *NullableRetryTaskResponse {
	return &NullableRetryTaskResponse{value: val, isSet: true}
}

func (v NullableRetryTaskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetryTaskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
