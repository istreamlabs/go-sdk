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

// PutClipRequest struct for PutClipRequest
type PutClipRequest struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string `json:"$schema,omitempty"`
	// Description of the clip being created
	Description *string `json:"description,omitempty"`
	// End timestamp in RFC3339Nano format
	End string `json:"end"`
	// Overwrite existing clip. Default: false
	Overwrite *bool `json:"overwrite,omitempty"`
	// Start timestamp in RFC3339Nano format
	Start string `json:"start"`
}

// NewPutClipRequest instantiates a new PutClipRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutClipRequest(end string, start string) *PutClipRequest {
	this := PutClipRequest{}
	this.End = end
	this.Start = start
	return &this
}

// NewPutClipRequestWithDefaults instantiates a new PutClipRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutClipRequestWithDefaults() *PutClipRequest {
	this := PutClipRequest{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *PutClipRequest) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutClipRequest) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *PutClipRequest) HasSchema() bool {
	if o != nil && o.Schema != nil {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *PutClipRequest) SetSchema(v string) {
	o.Schema = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PutClipRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutClipRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PutClipRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PutClipRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnd returns the End field value
func (o *PutClipRequest) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *PutClipRequest) GetEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *PutClipRequest) SetEnd(v string) {
	o.End = v
}

// GetOverwrite returns the Overwrite field value if set, zero value otherwise.
func (o *PutClipRequest) GetOverwrite() bool {
	if o == nil || o.Overwrite == nil {
		var ret bool
		return ret
	}
	return *o.Overwrite
}

// GetOverwriteOk returns a tuple with the Overwrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutClipRequest) GetOverwriteOk() (*bool, bool) {
	if o == nil || o.Overwrite == nil {
		return nil, false
	}
	return o.Overwrite, true
}

// HasOverwrite returns a boolean if a field has been set.
func (o *PutClipRequest) HasOverwrite() bool {
	if o != nil && o.Overwrite != nil {
		return true
	}

	return false
}

// SetOverwrite gets a reference to the given bool and assigns it to the Overwrite field.
func (o *PutClipRequest) SetOverwrite(v bool) {
	o.Overwrite = &v
}

// GetStart returns the Start field value
func (o *PutClipRequest) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *PutClipRequest) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *PutClipRequest) SetStart(v string) {
	o.Start = v
}

func (o PutClipRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Schema != nil {
		toSerialize["$schema"] = o.Schema
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["end"] = o.End
	}
	if o.Overwrite != nil {
		toSerialize["overwrite"] = o.Overwrite
	}
	if true {
		toSerialize["start"] = o.Start
	}
	return json.Marshal(toSerialize)
}

type NullablePutClipRequest struct {
	value *PutClipRequest
	isSet bool
}

func (v NullablePutClipRequest) Get() *PutClipRequest {
	return v.value
}

func (v *NullablePutClipRequest) Set(val *PutClipRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutClipRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutClipRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutClipRequest(val *PutClipRequest) *NullablePutClipRequest {
	return &NullablePutClipRequest{value: val, isSet: true}
}

func (v NullablePutClipRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutClipRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


