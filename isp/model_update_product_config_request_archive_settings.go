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

// checks if the UpdateProductConfigRequestArchiveSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProductConfigRequestArchiveSettings{}

// UpdateProductConfigRequestArchiveSettings Archive settings for collapses
type UpdateProductConfigRequestArchiveSettings struct {
	GlobalArchiveSettings *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings `json:"global_archive_settings,omitempty"`
	NetstorageConnectionSettings []UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner `json:"netstorage_connection_settings,omitempty"`
	RemoteCopySettings []UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInner `json:"remote_copy_settings,omitempty"`
	S3ArchiveSettings []UpdateProductConfigRequestArchiveSettingsS3ArchiveSettingsInner `json:"s3_archive_settings,omitempty"`
}

// NewUpdateProductConfigRequestArchiveSettings instantiates a new UpdateProductConfigRequestArchiveSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProductConfigRequestArchiveSettings() *UpdateProductConfigRequestArchiveSettings {
	this := UpdateProductConfigRequestArchiveSettings{}
	return &this
}

// NewUpdateProductConfigRequestArchiveSettingsWithDefaults instantiates a new UpdateProductConfigRequestArchiveSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProductConfigRequestArchiveSettingsWithDefaults() *UpdateProductConfigRequestArchiveSettings {
	this := UpdateProductConfigRequestArchiveSettings{}
	return &this
}

// GetGlobalArchiveSettings returns the GlobalArchiveSettings field value if set, zero value otherwise.
func (o *UpdateProductConfigRequestArchiveSettings) GetGlobalArchiveSettings() UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings {
	if o == nil || IsNil(o.GlobalArchiveSettings) {
		var ret UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings
		return ret
	}
	return *o.GlobalArchiveSettings
}

// GetGlobalArchiveSettingsOk returns a tuple with the GlobalArchiveSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettings) GetGlobalArchiveSettingsOk() (*UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings, bool) {
	if o == nil || IsNil(o.GlobalArchiveSettings) {
		return nil, false
	}
	return o.GlobalArchiveSettings, true
}

// HasGlobalArchiveSettings returns a boolean if a field has been set.
func (o *UpdateProductConfigRequestArchiveSettings) HasGlobalArchiveSettings() bool {
	if o != nil && !IsNil(o.GlobalArchiveSettings) {
		return true
	}

	return false
}

// SetGlobalArchiveSettings gets a reference to the given UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings and assigns it to the GlobalArchiveSettings field.
func (o *UpdateProductConfigRequestArchiveSettings) SetGlobalArchiveSettings(v UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) {
	o.GlobalArchiveSettings = &v
}

// GetNetstorageConnectionSettings returns the NetstorageConnectionSettings field value if set, zero value otherwise.
func (o *UpdateProductConfigRequestArchiveSettings) GetNetstorageConnectionSettings() []UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner {
	if o == nil || IsNil(o.NetstorageConnectionSettings) {
		var ret []UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner
		return ret
	}
	return o.NetstorageConnectionSettings
}

// GetNetstorageConnectionSettingsOk returns a tuple with the NetstorageConnectionSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettings) GetNetstorageConnectionSettingsOk() ([]UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner, bool) {
	if o == nil || IsNil(o.NetstorageConnectionSettings) {
		return nil, false
	}
	return o.NetstorageConnectionSettings, true
}

// HasNetstorageConnectionSettings returns a boolean if a field has been set.
func (o *UpdateProductConfigRequestArchiveSettings) HasNetstorageConnectionSettings() bool {
	if o != nil && !IsNil(o.NetstorageConnectionSettings) {
		return true
	}

	return false
}

// SetNetstorageConnectionSettings gets a reference to the given []UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner and assigns it to the NetstorageConnectionSettings field.
func (o *UpdateProductConfigRequestArchiveSettings) SetNetstorageConnectionSettings(v []UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) {
	o.NetstorageConnectionSettings = v
}

// GetRemoteCopySettings returns the RemoteCopySettings field value if set, zero value otherwise.
func (o *UpdateProductConfigRequestArchiveSettings) GetRemoteCopySettings() []UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInner {
	if o == nil || IsNil(o.RemoteCopySettings) {
		var ret []UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInner
		return ret
	}
	return o.RemoteCopySettings
}

