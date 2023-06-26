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

// checks if the UpdateProductConfigRequestEdcPartialPresentations type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProductConfigRequestEdcPartialPresentations{}

// UpdateProductConfigRequestEdcPartialPresentations Partial presentation definitions
type UpdateProductConfigRequestEdcPartialPresentations struct {
	Presentations []UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner `json:"presentations"`
}

// NewUpdateProductConfigRequestEdcPartialPresentations instantiates a new UpdateProductConfigRequestEdcPartialPresentations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProductConfigRequestEdcPartialPresentations(presentations []UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner) *UpdateProductConfigRequestEdcPartialPresentations {
	this := UpdateProductConfigRequestEdcPartialPresentations{}
	this.Presentations = presentations
	return &this
}

// NewUpdateProductConfigRequestEdcPartialPresentationsWithDefaults instantiates a new UpdateProductConfigRequestEdcPartialPresentations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProductConfigRequestEdcPartialPresentationsWithDefaults() *UpdateProductConfigRequestEdcPartialPresentations {
	this := UpdateProductConfigRequestEdcPartialPresentations{}
	return &this
}

// GetPresentations returns the Presentations field value
func (o *UpdateProductConfigRequestEdcPartialPresentations) GetPresentations() []UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner {
	if o == nil {
		var ret []UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner
		return ret
	}

	return o.Presentations
}

// GetPresentationsOk returns a tuple with the Presentations field value
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestEdcPartialPresentations) GetPresentationsOk() ([]UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Presentations, true
}

// SetPresentations sets field value
func (o *UpdateProductConfigRequestEdcPartialPresentations) SetPresentations(v []UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner) {
	o.Presentations = v
}

func (o UpdateProductConfigRequestEdcPartialPresentations) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProductConfigRequestEdcPartialPresentations) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["presentations"] = o.Presentations
	return toSerialize, nil
}

type NullableUpdateProductConfigRequestEdcPartialPresentations struct {
	value *UpdateProductConfigRequestEdcPartialPresentations
	isSet bool
}

func (v NullableUpdateProductConfigRequestEdcPartialPresentations) Get() *UpdateProductConfigRequestEdcPartialPresentations {
	return v.value
}

func (v *NullableUpdateProductConfigRequestEdcPartialPresentations) Set(val *UpdateProductConfigRequestEdcPartialPresentations) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProductConfigRequestEdcPartialPresentations) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProductConfigRequestEdcPartialPresentations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProductConfigRequestEdcPartialPresentations(val *UpdateProductConfigRequestEdcPartialPresentations) *NullableUpdateProductConfigRequestEdcPartialPresentations {
	return &NullableUpdateProductConfigRequestEdcPartialPresentations{value: val, isSet: true}
}

func (v NullableUpdateProductConfigRequestEdcPartialPresentations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProductConfigRequestEdcPartialPresentations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

