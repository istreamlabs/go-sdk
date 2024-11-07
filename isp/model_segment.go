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
	"time"
)

// checks if the Segment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Segment{}

// Segment struct for Segment
type Segment struct {
	// The program time when the segment ends (or ended).
	End *time.Time `json:"end,omitempty" format:"date-time" doc:"The program time when the segment ends (or ended)."`
	// The event_id used in the Signal START that resulted in this signaling segment.
	EventId *int32 `json:"event_id,omitempty" format:"int32" minimum:"0" doc:"The event_id used in the Signal START that resulted in this signaling segment."`
	// Uniquely identifies the signaling segment for a channel. This is assigned by the transcoder when the segment is created.
	SegmentId *int64 `json:"segment_id,omitempty" format:"int64" minimum:"0" doc:"Uniquely identifies the signaling segment for a channel. This is assigned by the transcoder when the segment is created."`
	// When set, contains the URL to the slate media asset that will play for the duration of the segment.
	SlateUrl *string `json:"slate_url,omitempty" doc:"When set, contains the URL to the slate media asset that will play for the duration of the segment."`
	// The program time when the segment starts (or started).
	Start *time.Time `json:"start,omitempty" format:"date-time" doc:"The program time when the segment starts (or started)."`
	// Timed metadata opaque payload data.
	TimedMetadataPayload *string `json:"timed_metadata_payload,omitempty" doc:"Timed metadata opaque payload data."`
	// The timed metadata signal's type.
	TimedMetadataType *string `json:"timed_metadata_type,omitempty" doc:"The timed metadata signal's type."`
	Type *string `json:"type,omitempty" enum:"SPLICE_INSERT,NOT_INDICATED,CONTENT_ID,PROGRAM,PROGRAM_BLACKOUT_OVERRIDE,PROGRAM_BREAKAWAY,CHAPTER,BREAK,OPENING_CREDIT,CLOSING_CREDIT,PROVIDER_PLACEMENT_OP,DISTRIBUTOR_PLACEMENT_OP,PROVIDER_OVERLAY_OP,DISTRIBUTOR_OVERLAY_OP,PROVIDER_AD,DISTRIBUTOR_AD,UNSCHEDULED_EVENT,NETWORK,SLATE,TIMED_METADATA"`
}

// NewSegment instantiates a new Segment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSegment() *Segment {
	this := Segment{}
	return &this
}

// NewSegmentWithDefaults instantiates a new Segment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSegmentWithDefaults() *Segment {
	this := Segment{}
	return &this
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *Segment) GetEnd() time.Time {
	if o == nil || IsNil(o.End) {
		var ret time.Time
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetEndOk() (*time.Time, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *Segment) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given time.Time and assigns it to the End field.
func (o *Segment) SetEnd(v time.Time) {
	o.End = &v
}

// GetEventId returns the EventId field value if set, zero value otherwise.
func (o *Segment) GetEventId() int32 {
	if o == nil || IsNil(o.EventId) {
		var ret int32
		return ret
	}
	return *o.EventId
}

// GetEventIdOk returns a tuple with the EventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetEventIdOk() (*int32, bool) {
	if o == nil || IsNil(o.EventId) {
		return nil, false
	}
	return o.EventId, true
}

// HasEventId returns a boolean if a field has been set.
func (o *Segment) HasEventId() bool {
	if o != nil && !IsNil(o.EventId) {
		return true
	}

	return false
}

// SetEventId gets a reference to the given int32 and assigns it to the EventId field.
func (o *Segment) SetEventId(v int32) {
	o.EventId = &v
}

// GetSegmentId returns the SegmentId field value if set, zero value otherwise.
func (o *Segment) GetSegmentId() int64 {
	if o == nil || IsNil(o.SegmentId) {
		var ret int64
		return ret
	}
	return *o.SegmentId
}

// GetSegmentIdOk returns a tuple with the SegmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetSegmentIdOk() (*int64, bool) {
	if o == nil || IsNil(o.SegmentId) {
		return nil, false
	}
	return o.SegmentId, true
}

// HasSegmentId returns a boolean if a field has been set.
func (o *Segment) HasSegmentId() bool {
	if o != nil && !IsNil(o.SegmentId) {
		return true
	}

	return false
}

