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

// PostClipResponse struct for PostClipResponse
type PostClipResponse struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string `json:"$schema,omitempty"`
	// List of clip identifiers paired with each VodID that was clipped
	Clips []PostClipResponseClipsInner `json:"clips"`
	// Reports failures of individual clip tasks. Empty if all are successful
	TaskErrors []ErrorModelErrorsInner `json:"task_errors"`
	// Identifies clipping tasks, each task producing N VodItems
	TaskIds []string `json:"task_ids"`
}

// NewPostClipResponse instantiates a new PostClipResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostClipResponse(clips []PostClipResponseClipsInner, taskErrors []ErrorModelErrorsInner, taskIds []string) *PostClipResponse {
	this := PostClipResponse{}
	this.Clips = clips
	this.TaskErrors = taskErrors
	this.TaskIds = taskIds
	return &this
}

// NewPostClipResponseWithDefaults instantiates a new PostClipResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostClipResponseWithDefaults() *PostClipResponse {
	this := PostClipResponse{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *PostClipResponse) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostClipResponse) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *PostClipResponse) HasSchema() bool {
	if o != nil && o.Schema != nil {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *PostClipResponse) SetSchema(v string) {
	o.Schema = &v
}

// GetClips returns the Clips field value
func (o *PostClipResponse) GetClips() []PostClipResponseClipsInner {
	if o == nil {
		var ret []PostClipResponseClipsInner
		return ret
	}

	return o.Clips
}

// GetClipsOk returns a tuple with the Clips field value
// and a boolean to check if the value has been set.
func (o *PostClipResponse) GetClipsOk() ([]PostClipResponseClipsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Clips, true
}

// SetClips sets field value
func (o *PostClipResponse) SetClips(v []PostClipResponseClipsInner) {
	o.Clips = v
}

// GetTaskErrors returns the TaskErrors field value
func (o *PostClipResponse) GetTaskErrors() []ErrorModelErrorsInner {
	if o == nil {
		var ret []ErrorModelErrorsInner
		return ret
	}

	return o.TaskErrors
}

// GetTaskErrorsOk returns a tuple with the TaskErrors field value
// and a boolean to check if the value has been set.
func (o *PostClipResponse) GetTaskErrorsOk() ([]ErrorModelErrorsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskErrors, true
}

// SetTaskErrors sets field value
func (o *PostClipResponse) SetTaskErrors(v []ErrorModelErrorsInner) {
	o.TaskErrors = v
}

// GetTaskIds returns the TaskIds field value
func (o *PostClipResponse) GetTaskIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TaskIds
}

// GetTaskIdsOk returns a tuple with the TaskIds field value
// and a boolean to check if the value has been set.
func (o *PostClipResponse) GetTaskIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskIds, true
}

// SetTaskIds sets field value
func (o *PostClipResponse) SetTaskIds(v []string) {
	o.TaskIds = v
}

func (o PostClipResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Schema != nil {
		toSerialize["$schema"] = o.Schema
	}
	if true {
		toSerialize["clips"] = o.Clips
	}
	if true {
		toSerialize["task_errors"] = o.TaskErrors
	}
	if true {
		toSerialize["task_ids"] = o.TaskIds
	}
	return json.Marshal(toSerialize)
}

type NullablePostClipResponse struct {
	value *PostClipResponse
	isSet bool
}

func (v NullablePostClipResponse) Get() *PostClipResponse {
	return v.value
}

func (v *NullablePostClipResponse) Set(val *PostClipResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePostClipResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePostClipResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostClipResponse(val *PostClipResponse) *NullablePostClipResponse {
	return &NullablePostClipResponse{value: val, isSet: true}
}

func (v NullablePostClipResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostClipResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


