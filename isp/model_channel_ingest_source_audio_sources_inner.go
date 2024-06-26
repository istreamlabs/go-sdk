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

// checks if the ChannelIngestSourceAudioSourcesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelIngestSourceAudioSourcesInner{}

// ChannelIngestSourceAudioSourcesInner struct for ChannelIngestSourceAudioSourcesInner
type ChannelIngestSourceAudioSourcesInner struct {
	Id *string `json:"id,omitempty" minLength:"1"`
	// RFC 5646, e.g. 'en' 'en-US'
	Language *string `json:"language,omitempty" minLength:"1" doc:"RFC 5646, e.g. 'en' 'en-US'"`
	// License specifies how the audio in this source is licensed.
	License *string `json:"license,omitempty" enum:"LIVE,REPLAY" doc:"License specifies how the audio in this source is licensed."`
	// Language fiendly name, e.g. 'English', 'Spanish'
	Name *string `json:"name,omitempty" minLength:"1" doc:"Language fiendly name, e.g. 'English', 'Spanish'"`
	// Expression for choosing an audio track in the stream for this AudioSource https://wbdstreaming.atlassian.net/wiki/spaces/LIVE/pages/250351679/Proposal+Audio+Track+Selection
	Selector *string `json:"selector,omitempty" doc:"Expression for choosing an audio track in the stream for this AudioSource https://wbdstreaming.atlassian.net/wiki/spaces/LIVE/pages/250351679/Proposal+Audio+Track+Selection"`
}

// NewChannelIngestSourceAudioSourcesInner instantiates a new ChannelIngestSourceAudioSourcesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelIngestSourceAudioSourcesInner() *ChannelIngestSourceAudioSourcesInner {
	this := ChannelIngestSourceAudioSourcesInner{}
	return &this
}

// NewChannelIngestSourceAudioSourcesInnerWithDefaults instantiates a new ChannelIngestSourceAudioSourcesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelIngestSourceAudioSourcesInnerWithDefaults() *ChannelIngestSourceAudioSourcesInner {
	this := ChannelIngestSourceAudioSourcesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChannelIngestSourceAudioSourcesInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelIngestSourceAudioSourcesInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChannelIngestSourceAudioSourcesInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChannelIngestSourceAudioSourcesInner) SetId(v string) {
	o.Id = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *ChannelIngestSourceAudioSourcesInner) GetLanguage() string {
	if o == nil || IsNil(o.Language) {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelIngestSourceAudioSourcesInner) GetLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.Language) {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *ChannelIngestSourceAudioSourcesInner) HasLanguage() bool {
	if o != nil && !IsNil(o.Language) {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *ChannelIngestSourceAudioSourcesInner) SetLanguage(v string) {
	o.Language = &v
}

// GetLicense returns the License field value if set, zero value otherwise.
func (o *ChannelIngestSourceAudioSourcesInner) GetLicense() string {
	if o == nil || IsNil(o.License) {
		var ret string
		return ret
	}
	return *o.License
}

// GetLicenseOk returns a tuple with the License field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelIngestSourceAudioSourcesInner) GetLicenseOk() (*string, bool) {
	if o == nil || IsNil(o.License) {
		return nil, false
	}
	return o.License, true
}

// HasLicense returns a boolean if a field has been set.
func (o *ChannelIngestSourceAudioSourcesInner) HasLicense() bool {
	if o != nil && !IsNil(o.License) {
		return true
	}

	return false
}

// SetLicense gets a reference to the given string and assigns it to the License field.
func (o *ChannelIngestSourceAudioSourcesInner) SetLicense(v string) {
	o.License = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ChannelIngestSourceAudioSourcesInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelIngestSourceAudioSourcesInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ChannelIngestSourceAudioSourcesInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ChannelIngestSourceAudioSourcesInner) SetName(v string) {
	o.Name = &v
}

// GetSelector returns the Selector field value if set, zero value otherwise.
func (o *ChannelIngestSourceAudioSourcesInner) GetSelector() string {
	if o == nil || IsNil(o.Selector) {
		var ret string
		return ret
	}
	return *o.Selector
}

// GetSelectorOk returns a tuple with the Selector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelIngestSourceAudioSourcesInner) GetSelectorOk() (*string, bool) {
	if o == nil || IsNil(o.Selector) {
		return nil, false
	}
	return o.Selector, true
}

// HasSelector returns a boolean if a field has been set.
func (o *ChannelIngestSourceAudioSourcesInner) HasSelector() bool {
	if o != nil && !IsNil(o.Selector) {
		return true
	}

	return false
}

// SetSelector gets a reference to the given string and assigns it to the Selector field.
func (o *ChannelIngestSourceAudioSourcesInner) SetSelector(v string) {
	o.Selector = &v
}

func (o ChannelIngestSourceAudioSourcesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelIngestSourceAudioSourcesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Language) {
		toSerialize["language"] = o.Language
	}
	if !IsNil(o.License) {
		toSerialize["license"] = o.License
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Selector) {
		toSerialize["selector"] = o.Selector
	}
	return toSerialize, nil
}

type NullableChannelIngestSourceAudioSourcesInner struct {
	value *ChannelIngestSourceAudioSourcesInner
	isSet bool
}

func (v NullableChannelIngestSourceAudioSourcesInner) Get() *ChannelIngestSourceAudioSourcesInner {
	return v.value
}

func (v *NullableChannelIngestSourceAudioSourcesInner) Set(val *ChannelIngestSourceAudioSourcesInner) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelIngestSourceAudioSourcesInner) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelIngestSourceAudioSourcesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelIngestSourceAudioSourcesInner(val *ChannelIngestSourceAudioSourcesInner) *NullableChannelIngestSourceAudioSourcesInner {
	return &NullableChannelIngestSourceAudioSourcesInner{value: val, isSet: true}
}

func (v NullableChannelIngestSourceAudioSourcesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelIngestSourceAudioSourcesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


