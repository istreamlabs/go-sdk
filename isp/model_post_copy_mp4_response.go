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

// checks if the PostCopyMP4Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostCopyMP4Response{}

// PostCopyMP4Response struct for PostCopyMP4Response
type PostCopyMP4Response struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string `json:"$schema,omitempty" format:"uri" doc:"An optional URL to a JSON Schema document describing this resource"`
	// Array of task ids for each copymp4 destination task successfully started by Live2VOD
	TaskIds []string `json:"task_ids" doc:"Array of task ids for each copymp4 destination task successfully started by Live2VOD"`
}

// NewPostCopyMP4Response instantiates a new PostCopyMP4Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostCopyMP4Response(taskIds []string) *PostCopyMP4Response {
	this := PostCopyMP4Response{}
	this.TaskIds = taskIds
	return &this
}

// NewPostCopyMP4ResponseWithDefaults instantiates a new PostCopyMP4Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostCopyMP4ResponseWithDefaults() *PostCopyMP4Response {
	this := PostCopyMP4Response{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *PostCopyMP4Response) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostCopyMP4Response) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *PostCopyMP4Response) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *PostCopyMP4Response) SetSchema(v string) {
	o.Schema = &v
}

// GetTaskIds returns the TaskIds field value
func (o *PostCopyMP4Response) GetTaskIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.TaskIds
}

// GetTaskIdsOk returns a tuple with the TaskIds field value
// and a boolean to check if the value has been set.
func (o *PostCopyMP4Response) GetTaskIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaskIds, true
}

// SetTaskIds sets field value
func (o *PostCopyMP4Response) SetTaskIds(v []string) {
	o.TaskIds = v
}

func (o PostCopyMP4Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostCopyMP4Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["task_ids"] = o.TaskIds
	return toSerialize, nil
}

type NullablePostCopyMP4Response struct {
	value *PostCopyMP4Response
	isSet bool
}

func (v NullablePostCopyMP4Response) Get() *PostCopyMP4Response {
	return v.value
}

func (v *NullablePostCopyMP4Response) Set(val *PostCopyMP4Response) {
	v.value = val
	v.isSet = true
}

func (v NullablePostCopyMP4Response) IsSet() bool {
	return v.isSet
}

func (v *NullablePostCopyMP4Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostCopyMP4Response(val *PostCopyMP4Response) *NullablePostCopyMP4Response {
	return &NullablePostCopyMP4Response{value: val, isSet: true}
}

func (v NullablePostCopyMP4Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostCopyMP4Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


