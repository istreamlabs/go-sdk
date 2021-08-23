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

// ChannelTranscodeLoudness Loudness normalization settings.
type ChannelTranscodeLoudness struct {
	// Enable Dialog Intelligence. Only supported for (E)AC-3 encoders.
	DialogIntel *bool `json:"dialog_intel,omitempty"`
	// Loudness normalization LKFS setting. Default value is -24.
	Lkfs *int32 `json:"lkfs,omitempty"`
	// Loudness Range. Only supported for non-(E)AC-3 encoders. Default value is 7.0.
	Lra *float32 `json:"lra,omitempty"`
	// Peak Limit. Default value is -2.0.
	PeakLimit *float32 `json:"peak_limit,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChannelTranscodeLoudness ChannelTranscodeLoudness

// NewChannelTranscodeLoudness instantiates a new ChannelTranscodeLoudness object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelTranscodeLoudness() *ChannelTranscodeLoudness {
	this := ChannelTranscodeLoudness{}
	return &this
}

// NewChannelTranscodeLoudnessWithDefaults instantiates a new ChannelTranscodeLoudness object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelTranscodeLoudnessWithDefaults() *ChannelTranscodeLoudness {
	this := ChannelTranscodeLoudness{}
	return &this
}

// GetDialogIntel returns the DialogIntel field value if set, zero value otherwise.
func (o *ChannelTranscodeLoudness) GetDialogIntel() bool {
	if o == nil || o.DialogIntel == nil {
		var ret bool
		return ret
	}
	return *o.DialogIntel
}

// GetDialogIntelOk returns a tuple with the DialogIntel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeLoudness) GetDialogIntelOk() (*bool, bool) {
	if o == nil || o.DialogIntel == nil {
		return nil, false
	}
	return o.DialogIntel, true
}

// HasDialogIntel returns a boolean if a field has been set.
func (o *ChannelTranscodeLoudness) HasDialogIntel() bool {
	if o != nil && o.DialogIntel != nil {
		return true
	}

	return false
}

// SetDialogIntel gets a reference to the given bool and assigns it to the DialogIntel field.
func (o *ChannelTranscodeLoudness) SetDialogIntel(v bool) {
	o.DialogIntel = &v
}

// GetLkfs returns the Lkfs field value if set, zero value otherwise.
func (o *ChannelTranscodeLoudness) GetLkfs() int32 {
	if o == nil || o.Lkfs == nil {
		var ret int32
		return ret
	}
	return *o.Lkfs
}

// GetLkfsOk returns a tuple with the Lkfs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeLoudness) GetLkfsOk() (*int32, bool) {
	if o == nil || o.Lkfs == nil {
		return nil, false
	}
	return o.Lkfs, true
}

// HasLkfs returns a boolean if a field has been set.
func (o *ChannelTranscodeLoudness) HasLkfs() bool {
	if o != nil && o.Lkfs != nil {
		return true
	}

	return false
}

// SetLkfs gets a reference to the given int32 and assigns it to the Lkfs field.
func (o *ChannelTranscodeLoudness) SetLkfs(v int32) {
	o.Lkfs = &v
}

// GetLra returns the Lra field value if set, zero value otherwise.
func (o *ChannelTranscodeLoudness) GetLra() float32 {
	if o == nil || o.Lra == nil {
		var ret float32
		return ret
	}
	return *o.Lra
}

// GetLraOk returns a tuple with the Lra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeLoudness) GetLraOk() (*float32, bool) {
	if o == nil || o.Lra == nil {
		return nil, false
	}
	return o.Lra, true
}

// HasLra returns a boolean if a field has been set.
func (o *ChannelTranscodeLoudness) HasLra() bool {
	if o != nil && o.Lra != nil {
		return true
	}

	return false
}

// SetLra gets a reference to the given float32 and assigns it to the Lra field.
func (o *ChannelTranscodeLoudness) SetLra(v float32) {
	o.Lra = &v
}

// GetPeakLimit returns the PeakLimit field value if set, zero value otherwise.
func (o *ChannelTranscodeLoudness) GetPeakLimit() float32 {
	if o == nil || o.PeakLimit == nil {
		var ret float32
		return ret
	}
	return *o.PeakLimit
}

// GetPeakLimitOk returns a tuple with the PeakLimit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeLoudness) GetPeakLimitOk() (*float32, bool) {
	if o == nil || o.PeakLimit == nil {
		return nil, false
	}
	return o.PeakLimit, true
}

// HasPeakLimit returns a boolean if a field has been set.
func (o *ChannelTranscodeLoudness) HasPeakLimit() bool {
	if o != nil && o.PeakLimit != nil {
		return true
	}

	return false
}

// SetPeakLimit gets a reference to the given float32 and assigns it to the PeakLimit field.
func (o *ChannelTranscodeLoudness) SetPeakLimit(v float32) {
	o.PeakLimit = &v
}

func (o ChannelTranscodeLoudness) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DialogIntel != nil {
		toSerialize["dialog_intel"] = o.DialogIntel
	}
	if o.Lkfs != nil {
		toSerialize["lkfs"] = o.Lkfs
	}
	if o.Lra != nil {
		toSerialize["lra"] = o.Lra
	}
	if o.PeakLimit != nil {
		toSerialize["peak_limit"] = o.PeakLimit
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChannelTranscodeLoudness) UnmarshalJSON(bytes []byte) (err error) {
	varChannelTranscodeLoudness := _ChannelTranscodeLoudness{}

	if err = json.Unmarshal(bytes, &varChannelTranscodeLoudness); err == nil {
		*o = ChannelTranscodeLoudness(varChannelTranscodeLoudness)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dialog_intel")
		delete(additionalProperties, "lkfs")
		delete(additionalProperties, "lra")
		delete(additionalProperties, "peak_limit")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChannelTranscodeLoudness struct {
	value *ChannelTranscodeLoudness
	isSet bool
}

func (v NullableChannelTranscodeLoudness) Get() *ChannelTranscodeLoudness {
	return v.value
}

func (v *NullableChannelTranscodeLoudness) Set(val *ChannelTranscodeLoudness) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelTranscodeLoudness) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelTranscodeLoudness) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelTranscodeLoudness(val *ChannelTranscodeLoudness) *NullableChannelTranscodeLoudness {
	return &NullableChannelTranscodeLoudness{value: val, isSet: true}
}

func (v NullableChannelTranscodeLoudness) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelTranscodeLoudness) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


