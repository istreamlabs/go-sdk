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

// checks if the ChannelPlaybackOriginInnerHls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelPlaybackOriginInnerHls{}

// ChannelPlaybackOriginInnerHls HLS can be an empty value if Aventus Origin does not have a valid HLS url for the packaged segments.
type ChannelPlaybackOriginInnerHls struct {
	// fallback_manifest is the name of the fallback playback manifest. ex. 'fallback.m3u8' or 'fallback.mpd'
	FallbackManifest *string `json:"fallback_manifest,omitempty" doc:"fallback_manifest is the name of the fallback playback manifest. ex. 'fallback.m3u8' or 'fallback.mpd'"`
	// fallback_url is the full playback url for the fallback. It is a composition of hostname, path and fallback_manifest fields. ex. 'https://foo.example.com/live/usw2/path/to/manifest/dir/fallback.m3u8'
	FallbackUrl *string `json:"fallback_url,omitempty" doc:"fallback_url is the full playback url for the fallback. It is a composition of hostname, path and fallback_manifest fields. ex. 'https://foo.example.com/live/usw2/path/to/manifest/dir/fallback.m3u8'"`
	// hostname for the playback url. ex. 'foo.example.com'
	Hostname *string `json:"hostname,omitempty" doc:"hostname for the playback url. ex. 'foo.example.com'"`
	// path identifies where manifests live at the host. ex. /live/usw2/path/to/manifest/dir
	Path *string `json:"path,omitempty" doc:"path identifies where manifests live at the host. ex. /live/usw2/path/to/manifest/dir"`
	// primary_manifest is the name of the primary playback manifest. ex. 'main.m3u8' or 'main.mpd'
	PrimaryManifest *string `json:"primary_manifest,omitempty" doc:"primary_manifest is the name of the primary playback manifest. ex. 'main.m3u8' or 'main.mpd'"`
	// primary_url is the full playback url for the primary. It is a composition of hostname, path and primary_manifest fields. ex. 'https://foo.example.com/live/usw2/path/to/manifest/dir/main.m3u8'
	PrimaryUrl *string `json:"primary_url,omitempty" doc:"primary_url is the full playback url for the primary. It is a composition of hostname, path and primary_manifest fields. ex. 'https://foo.example.com/live/usw2/path/to/manifest/dir/main.m3u8'"`
}

// NewChannelPlaybackOriginInnerHls instantiates a new ChannelPlaybackOriginInnerHls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPlaybackOriginInnerHls() *ChannelPlaybackOriginInnerHls {
	this := ChannelPlaybackOriginInnerHls{}
	return &this
}

// NewChannelPlaybackOriginInnerHlsWithDefaults instantiates a new ChannelPlaybackOriginInnerHls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPlaybackOriginInnerHlsWithDefaults() *ChannelPlaybackOriginInnerHls {
	this := ChannelPlaybackOriginInnerHls{}
	return &this
}

// GetFallbackManifest returns the FallbackManifest field value if set, zero value otherwise.
func (o *ChannelPlaybackOriginInnerHls) GetFallbackManifest() string {
	if o == nil || IsNil(o.FallbackManifest) {
		var ret string
		return ret
	}
	return *o.FallbackManifest
}

// GetFallbackManifestOk returns a tuple with the FallbackManifest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPlaybackOriginInnerHls) GetFallbackManifestOk() (*string, bool) {
	if o == nil || IsNil(o.FallbackManifest) {
		return nil, false
	}
	return o.FallbackManifest, true
}

// HasFallbackManifest returns a boolean if a field has been set.
func (o *ChannelPlaybackOriginInnerHls) HasFallbackManifest() bool {
	if o != nil && !IsNil(o.FallbackManifest) {
		return true
	}

	return false
}

// SetFallbackManifest gets a reference to the given string and assigns it to the FallbackManifest field.
func (o *ChannelPlaybackOriginInnerHls) SetFallbackManifest(v string) {
	o.FallbackManifest = &v
}

// GetFallbackUrl returns the FallbackUrl field value if set, zero value otherwise.
func (o *ChannelPlaybackOriginInnerHls) GetFallbackUrl() string {
	if o == nil || IsNil(o.FallbackUrl) {
		var ret string
		return ret
	}
	return *o.FallbackUrl
}

// GetFallbackUrlOk returns a tuple with the FallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPlaybackOriginInnerHls) GetFallbackUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FallbackUrl) {
		return nil, false
	}
	return o.FallbackUrl, true
}

