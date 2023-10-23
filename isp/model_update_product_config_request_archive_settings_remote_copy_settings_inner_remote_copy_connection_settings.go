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

// checks if the UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings{}

// UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings struct for UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings
type UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings struct {
	Basedir string `json:"basedir"`
	Hostname string `json:"hostname"`
	Password *string `json:"password,omitempty"`
	Port *int32 `json:"port,omitempty" format:"int32"`
	Username string `json:"username"`
}

// NewUpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings instantiates a new UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings(basedir string, hostname string, username string) *UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings {
	this := UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings{}
	this.Basedir = basedir
	this.Hostname = hostname
	this.Username = username
	return &this
}

// NewUpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettingsWithDefaults instantiates a new UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettingsWithDefaults() *UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings {
	this := UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings{}
	return &this
}

// GetBasedir returns the Basedir field value
func (o *UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) GetBasedir() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Basedir
}

// GetBasedirOk returns a tuple with the Basedir field value
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) GetBasedirOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Basedir, true
}

// SetBasedir sets field value
func (o *UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) SetBasedir(v string) {
	o.Basedir = v
}

// GetHostname returns the Hostname field value
func (o *UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) GetHostname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) GetHostnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hostname, true
}

// SetHostname sets field value
func (o *UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) SetHostname(v string) {
	o.Hostname = v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) SetPassword(v string) {
	o.Password = &v
}

// GetPort returns the Port field value if set, zero value otherwise.
func (o *UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) GetPort() int32 {
	if o == nil || IsNil(o.Port) {
		var ret int32
		return ret
	}
	return *o.Port
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) GetPortOk() (*int32, bool) {
	if o == nil || IsNil(o.Port) {
		return nil, false
	}
	return o.Port, true
}

// HasPort returns a boolean if a field has been set.
func (o *UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) HasPort() bool {
	if o != nil && !IsNil(o.Port) {
		return true
	}

	return false
}

// SetPort gets a reference to the given int32 and assigns it to the Port field.
func (o *UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) SetPort(v int32) {
	o.Port = &v
}

// GetUsername returns the Username field value
func (o *UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) SetUsername(v string) {
	o.Username = v
}

func (o UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["basedir"] = o.Basedir
	toSerialize["hostname"] = o.Hostname
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.Port) {
		toSerialize["port"] = o.Port
	}
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

type NullableUpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings struct {
	value *UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings
	isSet bool
}

func (v NullableUpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) Get() *UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings {
	return v.value
}

func (v *NullableUpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) Set(val *UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings(val *UpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) *NullableUpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings {
	return &NullableUpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings{value: val, isSet: true}
}

func (v NullableUpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProductConfigRequestArchiveSettingsRemoteCopySettingsInnerRemoteCopyConnectionSettings) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


