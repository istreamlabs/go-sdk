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

// checks if the DeprecatedGetClipResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeprecatedGetClipResponse{}

// DeprecatedGetClipResponse struct for DeprecatedGetClipResponse
type DeprecatedGetClipResponse struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string `json:"$schema,omitempty" format:"uri" doc:"An optional URL to a JSON Schema document describing this resource"`
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

// NewDeprecatedGetClipResponse instantiates a new DeprecatedGetClipResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeprecatedGetClipResponse(clipid int64, created time.Time, description string, duration int64, endTime time.Time, mp4path string, source string, startTime time.Time) *DeprecatedGetClipResponse {
	this := DeprecatedGetClipResponse{}
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

// NewDeprecatedGetClipResponseWithDefaults instantiates a new DeprecatedGetClipResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeprecatedGetClipResponseWithDefaults() *DeprecatedGetClipResponse {
	this := DeprecatedGetClipResponse{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *DeprecatedGetClipResponse) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedGetClipResponse) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *DeprecatedGetClipResponse) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *DeprecatedGetClipResponse) SetSchema(v string) {
	o.Schema = &v
}

// GetClipid returns the Clipid field value
func (o *DeprecatedGetClipResponse) GetClipid() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Clipid
}

// GetClipidOk returns a tuple with the Clipid field value
// and a boolean to check if the value has been set.
func (o *DeprecatedGetClipResponse) GetClipidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Clipid, true
}

// SetClipid sets field value
func (o *DeprecatedGetClipResponse) SetClipid(v int64) {
	o.Clipid = v
}

// GetCreated returns the Created field value
func (o *DeprecatedGetClipResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *DeprecatedGetClipResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *DeprecatedGetClipResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetDescription returns the Description field value
func (o *DeprecatedGetClipResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DeprecatedGetClipResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DeprecatedGetClipResponse) SetDescription(v string) {
	o.Description = v
}

// GetDuration returns the Duration field value
func (o *DeprecatedGetClipResponse) GetDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *DeprecatedGetClipResponse) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *DeprecatedGetClipResponse) SetDuration(v int64) {
	o.Duration = v
}

// GetEndTime returns the EndTime field value
func (o *DeprecatedGetClipResponse) GetEndTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EndTime
}

// GetEndTimeOk returns a tuple with the EndTime field value
// and a boolean to check if the value has been set.
func (o *DeprecatedGetClipResponse) GetEndTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EndTime, true
}

// SetEndTime sets field value
func (o *DeprecatedGetClipResponse) SetEndTime(v time.Time) {
	o.EndTime = v
}

// GetMp4path returns the Mp4path field value
func (o *DeprecatedGetClipResponse) GetMp4path() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Mp4path
}

// GetMp4pathOk returns a tuple with the Mp4path field value
// and a boolean to check if the value has been set.
func (o *DeprecatedGetClipResponse) GetMp4pathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Mp4path, true
}

// SetMp4path sets field value
func (o *DeprecatedGetClipResponse) SetMp4path(v string) {
	o.Mp4path = v
}

// GetSource returns the Source field value
func (o *DeprecatedGetClipResponse) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *DeprecatedGetClipResponse) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *DeprecatedGetClipResponse) SetSource(v string) {
	o.Source = v
}

// GetStartTime returns the StartTime field value
func (o *DeprecatedGetClipResponse) GetStartTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value
// and a boolean to check if the value has been set.
func (o *DeprecatedGetClipResponse) GetStartTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartTime, true
}

// SetStartTime sets field value
func (o *DeprecatedGetClipResponse) SetStartTime(v time.Time) {
	o.StartTime = v
}

func (o DeprecatedGetClipResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeprecatedGetClipResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
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

type NullableDeprecatedGetClipResponse struct {
	value *DeprecatedGetClipResponse
	isSet bool
}

func (v NullableDeprecatedGetClipResponse) Get() *DeprecatedGetClipResponse {
	return v.value
}

func (v *NullableDeprecatedGetClipResponse) Set(val *DeprecatedGetClipResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeprecatedGetClipResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeprecatedGetClipResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeprecatedGetClipResponse(val *DeprecatedGetClipResponse) *NullableDeprecatedGetClipResponse {
	return &NullableDeprecatedGetClipResponse{value: val, isSet: true}
}

func (v NullableDeprecatedGetClipResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeprecatedGetClipResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
