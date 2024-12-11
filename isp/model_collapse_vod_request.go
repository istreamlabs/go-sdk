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

// checks if the CollapseVODRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollapseVODRequest{}

// CollapseVODRequest struct for CollapseVODRequest
type CollapseVODRequest struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string `json:"$schema,omitempty" format:"uri" doc:"An optional URL to a JSON Schema document describing this resource"`
	// Prefix for clip
	ClipPrefix *string `json:"clip_prefix,omitempty" doc:"Prefix for clip"`
	// Suffix for clip
	ClipSuffix *string `json:"clip_suffix,omitempty" doc:"Suffix for clip"`
	// Disable any autoprefix
	DisableAutoPrefix *bool `json:"disable_auto_prefix,omitempty" doc:"Disable any autoprefix"`
	// EndTime of VOD
	EndTime      *time.Time                      `json:"end_time,omitempty" format:"date-time" doc:"EndTime of VOD"`
	Filterconfig *CollapseVODRequestFilterconfig `json:"filterconfig,omitempty"`
	// Description for new collapsed clip
	NewClipDescription string `json:"new_clip_description" doc:"Description for new collapsed clip"`
	// New Clip ID for collapsed clip
	NewClipId int64 `json:"new_clip_id" format:"int64" doc:"New Clip ID for collapsed clip"`
	// Publish newly created VOD
	PublishVod *bool `json:"publish_vod,omitempty" doc:"Publish newly created VOD"`
	// StartTime of VOD
	StartTime *time.Time `json:"start_time,omitempty" format:"date-time" doc:"StartTime of VOD"`
	// UPID for VOD
	Upid *string `json:"upid,omitempty" doc:"UPID for VOD"`
}

// NewCollapseVODRequest instantiates a new CollapseVODRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollapseVODRequest(newClipDescription string, newClipId int64) *CollapseVODRequest {
	this := CollapseVODRequest{}
	this.NewClipDescription = newClipDescription
	this.NewClipId = newClipId
	return &this
}

// NewCollapseVODRequestWithDefaults instantiates a new CollapseVODRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollapseVODRequestWithDefaults() *CollapseVODRequest {
	this := CollapseVODRequest{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *CollapseVODRequest) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollapseVODRequest) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *CollapseVODRequest) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *CollapseVODRequest) SetSchema(v string) {
	o.Schema = &v
}

// GetClipPrefix returns the ClipPrefix field value if set, zero value otherwise.
func (o *CollapseVODRequest) GetClipPrefix() string {
	if o == nil || IsNil(o.ClipPrefix) {
		var ret string
		return ret
	}
	return *o.ClipPrefix
}

// GetClipPrefixOk returns a tuple with the ClipPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollapseVODRequest) GetClipPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.ClipPrefix) {
		return nil, false
	}
	return o.ClipPrefix, true
}

// HasClipPrefix returns a boolean if a field has been set.
func (o *CollapseVODRequest) HasClipPrefix() bool {
	if o != nil && !IsNil(o.ClipPrefix) {
		return true
	}

	return false
}

// SetClipPrefix gets a reference to the given string and assigns it to the ClipPrefix field.
func (o *CollapseVODRequest) SetClipPrefix(v string) {
	o.ClipPrefix = &v
}

// GetClipSuffix returns the ClipSuffix field value if set, zero value otherwise.
func (o *CollapseVODRequest) GetClipSuffix() string {
	if o == nil || IsNil(o.ClipSuffix) {
		var ret string
		return ret
	}
	return *o.ClipSuffix
}

// GetClipSuffixOk returns a tuple with the ClipSuffix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollapseVODRequest) GetClipSuffixOk() (*string, bool) {
	if o == nil || IsNil(o.ClipSuffix) {
		return nil, false
	}
	return o.ClipSuffix, true
}

// HasClipSuffix returns a boolean if a field has been set.
func (o *CollapseVODRequest) HasClipSuffix() bool {
	if o != nil && !IsNil(o.ClipSuffix) {
		return true
	}

	return false
}

