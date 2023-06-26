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

// checks if the UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner{}

// UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner struct for UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner
type UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner struct {
	AutoArchiveOnCollapseTypes []string `json:"auto_archive_on_collapse_types,omitempty"`
	AutoCopyToNetstorage bool `json:"auto_copy_to_netstorage"`
	Basedir *string `json:"basedir,omitempty"`
	CopyCollapsedArchive bool `json:"copy_collapsed_archive"`
	CopyRawArchive bool `json:"copy_raw_archive"`
	Cpcode string `json:"cpcode"`
	Hostname string `json:"hostname"`
	Key string `json:"key"`
	Keyname string `json:"keyname"`
	ManifestUpdates *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInnerManifestUpdates `json:"manifest_updates,omitempty"`
}

// NewUpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner instantiates a new UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner(autoCopyToNetstorage bool, copyCollapsedArchive bool, copyRawArchive bool, cpcode string, hostname string, key string, keyname string) *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner {
	this := UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner{}
	this.AutoCopyToNetstorage = autoCopyToNetstorage
	this.CopyCollapsedArchive = copyCollapsedArchive
	this.CopyRawArchive = copyRawArchive
	this.Cpcode = cpcode
	this.Hostname = hostname
	this.Key = key
	this.Keyname = keyname
	return &this
}

// NewUpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInnerWithDefaults instantiates a new UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInnerWithDefaults() *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner {
	this := UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner{}
	return &this
}

// GetAutoArchiveOnCollapseTypes returns the AutoArchiveOnCollapseTypes field value if set, zero value otherwise.
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) GetAutoArchiveOnCollapseTypes() []string {
	if o == nil || IsNil(o.AutoArchiveOnCollapseTypes) {
		var ret []string
		return ret
	}
	return o.AutoArchiveOnCollapseTypes
}

// GetAutoArchiveOnCollapseTypesOk returns a tuple with the AutoArchiveOnCollapseTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) GetAutoArchiveOnCollapseTypesOk() ([]string, bool) {
	if o == nil || IsNil(o.AutoArchiveOnCollapseTypes) {
		return nil, false
	}
	return o.AutoArchiveOnCollapseTypes, true
}

// HasAutoArchiveOnCollapseTypes returns a boolean if a field has been set.
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) HasAutoArchiveOnCollapseTypes() bool {
	if o != nil && !IsNil(o.AutoArchiveOnCollapseTypes) {
		return true
	}

	return false
}

// SetAutoArchiveOnCollapseTypes gets a reference to the given []string and assigns it to the AutoArchiveOnCollapseTypes field.
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) SetAutoArchiveOnCollapseTypes(v []string) {
	o.AutoArchiveOnCollapseTypes = v
}

// GetAutoCopyToNetstorage returns the AutoCopyToNetstorage field value
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) GetAutoCopyToNetstorage() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AutoCopyToNetstorage
}

// GetAutoCopyToNetstorageOk returns a tuple with the AutoCopyToNetstorage field value
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) GetAutoCopyToNetstorageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AutoCopyToNetstorage, true
}

// SetAutoCopyToNetstorage sets field value
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) SetAutoCopyToNetstorage(v bool) {
	o.AutoCopyToNetstorage = v
}

// GetBasedir returns the Basedir field value if set, zero value otherwise.
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) GetBasedir() string {
	if o == nil || IsNil(o.Basedir) {
		var ret string
		return ret
	}
	return *o.Basedir
}

// GetBasedirOk returns a tuple with the Basedir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) GetBasedirOk() (*string, bool) {
	if o == nil || IsNil(o.Basedir) {
		return nil, false
	}
	return o.Basedir, true
}

// HasBasedir returns a boolean if a field has been set.
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) HasBasedir() bool {
	if o != nil && !IsNil(o.Basedir) {
		return true
	}

	return false
}

// SetBasedir gets a reference to the given string and assigns it to the Basedir field.
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) SetBasedir(v string) {
	o.Basedir = &v
}

