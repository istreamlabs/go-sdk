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

// checks if the PostClipArchiveRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostClipArchiveRequest{}

// PostClipArchiveRequest struct for PostClipArchiveRequest
type PostClipArchiveRequest struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string `json:"$schema,omitempty"`
	// Identifer that is carried through the archive request
	CorrelationId *string `json:"correlation_id,omitempty"`
}

// NewPostClipArchiveRequest instantiates a new PostClipArchiveRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostClipArchiveRequest() *PostClipArchiveRequest {
	this := PostClipArchiveRequest{}
	return &this
}

// NewPostClipArchiveRequestWithDefaults instantiates a new PostClipArchiveRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostClipArchiveRequestWithDefaults() *PostClipArchiveRequest {
	this := PostClipArchiveRequest{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *PostClipArchiveRequest) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostClipArchiveRequest) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *PostClipArchiveRequest) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *PostClipArchiveRequest) SetSchema(v string) {
	o.Schema = &v
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise.
func (o *PostClipArchiveRequest) GetCorrelationId() string {
	if o == nil || IsNil(o.CorrelationId) {
		var ret string
		return ret
	}
	return *o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostClipArchiveRequest) GetCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelationId) {
		return nil, false
	}
	return o.CorrelationId, true
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *PostClipArchiveRequest) HasCorrelationId() bool {
	if o != nil && !IsNil(o.CorrelationId) {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.
func (o *PostClipArchiveRequest) SetCorrelationId(v string) {
	o.CorrelationId = &v
}

func (o PostClipArchiveRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostClipArchiveRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	if !IsNil(o.CorrelationId) {
		toSerialize["correlation_id"] = o.CorrelationId
	}
	return toSerialize, nil
}

type NullablePostClipArchiveRequest struct {
	value *PostClipArchiveRequest
	isSet bool
}

func (v NullablePostClipArchiveRequest) Get() *PostClipArchiveRequest {
	return v.value
}

func (v *NullablePostClipArchiveRequest) Set(val *PostClipArchiveRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostClipArchiveRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostClipArchiveRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostClipArchiveRequest(val *PostClipArchiveRequest) *NullablePostClipArchiveRequest {
	return &NullablePostClipArchiveRequest{value: val, isSet: true}
}

func (v NullablePostClipArchiveRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostClipArchiveRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

