/*
 * Channel Lifecycle State API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package isp

import (
	"encoding/json"
)

// checks if the ErrorDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorDetail{}

// ErrorDetail struct for ErrorDetail
type ErrorDetail struct {
	// Where the error occurred, e.g. 'body.items[3].tags' or 'path.thing-id'
	Location interface{} `json:"location,omitempty" doc:"Where the error occurred, e.g. 'body.items[3].tags' or 'path.thing-id'"`
	// Error message text
	Message interface{} `json:"message,omitempty" doc:"Error message text"`
	// The value at the given location
	Value interface{} `json:"value,omitempty" doc:"The value at the given location"`
}

// NewErrorDetail instantiates a new ErrorDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorDetail() *ErrorDetail {
	this := ErrorDetail{}
	return &this
}

// NewErrorDetailWithDefaults instantiates a new ErrorDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorDetailWithDefaults() *ErrorDetail {
	this := ErrorDetail{}
	return &this
}

// GetLocation returns the Location field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErrorDetail) GetLocation() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorDetail) GetLocationOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return &o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *ErrorDetail) HasLocation() bool {
	if o != nil && IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given interface{} and assigns it to the Location field.
func (o *ErrorDetail) SetLocation(v interface{}) {
	o.Location = v
}

// GetMessage returns the Message field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErrorDetail) GetMessage() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorDetail) GetMessageOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return &o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ErrorDetail) HasMessage() bool {
	if o != nil && IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given interface{} and assigns it to the Message field.
func (o *ErrorDetail) SetMessage(v interface{}) {
	o.Message = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErrorDetail) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorDetail) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ErrorDetail) HasValue() bool {
	if o != nil && IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *ErrorDetail) SetValue(v interface{}) {
	o.Value = v
}

func (o ErrorDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorDetail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableErrorDetail struct {
	value *ErrorDetail
	isSet bool
}

func (v NullableErrorDetail) Get() *ErrorDetail {
	return v.value
}

func (v *NullableErrorDetail) Set(val *ErrorDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorDetail(val *ErrorDetail) *NullableErrorDetail {
	return &NullableErrorDetail{value: val, isSet: true}
}

func (v NullableErrorDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
