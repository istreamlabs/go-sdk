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

// checks if the UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings{}

// UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings struct for UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings
type UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings struct {
	AesKeyConversion *string `json:"aes_key_conversion,omitempty"`
	MasterUrlType *string `json:"master_url_type,omitempty"`
	MediaUrlType *string `json:"media_url_type,omitempty"`
	Origins *map[string]UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettingsOriginsValue `json:"origins,omitempty"`
	PreservePublishedDirectoryStructure *bool `json:"preserve_published_directory_structure,omitempty"`
	UrlType *string `json:"url_type,omitempty"`
}

// NewUpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings instantiates a new UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings() *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings {
	this := UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings{}
	return &this
}

// NewUpdateProductConfigRequestArchiveSettingsGlobalArchiveSettingsWithDefaults instantiates a new UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProductConfigRequestArchiveSettingsGlobalArchiveSettingsWithDefaults() *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings {
	this := UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings{}
	return &this
}

// GetAesKeyConversion returns the AesKeyConversion field value if set, zero value otherwise.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) GetAesKeyConversion() string {
	if o == nil || IsNil(o.AesKeyConversion) {
		var ret string
		return ret
	}
	return *o.AesKeyConversion
}

// GetAesKeyConversionOk returns a tuple with the AesKeyConversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) GetAesKeyConversionOk() (*string, bool) {
	if o == nil || IsNil(o.AesKeyConversion) {
		return nil, false
	}
	return o.AesKeyConversion, true
}

// HasAesKeyConversion returns a boolean if a field has been set.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) HasAesKeyConversion() bool {
	if o != nil && !IsNil(o.AesKeyConversion) {
		return true
	}

	return false
}

// SetAesKeyConversion gets a reference to the given string and assigns it to the AesKeyConversion field.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) SetAesKeyConversion(v string) {
	o.AesKeyConversion = &v
}

// GetMasterUrlType returns the MasterUrlType field value if set, zero value otherwise.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) GetMasterUrlType() string {
	if o == nil || IsNil(o.MasterUrlType) {
		var ret string
		return ret
	}
	return *o.MasterUrlType
}

// GetMasterUrlTypeOk returns a tuple with the MasterUrlType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) GetMasterUrlTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MasterUrlType) {
		return nil, false
	}
	return o.MasterUrlType, true
}

// HasMasterUrlType returns a boolean if a field has been set.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) HasMasterUrlType() bool {
	if o != nil && !IsNil(o.MasterUrlType) {
		return true
	}

	return false
}

// SetMasterUrlType gets a reference to the given string and assigns it to the MasterUrlType field.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) SetMasterUrlType(v string) {
	o.MasterUrlType = &v
}

// GetMediaUrlType returns the MediaUrlType field value if set, zero value otherwise.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) GetMediaUrlType() string {
	if o == nil || IsNil(o.MediaUrlType) {
		var ret string
		return ret
	}
	return *o.MediaUrlType
}

// GetMediaUrlTypeOk returns a tuple with the MediaUrlType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) GetMediaUrlTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MediaUrlType) {
		return nil, false
	}
	return o.MediaUrlType, true
}

// HasMediaUrlType returns a boolean if a field has been set.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) HasMediaUrlType() bool {
	if o != nil && !IsNil(o.MediaUrlType) {
		return true
	}

	return false
}

// SetMediaUrlType gets a reference to the given string and assigns it to the MediaUrlType field.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) SetMediaUrlType(v string) {
	o.MediaUrlType = &v
}

// GetOrigins returns the Origins field value if set, zero value otherwise.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) GetOrigins() map[string]UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettingsOriginsValue {
	if o == nil || IsNil(o.Origins) {
		var ret map[string]UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettingsOriginsValue
		return ret
	}
	return *o.Origins
}

// GetOriginsOk returns a tuple with the Origins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) GetOriginsOk() (*map[string]UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettingsOriginsValue, bool) {
	if o == nil || IsNil(o.Origins) {
		return nil, false
	}
	return o.Origins, true
}

// HasOrigins returns a boolean if a field has been set.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) HasOrigins() bool {
	if o != nil && !IsNil(o.Origins) {
		return true
	}

	return false
}

// SetOrigins gets a reference to the given map[string]UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettingsOriginsValue and assigns it to the Origins field.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) SetOrigins(v map[string]UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettingsOriginsValue) {
	o.Origins = &v
}

// GetPreservePublishedDirectoryStructure returns the PreservePublishedDirectoryStructure field value if set, zero value otherwise.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) GetPreservePublishedDirectoryStructure() bool {
	if o == nil || IsNil(o.PreservePublishedDirectoryStructure) {
		var ret bool
		return ret
	}
	return *o.PreservePublishedDirectoryStructure
}

// GetPreservePublishedDirectoryStructureOk returns a tuple with the PreservePublishedDirectoryStructure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) GetPreservePublishedDirectoryStructureOk() (*bool, bool) {
	if o == nil || IsNil(o.PreservePublishedDirectoryStructure) {
		return nil, false
	}
	return o.PreservePublishedDirectoryStructure, true
}

// HasPreservePublishedDirectoryStructure returns a boolean if a field has been set.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) HasPreservePublishedDirectoryStructure() bool {
	if o != nil && !IsNil(o.PreservePublishedDirectoryStructure) {
		return true
	}

	return false
}

// SetPreservePublishedDirectoryStructure gets a reference to the given bool and assigns it to the PreservePublishedDirectoryStructure field.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) SetPreservePublishedDirectoryStructure(v bool) {
	o.PreservePublishedDirectoryStructure = &v
}

// GetUrlType returns the UrlType field value if set, zero value otherwise.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) GetUrlType() string {
	if o == nil || IsNil(o.UrlType) {
		var ret string
		return ret
	}
	return *o.UrlType
}

// GetUrlTypeOk returns a tuple with the UrlType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) GetUrlTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UrlType) {
		return nil, false
	}
	return o.UrlType, true
}

// HasUrlType returns a boolean if a field has been set.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) HasUrlType() bool {
	if o != nil && !IsNil(o.UrlType) {
		return true
	}

	return false
}

// SetUrlType gets a reference to the given string and assigns it to the UrlType field.
func (o *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) SetUrlType(v string) {
	o.UrlType = &v
}

func (o UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AesKeyConversion) {
		toSerialize["aes_key_conversion"] = o.AesKeyConversion
	}
	if !IsNil(o.MasterUrlType) {
		toSerialize["master_url_type"] = o.MasterUrlType
	}
	if !IsNil(o.MediaUrlType) {
		toSerialize["media_url_type"] = o.MediaUrlType
	}
	if !IsNil(o.Origins) {
		toSerialize["origins"] = o.Origins
	}
	if !IsNil(o.PreservePublishedDirectoryStructure) {
		toSerialize["preserve_published_directory_structure"] = o.PreservePublishedDirectoryStructure
	}
	if !IsNil(o.UrlType) {
		toSerialize["url_type"] = o.UrlType
	}
	return toSerialize, nil
}

type NullableUpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings struct {
	value *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings
	isSet bool
}

func (v NullableUpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) Get() *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings {
	return v.value
}

func (v *NullableUpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) Set(val *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings(val *UpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) *NullableUpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings {
	return &NullableUpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings{value: val, isSet: true}
}

func (v NullableUpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProductConfigRequestArchiveSettingsGlobalArchiveSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

