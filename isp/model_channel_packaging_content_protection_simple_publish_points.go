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

// ChannelPackagingContentProtectionSimplePublishPoints struct for ChannelPackagingContentProtectionSimplePublishPoints
type ChannelPackagingContentProtectionSimplePublishPoints struct {
	BasicAuth *ChannelPackagingContentProtectionSimpleBasicAuth `json:"basic_auth,omitempty"`
	// Configures whether or not (and how) to compress manifests being published to the origin. If not specified, manifests will not be compressed.
	Compression *string `json:"compression,omitempty"`
	// Cross Playback Paths are playback paths that reference alternative content. These playback paths could reference publish points from the same publication or a completely different encoder and packager altogether. Content published to an endpoint referenced by one of these cross playback paths MUST be of the same Manifest.Type.
	CrossPlaybackPaths *[]string `json:"cross_playback_paths,omitempty"`
	// (Optional) Specifies if this pubpoint should not be monitored by PLM.
	DoNotMonitor *bool `json:"do_not_monitor,omitempty"`
	// Allows custom HTTP headers to be set via configuration for all HTTP requests.
	Headers *map[string]string `json:"headers,omitempty"`
	// Method overrides what HTTP method to specify in requests to the Publish Point. If not specified the service will default to POST.
	Method *string `json:"method,omitempty"`
	// The base URL where published playlists will be able to be obtained. This is often different than the publish_base_url for CDN publishing workflows.
	PlaybackBaseUrl *string `json:"playback_base_url,omitempty"`
	// Specifies any query parameters that will be added to playback urls. Should not include the initial '?' Example: 'foo=bar&q=golang'
	PlaybackQueryParams *string `json:"playback_query_params,omitempty"`
	// The base URL where generated playlists will be sent/published. Each publish point requires a unique 'publish_base_url'.
	PublishBaseUrl *string `json:"publish_base_url,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChannelPackagingContentProtectionSimplePublishPoints ChannelPackagingContentProtectionSimplePublishPoints

// NewChannelPackagingContentProtectionSimplePublishPoints instantiates a new ChannelPackagingContentProtectionSimplePublishPoints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPackagingContentProtectionSimplePublishPoints() *ChannelPackagingContentProtectionSimplePublishPoints {
	this := ChannelPackagingContentProtectionSimplePublishPoints{}
	return &this
}

// NewChannelPackagingContentProtectionSimplePublishPointsWithDefaults instantiates a new ChannelPackagingContentProtectionSimplePublishPoints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPackagingContentProtectionSimplePublishPointsWithDefaults() *ChannelPackagingContentProtectionSimplePublishPoints {
	this := ChannelPackagingContentProtectionSimplePublishPoints{}
	return &this
}

// GetBasicAuth returns the BasicAuth field value if set, zero value otherwise.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) GetBasicAuth() ChannelPackagingContentProtectionSimpleBasicAuth {
	if o == nil || o.BasicAuth == nil {
		var ret ChannelPackagingContentProtectionSimpleBasicAuth
		return ret
	}
	return *o.BasicAuth
}

// GetBasicAuthOk returns a tuple with the BasicAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) GetBasicAuthOk() (*ChannelPackagingContentProtectionSimpleBasicAuth, bool) {
	if o == nil || o.BasicAuth == nil {
		return nil, false
	}
	return o.BasicAuth, true
}

// HasBasicAuth returns a boolean if a field has been set.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) HasBasicAuth() bool {
	if o != nil && o.BasicAuth != nil {
		return true
	}

	return false
}

// SetBasicAuth gets a reference to the given ChannelPackagingContentProtectionSimpleBasicAuth and assigns it to the BasicAuth field.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) SetBasicAuth(v ChannelPackagingContentProtectionSimpleBasicAuth) {
	o.BasicAuth = &v
}

// GetCompression returns the Compression field value if set, zero value otherwise.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) GetCompression() string {
	if o == nil || o.Compression == nil {
		var ret string
		return ret
	}
	return *o.Compression
}

// GetCompressionOk returns a tuple with the Compression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) GetCompressionOk() (*string, bool) {
	if o == nil || o.Compression == nil {
		return nil, false
	}
	return o.Compression, true
}

// HasCompression returns a boolean if a field has been set.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) HasCompression() bool {
	if o != nil && o.Compression != nil {
		return true
	}

	return false
}

// SetCompression gets a reference to the given string and assigns it to the Compression field.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) SetCompression(v string) {
	o.Compression = &v
}

// GetCrossPlaybackPaths returns the CrossPlaybackPaths field value if set, zero value otherwise.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) GetCrossPlaybackPaths() []string {
	if o == nil || o.CrossPlaybackPaths == nil {
		var ret []string
		return ret
	}
	return *o.CrossPlaybackPaths
}

// GetCrossPlaybackPathsOk returns a tuple with the CrossPlaybackPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) GetCrossPlaybackPathsOk() (*[]string, bool) {
	if o == nil || o.CrossPlaybackPaths == nil {
		return nil, false
	}
	return o.CrossPlaybackPaths, true
}

// HasCrossPlaybackPaths returns a boolean if a field has been set.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) HasCrossPlaybackPaths() bool {
	if o != nil && o.CrossPlaybackPaths != nil {
		return true
	}

	return false
}

// SetCrossPlaybackPaths gets a reference to the given []string and assigns it to the CrossPlaybackPaths field.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) SetCrossPlaybackPaths(v []string) {
	o.CrossPlaybackPaths = &v
}

// GetDoNotMonitor returns the DoNotMonitor field value if set, zero value otherwise.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) GetDoNotMonitor() bool {
	if o == nil || o.DoNotMonitor == nil {
		var ret bool
		return ret
	}
	return *o.DoNotMonitor
}

// GetDoNotMonitorOk returns a tuple with the DoNotMonitor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) GetDoNotMonitorOk() (*bool, bool) {
	if o == nil || o.DoNotMonitor == nil {
		return nil, false
	}
	return o.DoNotMonitor, true
}

// HasDoNotMonitor returns a boolean if a field has been set.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) HasDoNotMonitor() bool {
	if o != nil && o.DoNotMonitor != nil {
		return true
	}

	return false
}

// SetDoNotMonitor gets a reference to the given bool and assigns it to the DoNotMonitor field.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) SetDoNotMonitor(v bool) {
	o.DoNotMonitor = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) GetHeaders() map[string]string {
	if o == nil || o.Headers == nil {
		var ret map[string]string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) SetHeaders(v map[string]string) {
	o.Headers = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) SetMethod(v string) {
	o.Method = &v
}

// GetPlaybackBaseUrl returns the PlaybackBaseUrl field value if set, zero value otherwise.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) GetPlaybackBaseUrl() string {
	if o == nil || o.PlaybackBaseUrl == nil {
		var ret string
		return ret
	}
	return *o.PlaybackBaseUrl
}

// GetPlaybackBaseUrlOk returns a tuple with the PlaybackBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) GetPlaybackBaseUrlOk() (*string, bool) {
	if o == nil || o.PlaybackBaseUrl == nil {
		return nil, false
	}
	return o.PlaybackBaseUrl, true
}

// HasPlaybackBaseUrl returns a boolean if a field has been set.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) HasPlaybackBaseUrl() bool {
	if o != nil && o.PlaybackBaseUrl != nil {
		return true
	}

	return false
}

// SetPlaybackBaseUrl gets a reference to the given string and assigns it to the PlaybackBaseUrl field.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) SetPlaybackBaseUrl(v string) {
	o.PlaybackBaseUrl = &v
}

// GetPlaybackQueryParams returns the PlaybackQueryParams field value if set, zero value otherwise.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) GetPlaybackQueryParams() string {
	if o == nil || o.PlaybackQueryParams == nil {
		var ret string
		return ret
	}
	return *o.PlaybackQueryParams
}

// GetPlaybackQueryParamsOk returns a tuple with the PlaybackQueryParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) GetPlaybackQueryParamsOk() (*string, bool) {
	if o == nil || o.PlaybackQueryParams == nil {
		return nil, false
	}
	return o.PlaybackQueryParams, true
}

// HasPlaybackQueryParams returns a boolean if a field has been set.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) HasPlaybackQueryParams() bool {
	if o != nil && o.PlaybackQueryParams != nil {
		return true
	}

	return false
}

// SetPlaybackQueryParams gets a reference to the given string and assigns it to the PlaybackQueryParams field.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) SetPlaybackQueryParams(v string) {
	o.PlaybackQueryParams = &v
}

// GetPublishBaseUrl returns the PublishBaseUrl field value if set, zero value otherwise.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) GetPublishBaseUrl() string {
	if o == nil || o.PublishBaseUrl == nil {
		var ret string
		return ret
	}
	return *o.PublishBaseUrl
}

// GetPublishBaseUrlOk returns a tuple with the PublishBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) GetPublishBaseUrlOk() (*string, bool) {
	if o == nil || o.PublishBaseUrl == nil {
		return nil, false
	}
	return o.PublishBaseUrl, true
}

// HasPublishBaseUrl returns a boolean if a field has been set.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) HasPublishBaseUrl() bool {
	if o != nil && o.PublishBaseUrl != nil {
		return true
	}

	return false
}

// SetPublishBaseUrl gets a reference to the given string and assigns it to the PublishBaseUrl field.
func (o *ChannelPackagingContentProtectionSimplePublishPoints) SetPublishBaseUrl(v string) {
	o.PublishBaseUrl = &v
}

func (o ChannelPackagingContentProtectionSimplePublishPoints) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BasicAuth != nil {
		toSerialize["basic_auth"] = o.BasicAuth
	}
	if o.Compression != nil {
		toSerialize["compression"] = o.Compression
	}
	if o.CrossPlaybackPaths != nil {
		toSerialize["cross_playback_paths"] = o.CrossPlaybackPaths
	}
	if o.DoNotMonitor != nil {
		toSerialize["do_not_monitor"] = o.DoNotMonitor
	}
	if o.Headers != nil {
		toSerialize["headers"] = o.Headers
	}
	if o.Method != nil {
		toSerialize["method"] = o.Method
	}
	if o.PlaybackBaseUrl != nil {
		toSerialize["playback_base_url"] = o.PlaybackBaseUrl
	}
	if o.PlaybackQueryParams != nil {
		toSerialize["playback_query_params"] = o.PlaybackQueryParams
	}
	if o.PublishBaseUrl != nil {
		toSerialize["publish_base_url"] = o.PublishBaseUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChannelPackagingContentProtectionSimplePublishPoints) UnmarshalJSON(bytes []byte) (err error) {
	varChannelPackagingContentProtectionSimplePublishPoints := _ChannelPackagingContentProtectionSimplePublishPoints{}

	if err = json.Unmarshal(bytes, &varChannelPackagingContentProtectionSimplePublishPoints); err == nil {
		*o = ChannelPackagingContentProtectionSimplePublishPoints(varChannelPackagingContentProtectionSimplePublishPoints)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "basic_auth")
		delete(additionalProperties, "compression")
		delete(additionalProperties, "cross_playback_paths")
		delete(additionalProperties, "do_not_monitor")
		delete(additionalProperties, "headers")
		delete(additionalProperties, "method")
		delete(additionalProperties, "playback_base_url")
		delete(additionalProperties, "playback_query_params")
		delete(additionalProperties, "publish_base_url")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChannelPackagingContentProtectionSimplePublishPoints struct {
	value *ChannelPackagingContentProtectionSimplePublishPoints
	isSet bool
}

func (v NullableChannelPackagingContentProtectionSimplePublishPoints) Get() *ChannelPackagingContentProtectionSimplePublishPoints {
	return v.value
}

func (v *NullableChannelPackagingContentProtectionSimplePublishPoints) Set(val *ChannelPackagingContentProtectionSimplePublishPoints) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPackagingContentProtectionSimplePublishPoints) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPackagingContentProtectionSimplePublishPoints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPackagingContentProtectionSimplePublishPoints(val *ChannelPackagingContentProtectionSimplePublishPoints) *NullableChannelPackagingContentProtectionSimplePublishPoints {
	return &NullableChannelPackagingContentProtectionSimplePublishPoints{value: val, isSet: true}
}

func (v NullableChannelPackagingContentProtectionSimplePublishPoints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPackagingContentProtectionSimplePublishPoints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


