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

// checks if the Summary2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Summary2{}

// Summary2 struct for Summary2
type Summary2 struct {
	// Desired state of channel
	DesiredState string `json:"desired_state" enum:"ON,OFF" doc:"Desired state of channel"`
	// Content hash
	Etag string `json:"etag" doc:"Content hash"`
	// Unique channel ID
	Id string `json:"id" doc:"Unique channel ID"`
	// Channel Labels
	Labels []string `json:"labels,omitempty" doc:"Channel Labels"`
	// Friendly channel description
	Name *string `json:"name,omitempty" doc:"Friendly channel description"`
	// Organization
	Org string `json:"org" doc:"Organization"`
	// Link to this resource
	Self   *string        `json:"self,omitempty" format:"uri" doc:"Link to this resource"`
	Source Summary2Source `json:"source"`
}

// NewSummary2 instantiates a new Summary2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSummary2(desiredState string, etag string, id string, org string, source Summary2Source) *Summary2 {
	this := Summary2{}
	this.DesiredState = desiredState
	this.Etag = etag
	this.Id = id
	this.Org = org
	this.Source = source
	return &this
}

// NewSummary2WithDefaults instantiates a new Summary2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSummary2WithDefaults() *Summary2 {
	this := Summary2{}
	return &this
}

// GetDesiredState returns the DesiredState field value
func (o *Summary2) GetDesiredState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DesiredState
}

// GetDesiredStateOk returns a tuple with the DesiredState field value
// and a boolean to check if the value has been set.
func (o *Summary2) GetDesiredStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DesiredState, true
}

// SetDesiredState sets field value
func (o *Summary2) SetDesiredState(v string) {
	o.DesiredState = v
}

// GetEtag returns the Etag field value
func (o *Summary2) GetEtag() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Etag
}

// GetEtagOk returns a tuple with the Etag field value
// and a boolean to check if the value has been set.
func (o *Summary2) GetEtagOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Etag, true
}

// SetEtag sets field value
func (o *Summary2) SetEtag(v string) {
	o.Etag = v
}

// GetId returns the Id field value
func (o *Summary2) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Summary2) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Summary2) SetId(v string) {
	o.Id = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *Summary2) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary2) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *Summary2) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *Summary2) SetLabels(v []string) {
	o.Labels = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Summary2) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary2) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Summary2) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Summary2) SetName(v string) {
	o.Name = &v
}

// GetOrg returns the Org field value
func (o *Summary2) GetOrg() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Org
}

// GetOrgOk returns a tuple with the Org field value
// and a boolean to check if the value has been set.
func (o *Summary2) GetOrgOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Org, true
}

// SetOrg sets field value
func (o *Summary2) SetOrg(v string) {
	o.Org = v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *Summary2) GetSelf() string {
	if o == nil || IsNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Summary2) GetSelfOk() (*string, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *Summary2) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *Summary2) SetSelf(v string) {
	o.Self = &v
}

// GetSource returns the Source field value
func (o *Summary2) GetSource() Summary2Source {
	if o == nil {
		var ret Summary2Source
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *Summary2) GetSourceOk() (*Summary2Source, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *Summary2) SetSource(v Summary2Source) {
	o.Source = v
}

func (o Summary2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Summary2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["desired_state"] = o.DesiredState
	toSerialize["etag"] = o.Etag
	toSerialize["id"] = o.Id
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["org"] = o.Org
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	toSerialize["source"] = o.Source
	return toSerialize, nil
}

type NullableSummary2 struct {
	value *Summary2
	isSet bool
}

func (v NullableSummary2) Get() *Summary2 {
	return v.value
}

func (v *NullableSummary2) Set(val *Summary2) {
	v.value = val
	v.isSet = true
}

func (v NullableSummary2) IsSet() bool {
	return v.isSet
}

func (v *NullableSummary2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSummary2(val *Summary2) *NullableSummary2 {
	return &NullableSummary2{value: val, isSet: true}
}

func (v NullableSummary2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSummary2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