// GetCopyCollapsedArchive returns the CopyCollapsedArchive field value
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) GetCopyCollapsedArchive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CopyCollapsedArchive
}

// GetCopyCollapsedArchiveOk returns a tuple with the CopyCollapsedArchive field value
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) GetCopyCollapsedArchiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CopyCollapsedArchive, true
}

// SetCopyCollapsedArchive sets field value
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) SetCopyCollapsedArchive(v bool) {
	o.CopyCollapsedArchive = v
}

// GetCopyRawArchive returns the CopyRawArchive field value
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) GetCopyRawArchive() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CopyRawArchive
}

// GetCopyRawArchiveOk returns a tuple with the CopyRawArchive field value
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) GetCopyRawArchiveOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CopyRawArchive, true
}

// SetCopyRawArchive sets field value
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) SetCopyRawArchive(v bool) {
	o.CopyRawArchive = v
}

// GetCpcode returns the Cpcode field value
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) GetCpcode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Cpcode
}

// GetCpcodeOk returns a tuple with the Cpcode field value
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) GetCpcodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cpcode, true
}

// SetCpcode sets field value
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) SetCpcode(v string) {
	o.Cpcode = v
}

// GetHostname returns the Hostname field value
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) SetHostname(v string) {
	o.Hostname = v
}

// GetKey returns the Key field value
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Key, true
}

// SetKey sets field value
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) SetKey(v string) {
	o.Key = v
}

// GetKeyname returns the Keyname field value
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) GetKeyname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Keyname
}

// GetKeynameOk returns a tuple with the Keyname field value
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) GetKeynameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Keyname, true
}

// SetKeyname sets field value
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) SetKeyname(v string) {
	o.Keyname = v
}

// GetManifestUpdates returns the ManifestUpdates field value if set, zero value otherwise.
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) GetManifestUpdates() UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInnerManifestUpdates {
	if o == nil || IsNil(o.ManifestUpdates) {
		var ret UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInnerManifestUpdates
		return ret
	}
	return *o.ManifestUpdates
}

// GetManifestUpdatesOk returns a tuple with the ManifestUpdates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) GetManifestUpdatesOk() (*UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInnerManifestUpdates, bool) {
	if o == nil || IsNil(o.ManifestUpdates) {
		return nil, false
	}
	return o.ManifestUpdates, true
}

// HasManifestUpdates returns a boolean if a field has been set.
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) HasManifestUpdates() bool {
	if o != nil && !IsNil(o.ManifestUpdates) {
		return true
	}

	return false
}

// SetManifestUpdates gets a reference to the given UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInnerManifestUpdates and assigns it to the ManifestUpdates field.
func (o *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) SetManifestUpdates(v UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInnerManifestUpdates) {
	o.ManifestUpdates = &v
}

func (o UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoArchiveOnCollapseTypes) {
		toSerialize["auto_archive_on_collapse_types"] = o.AutoArchiveOnCollapseTypes
	}
	toSerialize["auto_copy_to_netstorage"] = o.AutoCopyToNetstorage
	if !IsNil(o.Basedir) {
		toSerialize["basedir"] = o.Basedir
	}
	toSerialize["copy_collapsed_archive"] = o.CopyCollapsedArchive
	toSerialize["copy_raw_archive"] = o.CopyRawArchive
	toSerialize["cpcode"] = o.Cpcode
	toSerialize["hostname"] = o.Hostname
	toSerialize["key"] = o.Key
	toSerialize["keyname"] = o.Keyname
	if !IsNil(o.ManifestUpdates) {
		toSerialize["manifest_updates"] = o.ManifestUpdates
	}
	return toSerialize, nil
}

type NullableUpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner struct {
	value *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner
	isSet bool
}

func (v NullableUpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) Get() *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner {
	return v.value
}

func (v *NullableUpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) Set(val *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner(val *UpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) *NullableUpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner {
	return &NullableUpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner{value: val, isSet: true}
}

func (v NullableUpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProductConfigRequestArchiveSettingsNetstorageConnectionSettingsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


