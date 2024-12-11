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

// checks if the ChannelTimelineEntry type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelTimelineEntry{}

// ChannelTimelineEntry struct for ChannelTimelineEntry
type ChannelTimelineEntry struct {
	// Shortcode indicating what action was taken
	Action string `json:"action" doc:"Shortcode indicating what action was taken"`
	// Agent responsible for the action taken
	Agent string  `json:"agent" doc:"Agent responsible for the action taken"`
	Query *string `json:"query,omitempty"`
	// The request body, if any, of the original action
	RequestBody string `json:"request_body" doc:"The request body, if any, of the original action"`
	// HTTP Status code indicating outcome of the action.
	StatusCode int32 `json:"status_code" format:"int32" doc:"HTTP Status code indicating outcome of the action."`
	// Timestamp of the action in UTC
	Timestamp time.Time `json:"timestamp" format:"date-time" doc:"Timestamp of the action in UTC"`
	// Correlation identifier for tracing
	TraceId string `json:"trace_id" doc:"Correlation identifier for tracing"`
}

// NewChannelTimelineEntry instantiates a new ChannelTimelineEntry object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelTimelineEntry(action string, agent string, requestBody string, statusCode int32, timestamp time.Time, traceId string) *ChannelTimelineEntry {
	this := ChannelTimelineEntry{}
	this.Action = action
	this.Agent = agent
	this.RequestBody = requestBody
	this.StatusCode = statusCode
	this.Timestamp = timestamp
	this.TraceId = traceId
	return &this
}

// NewChannelTimelineEntryWithDefaults instantiates a new ChannelTimelineEntry object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelTimelineEntryWithDefaults() *ChannelTimelineEntry {
	this := ChannelTimelineEntry{}
	return &this
}

// GetAction returns the Action field value
func (o *ChannelTimelineEntry) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *ChannelTimelineEntry) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *ChannelTimelineEntry) SetAction(v string) {
	o.Action = v
}

// GetAgent returns the Agent field value
func (o *ChannelTimelineEntry) GetAgent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Agent
}

// GetAgentOk returns a tuple with the Agent field value
// and a boolean to check if the value has been set.
func (o *ChannelTimelineEntry) GetAgentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Agent, true
}

// SetAgent sets field value
func (o *ChannelTimelineEntry) SetAgent(v string) {
	o.Agent = v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *ChannelTimelineEntry) GetQuery() string {
	if o == nil || IsNil(o.Query) {
		var ret string
		return ret
	}
	return *o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTimelineEntry) GetQueryOk() (*string, bool) {
	if o == nil || IsNil(o.Query) {
		return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *ChannelTimelineEntry) HasQuery() bool {
	if o != nil && !IsNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given string and assigns it to the Query field.
func (o *ChannelTimelineEntry) SetQuery(v string) {
	o.Query = &v
}

// GetRequestBody returns the RequestBody field value
func (o *ChannelTimelineEntry) GetRequestBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestBody
}

// GetRequestBodyOk returns a tuple with the RequestBody field value
// and a boolean to check if the value has been set.
func (o *ChannelTimelineEntry) GetRequestBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestBody, true
}

// SetRequestBody sets field value
func (o *ChannelTimelineEntry) SetRequestBody(v string) {
	o.RequestBody = v
}

// GetStatusCode returns the StatusCode field value
func (o *ChannelTimelineEntry) GetStatusCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.StatusCode
}

// GetStatusCodeOk returns a tuple with the StatusCode field value
// and a boolean to check if the value has been set.
func (o *ChannelTimelineEntry) GetStatusCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StatusCode, true
}

// SetStatusCode sets field value
func (o *ChannelTimelineEntry) SetStatusCode(v int32) {
	o.StatusCode = v
}

// GetTimestamp returns the Timestamp field value
func (o *ChannelTimelineEntry) GetTimestamp() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *ChannelTimelineEntry) GetTimestampOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *ChannelTimelineEntry) SetTimestamp(v time.Time) {
	o.Timestamp = v
}

// GetTraceId returns the TraceId field value
func (o *ChannelTimelineEntry) GetTraceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TraceId
}

// GetTraceIdOk returns a tuple with the TraceId field value
// and a boolean to check if the value has been set.
func (o *ChannelTimelineEntry) GetTraceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TraceId, true
}

// SetTraceId sets field value
func (o *ChannelTimelineEntry) SetTraceId(v string) {
	o.TraceId = v
}

func (o ChannelTimelineEntry) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelTimelineEntry) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["action"] = o.Action
	toSerialize["agent"] = o.Agent
	if !IsNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	toSerialize["request_body"] = o.RequestBody
	toSerialize["status_code"] = o.StatusCode
	toSerialize["timestamp"] = o.Timestamp
	toSerialize["trace_id"] = o.TraceId
	return toSerialize, nil
}

type NullableChannelTimelineEntry struct {
	value *ChannelTimelineEntry
	isSet bool
}

func (v NullableChannelTimelineEntry) Get() *ChannelTimelineEntry {
	return v.value
}

func (v *NullableChannelTimelineEntry) Set(val *ChannelTimelineEntry) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelTimelineEntry) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelTimelineEntry) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelTimelineEntry(val *ChannelTimelineEntry) *NullableChannelTimelineEntry {
	return &NullableChannelTimelineEntry{value: val, isSet: true}
}

func (v NullableChannelTimelineEntry) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelTimelineEntry) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
