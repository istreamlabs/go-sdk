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

// checks if the PutChannelRequestIngest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutChannelRequestIngest{}

// PutChannelRequestIngest Ingest configures inputs for the transcoder.
type PutChannelRequestIngest struct {
	Slate *ChannelIngestSlate `json:"slate,omitempty"`
	Source PutChannelRequestIngestSource `json:"source"`
}

// NewPutChannelRequestIngest instantiates a new PutChannelRequestIngest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutChannelRequestIngest(source PutChannelRequestIngestSource) *PutChannelRequestIngest {
	this := PutChannelRequestIngest{}
	this.Source = source
	return &this
}

// NewPutChannelRequestIngestWithDefaults instantiates a new PutChannelRequestIngest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutChannelRequestIngestWithDefaults() *PutChannelRequestIngest {
	this := PutChannelRequestIngest{}
	return &this
}

// GetSlate returns the Slate field value if set, zero value otherwise.
func (o *PutChannelRequestIngest) GetSlate() ChannelIngestSlate {
	if o == nil || IsNil(o.Slate) {
		var ret ChannelIngestSlate
		return ret
	}
	return *o.Slate
}

// GetSlateOk returns a tuple with the Slate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutChannelRequestIngest) GetSlateOk() (*ChannelIngestSlate, bool) {
	if o == nil || IsNil(o.Slate) {
		return nil, false
	}
	return o.Slate, true
}

// HasSlate returns a boolean if a field has been set.
func (o *PutChannelRequestIngest) HasSlate() bool {
	if o != nil && !IsNil(o.Slate) {
		return true
	}

	return false
}

// SetSlate gets a reference to the given ChannelIngestSlate and assigns it to the Slate field.
func (o *PutChannelRequestIngest) SetSlate(v ChannelIngestSlate) {
	o.Slate = &v
}

// GetSource returns the Source field value
func (o *PutChannelRequestIngest) GetSource() PutChannelRequestIngestSource {
	if o == nil {
		var ret PutChannelRequestIngestSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *PutChannelRequestIngest) GetSourceOk() (*PutChannelRequestIngestSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *PutChannelRequestIngest) SetSource(v PutChannelRequestIngestSource) {
	o.Source = v
}

func (o PutChannelRequestIngest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutChannelRequestIngest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Slate) {
		toSerialize["slate"] = o.Slate
	}
	toSerialize["source"] = o.Source
	return toSerialize, nil
}

type NullablePutChannelRequestIngest struct {
	value *PutChannelRequestIngest
	isSet bool
}

func (v NullablePutChannelRequestIngest) Get() *PutChannelRequestIngest {
	return v.value
}

func (v *NullablePutChannelRequestIngest) Set(val *PutChannelRequestIngest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutChannelRequestIngest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutChannelRequestIngest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutChannelRequestIngest(val *PutChannelRequestIngest) *NullablePutChannelRequestIngest {
	return &NullablePutChannelRequestIngest{value: val, isSet: true}
}

func (v NullablePutChannelRequestIngest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutChannelRequestIngest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


