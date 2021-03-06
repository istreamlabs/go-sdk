/*
 * iStreamPlanet Channels API
 *
 * API version: 0.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package isp

import (
	"encoding/json"
)

// ChannelPublishingPublications struct for ChannelPublishingPublications
type ChannelPublishingPublications struct {
	Dash *ChannelPublishingDash `json:"dash,omitempty"`
	// Optionally specify which DRMs to advertise in the playlist. If specified, this must be a subset of the DRMs specified by the packager associated with this publication. If omitted or empty, all DRMs specified by the packager will be advertised.
	Drms *[]string `json:"drms,omitempty"`
	Hls *ChannelPublishingHls `json:"hls,omitempty"`
	// Determines how segments in this publication are packaged. Must reference a packager in 'packaging.packagers'. However, if this is a playlist-only publication (i.e. contains publish points that specify 'playlist_only_for'), this must remain unset as the packager will be inferred from the publication this one is providing playlists for.
	PackagerId *string `json:"packager_id,omitempty"`
	// Publish points specify where to output.
	PublishPoints *[]ChannelPublishingPublishPoints `json:"publish_points,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChannelPublishingPublications ChannelPublishingPublications

// NewChannelPublishingPublications instantiates a new ChannelPublishingPublications object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPublishingPublications() *ChannelPublishingPublications {
	this := ChannelPublishingPublications{}
	return &this
}

// NewChannelPublishingPublicationsWithDefaults instantiates a new ChannelPublishingPublications object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPublishingPublicationsWithDefaults() *ChannelPublishingPublications {
	this := ChannelPublishingPublications{}
	return &this
}

// GetDash returns the Dash field value if set, zero value otherwise.
func (o *ChannelPublishingPublications) GetDash() ChannelPublishingDash {
	if o == nil || o.Dash == nil {
		var ret ChannelPublishingDash
		return ret
	}
	return *o.Dash
}

// GetDashOk returns a tuple with the Dash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublications) GetDashOk() (*ChannelPublishingDash, bool) {
	if o == nil || o.Dash == nil {
		return nil, false
	}
	return o.Dash, true
}

// HasDash returns a boolean if a field has been set.
func (o *ChannelPublishingPublications) HasDash() bool {
	if o != nil && o.Dash != nil {
		return true
	}

	return false
}

// SetDash gets a reference to the given ChannelPublishingDash and assigns it to the Dash field.
func (o *ChannelPublishingPublications) SetDash(v ChannelPublishingDash) {
	o.Dash = &v
}

// GetDrms returns the Drms field value if set, zero value otherwise.
func (o *ChannelPublishingPublications) GetDrms() []string {
	if o == nil || o.Drms == nil {
		var ret []string
		return ret
	}
	return *o.Drms
}

// GetDrmsOk returns a tuple with the Drms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublications) GetDrmsOk() (*[]string, bool) {
	if o == nil || o.Drms == nil {
		return nil, false
	}
	return o.Drms, true
}

// HasDrms returns a boolean if a field has been set.
func (o *ChannelPublishingPublications) HasDrms() bool {
	if o != nil && o.Drms != nil {
		return true
	}

	return false
}

// SetDrms gets a reference to the given []string and assigns it to the Drms field.
func (o *ChannelPublishingPublications) SetDrms(v []string) {
	o.Drms = &v
}

// GetHls returns the Hls field value if set, zero value otherwise.
func (o *ChannelPublishingPublications) GetHls() ChannelPublishingHls {
	if o == nil || o.Hls == nil {
		var ret ChannelPublishingHls
		return ret
	}
	return *o.Hls
}

// GetHlsOk returns a tuple with the Hls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublications) GetHlsOk() (*ChannelPublishingHls, bool) {
	if o == nil || o.Hls == nil {
		return nil, false
	}
	return o.Hls, true
}

// HasHls returns a boolean if a field has been set.
func (o *ChannelPublishingPublications) HasHls() bool {
	if o != nil && o.Hls != nil {
		return true
	}

	return false
}

// SetHls gets a reference to the given ChannelPublishingHls and assigns it to the Hls field.
func (o *ChannelPublishingPublications) SetHls(v ChannelPublishingHls) {
	o.Hls = &v
}

// GetPackagerId returns the PackagerId field value if set, zero value otherwise.
func (o *ChannelPublishingPublications) GetPackagerId() string {
	if o == nil || o.PackagerId == nil {
		var ret string
		return ret
	}
	return *o.PackagerId
}

// GetPackagerIdOk returns a tuple with the PackagerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublications) GetPackagerIdOk() (*string, bool) {
	if o == nil || o.PackagerId == nil {
		return nil, false
	}
	return o.PackagerId, true
}

// HasPackagerId returns a boolean if a field has been set.
func (o *ChannelPublishingPublications) HasPackagerId() bool {
	if o != nil && o.PackagerId != nil {
		return true
	}

	return false
}

// SetPackagerId gets a reference to the given string and assigns it to the PackagerId field.
func (o *ChannelPublishingPublications) SetPackagerId(v string) {
	o.PackagerId = &v
}

// GetPublishPoints returns the PublishPoints field value if set, zero value otherwise.
func (o *ChannelPublishingPublications) GetPublishPoints() []ChannelPublishingPublishPoints {
	if o == nil || o.PublishPoints == nil {
		var ret []ChannelPublishingPublishPoints
		return ret
	}
	return *o.PublishPoints
}

// GetPublishPointsOk returns a tuple with the PublishPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublications) GetPublishPointsOk() (*[]ChannelPublishingPublishPoints, bool) {
	if o == nil || o.PublishPoints == nil {
		return nil, false
	}
	return o.PublishPoints, true
}

// HasPublishPoints returns a boolean if a field has been set.
func (o *ChannelPublishingPublications) HasPublishPoints() bool {
	if o != nil && o.PublishPoints != nil {
		return true
	}

	return false
}

// SetPublishPoints gets a reference to the given []ChannelPublishingPublishPoints and assigns it to the PublishPoints field.
func (o *ChannelPublishingPublications) SetPublishPoints(v []ChannelPublishingPublishPoints) {
	o.PublishPoints = &v
}

func (o ChannelPublishingPublications) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Dash != nil {
		toSerialize["dash"] = o.Dash
	}
	if o.Drms != nil {
		toSerialize["drms"] = o.Drms
	}
	if o.Hls != nil {
		toSerialize["hls"] = o.Hls
	}
	if o.PackagerId != nil {
		toSerialize["packager_id"] = o.PackagerId
	}
	if o.PublishPoints != nil {
		toSerialize["publish_points"] = o.PublishPoints
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChannelPublishingPublications) UnmarshalJSON(bytes []byte) (err error) {
	varChannelPublishingPublications := _ChannelPublishingPublications{}

	if err = json.Unmarshal(bytes, &varChannelPublishingPublications); err == nil {
		*o = ChannelPublishingPublications(varChannelPublishingPublications)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "dash")
		delete(additionalProperties, "drms")
		delete(additionalProperties, "hls")
		delete(additionalProperties, "packager_id")
		delete(additionalProperties, "publish_points")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChannelPublishingPublications struct {
	value *ChannelPublishingPublications
	isSet bool
}

func (v NullableChannelPublishingPublications) Get() *ChannelPublishingPublications {
	return v.value
}

func (v *NullableChannelPublishingPublications) Set(val *ChannelPublishingPublications) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPublishingPublications) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPublishingPublications) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPublishingPublications(val *ChannelPublishingPublications) *NullableChannelPublishingPublications {
	return &NullableChannelPublishingPublications{value: val, isSet: true}
}

func (v NullableChannelPublishingPublications) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPublishingPublications) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


