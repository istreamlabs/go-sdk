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

// checks if the DeprecatedGetPresentationsResponseItemRenditionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeprecatedGetPresentationsResponseItemRenditionsInner{}

// DeprecatedGetPresentationsResponseItemRenditionsInner struct for DeprecatedGetPresentationsResponseItemRenditionsInner
type DeprecatedGetPresentationsResponseItemRenditionsInner struct {
	Iframeonly bool `json:"iframeonly"`
	Metadata string `json:"metadata"`
	Path string `json:"path"`
	Ready bool `json:"ready"`
	Rendid int64 `json:"rendid" format:"int64"`
	Storepath string `json:"storepath"`
}

// NewDeprecatedGetPresentationsResponseItemRenditionsInner instantiates a new DeprecatedGetPresentationsResponseItemRenditionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeprecatedGetPresentationsResponseItemRenditionsInner(iframeonly bool, metadata string, path string, ready bool, rendid int64, storepath string) *DeprecatedGetPresentationsResponseItemRenditionsInner {
	this := DeprecatedGetPresentationsResponseItemRenditionsInner{}
	this.Iframeonly = iframeonly
	this.Metadata = metadata
	this.Path = path
	this.Ready = ready
	this.Rendid = rendid
	this.Storepath = storepath
	return &this
}

// NewDeprecatedGetPresentationsResponseItemRenditionsInnerWithDefaults instantiates a new DeprecatedGetPresentationsResponseItemRenditionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeprecatedGetPresentationsResponseItemRenditionsInnerWithDefaults() *DeprecatedGetPresentationsResponseItemRenditionsInner {
	this := DeprecatedGetPresentationsResponseItemRenditionsInner{}
	return &this
}

// GetIframeonly returns the Iframeonly field value
func (o *DeprecatedGetPresentationsResponseItemRenditionsInner) GetIframeonly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Iframeonly
}

// GetIframeonlyOk returns a tuple with the Iframeonly field value
// and a boolean to check if the value has been set.
func (o *DeprecatedGetPresentationsResponseItemRenditionsInner) GetIframeonlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Iframeonly, true
}

// SetIframeonly sets field value
func (o *DeprecatedGetPresentationsResponseItemRenditionsInner) SetIframeonly(v bool) {
	o.Iframeonly = v
}

// GetMetadata returns the Metadata field value
func (o *DeprecatedGetPresentationsResponseItemRenditionsInner) GetMetadata() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *DeprecatedGetPresentationsResponseItemRenditionsInner) GetMetadataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *DeprecatedGetPresentationsResponseItemRenditionsInner) SetMetadata(v string) {
	o.Metadata = v
}

// GetPath returns the Path field value
func (o *DeprecatedGetPresentationsResponseItemRenditionsInner) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *DeprecatedGetPresentationsResponseItemRenditionsInner) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *DeprecatedGetPresentationsResponseItemRenditionsInner) SetPath(v string) {
	o.Path = v
}

// GetReady returns the Ready field value
func (o *DeprecatedGetPresentationsResponseItemRenditionsInner) GetReady() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Ready
}

// GetReadyOk returns a tuple with the Ready field value
// and a boolean to check if the value has been set.
func (o *DeprecatedGetPresentationsResponseItemRenditionsInner) GetReadyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ready, true
}

// SetReady sets field value
func (o *DeprecatedGetPresentationsResponseItemRenditionsInner) SetReady(v bool) {
	o.Ready = v
}

// GetRendid returns the Rendid field value
func (o *DeprecatedGetPresentationsResponseItemRenditionsInner) GetRendid() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Rendid
}

// GetRendidOk returns a tuple with the Rendid field value
// and a boolean to check if the value has been set.
func (o *DeprecatedGetPresentationsResponseItemRenditionsInner) GetRendidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rendid, true
}

// SetRendid sets field value
func (o *DeprecatedGetPresentationsResponseItemRenditionsInner) SetRendid(v int64) {
	o.Rendid = v
}

// GetStorepath returns the Storepath field value
func (o *DeprecatedGetPresentationsResponseItemRenditionsInner) GetStorepath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Storepath
}

// GetStorepathOk returns a tuple with the Storepath field value
// and a boolean to check if the value has been set.
func (o *DeprecatedGetPresentationsResponseItemRenditionsInner) GetStorepathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storepath, true
}

// SetStorepath sets field value
func (o *DeprecatedGetPresentationsResponseItemRenditionsInner) SetStorepath(v string) {
	o.Storepath = v
}

func (o DeprecatedGetPresentationsResponseItemRenditionsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeprecatedGetPresentationsResponseItemRenditionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["iframeonly"] = o.Iframeonly
	toSerialize["metadata"] = o.Metadata
	toSerialize["path"] = o.Path
	toSerialize["ready"] = o.Ready
	toSerialize["rendid"] = o.Rendid
	toSerialize["storepath"] = o.Storepath
	return toSerialize, nil
}

type NullableDeprecatedGetPresentationsResponseItemRenditionsInner struct {
	value *DeprecatedGetPresentationsResponseItemRenditionsInner
	isSet bool
}

func (v NullableDeprecatedGetPresentationsResponseItemRenditionsInner) Get() *DeprecatedGetPresentationsResponseItemRenditionsInner {
	return v.value
}

func (v *NullableDeprecatedGetPresentationsResponseItemRenditionsInner) Set(val *DeprecatedGetPresentationsResponseItemRenditionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableDeprecatedGetPresentationsResponseItemRenditionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableDeprecatedGetPresentationsResponseItemRenditionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeprecatedGetPresentationsResponseItemRenditionsInner(val *DeprecatedGetPresentationsResponseItemRenditionsInner) *NullableDeprecatedGetPresentationsResponseItemRenditionsInner {
	return &NullableDeprecatedGetPresentationsResponseItemRenditionsInner{value: val, isSet: true}
}

func (v NullableDeprecatedGetPresentationsResponseItemRenditionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeprecatedGetPresentationsResponseItemRenditionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


