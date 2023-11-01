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

// checks if the GetProductConfigResponseProductConfigInnerCollapseConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProductConfigResponseProductConfigInnerCollapseConfig{}

// GetProductConfigResponseProductConfigInnerCollapseConfig Collapse filter options
type GetProductConfigResponseProductConfigInnerCollapseConfig struct {
	AutoCopyMp4 bool `json:"auto_copy_mp4"`
	AutoCreateMp4 bool `json:"auto_create_mp4"`
	M3u8FiltersConfig GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig `json:"m3u8_filters_config"`
	UsePerUrlCollapseNotification bool `json:"use_per_url_collapse_notification"`
	V2FiltersConfig GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig `json:"v2_filters_config"`
}

// NewGetProductConfigResponseProductConfigInnerCollapseConfig instantiates a new GetProductConfigResponseProductConfigInnerCollapseConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductConfigResponseProductConfigInnerCollapseConfig(autoCopyMp4 bool, autoCreateMp4 bool, m3u8FiltersConfig GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig, usePerUrlCollapseNotification bool, v2FiltersConfig GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig) *GetProductConfigResponseProductConfigInnerCollapseConfig {
	this := GetProductConfigResponseProductConfigInnerCollapseConfig{}
	this.AutoCopyMp4 = autoCopyMp4
	this.AutoCreateMp4 = autoCreateMp4
	this.M3u8FiltersConfig = m3u8FiltersConfig
	this.UsePerUrlCollapseNotification = usePerUrlCollapseNotification
	this.V2FiltersConfig = v2FiltersConfig
	return &this
}

// NewGetProductConfigResponseProductConfigInnerCollapseConfigWithDefaults instantiates a new GetProductConfigResponseProductConfigInnerCollapseConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductConfigResponseProductConfigInnerCollapseConfigWithDefaults() *GetProductConfigResponseProductConfigInnerCollapseConfig {
	this := GetProductConfigResponseProductConfigInnerCollapseConfig{}
	return &this
}

// GetAutoCopyMp4 returns the AutoCopyMp4 field value
func (o *GetProductConfigResponseProductConfigInnerCollapseConfig) GetAutoCopyMp4() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoCopyMp4
}

// GetAutoCopyMp4Ok returns a tuple with the AutoCopyMp4 field value
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseConfig) GetAutoCopyMp4Ok() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoCopyMp4, true
}

// SetAutoCopyMp4 sets field value
func (o *GetProductConfigResponseProductConfigInnerCollapseConfig) SetAutoCopyMp4(v bool) {
	o.AutoCopyMp4 = v
}

// GetAutoCreateMp4 returns the AutoCreateMp4 field value
func (o *GetProductConfigResponseProductConfigInnerCollapseConfig) GetAutoCreateMp4() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoCreateMp4
}

// GetAutoCreateMp4Ok returns a tuple with the AutoCreateMp4 field value
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseConfig) GetAutoCreateMp4Ok() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoCreateMp4, true
}

// SetAutoCreateMp4 sets field value
func (o *GetProductConfigResponseProductConfigInnerCollapseConfig) SetAutoCreateMp4(v bool) {
	o.AutoCreateMp4 = v
}

// GetM3u8FiltersConfig returns the M3u8FiltersConfig field value
func (o *GetProductConfigResponseProductConfigInnerCollapseConfig) GetM3u8FiltersConfig() GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig {
	if o == nil {
		var ret GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig
		return ret
	}

	return o.M3u8FiltersConfig
}

// GetM3u8FiltersConfigOk returns a tuple with the M3u8FiltersConfig field value
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseConfig) GetM3u8FiltersConfigOk() (*GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.M3u8FiltersConfig, true
}

// SetM3u8FiltersConfig sets field value
func (o *GetProductConfigResponseProductConfigInnerCollapseConfig) SetM3u8FiltersConfig(v GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig) {
	o.M3u8FiltersConfig = v
}

// GetUsePerUrlCollapseNotification returns the UsePerUrlCollapseNotification field value
func (o *GetProductConfigResponseProductConfigInnerCollapseConfig) GetUsePerUrlCollapseNotification() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.UsePerUrlCollapseNotification
}

// GetUsePerUrlCollapseNotificationOk returns a tuple with the UsePerUrlCollapseNotification field value
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseConfig) GetUsePerUrlCollapseNotificationOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UsePerUrlCollapseNotification, true
}

// SetUsePerUrlCollapseNotification sets field value
func (o *GetProductConfigResponseProductConfigInnerCollapseConfig) SetUsePerUrlCollapseNotification(v bool) {
	o.UsePerUrlCollapseNotification = v
}

// GetV2FiltersConfig returns the V2FiltersConfig field value
func (o *GetProductConfigResponseProductConfigInnerCollapseConfig) GetV2FiltersConfig() GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig {
	if o == nil {
		var ret GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig
		return ret
	}

	return o.V2FiltersConfig
}

// GetV2FiltersConfigOk returns a tuple with the V2FiltersConfig field value
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseConfig) GetV2FiltersConfigOk() (*GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.V2FiltersConfig, true
}

// SetV2FiltersConfig sets field value
func (o *GetProductConfigResponseProductConfigInnerCollapseConfig) SetV2FiltersConfig(v GetProductConfigResponseProductConfigInnerCollapseConfigV2FiltersConfig) {
	o.V2FiltersConfig = v
}

func (o GetProductConfigResponseProductConfigInnerCollapseConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProductConfigResponseProductConfigInnerCollapseConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auto_copy_mp4"] = o.AutoCopyMp4
	toSerialize["auto_create_mp4"] = o.AutoCreateMp4
	toSerialize["m3u8_filters_config"] = o.M3u8FiltersConfig
	toSerialize["use_per_url_collapse_notification"] = o.UsePerUrlCollapseNotification
	toSerialize["v2_filters_config"] = o.V2FiltersConfig
	return toSerialize, nil
}

type NullableGetProductConfigResponseProductConfigInnerCollapseConfig struct {
	value *GetProductConfigResponseProductConfigInnerCollapseConfig
	isSet bool
}

func (v NullableGetProductConfigResponseProductConfigInnerCollapseConfig) Get() *GetProductConfigResponseProductConfigInnerCollapseConfig {
	return v.value
}

func (v *NullableGetProductConfigResponseProductConfigInnerCollapseConfig) Set(val *GetProductConfigResponseProductConfigInnerCollapseConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductConfigResponseProductConfigInnerCollapseConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductConfigResponseProductConfigInnerCollapseConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductConfigResponseProductConfigInnerCollapseConfig(val *GetProductConfigResponseProductConfigInnerCollapseConfig) *NullableGetProductConfigResponseProductConfigInnerCollapseConfig {
	return &NullableGetProductConfigResponseProductConfigInnerCollapseConfig{value: val, isSet: true}
}

func (v NullableGetProductConfigResponseProductConfigInnerCollapseConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProductConfigResponseProductConfigInnerCollapseConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

