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
	"time"
)

// checks if the GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner{}

// GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner struct for GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner
type GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner struct {
	Endtime *time.Time `json:"endtime,omitempty" format:"date-time"`
	Starttime time.Time `json:"starttime" format:"date-time"`
}

// NewGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner instantiates a new GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner(starttime time.Time) *GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner {
	this := GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner{}
	this.Starttime = starttime
	return &this
}

// NewGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInnerWithDefaults instantiates a new GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInnerWithDefaults() *GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner {
	this := GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner{}
	return &this
}

// GetEndtime returns the Endtime field value if set, zero value otherwise.
func (o *GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner) GetEndtime() time.Time {
	if o == nil || IsNil(o.Endtime) {
		var ret time.Time
		return ret
	}
	return *o.Endtime
}

// GetEndtimeOk returns a tuple with the Endtime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner) GetEndtimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Endtime) {
		return nil, false
	}
	return o.Endtime, true
}

// HasEndtime returns a boolean if a field has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner) HasEndtime() bool {
	if o != nil && !IsNil(o.Endtime) {
		return true
	}

	return false
}

// SetEndtime gets a reference to the given time.Time and assigns it to the Endtime field.
func (o *GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner) SetEndtime(v time.Time) {
	o.Endtime = &v
}

// GetStarttime returns the Starttime field value
func (o *GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner) GetStarttime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Starttime
}

// GetStarttimeOk returns a tuple with the Starttime field value
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner) GetStarttimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Starttime, true
}

// SetStarttime sets field value
func (o *GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner) SetStarttime(v time.Time) {
	o.Starttime = v
}

func (o GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Endtime) {
		toSerialize["endtime"] = o.Endtime
	}
	toSerialize["starttime"] = o.Starttime
	return toSerialize, nil
}

type NullableGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner struct {
	value *GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner
	isSet bool
}

func (v NullableGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner) Get() *GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner {
	return v.value
}

func (v *NullableGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner) Set(val *GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner(val *GetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner) *NullableGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner {
	return &NullableGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner{value: val, isSet: true}
}

func (v NullableGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProductConfigResponseProductConfigInnerCollapseConfigM3u8FiltersConfigLinksInnerTimespanInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

