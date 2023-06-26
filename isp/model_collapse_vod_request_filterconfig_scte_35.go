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

// checks if the CollapseVODRequestFilterconfigScte35 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollapseVODRequestFilterconfigScte35{}

// CollapseVODRequestFilterconfigScte35 SCTE-based filtering
type CollapseVODRequestFilterconfigScte35 struct {
	// Filter out SCTE-35 break
	Break bool `json:"break"`
	// Filter out chapters
	Chapter bool `json:"chapter"`
	// Filter out distributor ads
	DistributorAds bool `json:"distributor_ads"`
	Options CollapseVODRequestFilterconfigScte35Options `json:"options"`
	// Filter out provider ads
	ProviderAds bool `json:"provider_ads"`
	// Filter out SCTE-35 message upid
	Upid string `json:"upid"`
}

// NewCollapseVODRequestFilterconfigScte35 instantiates a new CollapseVODRequestFilterconfigScte35 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollapseVODRequestFilterconfigScte35(break_ bool, chapter bool, distributorAds bool, options CollapseVODRequestFilterconfigScte35Options, providerAds bool, upid string) *CollapseVODRequestFilterconfigScte35 {
	this := CollapseVODRequestFilterconfigScte35{}
	this.Break = break_
	this.Chapter = chapter
	this.DistributorAds = distributorAds
	this.Options = options
	this.ProviderAds = providerAds
	this.Upid = upid
	return &this
}

// NewCollapseVODRequestFilterconfigScte35WithDefaults instantiates a new CollapseVODRequestFilterconfigScte35 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollapseVODRequestFilterconfigScte35WithDefaults() *CollapseVODRequestFilterconfigScte35 {
	this := CollapseVODRequestFilterconfigScte35{}
	return &this
}

// GetBreak returns the Break field value
func (o *CollapseVODRequestFilterconfigScte35) GetBreak() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Break
}

// GetBreakOk returns a tuple with the Break field value
// and a boolean to check if the value has been set.
func (o *CollapseVODRequestFilterconfigScte35) GetBreakOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Break, true
}

// SetBreak sets field value
func (o *CollapseVODRequestFilterconfigScte35) SetBreak(v bool) {
	o.Break = v
}

// GetChapter returns the Chapter field value
func (o *CollapseVODRequestFilterconfigScte35) GetChapter() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Chapter
}

// GetChapterOk returns a tuple with the Chapter field value
// and a boolean to check if the value has been set.
func (o *CollapseVODRequestFilterconfigScte35) GetChapterOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Chapter, true
}

// SetChapter sets field value
func (o *CollapseVODRequestFilterconfigScte35) SetChapter(v bool) {
	o.Chapter = v
}

// GetDistributorAds returns the DistributorAds field value
func (o *CollapseVODRequestFilterconfigScte35) GetDistributorAds() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.DistributorAds
}

// GetDistributorAdsOk returns a tuple with the DistributorAds field value
// and a boolean to check if the value has been set.
func (o *CollapseVODRequestFilterconfigScte35) GetDistributorAdsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DistributorAds, true
}

// SetDistributorAds sets field value
func (o *CollapseVODRequestFilterconfigScte35) SetDistributorAds(v bool) {
	o.DistributorAds = v
}

// GetOptions returns the Options field value
func (o *CollapseVODRequestFilterconfigScte35) GetOptions() CollapseVODRequestFilterconfigScte35Options {
	if o == nil {
		var ret CollapseVODRequestFilterconfigScte35Options
		return ret
	}

	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value
// and a boolean to check if the value has been set.
func (o *CollapseVODRequestFilterconfigScte35) GetOptionsOk() (*CollapseVODRequestFilterconfigScte35Options, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Options, true
}

// SetOptions sets field value
func (o *CollapseVODRequestFilterconfigScte35) SetOptions(v CollapseVODRequestFilterconfigScte35Options) {
	o.Options = v
}

// GetProviderAds returns the ProviderAds field value
func (o *CollapseVODRequestFilterconfigScte35) GetProviderAds() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ProviderAds
}

// GetProviderAdsOk returns a tuple with the ProviderAds field value
// and a boolean to check if the value has been set.
func (o *CollapseVODRequestFilterconfigScte35) GetProviderAdsOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderAds, true
}

// SetProviderAds sets field value
func (o *CollapseVODRequestFilterconfigScte35) SetProviderAds(v bool) {
	o.ProviderAds = v
}

// GetUpid returns the Upid field value
func (o *CollapseVODRequestFilterconfigScte35) GetUpid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Upid
}

// GetUpidOk returns a tuple with the Upid field value
// and a boolean to check if the value has been set.
func (o *CollapseVODRequestFilterconfigScte35) GetUpidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Upid, true
}

// SetUpid sets field value
func (o *CollapseVODRequestFilterconfigScte35) SetUpid(v string) {
	o.Upid = v
}

func (o CollapseVODRequestFilterconfigScte35) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollapseVODRequestFilterconfigScte35) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["break"] = o.Break
	toSerialize["chapter"] = o.Chapter
	toSerialize["distributor_ads"] = o.DistributorAds
	toSerialize["options"] = o.Options
	toSerialize["provider_ads"] = o.ProviderAds
	toSerialize["upid"] = o.Upid
	return toSerialize, nil
}

type NullableCollapseVODRequestFilterconfigScte35 struct {
	value *CollapseVODRequestFilterconfigScte35
	isSet bool
}

func (v NullableCollapseVODRequestFilterconfigScte35) Get() *CollapseVODRequestFilterconfigScte35 {
	return v.value
}

func (v *NullableCollapseVODRequestFilterconfigScte35) Set(val *CollapseVODRequestFilterconfigScte35) {
	v.value = val
	v.isSet = true
}

func (v NullableCollapseVODRequestFilterconfigScte35) IsSet() bool {
	return v.isSet
}

func (v *NullableCollapseVODRequestFilterconfigScte35) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollapseVODRequestFilterconfigScte35(val *CollapseVODRequestFilterconfigScte35) *NullableCollapseVODRequestFilterconfigScte35 {
	return &NullableCollapseVODRequestFilterconfigScte35{value: val, isSet: true}
}

func (v NullableCollapseVODRequestFilterconfigScte35) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollapseVODRequestFilterconfigScte35) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