// SetSegmentId gets a reference to the given int64 and assigns it to the SegmentId field.
func (o *Segment) SetSegmentId(v int64) {
	o.SegmentId = &v
}

// GetSlateUrl returns the SlateUrl field value if set, zero value otherwise.
func (o *Segment) GetSlateUrl() string {
	if o == nil || IsNil(o.SlateUrl) {
		var ret string
		return ret
	}
	return *o.SlateUrl
}

// GetSlateUrlOk returns a tuple with the SlateUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetSlateUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SlateUrl) {
		return nil, false
	}
	return o.SlateUrl, true
}

// HasSlateUrl returns a boolean if a field has been set.
func (o *Segment) HasSlateUrl() bool {
	if o != nil && !IsNil(o.SlateUrl) {
		return true
	}

	return false
}

// SetSlateUrl gets a reference to the given string and assigns it to the SlateUrl field.
func (o *Segment) SetSlateUrl(v string) {
	o.SlateUrl = &v
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *Segment) GetStart() time.Time {
	if o == nil || IsNil(o.Start) {
		var ret time.Time
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *Segment) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given time.Time and assigns it to the Start field.
func (o *Segment) SetStart(v time.Time) {
	o.Start = &v
}

// GetTimedMetadataPayload returns the TimedMetadataPayload field value if set, zero value otherwise.
func (o *Segment) GetTimedMetadataPayload() string {
	if o == nil || IsNil(o.TimedMetadataPayload) {
		var ret string
		return ret
	}
	return *o.TimedMetadataPayload
}

// GetTimedMetadataPayloadOk returns a tuple with the TimedMetadataPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetTimedMetadataPayloadOk() (*string, bool) {
	if o == nil || IsNil(o.TimedMetadataPayload) {
		return nil, false
	}
	return o.TimedMetadataPayload, true
}

// HasTimedMetadataPayload returns a boolean if a field has been set.
func (o *Segment) HasTimedMetadataPayload() bool {
	if o != nil && !IsNil(o.TimedMetadataPayload) {
		return true
	}

	return false
}

// SetTimedMetadataPayload gets a reference to the given string and assigns it to the TimedMetadataPayload field.
func (o *Segment) SetTimedMetadataPayload(v string) {
	o.TimedMetadataPayload = &v
}

// GetTimedMetadataType returns the TimedMetadataType field value if set, zero value otherwise.
func (o *Segment) GetTimedMetadataType() string {
	if o == nil || IsNil(o.TimedMetadataType) {
		var ret string
		return ret
	}
	return *o.TimedMetadataType
}

// GetTimedMetadataTypeOk returns a tuple with the TimedMetadataType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetTimedMetadataTypeOk() (*string, bool) {
	if o == nil || IsNil(o.TimedMetadataType) {
		return nil, false
	}
	return o.TimedMetadataType, true
}

// HasTimedMetadataType returns a boolean if a field has been set.
func (o *Segment) HasTimedMetadataType() bool {
	if o != nil && !IsNil(o.TimedMetadataType) {
		return true
	}

	return false
}

// SetTimedMetadataType gets a reference to the given string and assigns it to the TimedMetadataType field.
func (o *Segment) SetTimedMetadataType(v string) {
	o.TimedMetadataType = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Segment) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Segment) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Segment) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Segment) SetType(v string) {
	o.Type = &v
}

func (o Segment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Segment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !IsNil(o.EventId) {
		toSerialize["event_id"] = o.EventId
	}
	if !IsNil(o.SegmentId) {
		toSerialize["segment_id"] = o.SegmentId
	}
	if !IsNil(o.SlateUrl) {
		toSerialize["slate_url"] = o.SlateUrl
	}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.TimedMetadataPayload) {
		toSerialize["timed_metadata_payload"] = o.TimedMetadataPayload
	}
	if !IsNil(o.TimedMetadataType) {
		toSerialize["timed_metadata_type"] = o.TimedMetadataType
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableSegment struct {
	value *Segment
	isSet bool
}

func (v NullableSegment) Get() *Segment {
	return v.value
}

func (v *NullableSegment) Set(val *Segment) {
	v.value = val
	v.isSet = true
}

func (v NullableSegment) IsSet() bool {
	return v.isSet
}

func (v *NullableSegment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSegment(val *Segment) *NullableSegment {
	return &NullableSegment{value: val, isSet: true}
}

func (v NullableSegment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSegment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


