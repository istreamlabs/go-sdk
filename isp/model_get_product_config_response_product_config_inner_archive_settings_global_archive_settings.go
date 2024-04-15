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

// checks if the GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings{}

// GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings struct for GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings
type GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings struct {
	AesKeyConversion *string `json:"aes_key_conversion,omitempty"`
	MasterUrlType *string `json:"master_url_type,omitempty"`
	MediaUrlType *string `json:"media_url_type,omitempty"`
	Origins *map[string]GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettingsOriginsValue `json:"origins,omitempty"`
	PreservePublishedDirectoryStructure *bool `json:"preserve_published_directory_structure,omitempty"`
	UrlType *string `json:"url_type,omitempty"`
}

// NewGetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings instantiates a new GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings() *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings {
	this := GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings{}
	return &this
}

// NewGetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettingsWithDefaults instantiates a new GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettingsWithDefaults() *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings {
	this := GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings{}
	return &this
}

// GetAesKeyConversion returns the AesKeyConversion field value if set, zero value otherwise.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) GetAesKeyConversion() string {
	if o == nil || IsNil(o.AesKeyConversion) {
		var ret string
		return ret
	}
	return *o.AesKeyConversion
}

// GetAesKeyConversionOk returns a tuple with the AesKeyConversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) GetAesKeyConversionOk() (*string, bool) {
	if o == nil || IsNil(o.AesKeyConversion) {
		return nil, false
	}
	return o.AesKeyConversion, true
}

// HasAesKeyConversion returns a boolean if a field has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) HasAesKeyConversion() bool {
	if o != nil && !IsNil(o.AesKeyConversion) {
		return true
	}

	return false
}

// SetAesKeyConversion gets a reference to the given string and assigns it to the AesKeyConversion field.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) SetAesKeyConversion(v string) {
	o.AesKeyConversion = &v
}

// GetMasterUrlType returns the MasterUrlType field value if set, zero value otherwise.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) GetMasterUrlType() string {
	if o == nil || IsNil(o.MasterUrlType) {
		var ret string
		return ret
	}
	return *o.MasterUrlType
}

// GetMasterUrlTypeOk returns a tuple with the MasterUrlType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) GetMasterUrlTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MasterUrlType) {
		return nil, false
	}
	return o.MasterUrlType, true
}

// HasMasterUrlType returns a boolean if a field has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) HasMasterUrlType() bool {
	if o != nil && !IsNil(o.MasterUrlType) {
		return true
	}

	return false
}

// SetMasterUrlType gets a reference to the given string and assigns it to the MasterUrlType field.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) SetMasterUrlType(v string) {
	o.MasterUrlType = &v
}

// GetMediaUrlType returns the MediaUrlType field value if set, zero value otherwise.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) GetMediaUrlType() string {
	if o == nil || IsNil(o.MediaUrlType) {
		var ret string
		return ret
	}
	return *o.MediaUrlType
}

// GetMediaUrlTypeOk returns a tuple with the MediaUrlType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) GetMediaUrlTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MediaUrlType) {
		return nil, false
	}
	return o.MediaUrlType, true
}

// HasMediaUrlType returns a boolean if a field has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) HasMediaUrlType() bool {
	if o != nil && !IsNil(o.MediaUrlType) {
		return true
	}

	return false
}

// SetMediaUrlType gets a reference to the given string and assigns it to the MediaUrlType field.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) SetMediaUrlType(v string) {
	o.MediaUrlType = &v
}

// GetOrigins returns the Origins field value if set, zero value otherwise.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) GetOrigins() map[string]GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettingsOriginsValue {
	if o == nil || IsNil(o.Origins) {
		var ret map[string]GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettingsOriginsValue
		return ret
	}
	return *o.Origins
}

// GetOriginsOk returns a tuple with the Origins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) GetOriginsOk() (*map[string]GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettingsOriginsValue, bool) {
	if o == nil || IsNil(o.Origins) {
		return nil, false
	}
	return o.Origins, true
}

// HasOrigins returns a boolean if a field has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) HasOrigins() bool {
	if o != nil && !IsNil(o.Origins) {
		return true
	}

	return false
}

// SetOrigins gets a reference to the given map[string]GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettingsOriginsValue and assigns it to the Origins field.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) SetOrigins(v map[string]GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettingsOriginsValue) {
	o.Origins = &v
}

// GetPreservePublishedDirectoryStructure returns the PreservePublishedDirectoryStructure field value if set, zero value otherwise.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) GetPreservePublishedDirectoryStructure() bool {
	if o == nil || IsNil(o.PreservePublishedDirectoryStructure) {
		var ret bool
		return ret
	}
	return *o.PreservePublishedDirectoryStructure
}

// GetPreservePublishedDirectoryStructureOk returns a tuple with the PreservePublishedDirectoryStructure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) GetPreservePublishedDirectoryStructureOk() (*bool, bool) {
	if o == nil || IsNil(o.PreservePublishedDirectoryStructure) {
		return nil, false
	}
	return o.PreservePublishedDirectoryStructure, true
}

// HasPreservePublishedDirectoryStructure returns a boolean if a field has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) HasPreservePublishedDirectoryStructure() bool {
	if o != nil && !IsNil(o.PreservePublishedDirectoryStructure) {
		return true
	}

	return false
}

// SetPreservePublishedDirectoryStructure gets a reference to the given bool and assigns it to the PreservePublishedDirectoryStructure field.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) SetPreservePublishedDirectoryStructure(v bool) {
	o.PreservePublishedDirectoryStructure = &v
}

// GetUrlType returns the UrlType field value if set, zero value otherwise.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) GetUrlType() string {
	if o == nil || IsNil(o.UrlType) {
		var ret string
		return ret
	}
	return *o.UrlType
}

// GetUrlTypeOk returns a tuple with the UrlType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) GetUrlTypeOk() (*string, bool) {
	if o == nil || IsNil(o.UrlType) {
		return nil, false
	}
	return o.UrlType, true
}

// HasUrlType returns a boolean if a field has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) HasUrlType() bool {
	if o != nil && !IsNil(o.UrlType) {
		return true
	}

	return false
}

// SetUrlType gets a reference to the given string and assigns it to the UrlType field.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) SetUrlType(v string) {
	o.UrlType = &v
}

func (o GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) ToMap() (map[string]interface{}, error) {
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

type NullableGetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings struct {
	value *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings
	isSet bool
}

func (v NullableGetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) Get() *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings {
	return v.value
}

func (v *NullableGetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) Set(val *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings(val *GetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) *NullableGetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings {
	return &NullableGetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings{value: val, isSet: true}
}

func (v NullableGetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProductConfigResponseProductConfigInnerArchiveSettingsGlobalArchiveSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


