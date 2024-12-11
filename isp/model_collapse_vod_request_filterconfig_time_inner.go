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
	"time"
)

// checks if the CollapseVODRequestFilterconfigTimeInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollapseVODRequestFilterconfigTimeInner{}

// CollapseVODRequestFilterconfigTimeInner struct for CollapseVODRequestFilterconfigTimeInner
type CollapseVODRequestFilterconfigTimeInner struct {
	// end time of the time window of segments
	End time.Time `json:"end" format:"date-time" doc:"end time of the time window of segments"`
	// should include boundary within the VOD
	Include bool `json:"include" doc:"should include boundary within the VOD"`
	// start time of the time window of segments
	Start time.Time `json:"start" format:"date-time" doc:"start time of the time window of segments"`
}

// NewCollapseVODRequestFilterconfigTimeInner instantiates a new CollapseVODRequestFilterconfigTimeInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollapseVODRequestFilterconfigTimeInner(end time.Time, include bool, start time.Time) *CollapseVODRequestFilterconfigTimeInner {
	this := CollapseVODRequestFilterconfigTimeInner{}
	this.End = end
	this.Include = include
	this.Start = start
	return &this
}

// NewCollapseVODRequestFilterconfigTimeInnerWithDefaults instantiates a new CollapseVODRequestFilterconfigTimeInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollapseVODRequestFilterconfigTimeInnerWithDefaults() *CollapseVODRequestFilterconfigTimeInner {
	this := CollapseVODRequestFilterconfigTimeInner{}
	return &this
}

// GetEnd returns the End field value
func (o *CollapseVODRequestFilterconfigTimeInner) GetEnd() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *CollapseVODRequestFilterconfigTimeInner) GetEndOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *CollapseVODRequestFilterconfigTimeInner) SetEnd(v time.Time) {
	o.End = v
}

// GetInclude returns the Include field value
func (o *CollapseVODRequestFilterconfigTimeInner) GetInclude() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Include
}

// GetIncludeOk returns a tuple with the Include field value
// and a boolean to check if the value has been set.
func (o *CollapseVODRequestFilterconfigTimeInner) GetIncludeOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Include, true
}

// SetInclude sets field value
func (o *CollapseVODRequestFilterconfigTimeInner) SetInclude(v bool) {
	o.Include = v
}

// GetStart returns the Start field value
func (o *CollapseVODRequestFilterconfigTimeInner) GetStart() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *CollapseVODRequestFilterconfigTimeInner) GetStartOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *CollapseVODRequestFilterconfigTimeInner) SetStart(v time.Time) {
	o.Start = v
}

func (o CollapseVODRequestFilterconfigTimeInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollapseVODRequestFilterconfigTimeInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["end"] = o.End
	toSerialize["include"] = o.Include
	toSerialize["start"] = o.Start
	return toSerialize, nil
}

type NullableCollapseVODRequestFilterconfigTimeInner struct {
	value *CollapseVODRequestFilterconfigTimeInner
	isSet bool
}

func (v NullableCollapseVODRequestFilterconfigTimeInner) Get() *CollapseVODRequestFilterconfigTimeInner {
	return v.value
}

func (v *NullableCollapseVODRequestFilterconfigTimeInner) Set(val *CollapseVODRequestFilterconfigTimeInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCollapseVODRequestFilterconfigTimeInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCollapseVODRequestFilterconfigTimeInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollapseVODRequestFilterconfigTimeInner(val *CollapseVODRequestFilterconfigTimeInner) *NullableCollapseVODRequestFilterconfigTimeInner {
	return &NullableCollapseVODRequestFilterconfigTimeInner{value: val, isSet: true}
}

func (v NullableCollapseVODRequestFilterconfigTimeInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollapseVODRequestFilterconfigTimeInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
