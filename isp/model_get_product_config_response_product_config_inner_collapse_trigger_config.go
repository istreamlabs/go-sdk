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

// checks if the GetProductConfigResponseProductConfigInnerCollapseTriggerConfig type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProductConfigResponseProductConfigInnerCollapseTriggerConfig{}

// GetProductConfigResponseProductConfigInnerCollapseTriggerConfig Options for when to trigger collapses
type GetProductConfigResponseProductConfigInnerCollapseTriggerConfig struct {
	AllowMultiple   bool                                                                           `json:"allow_multiple"`
	CollapseProgram bool                                                                           `json:"collapse_program"`
	CollapseStream  bool                                                                           `json:"collapse_stream"`
	Scte35Upid      *string                                                                        `json:"scte35_upid,omitempty"`
	Triggers        []GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner `json:"triggers,omitempty"`
}

// NewGetProductConfigResponseProductConfigInnerCollapseTriggerConfig instantiates a new GetProductConfigResponseProductConfigInnerCollapseTriggerConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductConfigResponseProductConfigInnerCollapseTriggerConfig(allowMultiple bool, collapseProgram bool, collapseStream bool) *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig {
	this := GetProductConfigResponseProductConfigInnerCollapseTriggerConfig{}
	this.AllowMultiple = allowMultiple
	this.CollapseProgram = collapseProgram
	this.CollapseStream = collapseStream
	return &this
}

// NewGetProductConfigResponseProductConfigInnerCollapseTriggerConfigWithDefaults instantiates a new GetProductConfigResponseProductConfigInnerCollapseTriggerConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductConfigResponseProductConfigInnerCollapseTriggerConfigWithDefaults() *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig {
	this := GetProductConfigResponseProductConfigInnerCollapseTriggerConfig{}
	return &this
}

// GetAllowMultiple returns the AllowMultiple field value
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig) GetAllowMultiple() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AllowMultiple
}

// GetAllowMultipleOk returns a tuple with the AllowMultiple field value
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig) GetAllowMultipleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AllowMultiple, true
}

// SetAllowMultiple sets field value
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig) SetAllowMultiple(v bool) {
	o.AllowMultiple = v
}

// GetCollapseProgram returns the CollapseProgram field value
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig) GetCollapseProgram() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CollapseProgram
}

// GetCollapseProgramOk returns a tuple with the CollapseProgram field value
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig) GetCollapseProgramOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollapseProgram, true
}

// SetCollapseProgram sets field value
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig) SetCollapseProgram(v bool) {
	o.CollapseProgram = v
}

// GetCollapseStream returns the CollapseStream field value
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig) GetCollapseStream() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.CollapseStream
}

// GetCollapseStreamOk returns a tuple with the CollapseStream field value
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig) GetCollapseStreamOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CollapseStream, true
}

// SetCollapseStream sets field value
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig) SetCollapseStream(v bool) {
	o.CollapseStream = v
}

// GetScte35Upid returns the Scte35Upid field value if set, zero value otherwise.
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig) GetScte35Upid() string {
	if o == nil || IsNil(o.Scte35Upid) {
		var ret string
		return ret
	}
	return *o.Scte35Upid
}

// GetScte35UpidOk returns a tuple with the Scte35Upid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig) GetScte35UpidOk() (*string, bool) {
	if o == nil || IsNil(o.Scte35Upid) {
		return nil, false
	}
	return o.Scte35Upid, true
}

// HasScte35Upid returns a boolean if a field has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig) HasScte35Upid() bool {
	if o != nil && !IsNil(o.Scte35Upid) {
		return true
	}

	return false
}

// SetScte35Upid gets a reference to the given string and assigns it to the Scte35Upid field.
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig) SetScte35Upid(v string) {
	o.Scte35Upid = &v
}

// GetTriggers returns the Triggers field value if set, zero value otherwise.
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig) GetTriggers() []GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner {
	if o == nil || IsNil(o.Triggers) {
		var ret []GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner
		return ret
	}
	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig) GetTriggersOk() ([]GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner, bool) {
	if o == nil || IsNil(o.Triggers) {
		return nil, false
	}
	return o.Triggers, true
}

// HasTriggers returns a boolean if a field has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig) HasTriggers() bool {
	if o != nil && !IsNil(o.Triggers) {
		return true
	}

	return false
}

// SetTriggers gets a reference to the given []GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner and assigns it to the Triggers field.
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig) SetTriggers(v []GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner) {
	o.Triggers = v
}

func (o GetProductConfigResponseProductConfigInnerCollapseTriggerConfig) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProductConfigResponseProductConfigInnerCollapseTriggerConfig) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allow_multiple"] = o.AllowMultiple
	toSerialize["collapse_program"] = o.CollapseProgram
	toSerialize["collapse_stream"] = o.CollapseStream
	if !IsNil(o.Scte35Upid) {
		toSerialize["scte35_upid"] = o.Scte35Upid
	}
	if !IsNil(o.Triggers) {
		toSerialize["triggers"] = o.Triggers
	}
	return toSerialize, nil
}

type NullableGetProductConfigResponseProductConfigInnerCollapseTriggerConfig struct {
	value *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig
	isSet bool
}

func (v NullableGetProductConfigResponseProductConfigInnerCollapseTriggerConfig) Get() *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig {
	return v.value
}

func (v *NullableGetProductConfigResponseProductConfigInnerCollapseTriggerConfig) Set(val *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductConfigResponseProductConfigInnerCollapseTriggerConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductConfigResponseProductConfigInnerCollapseTriggerConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductConfigResponseProductConfigInnerCollapseTriggerConfig(val *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig) *NullableGetProductConfigResponseProductConfigInnerCollapseTriggerConfig {
	return &NullableGetProductConfigResponseProductConfigInnerCollapseTriggerConfig{value: val, isSet: true}
}

func (v NullableGetProductConfigResponseProductConfigInnerCollapseTriggerConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProductConfigResponseProductConfigInnerCollapseTriggerConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
