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

// checks if the CollapseVODRequestFilterconfigScte35Options type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollapseVODRequestFilterconfigScte35Options{}

// CollapseVODRequestFilterconfigScte35Options Additional SCTE filter options
type CollapseVODRequestFilterconfigScte35Options struct {
	// Signal removal of provider ads
	SignalRemovedProviderAd bool `json:"signal_removed_provider_ad"`
}

// NewCollapseVODRequestFilterconfigScte35Options instantiates a new CollapseVODRequestFilterconfigScte35Options object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollapseVODRequestFilterconfigScte35Options(signalRemovedProviderAd bool) *CollapseVODRequestFilterconfigScte35Options {
	this := CollapseVODRequestFilterconfigScte35Options{}
	this.SignalRemovedProviderAd = signalRemovedProviderAd
	return &this
}

// NewCollapseVODRequestFilterconfigScte35OptionsWithDefaults instantiates a new CollapseVODRequestFilterconfigScte35Options object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollapseVODRequestFilterconfigScte35OptionsWithDefaults() *CollapseVODRequestFilterconfigScte35Options {
	this := CollapseVODRequestFilterconfigScte35Options{}
	return &this
}

// GetSignalRemovedProviderAd returns the SignalRemovedProviderAd field value
func (o *CollapseVODRequestFilterconfigScte35Options) GetSignalRemovedProviderAd() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SignalRemovedProviderAd
}

// GetSignalRemovedProviderAdOk returns a tuple with the SignalRemovedProviderAd field value
// and a boolean to check if the value has been set.
func (o *CollapseVODRequestFilterconfigScte35Options) GetSignalRemovedProviderAdOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SignalRemovedProviderAd, true
}

// SetSignalRemovedProviderAd sets field value
func (o *CollapseVODRequestFilterconfigScte35Options) SetSignalRemovedProviderAd(v bool) {
	o.SignalRemovedProviderAd = v
}

func (o CollapseVODRequestFilterconfigScte35Options) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollapseVODRequestFilterconfigScte35Options) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["signal_removed_provider_ad"] = o.SignalRemovedProviderAd
	return toSerialize, nil
}

type NullableCollapseVODRequestFilterconfigScte35Options struct {
	value *CollapseVODRequestFilterconfigScte35Options
	isSet bool
}

func (v NullableCollapseVODRequestFilterconfigScte35Options) Get() *CollapseVODRequestFilterconfigScte35Options {
	return v.value
}

func (v *NullableCollapseVODRequestFilterconfigScte35Options) Set(val *CollapseVODRequestFilterconfigScte35Options) {
	v.value = val
	v.isSet = true
}

func (v NullableCollapseVODRequestFilterconfigScte35Options) IsSet() bool {
	return v.isSet
}

func (v *NullableCollapseVODRequestFilterconfigScte35Options) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollapseVODRequestFilterconfigScte35Options(val *CollapseVODRequestFilterconfigScte35Options) *NullableCollapseVODRequestFilterconfigScte35Options {
	return &NullableCollapseVODRequestFilterconfigScte35Options{value: val, isSet: true}
}

func (v NullableCollapseVODRequestFilterconfigScte35Options) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollapseVODRequestFilterconfigScte35Options) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

