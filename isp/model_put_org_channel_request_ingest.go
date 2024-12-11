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

// checks if the PutOrgChannelRequestIngest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PutOrgChannelRequestIngest{}

// PutOrgChannelRequestIngest Ingest configures inputs for the transcoder.
type PutOrgChannelRequestIngest struct {
	Slate  *ChannelIngestSlate              `json:"slate,omitempty"`
	Source PutOrgChannelRequestIngestSource `json:"source"`
}

// NewPutOrgChannelRequestIngest instantiates a new PutOrgChannelRequestIngest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutOrgChannelRequestIngest(source PutOrgChannelRequestIngestSource) *PutOrgChannelRequestIngest {
	this := PutOrgChannelRequestIngest{}
	this.Source = source
	return &this
}

// NewPutOrgChannelRequestIngestWithDefaults instantiates a new PutOrgChannelRequestIngest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutOrgChannelRequestIngestWithDefaults() *PutOrgChannelRequestIngest {
	this := PutOrgChannelRequestIngest{}
	return &this
}

// GetSlate returns the Slate field value if set, zero value otherwise.
func (o *PutOrgChannelRequestIngest) GetSlate() ChannelIngestSlate {
	if o == nil || IsNil(o.Slate) {
		var ret ChannelIngestSlate
		return ret
	}
	return *o.Slate
}

// GetSlateOk returns a tuple with the Slate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutOrgChannelRequestIngest) GetSlateOk() (*ChannelIngestSlate, bool) {
	if o == nil || IsNil(o.Slate) {
		return nil, false
	}
	return o.Slate, true
}

// HasSlate returns a boolean if a field has been set.
func (o *PutOrgChannelRequestIngest) HasSlate() bool {
	if o != nil && !IsNil(o.Slate) {
		return true
	}

	return false
}

// SetSlate gets a reference to the given ChannelIngestSlate and assigns it to the Slate field.
func (o *PutOrgChannelRequestIngest) SetSlate(v ChannelIngestSlate) {
	o.Slate = &v
}

// GetSource returns the Source field value
func (o *PutOrgChannelRequestIngest) GetSource() PutOrgChannelRequestIngestSource {
	if o == nil {
		var ret PutOrgChannelRequestIngestSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *PutOrgChannelRequestIngest) GetSourceOk() (*PutOrgChannelRequestIngestSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *PutOrgChannelRequestIngest) SetSource(v PutOrgChannelRequestIngestSource) {
	o.Source = v
}

func (o PutOrgChannelRequestIngest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PutOrgChannelRequestIngest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Slate) {
		toSerialize["slate"] = o.Slate
	}
	toSerialize["source"] = o.Source
	return toSerialize, nil
}

type NullablePutOrgChannelRequestIngest struct {
	value *PutOrgChannelRequestIngest
	isSet bool
}

func (v NullablePutOrgChannelRequestIngest) Get() *PutOrgChannelRequestIngest {
	return v.value
}

func (v *NullablePutOrgChannelRequestIngest) Set(val *PutOrgChannelRequestIngest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutOrgChannelRequestIngest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutOrgChannelRequestIngest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutOrgChannelRequestIngest(val *PutOrgChannelRequestIngest) *NullablePutOrgChannelRequestIngest {
	return &NullablePutOrgChannelRequestIngest{value: val, isSet: true}
}

func (v NullablePutOrgChannelRequestIngest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutOrgChannelRequestIngest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
