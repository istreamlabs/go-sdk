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

// checks if the DeprecatedGetPresentationsResponseItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeprecatedGetPresentationsResponseItem{}

// DeprecatedGetPresentationsResponseItem struct for DeprecatedGetPresentationsResponseItem
type DeprecatedGetPresentationsResponseItem struct {
	Path string `json:"path"`
	Presid int64 `json:"presid"`
	Renditions []DeprecatedGetPresentationsResponseItemRenditionsInner `json:"renditions"`
	Storepath string `json:"storepath"`
}

// NewDeprecatedGetPresentationsResponseItem instantiates a new DeprecatedGetPresentationsResponseItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeprecatedGetPresentationsResponseItem(path string, presid int64, renditions []DeprecatedGetPresentationsResponseItemRenditionsInner, storepath string) *DeprecatedGetPresentationsResponseItem {
	this := DeprecatedGetPresentationsResponseItem{}
	this.Path = path
	this.Presid = presid
	this.Renditions = renditions
	this.Storepath = storepath
	return &this
}

// NewDeprecatedGetPresentationsResponseItemWithDefaults instantiates a new DeprecatedGetPresentationsResponseItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeprecatedGetPresentationsResponseItemWithDefaults() *DeprecatedGetPresentationsResponseItem {
	this := DeprecatedGetPresentationsResponseItem{}
	return &this
}

// GetPath returns the Path field value
func (o *DeprecatedGetPresentationsResponseItem) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *DeprecatedGetPresentationsResponseItem) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *DeprecatedGetPresentationsResponseItem) SetPath(v string) {
	o.Path = v
}

// GetPresid returns the Presid field value
func (o *DeprecatedGetPresentationsResponseItem) GetPresid() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Presid
}

// GetPresidOk returns a tuple with the Presid field value
// and a boolean to check if the value has been set.
func (o *DeprecatedGetPresentationsResponseItem) GetPresidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Presid, true
}

// SetPresid sets field value
func (o *DeprecatedGetPresentationsResponseItem) SetPresid(v int64) {
	o.Presid = v
}

// GetRenditions returns the Renditions field value
func (o *DeprecatedGetPresentationsResponseItem) GetRenditions() []DeprecatedGetPresentationsResponseItemRenditionsInner {
	if o == nil {
		var ret []DeprecatedGetPresentationsResponseItemRenditionsInner
		return ret
	}

	return o.Renditions
}

// GetRenditionsOk returns a tuple with the Renditions field value
// and a boolean to check if the value has been set.
func (o *DeprecatedGetPresentationsResponseItem) GetRenditionsOk() ([]DeprecatedGetPresentationsResponseItemRenditionsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Renditions, true
}

// SetRenditions sets field value
func (o *DeprecatedGetPresentationsResponseItem) SetRenditions(v []DeprecatedGetPresentationsResponseItemRenditionsInner) {
	o.Renditions = v
}

// GetStorepath returns the Storepath field value
func (o *DeprecatedGetPresentationsResponseItem) GetStorepath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Storepath
}

// GetStorepathOk returns a tuple with the Storepath field value
// and a boolean to check if the value has been set.
func (o *DeprecatedGetPresentationsResponseItem) GetStorepathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storepath, true
}

// SetStorepath sets field value
func (o *DeprecatedGetPresentationsResponseItem) SetStorepath(v string) {
	o.Storepath = v
}

func (o DeprecatedGetPresentationsResponseItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeprecatedGetPresentationsResponseItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["path"] = o.Path
	toSerialize["presid"] = o.Presid
	toSerialize["renditions"] = o.Renditions
	toSerialize["storepath"] = o.Storepath
	return toSerialize, nil
}

type NullableDeprecatedGetPresentationsResponseItem struct {
	value *DeprecatedGetPresentationsResponseItem
	isSet bool
}

func (v NullableDeprecatedGetPresentationsResponseItem) Get() *DeprecatedGetPresentationsResponseItem {
	return v.value
}

func (v *NullableDeprecatedGetPresentationsResponseItem) Set(val *DeprecatedGetPresentationsResponseItem) {
	v.value = val
	v.isSet = true
}

func (v NullableDeprecatedGetPresentationsResponseItem) IsSet() bool {
	return v.isSet
}

func (v *NullableDeprecatedGetPresentationsResponseItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeprecatedGetPresentationsResponseItem(val *DeprecatedGetPresentationsResponseItem) *NullableDeprecatedGetPresentationsResponseItem {
	return &NullableDeprecatedGetPresentationsResponseItem{value: val, isSet: true}
}

func (v NullableDeprecatedGetPresentationsResponseItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeprecatedGetPresentationsResponseItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

