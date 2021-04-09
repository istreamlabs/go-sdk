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

// ChannelPublishingHttp HTTP destination where media segments and playlists will be published.
type ChannelPublishingHttp struct {
	// HTTP Basic Authentication.
	BasicAuth *map[string]interface{} `json:"basic_auth,omitempty"`
	// Configures whether or not (and how) to compress manifests being published to the origin. If not specified, manifests will not be compressed.
	Compression *string `json:"compression,omitempty"`
	// Cross Playback Paths are playback paths that reference alternative content. These playback paths could reference publish points from the same publication or a completely different encoder and packager altogether. Content published to an endpoint referenced by one of these cross playback paths MUST be of the same Manifest.Type.
	CrossPlaybackPaths *[]string `json:"cross_playback_paths,omitempty"`
	// DNS server IP address that should be used for DNS queries instead of the default DNS server (optional).
	DnsServerAddress *string `json:"dns_server_address,omitempty"`
	// Allows custom HTTP headers to be set via configuration for all HTTP requests.
	Headers *map[string]string `json:"headers,omitempty"`
	// Method overrides what HTTP method to specify in requests to the Publish Point. If not specified the service will default to POST.
	Method *string `json:"method,omitempty"`
	// The base URL where published playlists will be able to be obtained. This is often different than the publish_base_url for CDN publishing workflows.
	PlaybackBaseUrl *string `json:"playback_base_url,omitempty"`
	// Specifies any arguments that will be added to playback url. Will be used by PLM and players.
	PlaybackQueryParams *string `json:"playback_query_params,omitempty"`
	// The base URL where generated playlists will be sent/published.
	PublishBaseUrl *string `json:"publish_base_url,omitempty"`
	// Deprecated: Do not use. Whether or not to gzip manifests being published to the origin. Deprecated--use 'compression' instead.
	UseGzipCompression *bool `json:"use_gzip_compression,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChannelPublishingHttp ChannelPublishingHttp

// NewChannelPublishingHttp instantiates a new ChannelPublishingHttp object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPublishingHttp() *ChannelPublishingHttp {
	this := ChannelPublishingHttp{}
	return &this
}

// NewChannelPublishingHttpWithDefaults instantiates a new ChannelPublishingHttp object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPublishingHttpWithDefaults() *ChannelPublishingHttp {
	this := ChannelPublishingHttp{}
	return &this
}

// GetBasicAuth returns the BasicAuth field value if set, zero value otherwise.
func (o *ChannelPublishingHttp) GetBasicAuth() map[string]interface{} {
	if o == nil || o.BasicAuth == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.BasicAuth
}

// GetBasicAuthOk returns a tuple with the BasicAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHttp) GetBasicAuthOk() (*map[string]interface{}, bool) {
	if o == nil || o.BasicAuth == nil {
		return nil, false
	}
	return o.BasicAuth, true
}

// HasBasicAuth returns a boolean if a field has been set.
func (o *ChannelPublishingHttp) HasBasicAuth() bool {
	if o != nil && o.BasicAuth != nil {
		return true
	}

	return false
}

// SetBasicAuth gets a reference to the given map[string]interface{} and assigns it to the BasicAuth field.
func (o *ChannelPublishingHttp) SetBasicAuth(v map[string]interface{}) {
	o.BasicAuth = &v
}

// GetCompression returns the Compression field value if set, zero value otherwise.
func (o *ChannelPublishingHttp) GetCompression() string {
	if o == nil || o.Compression == nil {
		var ret string
		return ret
	}
	return *o.Compression
}

// GetCompressionOk returns a tuple with the Compression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHttp) GetCompressionOk() (*string, bool) {
	if o == nil || o.Compression == nil {
		return nil, false
	}
	return o.Compression, true
}

// HasCompression returns a boolean if a field has been set.
func (o *ChannelPublishingHttp) HasCompression() bool {
	if o != nil && o.Compression != nil {
		return true
	}

	return false
}

// SetCompression gets a reference to the given string and assigns it to the Compression field.
func (o *ChannelPublishingHttp) SetCompression(v string) {
	o.Compression = &v
}

// GetCrossPlaybackPaths returns the CrossPlaybackPaths field value if set, zero value otherwise.
func (o *ChannelPublishingHttp) GetCrossPlaybackPaths() []string {
	if o == nil || o.CrossPlaybackPaths == nil {
		var ret []string
		return ret
	}
	return *o.CrossPlaybackPaths
}

// GetCrossPlaybackPathsOk returns a tuple with the CrossPlaybackPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHttp) GetCrossPlaybackPathsOk() (*[]string, bool) {
	if o == nil || o.CrossPlaybackPaths == nil {
		return nil, false
	}
	return o.CrossPlaybackPaths, true
}

// HasCrossPlaybackPaths returns a boolean if a field has been set.
func (o *ChannelPublishingHttp) HasCrossPlaybackPaths() bool {
	if o != nil && o.CrossPlaybackPaths != nil {
		return true
	}

	return false
}

// SetCrossPlaybackPaths gets a reference to the given []string and assigns it to the CrossPlaybackPaths field.
func (o *ChannelPublishingHttp) SetCrossPlaybackPaths(v []string) {
	o.CrossPlaybackPaths = &v
}

// GetDnsServerAddress returns the DnsServerAddress field value if set, zero value otherwise.
func (o *ChannelPublishingHttp) GetDnsServerAddress() string {
	if o == nil || o.DnsServerAddress == nil {
		var ret string
		return ret
	}
	return *o.DnsServerAddress
}

// GetDnsServerAddressOk returns a tuple with the DnsServerAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHttp) GetDnsServerAddressOk() (*string, bool) {
	if o == nil || o.DnsServerAddress == nil {
		return nil, false
	}
	return o.DnsServerAddress, true
}

// HasDnsServerAddress returns a boolean if a field has been set.
func (o *ChannelPublishingHttp) HasDnsServerAddress() bool {
	if o != nil && o.DnsServerAddress != nil {
		return true
	}

	return false
}

// SetDnsServerAddress gets a reference to the given string and assigns it to the DnsServerAddress field.
func (o *ChannelPublishingHttp) SetDnsServerAddress(v string) {
	o.DnsServerAddress = &v
}

// GetHeaders returns the Headers field value if set, zero value otherwise.
func (o *ChannelPublishingHttp) GetHeaders() map[string]string {
	if o == nil || o.Headers == nil {
		var ret map[string]string
		return ret
	}
	return *o.Headers
}

// GetHeadersOk returns a tuple with the Headers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHttp) GetHeadersOk() (*map[string]string, bool) {
	if o == nil || o.Headers == nil {
		return nil, false
	}
	return o.Headers, true
}

// HasHeaders returns a boolean if a field has been set.
func (o *ChannelPublishingHttp) HasHeaders() bool {
	if o != nil && o.Headers != nil {
		return true
	}

	return false
}

// SetHeaders gets a reference to the given map[string]string and assigns it to the Headers field.
func (o *ChannelPublishingHttp) SetHeaders(v map[string]string) {
	o.Headers = &v
}

// GetMethod returns the Method field value if set, zero value otherwise.
func (o *ChannelPublishingHttp) GetMethod() string {
	if o == nil || o.Method == nil {
		var ret string
		return ret
	}
	return *o.Method
}

// GetMethodOk returns a tuple with the Method field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHttp) GetMethodOk() (*string, bool) {
	if o == nil || o.Method == nil {
		return nil, false
	}
	return o.Method, true
}

// HasMethod returns a boolean if a field has been set.
func (o *ChannelPublishingHttp) HasMethod() bool {
	if o != nil && o.Method != nil {
		return true
	}

	return false
}

// SetMethod gets a reference to the given string and assigns it to the Method field.
func (o *ChannelPublishingHttp) SetMethod(v string) {
	o.Method = &v
}

// GetPlaybackBaseUrl returns the PlaybackBaseUrl field value if set, zero value otherwise.
func (o *ChannelPublishingHttp) GetPlaybackBaseUrl() string {
	if o == nil || o.PlaybackBaseUrl == nil {
		var ret string
		return ret
	}
	return *o.PlaybackBaseUrl
}

// GetPlaybackBaseUrlOk returns a tuple with the PlaybackBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHttp) GetPlaybackBaseUrlOk() (*string, bool) {
	if o == nil || o.PlaybackBaseUrl == nil {
		return nil, false
	}
	return o.PlaybackBaseUrl, true
}

// HasPlaybackBaseUrl returns a boolean if a field has been set.
func (o *ChannelPublishingHttp) HasPlaybackBaseUrl() bool {
	if o != nil && o.PlaybackBaseUrl != nil {
		return true
	}

	return false
}

// SetPlaybackBaseUrl gets a reference to the given string and assigns it to the PlaybackBaseUrl field.
func (o *ChannelPublishingHttp) SetPlaybackBaseUrl(v string) {
	o.PlaybackBaseUrl = &v
}

// GetPlaybackQueryParams returns the PlaybackQueryParams field value if set, zero value otherwise.
func (o *ChannelPublishingHttp) GetPlaybackQueryParams() string {
	if o == nil || o.PlaybackQueryParams == nil {
		var ret string
		return ret
	}
	return *o.PlaybackQueryParams
}

// GetPlaybackQueryParamsOk returns a tuple with the PlaybackQueryParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHttp) GetPlaybackQueryParamsOk() (*string, bool) {
	if o == nil || o.PlaybackQueryParams == nil {
		return nil, false
	}
	return o.PlaybackQueryParams, true
}

// HasPlaybackQueryParams returns a boolean if a field has been set.
func (o *ChannelPublishingHttp) HasPlaybackQueryParams() bool {
	if o != nil && o.PlaybackQueryParams != nil {
		return true
	}

	return false
}

// SetPlaybackQueryParams gets a reference to the given string and assigns it to the PlaybackQueryParams field.
func (o *ChannelPublishingHttp) SetPlaybackQueryParams(v string) {
	o.PlaybackQueryParams = &v
}

// GetPublishBaseUrl returns the PublishBaseUrl field value if set, zero value otherwise.
func (o *ChannelPublishingHttp) GetPublishBaseUrl() string {
	if o == nil || o.PublishBaseUrl == nil {
		var ret string
		return ret
	}
	return *o.PublishBaseUrl
}

// GetPublishBaseUrlOk returns a tuple with the PublishBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHttp) GetPublishBaseUrlOk() (*string, bool) {
	if o == nil || o.PublishBaseUrl == nil {
		return nil, false
	}
	return o.PublishBaseUrl, true
}

// HasPublishBaseUrl returns a boolean if a field has been set.
func (o *ChannelPublishingHttp) HasPublishBaseUrl() bool {
	if o != nil && o.PublishBaseUrl != nil {
		return true
	}

	return false
}

// SetPublishBaseUrl gets a reference to the given string and assigns it to the PublishBaseUrl field.
func (o *ChannelPublishingHttp) SetPublishBaseUrl(v string) {
	o.PublishBaseUrl = &v
}

// GetUseGzipCompression returns the UseGzipCompression field value if set, zero value otherwise.
func (o *ChannelPublishingHttp) GetUseGzipCompression() bool {
	if o == nil || o.UseGzipCompression == nil {
		var ret bool
		return ret
	}
	return *o.UseGzipCompression
}

// GetUseGzipCompressionOk returns a tuple with the UseGzipCompression field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingHttp) GetUseGzipCompressionOk() (*bool, bool) {
	if o == nil || o.UseGzipCompression == nil {
		return nil, false
	}
	return o.UseGzipCompression, true
}

// HasUseGzipCompression returns a boolean if a field has been set.
func (o *ChannelPublishingHttp) HasUseGzipCompression() bool {
	if o != nil && o.UseGzipCompression != nil {
		return true
	}

	return false
}

// SetUseGzipCompression gets a reference to the given bool and assigns it to the UseGzipCompression field.
func (o *ChannelPublishingHttp) SetUseGzipCompression(v bool) {
	o.UseGzipCompression = &v
}

func (o ChannelPublishingHttp) MarshalJSON() ([]byte, error) {
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
	if o.DnsServerAddress != nil {
		toSerialize["dns_server_address"] = o.DnsServerAddress
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
	if o.UseGzipCompression != nil {
		toSerialize["use_gzip_compression"] = o.UseGzipCompression
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChannelPublishingHttp) UnmarshalJSON(bytes []byte) (err error) {
	varChannelPublishingHttp := _ChannelPublishingHttp{}

	if err = json.Unmarshal(bytes, &varChannelPublishingHttp); err == nil {
		*o = ChannelPublishingHttp(varChannelPublishingHttp)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "basic_auth")
		delete(additionalProperties, "compression")
		delete(additionalProperties, "cross_playback_paths")
		delete(additionalProperties, "dns_server_address")
		delete(additionalProperties, "headers")
		delete(additionalProperties, "method")
		delete(additionalProperties, "playback_base_url")
		delete(additionalProperties, "playback_query_params")
		delete(additionalProperties, "publish_base_url")
		delete(additionalProperties, "use_gzip_compression")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableChannelPublishingHttp struct {
	value *ChannelPublishingHttp
	isSet bool
}

func (v NullableChannelPublishingHttp) Get() *ChannelPublishingHttp {
	return v.value
}

func (v *NullableChannelPublishingHttp) Set(val *ChannelPublishingHttp) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPublishingHttp) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPublishingHttp) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPublishingHttp(val *ChannelPublishingHttp) *NullableChannelPublishingHttp {
	return &NullableChannelPublishingHttp{value: val, isSet: true}
}

func (v NullableChannelPublishingHttp) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPublishingHttp) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


