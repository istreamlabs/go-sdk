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

// checks if the Summary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Summary{}

// Summary struct for Summary
type Summary struct {
	// Station source ID
	Id string `json:"id"`
	// Source name
	Name *string `json:"name,omitempty"`
	// Link to this resource
	Self *string `json:"self,omitempty"`
}

// NewSummary instantiates a new Summary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSummary(id string) *Summary {
	this := Summary{}
	this.Id = id
	return &this
}

// NewSummaryWithDefaults instantiates a new Summary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSummaryWithDefaults() *Summary {
	this := Summary{}
	return &this
}

// GetId returns the Id field value
func (o *Summary) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Summary) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Summary) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Summary) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Summary) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Summary) SetName(v string) {
	o.Name = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *Summary) GetSelf() string {
	if o == nil || IsNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary) GetSelfOk() (*string, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *Summary) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *Summary) SetSelf(v string) {
	o.Self = &v
}

func (o Summary) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Summary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	return toSerialize, nil
}

type NullableSummary struct {
	value *Summary
	isSet bool
}

func (v NullableSummary) Get() *Summary {
	return v.value
}

func (v *NullableSummary) Set(val *Summary) {
	v.value = val
	v.isSet = true
}

func (v NullableSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSummary(val *Summary) *NullableSummary {
	return &NullableSummary{value: val, isSet: true}
}

func (v NullableSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


