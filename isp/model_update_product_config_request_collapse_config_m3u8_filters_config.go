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

// checks if the UpdateProductConfigRequestCollapseConfigM3u8FiltersConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProductConfigRequestCollapseConfigM3u8FiltersConfig{}

// UpdateProductConfigRequestCollapseConfigM3u8FiltersConfig struct for UpdateProductConfigRequestCollapseConfigM3u8FiltersConfig
type UpdateProductConfigRequestCollapseConfigM3u8FiltersConfig struct {
	Links []UpdateProductConfigRequestCollapseConfigM3u8FiltersConfigLinksInner `json:"links"`
}

// NewUpdateProductConfigRequestCollapseConfigM3u8FiltersConfig instantiates a new UpdateProductConfigRequestCollapseConfigM3u8FiltersConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProductConfigRequestCollapseConfigM3u8FiltersConfig(links []UpdateProductConfigRequestCollapseConfigM3u8FiltersConfigLinksInner) *UpdateProductConfigRequestCollapseConfigM3u8FiltersConfig {
	this := UpdateProductConfigRequestCollapseConfigM3u8FiltersConfig{}
	this.Links = links
	return &this
}

// NewUpdateProductConfigRequestCollapseConfigM3u8FiltersConfigWithDefaults instantiates a new UpdateProductConfigRequestCollapseConfigM3u8FiltersConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProductConfigRequestCollapseConfigM3u8FiltersConfigWithDefaults() *UpdateProductConfigRequestCollapseConfigM3u8FiltersConfig {
	this := UpdateProductConfigRequestCollapseConfigM3u8FiltersConfig{}
	return &this
}

// GetLinks returns the Links field value
func (o *UpdateProductConfigRequestCollapseConfigM3u8FiltersConfig) GetLinks() []UpdateProductConfigRequestCollapseConfigM3u8FiltersConfigLinksInner {
	if o == nil {
		var ret []UpdateProductConfigRequestCollapseConfigM3u8FiltersConfigLinksInner
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestCollapseConfigM3u8FiltersConfig) GetLinksOk() ([]UpdateProductConfigRequestCollapseConfigM3u8FiltersConfigLinksInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Links, true
}

// SetLinks sets field value
func (o *UpdateProductConfigRequestCollapseConfigM3u8FiltersConfig) SetLinks(v []UpdateProductConfigRequestCollapseConfigM3u8FiltersConfigLinksInner) {
	o.Links = v
}

func (o UpdateProductConfigRequestCollapseConfigM3u8FiltersConfig) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProductConfigRequestCollapseConfigM3u8FiltersConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableUpdateProductConfigRequestCollapseConfigM3u8FiltersConfig struct {
	value *UpdateProductConfigRequestCollapseConfigM3u8FiltersConfig
	isSet bool
}

func (v NullableUpdateProductConfigRequestCollapseConfigM3u8FiltersConfig) Get() *UpdateProductConfigRequestCollapseConfigM3u8FiltersConfig {
	return v.value
}

func (v *NullableUpdateProductConfigRequestCollapseConfigM3u8FiltersConfig) Set(val *UpdateProductConfigRequestCollapseConfigM3u8FiltersConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProductConfigRequestCollapseConfigM3u8FiltersConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProductConfigRequestCollapseConfigM3u8FiltersConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProductConfigRequestCollapseConfigM3u8FiltersConfig(val *UpdateProductConfigRequestCollapseConfigM3u8FiltersConfig) *NullableUpdateProductConfigRequestCollapseConfigM3u8FiltersConfig {
	return &NullableUpdateProductConfigRequestCollapseConfigM3u8FiltersConfig{value: val, isSet: true}
}

func (v NullableUpdateProductConfigRequestCollapseConfigM3u8FiltersConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProductConfigRequestCollapseConfigM3u8FiltersConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

