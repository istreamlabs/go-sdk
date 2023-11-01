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

// checks if the GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner{}

// GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner struct for GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner
type GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner struct {
	ArchiveTargetTypes []string `json:"archive_target_types,omitempty"`
	AutoArchiveOnCollapseTypes []string `json:"auto_archive_on_collapse_types,omitempty"`
	AutoCopyToRemoteHost bool `json:"auto_copy_to_remote_host"`
	RemoteAutoCopySettings *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings `json:"remote_auto_copy_settings,omitempty"`
	RemoteCopyConnectionSettings GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings `json:"remote_copy_connection_settings"`
	RemoteCopyType string `json:"remote_copy_type"`
}

// NewGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner instantiates a new GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner(autoCopyToRemoteHost bool, remoteCopyConnectionSettings GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings, remoteCopyType string) *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner {
	this := GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner{}
	this.AutoCopyToRemoteHost = autoCopyToRemoteHost
	this.RemoteCopyConnectionSettings = remoteCopyConnectionSettings
	this.RemoteCopyType = remoteCopyType
	return &this
}

// NewGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerWithDefaults instantiates a new GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerWithDefaults() *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner {
	this := GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner{}
	return &this
}

// GetArchiveTargetTypes returns the ArchiveTargetTypes field value if set, zero value otherwise.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) GetArchiveTargetTypes() []string {
	if o == nil || IsNil(o.ArchiveTargetTypes) {
		var ret []string
		return ret
	}
	return o.ArchiveTargetTypes
}

// GetArchiveTargetTypesOk returns a tuple with the ArchiveTargetTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) GetArchiveTargetTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.ArchiveTargetTypes) {
		return nil, false
	}
	return o.ArchiveTargetTypes, true
}

// HasArchiveTargetTypes returns a boolean if a field has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) HasArchiveTargetTypes() bool {
	if o != nil && !IsNil(o.ArchiveTargetTypes) {
		return true
	}

	return false
}

// SetArchiveTargetTypes gets a reference to the given []string and assigns it to the ArchiveTargetTypes field.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) SetArchiveTargetTypes(v []string) {
	o.ArchiveTargetTypes = v
}

// GetAutoArchiveOnCollapseTypes returns the AutoArchiveOnCollapseTypes field value if set, zero value otherwise.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) GetAutoArchiveOnCollapseTypes() []string {
	if o == nil || IsNil(o.AutoArchiveOnCollapseTypes) {
		var ret []string
		return ret
	}
	return o.AutoArchiveOnCollapseTypes
}

// GetAutoArchiveOnCollapseTypesOk returns a tuple with the AutoArchiveOnCollapseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) GetAutoArchiveOnCollapseTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.AutoArchiveOnCollapseTypes) {
		return nil, false
	}
	return o.AutoArchiveOnCollapseTypes, true
}

// HasAutoArchiveOnCollapseTypes returns a boolean if a field has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) HasAutoArchiveOnCollapseTypes() bool {
	if o != nil && !IsNil(o.AutoArchiveOnCollapseTypes) {
		return true
	}

	return false
}

// SetAutoArchiveOnCollapseTypes gets a reference to the given []string and assigns it to the AutoArchiveOnCollapseTypes field.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) SetAutoArchiveOnCollapseTypes(v []string) {
	o.AutoArchiveOnCollapseTypes = v
}

// GetAutoCopyToRemoteHost returns the AutoCopyToRemoteHost field value
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) GetAutoCopyToRemoteHost() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoCopyToRemoteHost
}

// GetAutoCopyToRemoteHostOk returns a tuple with the AutoCopyToRemoteHost field value
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) GetAutoCopyToRemoteHostOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoCopyToRemoteHost, true
}

// SetAutoCopyToRemoteHost sets field value
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) SetAutoCopyToRemoteHost(v bool) {
	o.AutoCopyToRemoteHost = v
}

// GetRemoteAutoCopySettings returns the RemoteAutoCopySettings field value if set, zero value otherwise.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) GetRemoteAutoCopySettings() GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings {
	if o == nil || IsNil(o.RemoteAutoCopySettings) {
		var ret GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings
		return ret
	}
	return *o.RemoteAutoCopySettings
}

// GetRemoteAutoCopySettingsOk returns a tuple with the RemoteAutoCopySettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) GetRemoteAutoCopySettingsOk() (*GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings, bool) {
	if o == nil || IsNil(o.RemoteAutoCopySettings) {
		return nil, false
	}
	return o.RemoteAutoCopySettings, true
}

// HasRemoteAutoCopySettings returns a boolean if a field has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) HasRemoteAutoCopySettings() bool {
	if o != nil && !IsNil(o.RemoteAutoCopySettings) {
		return true
	}

	return false
}

// SetRemoteAutoCopySettings gets a reference to the given GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings and assigns it to the RemoteAutoCopySettings field.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) SetRemoteAutoCopySettings(v GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteAutoCopySettings) {
	o.RemoteAutoCopySettings = &v
}

// GetRemoteCopyConnectionSettings returns the RemoteCopyConnectionSettings field value
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) GetRemoteCopyConnectionSettings() GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings {
	if o == nil {
		var ret GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings
		return ret
	}

	return o.RemoteCopyConnectionSettings
}

// GetRemoteCopyConnectionSettingsOk returns a tuple with the RemoteCopyConnectionSettings field value
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) GetRemoteCopyConnectionSettingsOk() (*GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteCopyConnectionSettings, true
}

// SetRemoteCopyConnectionSettings sets field value
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) SetRemoteCopyConnectionSettings(v GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) {
	o.RemoteCopyConnectionSettings = v
}

// GetRemoteCopyType returns the RemoteCopyType field value
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) GetRemoteCopyType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RemoteCopyType
}

// GetRemoteCopyTypeOk returns a tuple with the RemoteCopyType field value
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) GetRemoteCopyTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RemoteCopyType, true
}

// SetRemoteCopyType sets field value
func (o *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) SetRemoteCopyType(v string) {
	o.RemoteCopyType = v
}

func (o GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ArchiveTargetTypes) {
		toSerialize["archive_target_types"] = o.ArchiveTargetTypes
	}
	if !IsNil(o.AutoArchiveOnCollapseTypes) {
		toSerialize["auto_archive_on_collapse_types"] = o.AutoArchiveOnCollapseTypes
	}
	toSerialize["auto_copy_to_remote_host"] = o.AutoCopyToRemoteHost
	if !IsNil(o.RemoteAutoCopySettings) {
		toSerialize["remote_auto_copy_settings"] = o.RemoteAutoCopySettings
	}
	toSerialize["remote_copy_connection_settings"] = o.RemoteCopyConnectionSettings
	toSerialize["remote_copy_type"] = o.RemoteCopyType
	return toSerialize, nil
}

type NullableGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner struct {
	value *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner
	isSet bool
}

func (v NullableGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) Get() *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner {
	return v.value
}

func (v *NullableGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) Set(val *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner(val *GetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) *NullableGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner {
	return &NullableGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner{value: val, isSet: true}
}

func (v NullableGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProductConfigResponseProductConfigInnerArchiveSettingsRemoteCopySettingsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

