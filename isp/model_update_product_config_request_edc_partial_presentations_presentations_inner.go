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

// checks if the UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner{}

// UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner struct for UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner
type UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner struct {
	Name string `json:"name"`
	Renditions []int32 `json:"renditions" format:"int32"`
}

// NewUpdateProductConfigRequestEdcPartialPresentationsPresentationsInner instantiates a new UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProductConfigRequestEdcPartialPresentationsPresentationsInner(name string, renditions []int32) *UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner {
	this := UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner{}
	this.Name = name
	this.Renditions = renditions
	return &this
}

// NewUpdateProductConfigRequestEdcPartialPresentationsPresentationsInnerWithDefaults instantiates a new UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProductConfigRequestEdcPartialPresentationsPresentationsInnerWithDefaults() *UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner {
	this := UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner{}
	return &this
}

// GetName returns the Name field value
func (o *UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner) SetName(v string) {
	o.Name = v
}

// GetRenditions returns the Renditions field value
func (o *UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner) GetRenditions() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Renditions
}

// GetRenditionsOk returns a tuple with the Renditions field value
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner) GetRenditionsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Renditions, true
}

// SetRenditions sets field value
func (o *UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner) SetRenditions(v []int32) {
	o.Renditions = v
}

func (o UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["renditions"] = o.Renditions
	return toSerialize, nil
}

type NullableUpdateProductConfigRequestEdcPartialPresentationsPresentationsInner struct {
	value *UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner
	isSet bool
}

func (v NullableUpdateProductConfigRequestEdcPartialPresentationsPresentationsInner) Get() *UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner {
	return v.value
}

func (v *NullableUpdateProductConfigRequestEdcPartialPresentationsPresentationsInner) Set(val *UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProductConfigRequestEdcPartialPresentationsPresentationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProductConfigRequestEdcPartialPresentationsPresentationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProductConfigRequestEdcPartialPresentationsPresentationsInner(val *UpdateProductConfigRequestEdcPartialPresentationsPresentationsInner) *NullableUpdateProductConfigRequestEdcPartialPresentationsPresentationsInner {
	return &NullableUpdateProductConfigRequestEdcPartialPresentationsPresentationsInner{value: val, isSet: true}
}

func (v NullableUpdateProductConfigRequestEdcPartialPresentationsPresentationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProductConfigRequestEdcPartialPresentationsPresentationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


