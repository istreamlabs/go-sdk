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

// checks if the ChannelPackagingPackagersValueContentProtectionSimple type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelPackagingPackagersValueContentProtectionSimple{}

// ChannelPackagingPackagersValueContentProtectionSimple Only one of ['simple', 'atlas', 'cpix'] may be set.
type ChannelPackagingPackagersValueContentProtectionSimple struct {
	// Pub points where keys should be published. If multiple are specified, only one needs to succeed to consider the key successfully published.
	PublishPoints []ChannelPackagingPackagersValueContentProtectionSimplePublishPointsInner `json:"publish_points,omitempty"`
	// Indicates which publish points must succeed for segment publishing to use the keys.
	RequirePublish *string `json:"require_publish,omitempty"`
}

// NewChannelPackagingPackagersValueContentProtectionSimple instantiates a new ChannelPackagingPackagersValueContentProtectionSimple object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPackagingPackagersValueContentProtectionSimple() *ChannelPackagingPackagersValueContentProtectionSimple {
	this := ChannelPackagingPackagersValueContentProtectionSimple{}
	return &this
}

// NewChannelPackagingPackagersValueContentProtectionSimpleWithDefaults instantiates a new ChannelPackagingPackagersValueContentProtectionSimple object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPackagingPackagersValueContentProtectionSimpleWithDefaults() *ChannelPackagingPackagersValueContentProtectionSimple {
	this := ChannelPackagingPackagersValueContentProtectionSimple{}
	return &this
}

// GetPublishPoints returns the PublishPoints field value if set, zero value otherwise.
func (o *ChannelPackagingPackagersValueContentProtectionSimple) GetPublishPoints() []ChannelPackagingPackagersValueContentProtectionSimplePublishPointsInner {
	if o == nil || IsNil(o.PublishPoints) {
		var ret []ChannelPackagingPackagersValueContentProtectionSimplePublishPointsInner
		return ret
	}
	return o.PublishPoints
}

// GetPublishPointsOk returns a tuple with the PublishPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPackagingPackagersValueContentProtectionSimple) GetPublishPointsOk() ([]ChannelPackagingPackagersValueContentProtectionSimplePublishPointsInner, bool) {
	if o == nil || IsNil(o.PublishPoints) {
		return nil, false
	}
	return o.PublishPoints, true
}

// HasPublishPoints returns a boolean if a field has been set.
func (o *ChannelPackagingPackagersValueContentProtectionSimple) HasPublishPoints() bool {
	if o != nil && !IsNil(o.PublishPoints) {
		return true
	}

	return false
}

// SetPublishPoints gets a reference to the given []ChannelPackagingPackagersValueContentProtectionSimplePublishPointsInner and assigns it to the PublishPoints field.
func (o *ChannelPackagingPackagersValueContentProtectionSimple) SetPublishPoints(v []ChannelPackagingPackagersValueContentProtectionSimplePublishPointsInner) {
	o.PublishPoints = v
}

// GetRequirePublish returns the RequirePublish field value if set, zero value otherwise.
func (o *ChannelPackagingPackagersValueContentProtectionSimple) GetRequirePublish() string {
	if o == nil || IsNil(o.RequirePublish) {
		var ret string
		return ret
	}
	return *o.RequirePublish
}

// GetRequirePublishOk returns a tuple with the RequirePublish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPackagingPackagersValueContentProtectionSimple) GetRequirePublishOk() (*string, bool) {
	if o == nil || IsNil(o.RequirePublish) {
		return nil, false
	}
	return o.RequirePublish, true
}

// HasRequirePublish returns a boolean if a field has been set.
func (o *ChannelPackagingPackagersValueContentProtectionSimple) HasRequirePublish() bool {
	if o != nil && !IsNil(o.RequirePublish) {
		return true
	}

	return false
}

// SetRequirePublish gets a reference to the given string and assigns it to the RequirePublish field.
func (o *ChannelPackagingPackagersValueContentProtectionSimple) SetRequirePublish(v string) {
	o.RequirePublish = &v
}

func (o ChannelPackagingPackagersValueContentProtectionSimple) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelPackagingPackagersValueContentProtectionSimple) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PublishPoints) {
		toSerialize["publish_points"] = o.PublishPoints
	}
	if !IsNil(o.RequirePublish) {
		toSerialize["require_publish"] = o.RequirePublish
	}
	return toSerialize, nil
}

type NullableChannelPackagingPackagersValueContentProtectionSimple struct {
	value *ChannelPackagingPackagersValueContentProtectionSimple
	isSet bool
}

func (v NullableChannelPackagingPackagersValueContentProtectionSimple) Get() *ChannelPackagingPackagersValueContentProtectionSimple {
	return v.value
}

func (v *NullableChannelPackagingPackagersValueContentProtectionSimple) Set(val *ChannelPackagingPackagersValueContentProtectionSimple) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPackagingPackagersValueContentProtectionSimple) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPackagingPackagersValueContentProtectionSimple) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPackagingPackagersValueContentProtectionSimple(val *ChannelPackagingPackagersValueContentProtectionSimple) *NullableChannelPackagingPackagersValueContentProtectionSimple {
	return &NullableChannelPackagingPackagersValueContentProtectionSimple{value: val, isSet: true}
}

func (v NullableChannelPackagingPackagersValueContentProtectionSimple) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPackagingPackagersValueContentProtectionSimple) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


