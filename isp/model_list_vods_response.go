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

// ListVODsResponse struct for ListVODsResponse
type ListVODsResponse struct {
	Description string `json:"description"`
	FileCount int64 `json:"file_count"`
	Id string `json:"id"`
	LastPublished *time.Time `json:"last_published,omitempty"`
	PackagingFormat *string `json:"packaging_format,omitempty"`
	Store string `json:"store"`
	StorePrefix string `json:"store_prefix"`
	TotalBytes int64 `json:"total_bytes"`
	Version string `json:"version"`
}

// NewListVODsResponse instantiates a new ListVODsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListVODsResponse(description string, fileCount int64, id string, store string, storePrefix string, totalBytes int64, version string) *ListVODsResponse {
	this := ListVODsResponse{}
	this.Description = description
	this.FileCount = fileCount
	this.Id = id
	this.Store = store
	this.StorePrefix = storePrefix
	this.TotalBytes = totalBytes
	this.Version = version
	return &this
}

// NewListVODsResponseWithDefaults instantiates a new ListVODsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListVODsResponseWithDefaults() *ListVODsResponse {
	this := ListVODsResponse{}
	return &this
}

// GetDescription returns the Description field value
func (o *ListVODsResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *ListVODsResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *ListVODsResponse) SetDescription(v string) {
	o.Description = v
}

// GetFileCount returns the FileCount field value
func (o *ListVODsResponse) GetFileCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FileCount
}

// GetFileCountOk returns a tuple with the FileCount field value
// and a boolean to check if the value has been set.
func (o *ListVODsResponse) GetFileCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileCount, true
}

// SetFileCount sets field value
func (o *ListVODsResponse) SetFileCount(v int64) {
	o.FileCount = v
}

// GetId returns the Id field value
func (o *ListVODsResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ListVODsResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ListVODsResponse) SetId(v string) {
	o.Id = v
}

// GetLastPublished returns the LastPublished field value if set, zero value otherwise.
func (o *ListVODsResponse) GetLastPublished() time.Time {
	if o == nil || o.LastPublished == nil {
		var ret time.Time
		return ret
	}
	return *o.LastPublished
}

// GetLastPublishedOk returns a tuple with the LastPublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVODsResponse) GetLastPublishedOk() (*time.Time, bool) {
	if o == nil || o.LastPublished == nil {
		return nil, false
	}
	return o.LastPublished, true
}

// HasLastPublished returns a boolean if a field has been set.
func (o *ListVODsResponse) HasLastPublished() bool {
	if o != nil && o.LastPublished != nil {
		return true
	}

	return false
}

// SetLastPublished gets a reference to the given time.Time and assigns it to the LastPublished field.
func (o *ListVODsResponse) SetLastPublished(v time.Time) {
	o.LastPublished = &v
}

// GetPackagingFormat returns the PackagingFormat field value if set, zero value otherwise.
func (o *ListVODsResponse) GetPackagingFormat() string {
	if o == nil || o.PackagingFormat == nil {
		var ret string
		return ret
	}
	return *o.PackagingFormat
}

// GetPackagingFormatOk returns a tuple with the PackagingFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVODsResponse) GetPackagingFormatOk() (*string, bool) {
	if o == nil || o.PackagingFormat == nil {
		return nil, false
	}
	return o.PackagingFormat, true
}

// HasPackagingFormat returns a boolean if a field has been set.
func (o *ListVODsResponse) HasPackagingFormat() bool {
	if o != nil && o.PackagingFormat != nil {
		return true
	}

	return false
}

// SetPackagingFormat gets a reference to the given string and assigns it to the PackagingFormat field.
func (o *ListVODsResponse) SetPackagingFormat(v string) {
	o.PackagingFormat = &v
}

// GetStore returns the Store field value
func (o *ListVODsResponse) GetStore() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Store
}

// GetStoreOk returns a tuple with the Store field value
// and a boolean to check if the value has been set.
func (o *ListVODsResponse) GetStoreOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Store, true
}

// SetStore sets field value
func (o *ListVODsResponse) SetStore(v string) {
	o.Store = v
}

// GetStorePrefix returns the StorePrefix field value
func (o *ListVODsResponse) GetStorePrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StorePrefix
}

// GetStorePrefixOk returns a tuple with the StorePrefix field value
// and a boolean to check if the value has been set.
func (o *ListVODsResponse) GetStorePrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StorePrefix, true
}

// SetStorePrefix sets field value
func (o *ListVODsResponse) SetStorePrefix(v string) {
	o.StorePrefix = v
}

// GetTotalBytes returns the TotalBytes field value
func (o *ListVODsResponse) GetTotalBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalBytes
}

// GetTotalBytesOk returns a tuple with the TotalBytes field value
// and a boolean to check if the value has been set.
func (o *ListVODsResponse) GetTotalBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalBytes, true
}

// SetTotalBytes sets field value
func (o *ListVODsResponse) SetTotalBytes(v int64) {
	o.TotalBytes = v
}

// GetVersion returns the Version field value
func (o *ListVODsResponse) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *ListVODsResponse) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *ListVODsResponse) SetVersion(v string) {
	o.Version = v
}

func (o ListVODsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["file_count"] = o.FileCount
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.LastPublished != nil {
		toSerialize["last_published"] = o.LastPublished
	}
	if o.PackagingFormat != nil {
		toSerialize["packaging_format"] = o.PackagingFormat
	}
	if true {
		toSerialize["store"] = o.Store
	}
	if true {
		toSerialize["store_prefix"] = o.StorePrefix
	}
	if true {
		toSerialize["total_bytes"] = o.TotalBytes
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableListVODsResponse struct {
	value *ListVODsResponse
	isSet bool
}

func (v NullableListVODsResponse) Get() *ListVODsResponse {
	return v.value
}

func (v *NullableListVODsResponse) Set(val *ListVODsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListVODsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListVODsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListVODsResponse(val *ListVODsResponse) *NullableListVODsResponse {
	return &NullableListVODsResponse{value: val, isSet: true}
}

func (v NullableListVODsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListVODsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


