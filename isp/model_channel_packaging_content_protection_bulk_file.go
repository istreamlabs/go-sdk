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

// ChannelPackagingContentProtectionBulkFile Only one of ['bulk_file', 'sample_aes', 'common'] may be set.
type ChannelPackagingContentProtectionBulkFile struct {
	// How often the IV should be rotated and how it should be created
	IvRotation *string `json:"iv_rotation,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChannelPackagingContentProtectionBulkFile ChannelPackagingContentProtectionBulkFile

// NewChannelPackagingContentProtectionBulkFile instantiates a new ChannelPackagingContentProtectionBulkFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPackagingContentProtectionBulkFile() *ChannelPackagingContentProtectionBulkFile {
	this := ChannelPackagingContentProtectionBulkFile{}
	return &this
}

// NewChannelPackagingContentProtectionBulkFileWithDefaults instantiates a new ChannelPackagingContentProtectionBulkFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPackagingContentProtectionBulkFileWithDefaults() *ChannelPackagingContentProtectionBulkFile {
	this := ChannelPackagingContentProtectionBulkFile{}
	return &this
}

// GetIvRotation returns the IvRotation field value if set, zero value otherwise.
func (o *ChannelPackagingContentProtectionBulkFile) GetIvRotation() string {
	if o == nil || o.IvRotation == nil {
		var ret string
		return ret
	}
	return *o.IvRotation
}

// GetIvRotationOk returns a tuple with the IvRotation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPackagingContentProtectionBulkFile) GetIvRotationOk() (*string, bool) {
	if o == nil || o.IvRotation == nil {
		return nil, false
	}
	return o.IvRotation, true
}

// HasIvRotation returns a boolean if a field has been set.
func (o *ChannelPackagingContentProtectionBulkFile) HasIvRotation() bool {
	if o != nil && o.IvRotation != nil {
		return true
	}

	return false
}

// SetIvRotation gets a reference to the given string and assigns it to the IvRotation field.
func (o *ChannelPackagingContentProtectionBulkFile) SetIvRotation(v string) {
	o.IvRotation = &v
}

func (o ChannelPackagingContentProtectionBulkFile) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.IvRotation != nil {
		toSerialize["iv_rotation"] = o.IvRotation
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChannelPackagingContentProtectionBulkFile) UnmarshalJSON(bytes []byte) (err error) {
	varChannelPackagingContentProtectionBulkFile := _ChannelPackagingContentProtectionBulkFile{}

	if err = json.Unmarshal(bytes, &varChannelPackagingContentProtectionBulkFile); err == nil {
		*o = ChannelPackagingContentProtectionBulkFile(varChannelPackagingContentProtectionBulkFile)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "iv_rotation")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChannelPackagingContentProtectionBulkFile struct {
	value *ChannelPackagingContentProtectionBulkFile
	isSet bool
}

func (v NullableChannelPackagingContentProtectionBulkFile) Get() *ChannelPackagingContentProtectionBulkFile {
	return v.value
}

func (v *NullableChannelPackagingContentProtectionBulkFile) Set(val *ChannelPackagingContentProtectionBulkFile) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPackagingContentProtectionBulkFile) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPackagingContentProtectionBulkFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPackagingContentProtectionBulkFile(val *ChannelPackagingContentProtectionBulkFile) *NullableChannelPackagingContentProtectionBulkFile {
	return &NullableChannelPackagingContentProtectionBulkFile{value: val, isSet: true}
}

func (v NullableChannelPackagingContentProtectionBulkFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPackagingContentProtectionBulkFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