// GetRemoteCopySettingsOk returns a tuple with the RemoteCopySettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettings) GetRemoteCopySettingsOk() ([]UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInner, bool) {
	if o == nil || IsNil(o.RemoteCopySettings) {
		return nil, false
	}
	return o.RemoteCopySettings, true
}

// HasRemoteCopySettings returns a boolean if a field has been set.
func (o *UpdateProductConfigRequestArchiveSettings) HasRemoteCopySettings() bool {
	if o != nil && !IsNil(o.RemoteCopySettings) {
		return true
	}

	return false
}

// SetRemoteCopySettings gets a reference to the given []UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInner and assigns it to the RemoteCopySettings field.
func (o *UpdateProductConfigRequestArchiveSettings) SetRemoteCopySettings(v []UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInner) {
	o.RemoteCopySettings = v
}

// GetS3ArchiveSettings returns the S3ArchiveSettings field value if set, zero value otherwise.
func (o *UpdateProductConfigRequestArchiveSettings) GetS3ArchiveSettings() []UpdateProductConfigRequestArchiveSettingsS3ArchiveSettingsInner {
	if o == nil || IsNil(o.S3ArchiveSettings) {
		var ret []UpdateProductConfigRequestArchiveSettingsS3ArchiveSettingsInner
		return ret
	}
	return o.S3ArchiveSettings
}

// GetS3ArchiveSettingsOk returns a tuple with the S3ArchiveSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettings) GetS3ArchiveSettingsOk() ([]UpdateProductConfigRequestArchiveSettingsS3ArchiveSettingsInner, bool) {
	if o == nil || IsNil(o.S3ArchiveSettings) {
		return nil, false
	}
	return o.S3ArchiveSettings, true
}

// HasS3ArchiveSettings returns a boolean if a field has been set.
func (o *UpdateProductConfigRequestArchiveSettings) HasS3ArchiveSettings() bool {
	if o != nil && !IsNil(o.S3ArchiveSettings) {
		return true
	}

	return false
}

// SetS3ArchiveSettings gets a reference to the given []UpdateProductConfigRequestArchiveSettingsS3ArchiveSettingsInner and assigns it to the S3ArchiveSettings field.
func (o *UpdateProductConfigRequestArchiveSettings) SetS3ArchiveSettings(v []UpdateProductConfigRequestArchiveSettingsS3ArchiveSettingsInner) {
	o.S3ArchiveSettings = v
}

func (o UpdateProductConfigRequestArchiveSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProductConfigRequestArchiveSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GlobalArchiveSettings) {
		toSerialize["global_archive_settings"] = o.GlobalArchiveSettings
	}
	if !IsNil(o.NetstorageConnectionSettings) {
		toSerialize["netstorage_connection_settings"] = o.NetstorageConnectionSettings
	}
	if !IsNil(o.RemoteCopySettings) {
		toSerialize["remote_copy_settings"] = o.RemoteCopySettings
	}
	if !IsNil(o.S3ArchiveSettings) {
		toSerialize["s3_archive_settings"] = o.S3ArchiveSettings
	}
	return toSerialize, nil
}

type NullableUpdateProductConfigRequestArchiveSettings struct {
	value *UpdateProductConfigRequestArchiveSettings
	isSet bool
}

func (v NullableUpdateProductConfigRequestArchiveSettings) Get() *UpdateProductConfigRequestArchiveSettings {
	return v.value
}

func (v *NullableUpdateProductConfigRequestArchiveSettings) Set(val *UpdateProductConfigRequestArchiveSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProductConfigRequestArchiveSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProductConfigRequestArchiveSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProductConfigRequestArchiveSettings(val *UpdateProductConfigRequestArchiveSettings) *NullableUpdateProductConfigRequestArchiveSettings {
	return &NullableUpdateProductConfigRequestArchiveSettings{value: val, isSet: true}
}

func (v NullableUpdateProductConfigRequestArchiveSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProductConfigRequestArchiveSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