// SetClipSuffix gets a reference to the given string and assigns it to the ClipSuffix field.
func (o *CollapseVODRequest) SetClipSuffix(v string) {
	o.ClipSuffix = &v
}

// GetDisableAutoPrefix returns the DisableAutoPrefix field value if set, zero value otherwise.
func (o *CollapseVODRequest) GetDisableAutoPrefix() bool {
	if o == nil || IsNil(o.DisableAutoPrefix) {
		var ret bool
		return ret
	}
	return *o.DisableAutoPrefix
}

// GetDisableAutoPrefixOk returns a tuple with the DisableAutoPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollapseVODRequest) GetDisableAutoPrefixOk() (*bool, bool) {
	if o == nil || IsNil(o.DisableAutoPrefix) {
		return nil, false
	}
	return o.DisableAutoPrefix, true
}

// HasDisableAutoPrefix returns a boolean if a field has been set.
func (o *CollapseVODRequest) HasDisableAutoPrefix() bool {
	if o != nil && !IsNil(o.DisableAutoPrefix) {
		return true
	}

	return false
}

// SetDisableAutoPrefix gets a reference to the given bool and assigns it to the DisableAutoPrefix field.
func (o *CollapseVODRequest) SetDisableAutoPrefix(v bool) {
	o.DisableAutoPrefix = &v
}

// GetEndTime returns the EndTime field value if set, zero value otherwise.
func (o *CollapseVODRequest) GetEndTime() time.Time {
	if o == nil || IsNil(o.EndTime) {
		var ret time.Time
		return ret
	}
	return *o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollapseVODRequest) GetEndTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndTime) {
		return nil, false
	}
	return o.EndTime, true
}

// HasEndTime returns a boolean if a field has been set.
func (o *CollapseVODRequest) HasEndTime() bool {
	if o != nil && !IsNil(o.EndTime) {
		return true
	}

	return false
}

// SetEndTime gets a reference to the given time.Time and assigns it to the EndTime field.
func (o *CollapseVODRequest) SetEndTime(v time.Time) {
	o.EndTime = &v
}

// GetFilterconfig returns the Filterconfig field value if set, zero value otherwise.
func (o *CollapseVODRequest) GetFilterconfig() CollapseVODRequestFilterconfig {
	if o == nil || IsNil(o.Filterconfig) {
		var ret CollapseVODRequestFilterconfig
		return ret
	}
	return *o.Filterconfig
}

// GetFilterconfigOk returns a tuple with the Filterconfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollapseVODRequest) GetFilterconfigOk() (*CollapseVODRequestFilterconfig, bool) {
	if o == nil || IsNil(o.Filterconfig) {
		return nil, false
	}
	return o.Filterconfig, true
}

// HasFilterconfig returns a boolean if a field has been set.
func (o *CollapseVODRequest) HasFilterconfig() bool {
	if o != nil && !IsNil(o.Filterconfig) {
		return true
	}

	return false
}

// SetFilterconfig gets a reference to the given CollapseVODRequestFilterconfig and assigns it to the Filterconfig field.
func (o *CollapseVODRequest) SetFilterconfig(v CollapseVODRequestFilterconfig) {
	o.Filterconfig = &v
}

// GetNewClipDescription returns the NewClipDescription field value
func (o *CollapseVODRequest) GetNewClipDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NewClipDescription
}

// GetNewClipDescriptionOk returns a tuple with the NewClipDescription field value
// and a boolean to check if the value has been set.
func (o *CollapseVODRequest) GetNewClipDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewClipDescription, true
}

// SetNewClipDescription sets field value
func (o *CollapseVODRequest) SetNewClipDescription(v string) {
	o.NewClipDescription = v
}

// GetNewClipId returns the NewClipId field value
func (o *CollapseVODRequest) GetNewClipId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NewClipId
}

