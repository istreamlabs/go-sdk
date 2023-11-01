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

// checks if the GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner{}

// GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner struct for GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner
type GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner struct {
	Scte35Upid *string `json:"scte35_upid,omitempty"`
	SegmentationEventId *int32 `json:"segmentation_event_id,omitempty" format:"int32" minimum:"0"`
	SegmentationTypeId int32 `json:"segmentation_type_id" format:"int32"`
}

// NewGetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner instantiates a new GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner(segmentationTypeId int32) *GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner {
	this := GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner{}
	this.SegmentationTypeId = segmentationTypeId
	return &this
}

// NewGetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInnerWithDefaults instantiates a new GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInnerWithDefaults() *GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner {
	this := GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner{}
	return &this
}

// GetScte35Upid returns the Scte35Upid field value if set, zero value otherwise.
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner) GetScte35Upid() string {
	if o == nil || IsNil(o.Scte35Upid) {
		var ret string
		return ret
	}
	return *o.Scte35Upid
}

// GetScte35UpidOk returns a tuple with the Scte35Upid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner) GetScte35UpidOk() (*string, bool) {
	if o == nil || IsNil(o.Scte35Upid) {
		return nil, false
	}
	return o.Scte35Upid, true
}

// HasScte35Upid returns a boolean if a field has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner) HasScte35Upid() bool {
	if o != nil && !IsNil(o.Scte35Upid) {
		return true
	}

	return false
}

// SetScte35Upid gets a reference to the given string and assigns it to the Scte35Upid field.
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner) SetScte35Upid(v string) {
	o.Scte35Upid = &v
}

// GetSegmentationEventId returns the SegmentationEventId field value if set, zero value otherwise.
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner) GetSegmentationEventId() int32 {
	if o == nil || IsNil(o.SegmentationEventId) {
		var ret int32
		return ret
	}
	return *o.SegmentationEventId
}

// GetSegmentationEventIdOk returns a tuple with the SegmentationEventId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner) GetSegmentationEventIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SegmentationEventId) {
		return nil, false
	}
	return o.SegmentationEventId, true
}

// HasSegmentationEventId returns a boolean if a field has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner) HasSegmentationEventId() bool {
	if o != nil && !IsNil(o.SegmentationEventId) {
		return true
	}

	return false
}

// SetSegmentationEventId gets a reference to the given int32 and assigns it to the SegmentationEventId field.
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner) SetSegmentationEventId(v int32) {
	o.SegmentationEventId = &v
}

// GetSegmentationTypeId returns the SegmentationTypeId field value
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner) GetSegmentationTypeId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SegmentationTypeId
}

// GetSegmentationTypeIdOk returns a tuple with the SegmentationTypeId field value
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner) GetSegmentationTypeIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SegmentationTypeId, true
}

// SetSegmentationTypeId sets field value
func (o *GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner) SetSegmentationTypeId(v int32) {
	o.SegmentationTypeId = v
}

func (o GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Scte35Upid) {
		toSerialize["scte35_upid"] = o.Scte35Upid
	}
	if !IsNil(o.SegmentationEventId) {
		toSerialize["segmentation_event_id"] = o.SegmentationEventId
	}
	toSerialize["segmentation_type_id"] = o.SegmentationTypeId
	return toSerialize, nil
}

type NullableGetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner struct {
	value *GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner
	isSet bool
}

func (v NullableGetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner) Get() *GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner {
	return v.value
}

func (v *NullableGetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner) Set(val *GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner(val *GetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner) *NullableGetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner {
	return &NullableGetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner{value: val, isSet: true}
}

func (v NullableGetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProductConfigResponseProductConfigInnerCollapseTriggerConfigTriggersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


