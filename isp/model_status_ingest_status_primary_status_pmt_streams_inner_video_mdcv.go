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

// checks if the StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv{}

// StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv struct for StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv
type StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv struct {
	Blue *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue `json:"blue,omitempty"`
	Green *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue `json:"green,omitempty"`
	MaxDisplayMasteringLuminance *float64 `json:"max_display_mastering_luminance,omitempty" format:"double"`
	MinDisplayMasteringLuminance *float64 `json:"min_display_mastering_luminance,omitempty" format:"double"`
	Red *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue `json:"red,omitempty"`
	WhitePoint *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue `json:"white_point,omitempty"`
}

// NewStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv instantiates a new StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv() *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv {
	this := StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv{}
	return &this
}

// NewStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvWithDefaults instantiates a new StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvWithDefaults() *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv {
	this := StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv{}
	return &this
}

// GetBlue returns the Blue field value if set, zero value otherwise.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) GetBlue() StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue {
	if o == nil || IsNil(o.Blue) {
		var ret StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue
		return ret
	}
	return *o.Blue
}

// GetBlueOk returns a tuple with the Blue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) GetBlueOk() (*StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue, bool) {
	if o == nil || IsNil(o.Blue) {
		return nil, false
	}
	return o.Blue, true
}

// HasBlue returns a boolean if a field has been set.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) HasBlue() bool {
	if o != nil && !IsNil(o.Blue) {
		return true
	}

	return false
}

// SetBlue gets a reference to the given StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue and assigns it to the Blue field.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) SetBlue(v StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue) {
	o.Blue = &v
}

// GetGreen returns the Green field value if set, zero value otherwise.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) GetGreen() StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue {
	if o == nil || IsNil(o.Green) {
		var ret StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue
		return ret
	}
	return *o.Green
}

// GetGreenOk returns a tuple with the Green field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) GetGreenOk() (*StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue, bool) {
	if o == nil || IsNil(o.Green) {
		return nil, false
	}
	return o.Green, true
}

// HasGreen returns a boolean if a field has been set.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) HasGreen() bool {
	if o != nil && !IsNil(o.Green) {
		return true
	}

	return false
}

// SetGreen gets a reference to the given StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue and assigns it to the Green field.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) SetGreen(v StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue) {
	o.Green = &v
}

// GetMaxDisplayMasteringLuminance returns the MaxDisplayMasteringLuminance field value if set, zero value otherwise.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) GetMaxDisplayMasteringLuminance() float64 {
	if o == nil || IsNil(o.MaxDisplayMasteringLuminance) {
		var ret float64
		return ret
	}
	return *o.MaxDisplayMasteringLuminance
}

// GetMaxDisplayMasteringLuminanceOk returns a tuple with the MaxDisplayMasteringLuminance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) GetMaxDisplayMasteringLuminanceOk() (*float64, bool) {
	if o == nil || IsNil(o.MaxDisplayMasteringLuminance) {
		return nil, false
	}
	return o.MaxDisplayMasteringLuminance, true
}

// HasMaxDisplayMasteringLuminance returns a boolean if a field has been set.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) HasMaxDisplayMasteringLuminance() bool {
	if o != nil && !IsNil(o.MaxDisplayMasteringLuminance) {
		return true
	}

	return false
}

// SetMaxDisplayMasteringLuminance gets a reference to the given float64 and assigns it to the MaxDisplayMasteringLuminance field.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) SetMaxDisplayMasteringLuminance(v float64) {
	o.MaxDisplayMasteringLuminance = &v
}

// GetMinDisplayMasteringLuminance returns the MinDisplayMasteringLuminance field value if set, zero value otherwise.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) GetMinDisplayMasteringLuminance() float64 {
	if o == nil || IsNil(o.MinDisplayMasteringLuminance) {
		var ret float64
		return ret
	}
	return *o.MinDisplayMasteringLuminance
}

// GetMinDisplayMasteringLuminanceOk returns a tuple with the MinDisplayMasteringLuminance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) GetMinDisplayMasteringLuminanceOk() (*float64, bool) {
	if o == nil || IsNil(o.MinDisplayMasteringLuminance) {
		return nil, false
	}
	return o.MinDisplayMasteringLuminance, true
}

// HasMinDisplayMasteringLuminance returns a boolean if a field has been set.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) HasMinDisplayMasteringLuminance() bool {
	if o != nil && !IsNil(o.MinDisplayMasteringLuminance) {
		return true
	}

	return false
}

// SetMinDisplayMasteringLuminance gets a reference to the given float64 and assigns it to the MinDisplayMasteringLuminance field.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) SetMinDisplayMasteringLuminance(v float64) {
	o.MinDisplayMasteringLuminance = &v
}

// GetRed returns the Red field value if set, zero value otherwise.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) GetRed() StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue {
	if o == nil || IsNil(o.Red) {
		var ret StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue
		return ret
	}
	return *o.Red
}

// GetRedOk returns a tuple with the Red field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) GetRedOk() (*StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue, bool) {
	if o == nil || IsNil(o.Red) {
		return nil, false
	}
	return o.Red, true
}

// HasRed returns a boolean if a field has been set.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) HasRed() bool {
	if o != nil && !IsNil(o.Red) {
		return true
	}

	return false
}

// SetRed gets a reference to the given StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue and assigns it to the Red field.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) SetRed(v StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue) {
	o.Red = &v
}

// GetWhitePoint returns the WhitePoint field value if set, zero value otherwise.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) GetWhitePoint() StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue {
	if o == nil || IsNil(o.WhitePoint) {
		var ret StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue
		return ret
	}
	return *o.WhitePoint
}

// GetWhitePointOk returns a tuple with the WhitePoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) GetWhitePointOk() (*StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue, bool) {
	if o == nil || IsNil(o.WhitePoint) {
		return nil, false
	}
	return o.WhitePoint, true
}

// HasWhitePoint returns a boolean if a field has been set.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) HasWhitePoint() bool {
	if o != nil && !IsNil(o.WhitePoint) {
		return true
	}

	return false
}

// SetWhitePoint gets a reference to the given StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue and assigns it to the WhitePoint field.
func (o *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) SetWhitePoint(v StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcvBlue) {
	o.WhitePoint = &v
}

func (o StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Blue) {
		toSerialize["blue"] = o.Blue
	}
	if !IsNil(o.Green) {
		toSerialize["green"] = o.Green
	}
	if !IsNil(o.MaxDisplayMasteringLuminance) {
		toSerialize["max_display_mastering_luminance"] = o.MaxDisplayMasteringLuminance
	}
	if !IsNil(o.MinDisplayMasteringLuminance) {
		toSerialize["min_display_mastering_luminance"] = o.MinDisplayMasteringLuminance
	}
	if !IsNil(o.Red) {
		toSerialize["red"] = o.Red
	}
	if !IsNil(o.WhitePoint) {
		toSerialize["white_point"] = o.WhitePoint
	}
	return toSerialize, nil
}

type NullableStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv struct {
	value *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv
	isSet bool
}

func (v NullableStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) Get() *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv {
	return v.value
}

func (v *NullableStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) Set(val *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv(val *StatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) *NullableStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv {
	return &NullableStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv{value: val, isSet: true}
}

func (v NullableStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusIngestStatusPrimaryStatusPmtStreamsInnerVideoMdcv) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


