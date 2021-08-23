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

// ProgramSignal struct for ProgramSignal
type ProgramSignal struct {
	// Event ID
	EventId int32 `json:"event_id"`
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
	if o == nil  {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *ProgramSignal) SetEventId(v int32) {
	o.EventId = v
}

func (o ProgramSignal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["event_id"] = o.EventId
	}
	return json.Marshal(toSerialize)
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


