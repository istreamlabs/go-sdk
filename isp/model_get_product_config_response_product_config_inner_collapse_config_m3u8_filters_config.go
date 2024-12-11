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

// checks if the GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig{}

// GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig struct for GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig
type GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig struct {
	Links []GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInner `json:"links"`
}

// NewGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig instantiates a new GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig(links []GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInner) *GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig {
	this := GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig{}
	this.Links = links
	return &this
}

// NewGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigWithDefaults instantiates a new GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigWithDefaults() *GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig {
	this := GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig{}
	return &this
}

// GetLinks returns the Links field value
func (o *GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig) GetLinks() []GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInner {
	if o == nil {
		var ret []GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInner
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig) GetLinksOk() ([]GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig) SetLinks(v []GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInner) {
	o.Links = v
}

func (o GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig struct {
	value *GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig
	isSet bool
}

func (v NullableGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig) Get() *GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig {
	return v.value
}

func (v *NullableGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig) Set(val *GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig(val *GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig) *NullableGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig {
	return &NullableGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig{value: val, isSet: true}
}

func (v NullableGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
