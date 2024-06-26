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

// checks if the Summary2Source type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Summary2Source{}

// Summary2Source Channel source
type Summary2Source struct {
	// Station source ID
	Id string `json:"id" doc:"Station source ID"`
	// Source name
	Name *string `json:"name,omitempty" doc:"Source name"`
	// Link to this resource
	Self *string `json:"self,omitempty" format:"uri" doc:"Link to this resource"`
}

// NewSummary2Source instantiates a new Summary2Source object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSummary2Source(id string) *Summary2Source {
	this := Summary2Source{}
	this.Id = id
	return &this
}

// NewSummary2SourceWithDefaults instantiates a new Summary2Source object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSummary2SourceWithDefaults() *Summary2Source {
	this := Summary2Source{}
	return &this
}

// GetId returns the Id field value
func (o *Summary2Source) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Summary2Source) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Summary2Source) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Summary2Source) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary2Source) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Summary2Source) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Summary2Source) SetName(v string) {
	o.Name = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *Summary2Source) GetSelf() string {
	if o == nil || IsNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary2Source) GetSelfOk() (*string, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *Summary2Source) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *Summary2Source) SetSelf(v string) {
	o.Self = &v
}

func (o Summary2Source) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Summary2Source) ToMap() (map[string]interface{}, error) {
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

type NullableSummary2Source struct {
	value *Summary2Source
	isSet bool
}

func (v NullableSummary2Source) Get() *Summary2Source {
	return v.value
}

func (v *NullableSummary2Source) Set(val *Summary2Source) {
	v.value = val
	v.isSet = true
}

func (v NullableSummary2Source) IsSet() bool {
	return v.isSet
}

func (v *NullableSummary2Source) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSummary2Source(val *Summary2Source) *NullableSummary2Source {
	return &NullableSummary2Source{value: val, isSet: true}
}

func (v NullableSummary2Source) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSummary2Source) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


