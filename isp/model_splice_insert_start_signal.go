/*
 * iStreamPlanet Channels API
 *
 * API version: 0.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package isp

import (
	"encoding/json"
)

// SpliceInsertStartSignal struct for SpliceInsertStartSignal
type SpliceInsertStartSignal struct {
	// Splice duration (ms).  If no duration or a duration of 0 is passed then the splice is indifinite.
	Duration *int32 `json:"duration,omitempty"`
	// Event ID
	EventId int32 `json:"event_id"`
	// Slate url
	SlateUri *string `json:"slate_uri,omitempty"`
	// UPIDs
	Upids *[]string `json:"upids,omitempty"`
}

// NewSpliceInsertStartSignal instantiates a new SpliceInsertStartSignal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSpliceInsertStartSignal(eventId int32, ) *SpliceInsertStartSignal {
	this := SpliceInsertStartSignal{}
	var duration int32 = 0
	this.Duration = &duration
	this.EventId = eventId
	return &this
}

// NewSpliceInsertStartSignalWithDefaults instantiates a new SpliceInsertStartSignal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSpliceInsertStartSignalWithDefaults() *SpliceInsertStartSignal {
	this := SpliceInsertStartSignal{}
	var duration int32 = 0
	this.Duration = &duration
	return &this
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SpliceInsertStartSignal) GetDuration() int32 {
	if o == nil || o.Duration == nil {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpliceInsertStartSignal) GetDurationOk() (*int32, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SpliceInsertStartSignal) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *SpliceInsertStartSignal) SetDuration(v int32) {
	o.Duration = &v
}

// GetEventId returns the EventId field value
func (o *SpliceInsertStartSignal) GetEventId() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value
// and a boolean to check if the value has been set.
func (o *SpliceInsertStartSignal) GetEventIdOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EventId, true
}

// SetEventId sets field value
func (o *SpliceInsertStartSignal) SetEventId(v int32) {
	o.EventId = v
}

// GetSlateUri returns the SlateUri field value if set, zero value otherwise.
func (o *SpliceInsertStartSignal) GetSlateUri() string {
	if o == nil || o.SlateUri == nil {
		var ret string
		return ret
	}
	return *o.SlateUri
}

// GetSlateUriOk returns a tuple with the SlateUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpliceInsertStartSignal) GetSlateUriOk() (*string, bool) {
	if o == nil || o.SlateUri == nil {
		return nil, false
	}
	return o.SlateUri, true
}

// HasSlateUri returns a boolean if a field has been set.
func (o *SpliceInsertStartSignal) HasSlateUri() bool {
	if o != nil && o.SlateUri != nil {
		return true
	}

	return false
}

// SetSlateUri gets a reference to the given string and assigns it to the SlateUri field.
func (o *SpliceInsertStartSignal) SetSlateUri(v string) {
	o.SlateUri = &v
}

// GetUpids returns the Upids field value if set, zero value otherwise.
func (o *SpliceInsertStartSignal) GetUpids() []string {
	if o == nil || o.Upids == nil {
		var ret []string
		return ret
	}
	return *o.Upids
}

// GetUpidsOk returns a tuple with the Upids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SpliceInsertStartSignal) GetUpidsOk() (*[]string, bool) {
	if o == nil || o.Upids == nil {
		return nil, false
	}
	return o.Upids, true
}

// HasUpids returns a boolean if a field has been set.
func (o *SpliceInsertStartSignal) HasUpids() bool {
	if o != nil && o.Upids != nil {
		return true
	}

	return false
}

// SetUpids gets a reference to the given []string and assigns it to the Upids field.
func (o *SpliceInsertStartSignal) SetUpids(v []string) {
	o.Upids = &v
}

func (o SpliceInsertStartSignal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if true {
		toSerialize["event_id"] = o.EventId
	}
	if o.SlateUri != nil {
		toSerialize["slate_uri"] = o.SlateUri
	}
	if o.Upids != nil {
		toSerialize["upids"] = o.Upids
	}
	return json.Marshal(toSerialize)
}

type NullableSpliceInsertStartSignal struct {
	value *SpliceInsertStartSignal
	isSet bool
}

func (v NullableSpliceInsertStartSignal) Get() *SpliceInsertStartSignal {
	return v.value
}

func (v *NullableSpliceInsertStartSignal) Set(val *SpliceInsertStartSignal) {
	v.value = val
	v.isSet = true
}

func (v NullableSpliceInsertStartSignal) IsSet() bool {
	return v.isSet
}

func (v *NullableSpliceInsertStartSignal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSpliceInsertStartSignal(val *SpliceInsertStartSignal) *NullableSpliceInsertStartSignal {
	return &NullableSpliceInsertStartSignal{value: val, isSet: true}
}

func (v NullableSpliceInsertStartSignal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSpliceInsertStartSignal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


