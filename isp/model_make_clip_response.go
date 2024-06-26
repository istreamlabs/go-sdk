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

// checks if the MakeClipResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MakeClipResponse{}

// MakeClipResponse struct for MakeClipResponse
type MakeClipResponse struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string `json:"$schema,omitempty" format:"uri" doc:"An optional URL to a JSON Schema document describing this resource"`
	// Identifier for the clip created from the VOD
	ClipId string `json:"clip_id" doc:"Identifier for the clip created from the VOD"`
	// Identifier for the task from which the clip was created
	TaskId string `json:"task_id" doc:"Identifier for the task from which the clip was created"`
}

// NewMakeClipResponse instantiates a new MakeClipResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMakeClipResponse(clipId string, taskId string) *MakeClipResponse {
	this := MakeClipResponse{}
	this.ClipId = clipId
	this.TaskId = taskId
	return &this
}

// NewMakeClipResponseWithDefaults instantiates a new MakeClipResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMakeClipResponseWithDefaults() *MakeClipResponse {
	this := MakeClipResponse{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *MakeClipResponse) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MakeClipResponse) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *MakeClipResponse) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *MakeClipResponse) SetSchema(v string) {
	o.Schema = &v
}

// GetClipId returns the ClipId field value
func (o *MakeClipResponse) GetClipId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClipId
}

// GetClipIdOk returns a tuple with the ClipId field value
// and a boolean to check if the value has been set.
func (o *MakeClipResponse) GetClipIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClipId, true
}

// SetClipId sets field value
func (o *MakeClipResponse) SetClipId(v string) {
	o.ClipId = v
}

// GetTaskId returns the TaskId field value
func (o *MakeClipResponse) GetTaskId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value
// and a boolean to check if the value has been set.
func (o *MakeClipResponse) GetTaskIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TaskId, true
}

// SetTaskId sets field value
func (o *MakeClipResponse) SetTaskId(v string) {
	o.TaskId = v
}

func (o MakeClipResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MakeClipResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["clip_id"] = o.ClipId
	toSerialize["task_id"] = o.TaskId
	return toSerialize, nil
}

type NullableMakeClipResponse struct {
	value *MakeClipResponse
	isSet bool
}

func (v NullableMakeClipResponse) Get() *MakeClipResponse {
	return v.value
}

func (v *NullableMakeClipResponse) Set(val *MakeClipResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMakeClipResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMakeClipResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMakeClipResponse(val *MakeClipResponse) *NullableMakeClipResponse {
	return &NullableMakeClipResponse{value: val, isSet: true}
}

func (v NullableMakeClipResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMakeClipResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


