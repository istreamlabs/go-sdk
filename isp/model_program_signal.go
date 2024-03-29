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

// checks if the ProgramSignal type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProgramSignal{}

// ProgramSignal struct for ProgramSignal
type ProgramSignal struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string `json:"$schema,omitempty" format:"uri" doc:"An optional URL to a JSON Schema document describing this resource"`
	// Event ID
	EventId int32 `json:"event_id" format:"int32" minimum:"0" doc:"Event ID"`
}

// NewProgramSignal instantiates a new ProgramSignal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProgramSignal(eventId int32) *ProgramSignal {
	this := ProgramSignal{}
	this.EventId = eventId
	return &this
}

// NewProgramSignalWithDefaults instantiates a new ProgramSignal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProgramSignalWithDefaults() *ProgramSignal {
	this := ProgramSignal{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *ProgramSignal) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProgramSignal) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *ProgramSignal) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *ProgramSignal) SetSchema(v string) {
	o.Schema = &v
}

// GetEventId returns the EventId field value
func (o *ProgramSignal) GetEventId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *ProgramSignal) GetEventIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *ProgramSignal) SetEventId(v int32) {
	o.EventId = v
}

func (o ProgramSignal) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProgramSignal) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["event_id"] = o.EventId
	return toSerialize, nil
}

type NullableProgramSignal struct {
	value *ProgramSignal
	isSet bool
}

func (v NullableProgramSignal) Get() *ProgramSignal {
	return v.value
}

func (v *NullableProgramSignal) Set(val *ProgramSignal) {
	v.value = val
	v.isSet = true
}

func (v NullableProgramSignal) IsSet() bool {
	return v.isSet
}

func (v *NullableProgramSignal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProgramSignal(val *ProgramSignal) *NullableProgramSignal {
	return &NullableProgramSignal{value: val, isSet: true}
}

func (v NullableProgramSignal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProgramSignal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


