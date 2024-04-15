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
	"time"
)

// checks if the DeprecatedListVODsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeprecatedListVODsResponse{}

// DeprecatedListVODsResponse struct for DeprecatedListVODsResponse
type DeprecatedListVODsResponse struct {
	// description of the vod
	Description string `json:"description" doc:"description of the vod"`
	// number of files for the vod
	FileCount int64 `json:"file_count" format:"int64" doc:"number of files for the vod"`
	// date last published
	LastPublished *time.Time `json:"last_published,omitempty" format:"date-time" doc:"date last published"`
	// format
	PackagingFormat *string `json:"packaging_format,omitempty" doc:"format"`
	// store location of the vod
	Store string `json:"store" doc:"store location of the vod"`
	// store prefix for the vod
	Storeprefix string `json:"storeprefix" doc:"store prefix for the vod"`
	// total bytes for the vod
	TotalBytes int64 `json:"total_bytes" format:"int64" doc:"total bytes for the vod"`
	// version of the vod
	Version string `json:"version" doc:"version of the vod"`
	// id of the vod
	Vodid int64 `json:"vodid" format:"int64" doc:"id of the vod"`
}

// NewDeprecatedListVODsResponse instantiates a new DeprecatedListVODsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeprecatedListVODsResponse(description string, fileCount int64, store string, storeprefix string, totalBytes int64, version string, vodid int64) *DeprecatedListVODsResponse {
	this := DeprecatedListVODsResponse{}
	this.Description = description
	this.FileCount = fileCount
	this.Store = store
	this.Storeprefix = storeprefix
	this.TotalBytes = totalBytes
	this.Version = version
	this.Vodid = vodid
	return &this
}

// NewDeprecatedListVODsResponseWithDefaults instantiates a new DeprecatedListVODsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeprecatedListVODsResponseWithDefaults() *DeprecatedListVODsResponse {
	this := DeprecatedListVODsResponse{}
	return &this
}

// GetDescription returns the Description field value
func (o *DeprecatedListVODsResponse) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *DeprecatedListVODsResponse) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *DeprecatedListVODsResponse) SetDescription(v string) {
	o.Description = v
}

// GetFileCount returns the FileCount field value
func (o *DeprecatedListVODsResponse) GetFileCount() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.FileCount
}

// GetFileCountOk returns a tuple with the FileCount field value
// and a boolean to check if the value has been set.
func (o *DeprecatedListVODsResponse) GetFileCountOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileCount, true
}

// SetFileCount sets field value
func (o *DeprecatedListVODsResponse) SetFileCount(v int64) {
	o.FileCount = v
}

// GetLastPublished returns the LastPublished field value if set, zero value otherwise.
func (o *DeprecatedListVODsResponse) GetLastPublished() time.Time {
	if o == nil || IsNil(o.LastPublished) {
		var ret time.Time
		return ret
	}
	return *o.LastPublished
}

// GetLastPublishedOk returns a tuple with the LastPublished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedListVODsResponse) GetLastPublishedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastPublished) {
		return nil, false
	}
	return o.LastPublished, true
}

// HasLastPublished returns a boolean if a field has been set.
func (o *DeprecatedListVODsResponse) HasLastPublished() bool {
	if o != nil && !IsNil(o.LastPublished) {
		return true
	}

	return false
}

// SetLastPublished gets a reference to the given time.Time and assigns it to the LastPublished field.
func (o *DeprecatedListVODsResponse) SetLastPublished(v time.Time) {
	o.LastPublished = &v
}

// GetPackagingFormat returns the PackagingFormat field value if set, zero value otherwise.
func (o *DeprecatedListVODsResponse) GetPackagingFormat() string {
	if o == nil || IsNil(o.PackagingFormat) {
		var ret string
		return ret
	}
	return *o.PackagingFormat
}

// GetPackagingFormatOk returns a tuple with the PackagingFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeprecatedListVODsResponse) GetPackagingFormatOk() (*string, bool) {
	if o == nil || IsNil(o.PackagingFormat) {
		return nil, false
	}
	return o.PackagingFormat, true
}

