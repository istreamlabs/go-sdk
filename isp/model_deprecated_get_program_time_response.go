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

// checks if the DeprecatedGetProgramTimeResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeprecatedGetProgramTimeResponse{}

// DeprecatedGetProgramTimeResponse struct for DeprecatedGetProgramTimeResponse
type DeprecatedGetProgramTimeResponse struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema       *string `json:"$schema,omitempty" format:"uri" doc:"An optional URL to a JSON Schema document describing this resource"`
	Duration     int64   `json:"duration" format:"int64"`
	ProgramEnd   string  `json:"program_end"`
	ProgramStart string  `json:"program_start"`
}

// NewDeprecatedGetProgramTimeResponse instantiates a new DeprecatedGetProgramTimeResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeprecatedGetProgramTimeResponse(duration int64, programEnd string, programStart string) *DeprecatedGetProgramTimeResponse {
	this := DeprecatedGetProgramTimeResponse{}
	this.Duration = duration
	this.ProgramEnd = programEnd
	this.ProgramStart = programStart
	return &this
}

// NewDeprecatedGetProgramTimeResponseWithDefaults instantiates a new DeprecatedGetProgramTimeResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeprecatedGetProgramTimeResponseWithDefaults() *DeprecatedGetProgramTimeResponse {
	this := DeprecatedGetProgramTimeResponse{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *DeprecatedGetProgramTimeResponse) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedGetProgramTimeResponse) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *DeprecatedGetProgramTimeResponse) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *DeprecatedGetProgramTimeResponse) SetSchema(v string) {
	o.Schema = &v
}

// GetDuration returns the Duration field value
func (o *DeprecatedGetProgramTimeResponse) GetDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *DeprecatedGetProgramTimeResponse) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *DeprecatedGetProgramTimeResponse) SetDuration(v int64) {
	o.Duration = v
}

// GetProgramEnd returns the ProgramEnd field value
func (o *DeprecatedGetProgramTimeResponse) GetProgramEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProgramEnd
}

// GetProgramEndOk returns a tuple with the ProgramEnd field value
// and a boolean to check if the value has been set.
func (o *DeprecatedGetProgramTimeResponse) GetProgramEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProgramEnd, true
}

// SetProgramEnd sets field value
func (o *DeprecatedGetProgramTimeResponse) SetProgramEnd(v string) {
	o.ProgramEnd = v
}

// GetProgramStart returns the ProgramStart field value
func (o *DeprecatedGetProgramTimeResponse) GetProgramStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProgramStart
}

// GetProgramStartOk returns a tuple with the ProgramStart field value
// and a boolean to check if the value has been set.
func (o *DeprecatedGetProgramTimeResponse) GetProgramStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProgramStart, true
}

// SetProgramStart sets field value
func (o *DeprecatedGetProgramTimeResponse) SetProgramStart(v string) {
	o.ProgramStart = v
}

func (o DeprecatedGetProgramTimeResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeprecatedGetProgramTimeResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["duration"] = o.Duration
	toSerialize["program_end"] = o.ProgramEnd
	toSerialize["program_start"] = o.ProgramStart
	return toSerialize, nil
}

type NullableDeprecatedGetProgramTimeResponse struct {
	value *DeprecatedGetProgramTimeResponse
	isSet bool
}

func (v NullableDeprecatedGetProgramTimeResponse) Get() *DeprecatedGetProgramTimeResponse {
	return v.value
}

func (v *NullableDeprecatedGetProgramTimeResponse) Set(val *DeprecatedGetProgramTimeResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeprecatedGetProgramTimeResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeprecatedGetProgramTimeResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeprecatedGetProgramTimeResponse(val *DeprecatedGetProgramTimeResponse) *NullableDeprecatedGetProgramTimeResponse {
	return &NullableDeprecatedGetProgramTimeResponse{value: val, isSet: true}
}

func (v NullableDeprecatedGetProgramTimeResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeprecatedGetProgramTimeResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
