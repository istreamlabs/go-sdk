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

// checks if the GetProductConfigResponseProductConfigInnerEdcPartialPresentations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProductConfigResponseProductConfigInnerEdcPartialPresentations{}

// GetProductConfigResponseProductConfigInnerEdcPartialPresentations Partial presentation definitions
type GetProductConfigResponseProductConfigInnerEdcPartialPresentations struct {
	Presentations []GetProductConfigResponseProductConfigInnerEdcPartialPresentationsPresentationsInner `json:"presentations"`
}

// NewGetProductConfigResponseProductConfigInnerEdcPartialPresentations instantiates a new GetProductConfigResponseProductConfigInnerEdcPartialPresentations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductConfigResponseProductConfigInnerEdcPartialPresentations(presentations []GetProductConfigResponseProductConfigInnerEdcPartialPresentationsPresentationsInner) *GetProductConfigResponseProductConfigInnerEdcPartialPresentations {
	this := GetProductConfigResponseProductConfigInnerEdcPartialPresentations{}
	this.Presentations = presentations
	return &this
}

// NewGetProductConfigResponseProductConfigInnerEdcPartialPresentationsWithDefaults instantiates a new GetProductConfigResponseProductConfigInnerEdcPartialPresentations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductConfigResponseProductConfigInnerEdcPartialPresentationsWithDefaults() *GetProductConfigResponseProductConfigInnerEdcPartialPresentations {
	this := GetProductConfigResponseProductConfigInnerEdcPartialPresentations{}
	return &this
}

// GetPresentations returns the Presentations field value
func (o *GetProductConfigResponseProductConfigInnerEdcPartialPresentations) GetPresentations() []GetProductConfigResponseProductConfigInnerEdcPartialPresentationsPresentationsInner {
	if o == nil {
		var ret []GetProductConfigResponseProductConfigInnerEdcPartialPresentationsPresentationsInner
		return ret
	}

	return o.Presentations
}

// GetPresentationsOk returns a tuple with the Presentations field value
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerEdcPartialPresentations) GetPresentationsOk() ([]GetProductConfigResponseProductConfigInnerEdcPartialPresentationsPresentationsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Presentations, true
}

// SetPresentations sets field value
func (o *GetProductConfigResponseProductConfigInnerEdcPartialPresentations) SetPresentations(v []GetProductConfigResponseProductConfigInnerEdcPartialPresentationsPresentationsInner) {
	o.Presentations = v
}

func (o GetProductConfigResponseProductConfigInnerEdcPartialPresentations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProductConfigResponseProductConfigInnerEdcPartialPresentations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["presentations"] = o.Presentations
	return toSerialize, nil
}

type NullableGetProductConfigResponseProductConfigInnerEdcPartialPresentations struct {
	value *GetProductConfigResponseProductConfigInnerEdcPartialPresentations
	isSet bool
}

func (v NullableGetProductConfigResponseProductConfigInnerEdcPartialPresentations) Get() *GetProductConfigResponseProductConfigInnerEdcPartialPresentations {
	return v.value
}

func (v *NullableGetProductConfigResponseProductConfigInnerEdcPartialPresentations) Set(val *GetProductConfigResponseProductConfigInnerEdcPartialPresentations) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductConfigResponseProductConfigInnerEdcPartialPresentations) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductConfigResponseProductConfigInnerEdcPartialPresentations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductConfigResponseProductConfigInnerEdcPartialPresentations(val *GetProductConfigResponseProductConfigInnerEdcPartialPresentations) *NullableGetProductConfigResponseProductConfigInnerEdcPartialPresentations {
	return &NullableGetProductConfigResponseProductConfigInnerEdcPartialPresentations{value: val, isSet: true}
}

func (v NullableGetProductConfigResponseProductConfigInnerEdcPartialPresentations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProductConfigResponseProductConfigInnerEdcPartialPresentations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

