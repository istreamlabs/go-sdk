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

// checks if the StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue{}

// StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue struct for StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue
type StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue struct {
	X *float64 `json:"x,omitempty" format:"double"`
	Y *float64 `json:"y,omitempty" format:"double"`
}

// NewStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue instantiates a new StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue() *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue {
	this := StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue{}
	return &this
}

// NewStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlueWithDefaults instantiates a new StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlueWithDefaults() *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue {
	this := StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue{}
	return &this
}

// GetX returns the X field value if set, zero value otherwise.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue) GetX() float64 {
	if o == nil || IsNil(o.X) {
		var ret float64
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue) GetXOk() (*float64, bool) {
	if o == nil || IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue) HasX() bool {
	if o != nil && !IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given float64 and assigns it to the X field.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue) SetX(v float64) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue) GetY() float64 {
	if o == nil || IsNil(o.Y) {
		var ret float64
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue) GetYOk() (*float64, bool) {
	if o == nil || IsNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue) HasY() bool {
	if o != nil && !IsNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given float64 and assigns it to the Y field.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue) SetY(v float64) {
	o.Y = &v
}

func (o StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !IsNil(o.Y) {
		toSerialize["y"] = o.Y
	}
	return toSerialize, nil
}

type NullableStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue struct {
	value *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue
	isSet bool
}

func (v NullableStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue) Get() *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue {
	return v.value
}

func (v *NullableStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue) Set(val *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue(val *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue) *NullableStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue {
	return &NullableStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue{value: val, isSet: true}
}

func (v NullableStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
