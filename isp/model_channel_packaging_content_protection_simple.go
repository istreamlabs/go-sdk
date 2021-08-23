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

// ChannelPackagingContentProtectionSimple Only one of ['simple', 'atlas'] may be set.
type ChannelPackagingContentProtectionSimple struct {
	// Pub points where keys should be published. If multiple are specified, only one needs to succeed to consider the key successfully published.
	PublishPoints *[]ChannelPackagingContentProtectionSimplePublishPoints `json:"publish_points,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChannelPackagingContentProtectionSimple ChannelPackagingContentProtectionSimple

// NewChannelPackagingContentProtectionSimple instantiates a new ChannelPackagingContentProtectionSimple object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPackagingContentProtectionSimple() *ChannelPackagingContentProtectionSimple {
	this := ChannelPackagingContentProtectionSimple{}
	return &this
}

// NewChannelPackagingContentProtectionSimpleWithDefaults instantiates a new ChannelPackagingContentProtectionSimple object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPackagingContentProtectionSimpleWithDefaults() *ChannelPackagingContentProtectionSimple {
	this := ChannelPackagingContentProtectionSimple{}
	return &this
}

// GetPublishPoints returns the PublishPoints field value if set, zero value otherwise.
func (o *ChannelPackagingContentProtectionSimple) GetPublishPoints() []ChannelPackagingContentProtectionSimplePublishPoints {
	if o == nil || o.PublishPoints == nil {
		var ret []ChannelPackagingContentProtectionSimplePublishPoints
		return ret
	}
	return *o.PublishPoints
}

// GetPublishPointsOk returns a tuple with the PublishPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPackagingContentProtectionSimple) GetPublishPointsOk() (*[]ChannelPackagingContentProtectionSimplePublishPoints, bool) {
	if o == nil || o.PublishPoints == nil {
		return nil, false
	}
	return o.PublishPoints, true
}

// HasPublishPoints returns a boolean if a field has been set.
func (o *ChannelPackagingContentProtectionSimple) HasPublishPoints() bool {
	if o != nil && o.PublishPoints != nil {
		return true
	}

	return false
}

// SetPublishPoints gets a reference to the given []ChannelPackagingContentProtectionSimplePublishPoints and assigns it to the PublishPoints field.
func (o *ChannelPackagingContentProtectionSimple) SetPublishPoints(v []ChannelPackagingContentProtectionSimplePublishPoints) {
	o.PublishPoints = &v
}

func (o ChannelPackagingContentProtectionSimple) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PublishPoints != nil {
		toSerialize["publish_points"] = o.PublishPoints
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChannelPackagingContentProtectionSimple) UnmarshalJSON(bytes []byte) (err error) {
	varChannelPackagingContentProtectionSimple := _ChannelPackagingContentProtectionSimple{}

	if err = json.Unmarshal(bytes, &varChannelPackagingContentProtectionSimple); err == nil {
		*o = ChannelPackagingContentProtectionSimple(varChannelPackagingContentProtectionSimple)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "publish_points")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChannelPackagingContentProtectionSimple struct {
	value *ChannelPackagingContentProtectionSimple
	isSet bool
}

func (v NullableChannelPackagingContentProtectionSimple) Get() *ChannelPackagingContentProtectionSimple {
	return v.value
}

func (v *NullableChannelPackagingContentProtectionSimple) Set(val *ChannelPackagingContentProtectionSimple) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPackagingContentProtectionSimple) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPackagingContentProtectionSimple) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPackagingContentProtectionSimple(val *ChannelPackagingContentProtectionSimple) *NullableChannelPackagingContentProtectionSimple {
	return &NullableChannelPackagingContentProtectionSimple{value: val, isSet: true}
}

func (v NullableChannelPackagingContentProtectionSimple) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPackagingContentProtectionSimple) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

