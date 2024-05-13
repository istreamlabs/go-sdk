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

// checks if the PatchChannelRequest2Inner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchChannelRequest2Inner{}

// PatchChannelRequest2Inner struct for PatchChannelRequest2Inner
type PatchChannelRequest2Inner struct {
	// JSON Pointer for the source of a move or copy
	From *string `json:"from,omitempty" doc:"JSON Pointer for the source of a move or copy"`
	// Operation name
	Op string `json:"op" enum:"add,remove,replace,move,copy,test" doc:"Operation name"`
	// JSON Pointer to the field being operated on, or the destination of a move/copy operation
	Path string `json:"path" doc:"JSON Pointer to the field being operated on, or the destination of a move/copy operation"`
	// The value to set
	Value interface{} `json:"value,omitempty" doc:"The value to set"`
}

// NewPatchChannelRequest2Inner instantiates a new PatchChannelRequest2Inner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchChannelRequest2Inner(op string, path string) *PatchChannelRequest2Inner {
	this := PatchChannelRequest2Inner{}
	this.Op = op
	this.Path = path
	return &this
}

// NewPatchChannelRequest2InnerWithDefaults instantiates a new PatchChannelRequest2Inner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchChannelRequest2InnerWithDefaults() *PatchChannelRequest2Inner {
	this := PatchChannelRequest2Inner{}
	return &this
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *PatchChannelRequest2Inner) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchChannelRequest2Inner) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *PatchChannelRequest2Inner) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *PatchChannelRequest2Inner) SetFrom(v string) {
	o.From = &v
}

// GetOp returns the Op field value
func (o *PatchChannelRequest2Inner) GetOp() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Op
}

// GetOpOk returns a tuple with the Op field value
// and a boolean to check if the value has been set.
func (o *PatchChannelRequest2Inner) GetOpOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Op, true
}

// SetOp sets field value
func (o *PatchChannelRequest2Inner) SetOp(v string) {
	o.Op = v
}

// GetPath returns the Path field value
func (o *PatchChannelRequest2Inner) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *PatchChannelRequest2Inner) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *PatchChannelRequest2Inner) SetPath(v string) {
	o.Path = v
}

// GetValue returns the Value field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PatchChannelRequest2Inner) GetValue() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PatchChannelRequest2Inner) GetValueOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return &o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PatchChannelRequest2Inner) HasValue() bool {
	if o != nil && IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given interface{} and assigns it to the Value field.
func (o *PatchChannelRequest2Inner) SetValue(v interface{}) {
	o.Value = v
}

func (o PatchChannelRequest2Inner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchChannelRequest2Inner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	toSerialize["op"] = o.Op
	toSerialize["path"] = o.Path
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullablePatchChannelRequest2Inner struct {
	value *PatchChannelRequest2Inner
	isSet bool
}

func (v NullablePatchChannelRequest2Inner) Get() *PatchChannelRequest2Inner {
	return v.value
}

func (v *NullablePatchChannelRequest2Inner) Set(val *PatchChannelRequest2Inner) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchChannelRequest2Inner) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchChannelRequest2Inner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchChannelRequest2Inner(val *PatchChannelRequest2Inner) *NullablePatchChannelRequest2Inner {
	return &NullablePatchChannelRequest2Inner{value: val, isSet: true}
}

func (v NullablePatchChannelRequest2Inner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchChannelRequest2Inner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

