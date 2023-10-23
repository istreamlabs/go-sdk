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

// checks if the DeprecatedGetMp4UrlResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeprecatedGetMp4UrlResponse{}

// DeprecatedGetMp4UrlResponse struct for DeprecatedGetMp4UrlResponse
type DeprecatedGetMp4UrlResponse struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string `json:"$schema,omitempty" format:"uri" doc:"An optional URL to a JSON Schema document describing this resource"`
	PresignedUrl string `json:"presigned_url"`
}

// NewDeprecatedGetMp4UrlResponse instantiates a new DeprecatedGetMp4UrlResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeprecatedGetMp4UrlResponse(presignedUrl string) *DeprecatedGetMp4UrlResponse {
	this := DeprecatedGetMp4UrlResponse{}
	this.PresignedUrl = presignedUrl
	return &this
}

// NewDeprecatedGetMp4UrlResponseWithDefaults instantiates a new DeprecatedGetMp4UrlResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeprecatedGetMp4UrlResponseWithDefaults() *DeprecatedGetMp4UrlResponse {
	this := DeprecatedGetMp4UrlResponse{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *DeprecatedGetMp4UrlResponse) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedGetMp4UrlResponse) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *DeprecatedGetMp4UrlResponse) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *DeprecatedGetMp4UrlResponse) SetSchema(v string) {
	o.Schema = &v
}

// GetPresignedUrl returns the PresignedUrl field value
func (o *DeprecatedGetMp4UrlResponse) GetPresignedUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PresignedUrl
}

// GetPresignedUrlOk returns a tuple with the PresignedUrl field value
// and a boolean to check if the value has been set.
func (o *DeprecatedGetMp4UrlResponse) GetPresignedUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PresignedUrl, true
}

// SetPresignedUrl sets field value
func (o *DeprecatedGetMp4UrlResponse) SetPresignedUrl(v string) {
	o.PresignedUrl = v
}

func (o DeprecatedGetMp4UrlResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeprecatedGetMp4UrlResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["presigned_url"] = o.PresignedUrl
	return toSerialize, nil
}

type NullableDeprecatedGetMp4UrlResponse struct {
	value *DeprecatedGetMp4UrlResponse
	isSet bool
}

func (v NullableDeprecatedGetMp4UrlResponse) Get() *DeprecatedGetMp4UrlResponse {
	return v.value
}

func (v *NullableDeprecatedGetMp4UrlResponse) Set(val *DeprecatedGetMp4UrlResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeprecatedGetMp4UrlResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeprecatedGetMp4UrlResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeprecatedGetMp4UrlResponse(val *DeprecatedGetMp4UrlResponse) *NullableDeprecatedGetMp4UrlResponse {
	return &NullableDeprecatedGetMp4UrlResponse{value: val, isSet: true}
}

func (v NullableDeprecatedGetMp4UrlResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeprecatedGetMp4UrlResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


