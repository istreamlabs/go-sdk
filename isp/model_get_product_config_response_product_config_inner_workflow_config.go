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

// checks if the GetProductConfigResponseProductConfigInnerWorkflowConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProductConfigResponseProductConfigInnerWorkflowConfig{}

// GetProductConfigResponseProductConfigInnerWorkflowConfig Config variables for workflows
type GetProductConfigResponseProductConfigInnerWorkflowConfig struct {
	CutprogramPaddingInSeconds *int32 `json:"cutprogram_padding_in_seconds,omitempty" format:"int32"`
	MakeMp4 *bool `json:"make_mp4,omitempty"`
	Mp4FeatureFlag *string `json:"mp4-feature-flag,omitempty"`
}

// NewGetProductConfigResponseProductConfigInnerWorkflowConfig instantiates a new GetProductConfigResponseProductConfigInnerWorkflowConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductConfigResponseProductConfigInnerWorkflowConfig() *GetProductConfigResponseProductConfigInnerWorkflowConfig {
	this := GetProductConfigResponseProductConfigInnerWorkflowConfig{}
	return &this
}

// NewGetProductConfigResponseProductConfigInnerWorkflowConfigWithDefaults instantiates a new GetProductConfigResponseProductConfigInnerWorkflowConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductConfigResponseProductConfigInnerWorkflowConfigWithDefaults() *GetProductConfigResponseProductConfigInnerWorkflowConfig {
	this := GetProductConfigResponseProductConfigInnerWorkflowConfig{}
	return &this
}

// GetCutprogramPaddingInSeconds returns the CutprogramPaddingInSeconds field value if set, zero value otherwise.
func (o *GetProductConfigResponseProductConfigInnerWorkflowConfig) GetCutprogramPaddingInSeconds() int32 {
	if o == nil || IsNil(o.CutprogramPaddingInSeconds) {
		var ret int32
		return ret
	}
	return *o.CutprogramPaddingInSeconds
}

// GetCutprogramPaddingInSecondsOk returns a tuple with the CutprogramPaddingInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerWorkflowConfig) GetCutprogramPaddingInSecondsOk() (*int32, bool) {
	if o == nil || IsNil(o.CutprogramPaddingInSeconds) {
		return nil, false
	}
	return o.CutprogramPaddingInSeconds, true
}

// HasCutprogramPaddingInSeconds returns a boolean if a field has been set.
func (o *GetProductConfigResponseProductConfigInnerWorkflowConfig) HasCutprogramPaddingInSeconds() bool {
	if o != nil && !IsNil(o.CutprogramPaddingInSeconds) {
		return true
	}

	return false
}

// SetCutprogramPaddingInSeconds gets a reference to the given int32 and assigns it to the CutprogramPaddingInSeconds field.
func (o *GetProductConfigResponseProductConfigInnerWorkflowConfig) SetCutprogramPaddingInSeconds(v int32) {
	o.CutprogramPaddingInSeconds = &v
}

// GetMakeMp4 returns the MakeMp4 field value if set, zero value otherwise.
func (o *GetProductConfigResponseProductConfigInnerWorkflowConfig) GetMakeMp4() bool {
	if o == nil || IsNil(o.MakeMp4) {
		var ret bool
		return ret
	}
	return *o.MakeMp4
}

// GetMakeMp4Ok returns a tuple with the MakeMp4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerWorkflowConfig) GetMakeMp4Ok() (*bool, bool) {
	if o == nil || IsNil(o.MakeMp4) {
		return nil, false
	}
	return o.MakeMp4, true
}

// HasMakeMp4 returns a boolean if a field has been set.
func (o *GetProductConfigResponseProductConfigInnerWorkflowConfig) HasMakeMp4() bool {
	if o != nil && !IsNil(o.MakeMp4) {
		return true
	}

	return false
}

// SetMakeMp4 gets a reference to the given bool and assigns it to the MakeMp4 field.
func (o *GetProductConfigResponseProductConfigInnerWorkflowConfig) SetMakeMp4(v bool) {
	o.MakeMp4 = &v
}

// GetMp4FeatureFlag returns the Mp4FeatureFlag field value if set, zero value otherwise.
func (o *GetProductConfigResponseProductConfigInnerWorkflowConfig) GetMp4FeatureFlag() string {
	if o == nil || IsNil(o.Mp4FeatureFlag) {
		var ret string
		return ret
	}
	return *o.Mp4FeatureFlag
}

// GetMp4FeatureFlagOk returns a tuple with the Mp4FeatureFlag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerWorkflowConfig) GetMp4FeatureFlagOk() (*string, bool) {
	if o == nil || IsNil(o.Mp4FeatureFlag) {
		return nil, false
	}
	return o.Mp4FeatureFlag, true
}

// HasMp4FeatureFlag returns a boolean if a field has been set.
func (o *GetProductConfigResponseProductConfigInnerWorkflowConfig) HasMp4FeatureFlag() bool {
	if o != nil && !IsNil(o.Mp4FeatureFlag) {
		return true
	}

	return false
}

// SetMp4FeatureFlag gets a reference to the given string and assigns it to the Mp4FeatureFlag field.
func (o *GetProductConfigResponseProductConfigInnerWorkflowConfig) SetMp4FeatureFlag(v string) {
	o.Mp4FeatureFlag = &v
}

func (o GetProductConfigResponseProductConfigInnerWorkflowConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProductConfigResponseProductConfigInnerWorkflowConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CutprogramPaddingInSeconds) {
		toSerialize["cutprogram_padding_in_seconds"] = o.CutprogramPaddingInSeconds
	}
	if !IsNil(o.MakeMp4) {
		toSerialize["make_mp4"] = o.MakeMp4
	}
	if !IsNil(o.Mp4FeatureFlag) {
		toSerialize["mp4-feature-flag"] = o.Mp4FeatureFlag
	}
	return toSerialize, nil
}

type NullableGetProductConfigResponseProductConfigInnerWorkflowConfig struct {
	value *GetProductConfigResponseProductConfigInnerWorkflowConfig
	isSet bool
}

func (v NullableGetProductConfigResponseProductConfigInnerWorkflowConfig) Get() *GetProductConfigResponseProductConfigInnerWorkflowConfig {
	return v.value
}

func (v *NullableGetProductConfigResponseProductConfigInnerWorkflowConfig) Set(val *GetProductConfigResponseProductConfigInnerWorkflowConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductConfigResponseProductConfigInnerWorkflowConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductConfigResponseProductConfigInnerWorkflowConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductConfigResponseProductConfigInnerWorkflowConfig(val *GetProductConfigResponseProductConfigInnerWorkflowConfig) *NullableGetProductConfigResponseProductConfigInnerWorkflowConfig {
	return &NullableGetProductConfigResponseProductConfigInnerWorkflowConfig{value: val, isSet: true}
}

func (v NullableGetProductConfigResponseProductConfigInnerWorkflowConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProductConfigResponseProductConfigInnerWorkflowConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