// HasPackagingFormat returns a boolean if a field has been set.
func (o *DeprecatedListVODsResponse) HasPackagingFormat() bool {
	if o != nil && !IsNil(o.PackagingFormat) {
		return true
	}

	return false
}

// SetPackagingFormat gets a reference to the given string and assigns it to the PackagingFormat field.
func (o *DeprecatedListVODsResponse) SetPackagingFormat(v string) {
	o.PackagingFormat = &v
}

// GetStore returns the Store field value
func (o *DeprecatedListVODsResponse) GetStore() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Store
}

// GetStoreOk returns a tuple with the Store field value
// and a boolean to check if the value has been set.
func (o *DeprecatedListVODsResponse) GetStoreOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Store, true
}

// SetStore sets field value
func (o *DeprecatedListVODsResponse) SetStore(v string) {
	o.Store = v
}

// GetStoreprefix returns the Storeprefix field value
func (o *DeprecatedListVODsResponse) GetStoreprefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Storeprefix
}

// GetStoreprefixOk returns a tuple with the Storeprefix field value
// and a boolean to check if the value has been set.
func (o *DeprecatedListVODsResponse) GetStoreprefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Storeprefix, true
}

// SetStoreprefix sets field value
func (o *DeprecatedListVODsResponse) SetStoreprefix(v string) {
	o.Storeprefix = v
}

// GetTotalBytes returns the TotalBytes field value
func (o *DeprecatedListVODsResponse) GetTotalBytes() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TotalBytes
}

// GetTotalBytesOk returns a tuple with the TotalBytes field value
// and a boolean to check if the value has been set.
func (o *DeprecatedListVODsResponse) GetTotalBytesOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TotalBytes, true
}

// SetTotalBytes sets field value
func (o *DeprecatedListVODsResponse) SetTotalBytes(v int64) {
	o.TotalBytes = v
}

// GetVersion returns the Version field value
func (o *DeprecatedListVODsResponse) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *DeprecatedListVODsResponse) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *DeprecatedListVODsResponse) SetVersion(v string) {
	o.Version = v
}

// GetVodid returns the Vodid field value
func (o *DeprecatedListVODsResponse) GetVodid() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Vodid
}

// GetVodidOk returns a tuple with the Vodid field value
// and a boolean to check if the value has been set.
func (o *DeprecatedListVODsResponse) GetVodidOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vodid, true
}

// SetVodid sets field value
func (o *DeprecatedListVODsResponse) SetVodid(v int64) {
	o.Vodid = v
}

func (o DeprecatedListVODsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeprecatedListVODsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["description"] = o.Description
	toSerialize["file_count"] = o.FileCount
	if !IsNil(o.LastPublished) {
		toSerialize["last_published"] = o.LastPublished
	}
	if !IsNil(o.PackagingFormat) {
		toSerialize["packaging_format"] = o.PackagingFormat
	}
	toSerialize["store"] = o.Store
	toSerialize["storeprefix"] = o.Storeprefix
	toSerialize["total_bytes"] = o.TotalBytes
	toSerialize["version"] = o.Version
	toSerialize["vodid"] = o.Vodid
	return toSerialize, nil
}

type NullableDeprecatedListVODsResponse struct {
	value *DeprecatedListVODsResponse
	isSet bool
}

func (v NullableDeprecatedListVODsResponse) Get() *DeprecatedListVODsResponse {
	return v.value
}

func (v *NullableDeprecatedListVODsResponse) Set(val *DeprecatedListVODsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeprecatedListVODsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeprecatedListVODsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeprecatedListVODsResponse(val *DeprecatedListVODsResponse) *NullableDeprecatedListVODsResponse {
	return &NullableDeprecatedListVODsResponse{value: val, isSet: true}
}

func (v NullableDeprecatedListVODsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeprecatedListVODsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


