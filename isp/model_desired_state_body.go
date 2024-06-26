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

// checks if the DesiredStateBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DesiredStateBody{}

// DesiredStateBody struct for DesiredStateBody
type DesiredStateBody struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string `json:"$schema,omitempty" format:"uri" doc:"An optional URL to a JSON Schema document describing this resource"`
	// Desired state
	DesiredState string `json:"desired_state" enum:"ON,OFF" doc:"Desired state"`
}

// NewDesiredStateBody instantiates a new DesiredStateBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDesiredStateBody(desiredState string) *DesiredStateBody {
	this := DesiredStateBody{}
	this.DesiredState = desiredState
	return &this
}

// NewDesiredStateBodyWithDefaults instantiates a new DesiredStateBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDesiredStateBodyWithDefaults() *DesiredStateBody {
	this := DesiredStateBody{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *DesiredStateBody) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DesiredStateBody) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *DesiredStateBody) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *DesiredStateBody) SetSchema(v string) {
	o.Schema = &v
}

// GetDesiredState returns the DesiredState field value
func (o *DesiredStateBody) GetDesiredState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DesiredState
}

// GetDesiredStateOk returns a tuple with the DesiredState field value
// and a boolean to check if the value has been set.
func (o *DesiredStateBody) GetDesiredStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DesiredState, true
}

// SetDesiredState sets field value
func (o *DesiredStateBody) SetDesiredState(v string) {
	o.DesiredState = v
}

func (o DesiredStateBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DesiredStateBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["desired_state"] = o.DesiredState
	return toSerialize, nil
}

type NullableDesiredStateBody struct {
	value *DesiredStateBody
	isSet bool
}

func (v NullableDesiredStateBody) Get() *DesiredStateBody {
	return v.value
}

func (v *NullableDesiredStateBody) Set(val *DesiredStateBody) {
	v.value = val
	v.isSet = true
}

func (v NullableDesiredStateBody) IsSet() bool {
	return v.isSet
}

func (v *NullableDesiredStateBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDesiredStateBody(val *DesiredStateBody) *NullableDesiredStateBody {
	return &NullableDesiredStateBody{value: val, isSet: true}
}

func (v NullableDesiredStateBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDesiredStateBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


