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

// checks if the GetProductConfigResponseProductConfigInnerArchiveSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProductConfigResponseProductConfigInnerArchiveSettings{}

// GetProductConfigResponseProductConfigInnerArchiveSettings Archive settings for collapses
type GetProductConfigResponseProductConfigInnerArchiveSettings struct {
	GlobalArchiveSettings *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings `json:"global_archive_settings,omitempty"`
	NetstorageConnectionSettings []GetProductConfigResponseProductConfigInnerArchiveSettingsNetstorageConnectionSettingsInner `json:"netstorage_connection_settings,omitempty"`
	S3ArchiveSettings []GetProductConfigResponseProductConfigInnerArchiveSettingsS3ArchiveSettingsInner `json:"s3_archive_settings,omitempty"`
}

// NewGetProductConfigResponseProductConfigInnerArchiveSettings instantiates a new GetProductConfigResponseProductConfigInnerArchiveSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductConfigResponseProductConfigInnerArchiveSettings() *GetProductConfigResponseProductConfigInnerArchiveSettings {
	this := GetProductConfigResponseProductConfigInnerArchiveSettings{}
	return &this
}

// NewGetProductConfigResponseProductConfigInnerArchiveSettingsWithDefaults instantiates a new GetProductConfigResponseProductConfigInnerArchiveSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductConfigResponseProductConfigInnerArchiveSettingsWithDefaults() *GetProductConfigResponseProductConfigInnerArchiveSettings {
	this := GetProductConfigResponseProductConfigInnerArchiveSettings{}
	return &this
}

// GetGlobalArchiveSettings returns the GlobalArchiveSettings field value if set, zero value otherwise.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettings) GetGlobalArchiveSettings() GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings {
	if o == nil || IsNil(o.GlobalArchiveSettings) {
		var ret GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings
		return ret
	}
	return *o.GlobalArchiveSettings
}

// GetGlobalArchiveSettingsOk returns a tuple with the GlobalArchiveSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettings) GetGlobalArchiveSettingsOk() (*GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings, bool) {
	if o == nil || IsNil(o.GlobalArchiveSettings) {
		return nil, false
	}
	return o.GlobalArchiveSettings, true
}

// HasGlobalArchiveSettings returns a boolean if a field has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettings) HasGlobalArchiveSettings() bool {
	if o != nil && !IsNil(o.GlobalArchiveSettings) {
		return true
	}

	return false
}

// SetGlobalArchiveSettings gets a reference to the given GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings and assigns it to the GlobalArchiveSettings field.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettings) SetGlobalArchiveSettings(v GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) {
	o.GlobalArchiveSettings = &v
}

// GetNetstorageConnectionSettings returns the NetstorageConnectionSettings field value if set, zero value otherwise.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettings) GetNetstorageConnectionSettings() []GetProductConfigResponseProductConfigInnerArchiveSettingsNetstorageConnectionSettingsInner {
	if o == nil || IsNil(o.NetstorageConnectionSettings) {
		var ret []GetProductConfigResponseProductConfigInnerArchiveSettingsNetstorageConnectionSettingsInner
		return ret
	}
	return o.NetstorageConnectionSettings
}

// GetNetstorageConnectionSettingsOk returns a tuple with the NetstorageConnectionSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettings) GetNetstorageConnectionSettingsOk() ([]GetProductConfigResponseProductConfigInnerArchiveSettingsNetstorageConnectionSettingsInner, bool) {
	if o == nil || IsNil(o.NetstorageConnectionSettings) {
		return nil, false
	}
	return o.NetstorageConnectionSettings, true
}

// HasNetstorageConnectionSettings returns a boolean if a field has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettings) HasNetstorageConnectionSettings() bool {
	if o != nil && !IsNil(o.NetstorageConnectionSettings) {
		return true
	}

	return false
}

// SetNetstorageConnectionSettings gets a reference to the given []GetProductConfigResponseProductConfigInnerArchiveSettingsNetstorageConnectionSettingsInner and assigns it to the NetstorageConnectionSettings field.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettings) SetNetstorageConnectionSettings(v []GetProductConfigResponseProductConfigInnerArchiveSettingsNetstorageConnectionSettingsInner) {
	o.NetstorageConnectionSettings = v
}

// GetS3ArchiveSettings returns the S3ArchiveSettings field value if set, zero value otherwise.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettings) GetS3ArchiveSettings() []GetProductConfigResponseProductConfigInnerArchiveSettingsS3ArchiveSettingsInner {
	if o == nil || IsNil(o.S3ArchiveSettings) {
		var ret []GetProductConfigResponseProductConfigInnerArchiveSettingsS3ArchiveSettingsInner
		return ret
	}
	return o.S3ArchiveSettings
}

// GetS3ArchiveSettingsOk returns a tuple with the S3ArchiveSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettings) GetS3ArchiveSettingsOk() ([]GetProductConfigResponseProductConfigInnerArchiveSettingsS3ArchiveSettingsInner, bool) {
	if o == nil || IsNil(o.S3ArchiveSettings) {
		return nil, false
	}
	return o.S3ArchiveSettings, true
}

// HasS3ArchiveSettings returns a boolean if a field has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettings) HasS3ArchiveSettings() bool {
	if o != nil && !IsNil(o.S3ArchiveSettings) {
		return true
	}

	return false
}

// SetS3ArchiveSettings gets a reference to the given []GetProductConfigResponseProductConfigInnerArchiveSettingsS3ArchiveSettingsInner and assigns it to the S3ArchiveSettings field.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettings) SetS3ArchiveSettings(v []GetProductConfigResponseProductConfigInnerArchiveSettingsS3ArchiveSettingsInner) {
	o.S3ArchiveSettings = v
}

func (o GetProductConfigResponseProductConfigInnerArchiveSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProductConfigResponseProductConfigInnerArchiveSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GlobalArchiveSettings) {
		toSerialize["global_archive_settings"] = o.GlobalArchiveSettings
	}
	if !IsNil(o.NetstorageConnectionSettings) {
		toSerialize["netstorage_connection_settings"] = o.NetstorageConnectionSettings
	}
	if !IsNil(o.S3ArchiveSettings) {
		toSerialize["s3_archive_settings"] = o.S3ArchiveSettings
	}
	return toSerialize, nil
}

type NullableGetProductConfigResponseProductConfigInnerArchiveSettings struct {
	value *GetProductConfigResponseProductConfigInnerArchiveSettings
	isSet bool
}

func (v NullableGetProductConfigResponseProductConfigInnerArchiveSettings) Get() *GetProductConfigResponseProductConfigInnerArchiveSettings {
	return v.value
}

func (v *NullableGetProductConfigResponseProductConfigInnerArchiveSettings) Set(val *GetProductConfigResponseProductConfigInnerArchiveSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductConfigResponseProductConfigInnerArchiveSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductConfigResponseProductConfigInnerArchiveSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductConfigResponseProductConfigInnerArchiveSettings(val *GetProductConfigResponseProductConfigInnerArchiveSettings) *NullableGetProductConfigResponseProductConfigInnerArchiveSettings {
	return &NullableGetProductConfigResponseProductConfigInnerArchiveSettings{value: val, isSet: true}
}

func (v NullableGetProductConfigResponseProductConfigInnerArchiveSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProductConfigResponseProductConfigInnerArchiveSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


