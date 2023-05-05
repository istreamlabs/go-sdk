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

// checks if the InsertMetadataRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InsertMetadataRequest{}

// InsertMetadataRequest struct for InsertMetadataRequest
type InsertMetadataRequest struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string `json:"$schema,omitempty"`
	// ID3 payload as UTF-8 text
	Payload string `json:"payload"`
}

// NewInsertMetadataRequest instantiates a new InsertMetadataRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInsertMetadataRequest(payload string) *InsertMetadataRequest {
	this := InsertMetadataRequest{}
	this.Payload = payload
	return &this
}

// NewInsertMetadataRequestWithDefaults instantiates a new InsertMetadataRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInsertMetadataRequestWithDefaults() *InsertMetadataRequest {
	this := InsertMetadataRequest{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *InsertMetadataRequest) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InsertMetadataRequest) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *InsertMetadataRequest) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *InsertMetadataRequest) SetSchema(v string) {
	o.Schema = &v
}

// GetPayload returns the Payload field value
func (o *InsertMetadataRequest) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *InsertMetadataRequest) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *InsertMetadataRequest) SetPayload(v string) {
	o.Payload = v
}

func (o InsertMetadataRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InsertMetadataRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["payload"] = o.Payload
	return toSerialize, nil
}

type NullableInsertMetadataRequest struct {
	value *InsertMetadataRequest
	isSet bool
}

func (v NullableInsertMetadataRequest) Get() *InsertMetadataRequest {
	return v.value
}

func (v *NullableInsertMetadataRequest) Set(val *InsertMetadataRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInsertMetadataRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInsertMetadataRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsertMetadataRequest(val *InsertMetadataRequest) *NullableInsertMetadataRequest {
	return &NullableInsertMetadataRequest{value: val, isSet: true}
}

func (v NullableInsertMetadataRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsertMetadataRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


