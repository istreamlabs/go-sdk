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

// checks if the DeprecatedListClipsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeprecatedListClipsResponse{}

// DeprecatedListClipsResponse struct for DeprecatedListClipsResponse
type DeprecatedListClipsResponse struct {
	// id for this clip
	Clipid int64 `json:"clipid" format:"int64" doc:"id for this clip"`
	// date and time when clip was created
	Created time.Time `json:"created" format:"date-time" doc:"date and time when clip was created"`
	// description of the clip
	Description string `json:"description" doc:"description of the clip"`
	// clip duration in seconds
	Duration int64 `json:"duration" format:"int64" doc:"clip duration in seconds"`
	// clip end time
	EndTime time.Time `json:"endTime" format:"date-time" doc:"clip end time"`
	// path to the mp4
	Mp4path string `json:"mp4path" doc:"path to the mp4"`
	// source of the clip
	Source string `json:"source" doc:"source of the clip"`
	// clip start time
	StartTime time.Time `json:"startTime" format:"date-time" doc:"clip start time"`
}

// NewDeprecatedListClipsResponse instantiates a new DeprecatedListClipsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeprecatedListClipsResponse(clipid int64, created time.Time, description string, duration int64, endTime time.Time, mp4path string, source string, startTime time.Time) *DeprecatedListClipsResponse {
	this := DeprecatedListClipsResponse{}
	this.Clipid = clipid
	this.Created = created
	this.Description = description
	this.Duration = duration
	this.EndTime = endTime
	this.Mp4path = mp4path
	this.Source = source
	this.StartTime = startTime
	return &this
}

// NewDeprecatedListClipsResponseWithDefaults instantiates a new DeprecatedListClipsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeprecatedListClipsResponseWithDefaults() *DeprecatedListClipsResponse {
	this := DeprecatedListClipsResponse{}
	return &this
}

// GetClipid returns the Clipid field value
func (o *DeprecatedListClipsResponse) GetClipid() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Clipid
}

// GetClipidOk returns a tuple with the Clipid field value
// and a boolean to check if the value has been set.
func (o *DeprecatedListClipsResponse) GetClipidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Clipid, true
}

// SetClipid sets field value
func (o *DeprecatedListClipsResponse) SetClipid(v int64) {
	o.Clipid = v
}

// GetCreated returns the Created field value
func (o *DeprecatedListClipsResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *DeprecatedListClipsResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *DeprecatedListClipsResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetDescription returns the Description field value
func (o *DeprecatedListClipsResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DeprecatedListClipsResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DeprecatedListClipsResponse) SetDescription(v string) {
	o.Description = v
}

// GetDuration returns the Duration field value
func (o *DeprecatedListClipsResponse) GetDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *DeprecatedListClipsResponse) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *DeprecatedListClipsResponse) SetDuration(v int64) {
	o.Duration = v
}

// GetEndTime returns the EndTime field value
func (o *DeprecatedListClipsResponse) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *DeprecatedListClipsResponse) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *DeprecatedListClipsResponse) SetEndTime(v time.Time) {
	o.EndTime = v
}

// GetMp4path returns the Mp4path field value
func (o *DeprecatedListClipsResponse) GetMp4path() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mp4path
}

// GetMp4pathOk returns a tuple with the Mp4path field value
// and a boolean to check if the value has been set.
func (o *DeprecatedListClipsResponse) GetMp4pathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mp4path, true
}

// SetMp4path sets field value
func (o *DeprecatedListClipsResponse) SetMp4path(v string) {
	o.Mp4path = v
}

// GetSource returns the Source field value
func (o *DeprecatedListClipsResponse) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *DeprecatedListClipsResponse) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *DeprecatedListClipsResponse) SetSource(v string) {
	o.Source = v
}

// GetStartTime returns the StartTime field value
func (o *DeprecatedListClipsResponse) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *DeprecatedListClipsResponse) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *DeprecatedListClipsResponse) SetStartTime(v time.Time) {
	o.StartTime = v
}

func (o DeprecatedListClipsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeprecatedListClipsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["clipid"] = o.Clipid
	toSerialize["created"] = o.Created
	toSerialize["description"] = o.Description
	toSerialize["duration"] = o.Duration
	toSerialize["endTime"] = o.EndTime
	toSerialize["mp4path"] = o.Mp4path
	toSerialize["source"] = o.Source
	toSerialize["startTime"] = o.StartTime
	return toSerialize, nil
}

type NullableDeprecatedListClipsResponse struct {
	value *DeprecatedListClipsResponse
	isSet bool
}

func (v NullableDeprecatedListClipsResponse) Get() *DeprecatedListClipsResponse {
	return v.value
}

func (v *NullableDeprecatedListClipsResponse) Set(val *DeprecatedListClipsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeprecatedListClipsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeprecatedListClipsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeprecatedListClipsResponse(val *DeprecatedListClipsResponse) *NullableDeprecatedListClipsResponse {
	return &NullableDeprecatedListClipsResponse{value: val, isSet: true}
}

func (v NullableDeprecatedListClipsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeprecatedListClipsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