// HasFallbackUrl returns a boolean if a field has been set.
func (o *ChannelPlaybackOriginInnerHls) HasFallbackUrl() bool {
	if o != nil && !IsNil(o.FallbackUrl) {
		return true
	}

	return false
}

// SetFallbackUrl gets a reference to the given string and assigns it to the FallbackUrl field.
func (o *ChannelPlaybackOriginInnerHls) SetFallbackUrl(v string) {
	o.FallbackUrl = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *ChannelPlaybackOriginInnerHls) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPlaybackOriginInnerHls) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *ChannelPlaybackOriginInnerHls) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *ChannelPlaybackOriginInnerHls) SetHostname(v string) {
	o.Hostname = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ChannelPlaybackOriginInnerHls) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPlaybackOriginInnerHls) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ChannelPlaybackOriginInnerHls) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ChannelPlaybackOriginInnerHls) SetPath(v string) {
	o.Path = &v
}

// GetPrimaryManifest returns the PrimaryManifest field value if set, zero value otherwise.
func (o *ChannelPlaybackOriginInnerHls) GetPrimaryManifest() string {
	if o == nil || IsNil(o.PrimaryManifest) {
		var ret string
		return ret
	}
	return *o.PrimaryManifest
}

// GetPrimaryManifestOk returns a tuple with the PrimaryManifest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPlaybackOriginInnerHls) GetPrimaryManifestOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryManifest) {
		return nil, false
	}
	return o.PrimaryManifest, true
}

// HasPrimaryManifest returns a boolean if a field has been set.
func (o *ChannelPlaybackOriginInnerHls) HasPrimaryManifest() bool {
	if o != nil && !IsNil(o.PrimaryManifest) {
		return true
	}

	return false
}

// SetPrimaryManifest gets a reference to the given string and assigns it to the PrimaryManifest field.
func (o *ChannelPlaybackOriginInnerHls) SetPrimaryManifest(v string) {
	o.PrimaryManifest = &v
}

// GetPrimaryUrl returns the PrimaryUrl field value if set, zero value otherwise.
func (o *ChannelPlaybackOriginInnerHls) GetPrimaryUrl() string {
	if o == nil || IsNil(o.PrimaryUrl) {
		var ret string
		return ret
	}
	return *o.PrimaryUrl
}

// GetPrimaryUrlOk returns a tuple with the PrimaryUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPlaybackOriginInnerHls) GetPrimaryUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryUrl) {
		return nil, false
	}
	return o.PrimaryUrl, true
}

// HasPrimaryUrl returns a boolean if a field has been set.
func (o *ChannelPlaybackOriginInnerHls) HasPrimaryUrl() bool {
	if o != nil && !IsNil(o.PrimaryUrl) {
		return true
	}

	return false
}

// SetPrimaryUrl gets a reference to the given string and assigns it to the PrimaryUrl field.
func (o *ChannelPlaybackOriginInnerHls) SetPrimaryUrl(v string) {
	o.PrimaryUrl = &v
}

func (o ChannelPlaybackOriginInnerHls) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelPlaybackOriginInnerHls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FallbackManifest) {
		toSerialize["fallback_manifest"] = o.FallbackManifest
	}
	if !IsNil(o.FallbackUrl) {
		toSerialize["fallback_url"] = o.FallbackUrl
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.PrimaryManifest) {
		toSerialize["primary_manifest"] = o.PrimaryManifest
	}
	if !IsNil(o.PrimaryUrl) {
		toSerialize["primary_url"] = o.PrimaryUrl
	}
	return toSerialize, nil
}

type NullableChannelPlaybackOriginInnerHls struct {
	value *ChannelPlaybackOriginInnerHls
	isSet bool
}

func (v NullableChannelPlaybackOriginInnerHls) Get() *ChannelPlaybackOriginInnerHls {
	return v.value
}

func (v *NullableChannelPlaybackOriginInnerHls) Set(val *ChannelPlaybackOriginInnerHls) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPlaybackOriginInnerHls) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPlaybackOriginInnerHls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPlaybackOriginInnerHls(val *ChannelPlaybackOriginInnerHls) *NullableChannelPlaybackOriginInnerHls {
	return &NullableChannelPlaybackOriginInnerHls{value: val, isSet: true}
}

func (v NullableChannelPlaybackOriginInnerHls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPlaybackOriginInnerHls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


