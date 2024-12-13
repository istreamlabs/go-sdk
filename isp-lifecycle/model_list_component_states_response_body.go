/*
 * Channel Lifecycle State API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package isp

import (
	"encoding/json"
)

// checks if the ListComponentStatesResponseBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListComponentStatesResponseBody{}

// ListComponentStatesResponseBody struct for ListComponentStatesResponseBody
type ListComponentStatesResponseBody struct {
	// A URL to the JSON Schema for this object.
	Schema *string `json:"$schema,omitempty" format:"uri" doc:"A URL to the JSON Schema for this object."`
	// The list of components states.
	Components []ListComponentStatesResponseEntry `json:"components" doc:"The list of components states."`
}

// NewListComponentStatesResponseBody instantiates a new ListComponentStatesResponseBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListComponentStatesResponseBody(components []ListComponentStatesResponseEntry) *ListComponentStatesResponseBody {
	this := ListComponentStatesResponseBody{}
	this.Components = components
	return &this
}

// NewListComponentStatesResponseBodyWithDefaults instantiates a new ListComponentStatesResponseBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListComponentStatesResponseBodyWithDefaults() *ListComponentStatesResponseBody {
	this := ListComponentStatesResponseBody{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *ListComponentStatesResponseBody) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListComponentStatesResponseBody) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *ListComponentStatesResponseBody) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *ListComponentStatesResponseBody) SetSchema(v string) {
	o.Schema = &v
}

// GetComponents returns the Components field value
// If the value is explicit nil, the zero value for []ListComponentStatesResponseEntry will be returned
func (o *ListComponentStatesResponseBody) GetComponents() []ListComponentStatesResponseEntry {
	if o == nil {
		var ret []ListComponentStatesResponseEntry
		return ret
	}

	return o.Components
}

// GetComponentsOk returns a tuple with the Components field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ListComponentStatesResponseBody) GetComponentsOk() ([]ListComponentStatesResponseEntry, bool) {
	if o == nil || IsNil(o.Components) {
		return nil, false
	}
	return o.Components, true
}

// SetComponents sets field value
func (o *ListComponentStatesResponseBody) SetComponents(v []ListComponentStatesResponseEntry) {
	o.Components = v
}

func (o ListComponentStatesResponseBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListComponentStatesResponseBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	if o.Components != nil {
		toSerialize["components"] = o.Components
	}
	return toSerialize, nil
}

type NullableListComponentStatesResponseBody struct {
	value *ListComponentStatesResponseBody
	isSet bool
}

func (v NullableListComponentStatesResponseBody) Get() *ListComponentStatesResponseBody {
	return v.value
}

func (v *NullableListComponentStatesResponseBody) Set(val *ListComponentStatesResponseBody) {
	v.value = val
	v.isSet = true
}

func (v NullableListComponentStatesResponseBody) IsSet() bool {
	return v.isSet
}

func (v *NullableListComponentStatesResponseBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListComponentStatesResponseBody(val *ListComponentStatesResponseBody) *NullableListComponentStatesResponseBody {
	return &NullableListComponentStatesResponseBody{value: val, isSet: true}
}

func (v NullableListComponentStatesResponseBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListComponentStatesResponseBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

