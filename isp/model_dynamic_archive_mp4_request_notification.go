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

// checks if the DynamicArchiveMP4RequestNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DynamicArchiveMP4RequestNotification{}

// DynamicArchiveMP4RequestNotification Notification Settings
type DynamicArchiveMP4RequestNotification struct {
	Sns ArchiveFERRequestNotificationSns `json:"sns"`
}

// NewDynamicArchiveMP4RequestNotification instantiates a new DynamicArchiveMP4RequestNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicArchiveMP4RequestNotification(sns ArchiveFERRequestNotificationSns) *DynamicArchiveMP4RequestNotification {
	this := DynamicArchiveMP4RequestNotification{}
	this.Sns = sns
	return &this
}

// NewDynamicArchiveMP4RequestNotificationWithDefaults instantiates a new DynamicArchiveMP4RequestNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicArchiveMP4RequestNotificationWithDefaults() *DynamicArchiveMP4RequestNotification {
	this := DynamicArchiveMP4RequestNotification{}
	return &this
}

// GetSns returns the Sns field value
func (o *DynamicArchiveMP4RequestNotification) GetSns() ArchiveFERRequestNotificationSns {
	if o == nil {
		var ret ArchiveFERRequestNotificationSns
		return ret
	}

	return o.Sns
}

// GetSnsOk returns a tuple with the Sns field value
// and a boolean to check if the value has been set.
func (o *DynamicArchiveMP4RequestNotification) GetSnsOk() (*ArchiveFERRequestNotificationSns, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Sns, true
}

// SetSns sets field value
func (o *DynamicArchiveMP4RequestNotification) SetSns(v ArchiveFERRequestNotificationSns) {
	o.Sns = v
}

func (o DynamicArchiveMP4RequestNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DynamicArchiveMP4RequestNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sns"] = o.Sns
	return toSerialize, nil
}

type NullableDynamicArchiveMP4RequestNotification struct {
	value *DynamicArchiveMP4RequestNotification
	isSet bool
}

func (v NullableDynamicArchiveMP4RequestNotification) Get() *DynamicArchiveMP4RequestNotification {
	return v.value
}

func (v *NullableDynamicArchiveMP4RequestNotification) Set(val *DynamicArchiveMP4RequestNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicArchiveMP4RequestNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicArchiveMP4RequestNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicArchiveMP4RequestNotification(val *DynamicArchiveMP4RequestNotification) *NullableDynamicArchiveMP4RequestNotification {
	return &NullableDynamicArchiveMP4RequestNotification{value: val, isSet: true}
}

func (v NullableDynamicArchiveMP4RequestNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicArchiveMP4RequestNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