// GetNewClipIdOk returns a tuple with the NewClipId field value
// and a boolean to check if the value has been set.
func (o *CollapseVODRequest) GetNewClipIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewClipId, true
}

// SetNewClipId sets field value
func (o *CollapseVODRequest) SetNewClipId(v int64) {
	o.NewClipId = v
}

// GetPublishVod returns the PublishVod field value if set, zero value otherwise.
func (o *CollapseVODRequest) GetPublishVod() bool {
	if o == nil || IsNil(o.PublishVod) {
		var ret bool
		return ret
	}
	return *o.PublishVod
}

// GetPublishVodOk returns a tuple with the PublishVod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollapseVODRequest) GetPublishVodOk() (*bool, bool) {
	if o == nil || IsNil(o.PublishVod) {
		return nil, false
	}
	return o.PublishVod, true
}

// HasPublishVod returns a boolean if a field has been set.
func (o *CollapseVODRequest) HasPublishVod() bool {
	if o != nil && !IsNil(o.PublishVod) {
		return true
	}

	return false
}

// SetPublishVod gets a reference to the given bool and assigns it to the PublishVod field.
func (o *CollapseVODRequest) SetPublishVod(v bool) {
	o.PublishVod = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *CollapseVODRequest) GetStartTime() time.Time {
	if o == nil || IsNil(o.StartTime) {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollapseVODRequest) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartTime) {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *CollapseVODRequest) HasStartTime() bool {
	if o != nil && !IsNil(o.StartTime) {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *CollapseVODRequest) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetUpid returns the Upid field value if set, zero value otherwise.
func (o *CollapseVODRequest) GetUpid() string {
	if o == nil || IsNil(o.Upid) {
		var ret string
		return ret
	}
	return *o.Upid
}

// GetUpidOk returns a tuple with the Upid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollapseVODRequest) GetUpidOk() (*string, bool) {
	if o == nil || IsNil(o.Upid) {
		return nil, false
	}
	return o.Upid, true
}

// HasUpid returns a boolean if a field has been set.
func (o *CollapseVODRequest) HasUpid() bool {
	if o != nil && !IsNil(o.Upid) {
		return true
	}

	return false
}

// SetUpid gets a reference to the given string and assigns it to the Upid field.
func (o *CollapseVODRequest) SetUpid(v string) {
	o.Upid = &v
}

func (o CollapseVODRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollapseVODRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	if !IsNil(o.ClipPrefix) {
		toSerialize["clip_prefix"] = o.ClipPrefix
	}
	if !IsNil(o.ClipSuffix) {
		toSerialize["clip_suffix"] = o.ClipSuffix
	}
	if !IsNil(o.DisableAutoPrefix) {
		toSerialize["disable_auto_prefix"] = o.DisableAutoPrefix
	}
	if !IsNil(o.EndTime) {
		toSerialize["end_time"] = o.EndTime
	}
	if !IsNil(o.Filterconfig) {
		toSerialize["filterconfig"] = o.Filterconfig
	}
	toSerialize["new_clip_description"] = o.NewClipDescription
	toSerialize["new_clip_id"] = o.NewClipId
	if !IsNil(o.PublishVod) {
		toSerialize["publish_vod"] = o.PublishVod
	}
	if !IsNil(o.StartTime) {
		toSerialize["start_time"] = o.StartTime
	}
	if !IsNil(o.Upid) {
		toSerialize["upid"] = o.Upid
	}
	return toSerialize, nil
}

type NullableCollapseVODRequest struct {
	value *CollapseVODRequest
	isSet bool
}

func (v NullableCollapseVODRequest) Get() *CollapseVODRequest {
	return v.value
}

func (v *NullableCollapseVODRequest) Set(val *CollapseVODRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCollapseVODRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCollapseVODRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollapseVODRequest(val *CollapseVODRequest) *NullableCollapseVODRequest {
	return &NullableCollapseVODRequest{value: val, isSet: true}
}

func (v NullableCollapseVODRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollapseVODRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
