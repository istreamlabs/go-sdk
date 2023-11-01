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

// checks if the GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings{}

// GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings struct for GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings
type GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings struct {
	CopyCollapsedArchive bool `json:"copy_collapsed_archive"`
	CopyRawArchive bool `json:"copy_raw_archive"`
}

// NewGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings instantiates a new GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings(copyCollapsedArchive bool, copyRawArchive bool) *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings {
	this := GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings{}
	this.CopyCollapsedArchive = copyCollapsedArchive
	this.CopyRawArchive = copyRawArchive
	return &this
}

// NewGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettingsWithDefaults instantiates a new GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettingsWithDefaults() *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings {
	this := GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings{}
	return &this
}

// GetCopyCollapsedArchive returns the CopyCollapsedArchive field value
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings) GetCopyCollapsedArchive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CopyCollapsedArchive
}

// GetCopyCollapsedArchiveOk returns a tuple with the CopyCollapsedArchive field value
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings) GetCopyCollapsedArchiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CopyCollapsedArchive, true
}

// SetCopyCollapsedArchive sets field value
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings) SetCopyCollapsedArchive(v bool) {
	o.CopyCollapsedArchive = v
}

// GetCopyRawArchive returns the CopyRawArchive field value
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings) GetCopyRawArchive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CopyRawArchive
}

// GetCopyRawArchiveOk returns a tuple with the CopyRawArchive field value
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings) GetCopyRawArchiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CopyRawArchive, true
}

// SetCopyRawArchive sets field value
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings) SetCopyRawArchive(v bool) {
	o.CopyRawArchive = v
}

func (o GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["copy_collapsed_archive"] = o.CopyCollapsedArchive
	toSerialize["copy_raw_archive"] = o.CopyRawArchive
	return toSerialize, nil
}

type NullableGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings struct {
	value *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings
	isSet bool
}

func (v NullableGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings) Get() *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings {
	return v.value
}

func (v *NullableGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings) Set(val *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings(val *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings) *NullableGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings {
	return &NullableGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings{value: val, isSet: true}
}

func (v NullableGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


