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

// checks if the Scte35 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Scte35{}

// Scte35 struct for Scte35
type Scte35 struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string `json:"$schema,omitempty" format:"uri" doc:"An optional URL to a JSON Schema document describing this resource"`
	// The SCTE-35 payload, encoded as base-64 in JSON or binary data in CBOR
	Payload string `json:"payload" doc:"The SCTE-35 payload, encoded as base-64 in JSON or binary data in CBOR"`
}

// NewScte35 instantiates a new Scte35 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScte35(payload string) *Scte35 {
	this := Scte35{}
	this.Payload = payload
	return &this
}

// NewScte35WithDefaults instantiates a new Scte35 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScte35WithDefaults() *Scte35 {
	this := Scte35{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *Scte35) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Scte35) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *Scte35) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *Scte35) SetSchema(v string) {
	o.Schema = &v
}

// GetPayload returns the Payload field value
func (o *Scte35) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *Scte35) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *Scte35) SetPayload(v string) {
	o.Payload = v
}

func (o Scte35) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Scte35) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["payload"] = o.Payload
	return toSerialize, nil
}

type NullableScte35 struct {
	value *Scte35
	isSet bool
}

func (v NullableScte35) Get() *Scte35 {
	return v.value
}

func (v *NullableScte35) Set(val *Scte35) {
	v.value = val
	v.isSet = true
}

func (v NullableScte35) IsSet() bool {
	return v.isSet
}

func (v *NullableScte35) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScte35(val *Scte35) *NullableScte35 {
	return &NullableScte35{value: val, isSet: true}
}

func (v NullableScte35) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScte35) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
