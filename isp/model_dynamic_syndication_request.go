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

// checks if the DynamicSyndicationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DynamicSyndicationRequest{}

// DynamicSyndicationRequest struct for DynamicSyndicationRequest
type DynamicSyndicationRequest struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string `json:"$schema,omitempty" format:"uri" doc:"An optional URL to a JSON Schema document describing this resource"`
	Archive DynamicSyndicationRequestArchive `json:"archive"`
	// Correlation ID for this FER archive request
	CorrelationId string `json:"correlation_id" doc:"Correlation ID for this FER archive request"`
	// URL of the main manifest to reference for the mp4
	ManifestUrl string `json:"manifest_url" doc:"URL of the main manifest to reference for the mp4"`
	Notification DynamicSyndicationRequestNotification `json:"notification"`
	// Query string containing params for the manifest url
	QueryString string `json:"query_string" doc:"Query string containing params for the manifest url"`
	// List of files to be created by Syndication
	SyndicationFiles []DynamicSyndicationRequestSyndicationFilesInner `json:"syndication_files" minItems:"1" doc:"List of files to be created by Syndication"`
}

// NewDynamicSyndicationRequest instantiates a new DynamicSyndicationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDynamicSyndicationRequest(archive DynamicSyndicationRequestArchive, correlationId string, manifestUrl string, notification DynamicSyndicationRequestNotification, queryString string, syndicationFiles []DynamicSyndicationRequestSyndicationFilesInner) *DynamicSyndicationRequest {
	this := DynamicSyndicationRequest{}
	this.Archive = archive
	this.CorrelationId = correlationId
	this.ManifestUrl = manifestUrl
	this.Notification = notification
	this.QueryString = queryString
	this.SyndicationFiles = syndicationFiles
	return &this
}

// NewDynamicSyndicationRequestWithDefaults instantiates a new DynamicSyndicationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDynamicSyndicationRequestWithDefaults() *DynamicSyndicationRequest {
	this := DynamicSyndicationRequest{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *DynamicSyndicationRequest) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DynamicSyndicationRequest) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *DynamicSyndicationRequest) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *DynamicSyndicationRequest) SetSchema(v string) {
	o.Schema = &v
}

// GetArchive returns the Archive field value
func (o *DynamicSyndicationRequest) GetArchive() DynamicSyndicationRequestArchive {
	if o == nil {
		var ret DynamicSyndicationRequestArchive
		return ret
	}

	return o.Archive
}

// GetArchiveOk returns a tuple with the Archive field value
// and a boolean to check if the value has been set.
func (o *DynamicSyndicationRequest) GetArchiveOk() (*DynamicSyndicationRequestArchive, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archive, true
}

// SetArchive sets field value
func (o *DynamicSyndicationRequest) SetArchive(v DynamicSyndicationRequestArchive) {
	o.Archive = v
}

// GetCorrelationId returns the CorrelationId field value
func (o *DynamicSyndicationRequest) GetCorrelationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value
// and a boolean to check if the value has been set.
func (o *DynamicSyndicationRequest) GetCorrelationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CorrelationId, true
}

// SetCorrelationId sets field value
func (o *DynamicSyndicationRequest) SetCorrelationId(v string) {
	o.CorrelationId = v
}

// GetManifestUrl returns the ManifestUrl field value
func (o *DynamicSyndicationRequest) GetManifestUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ManifestUrl
}

// GetManifestUrlOk returns a tuple with the ManifestUrl field value
// and a boolean to check if the value has been set.
func (o *DynamicSyndicationRequest) GetManifestUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ManifestUrl, true
}

// SetManifestUrl sets field value
func (o *DynamicSyndicationRequest) SetManifestUrl(v string) {
	o.ManifestUrl = v
}

// GetNotification returns the Notification field value
func (o *DynamicSyndicationRequest) GetNotification() DynamicSyndicationRequestNotification {
	if o == nil {
		var ret DynamicSyndicationRequestNotification
		return ret
	}

	return o.Notification
}

// GetNotificationOk returns a tuple with the Notification field value
// and a boolean to check if the value has been set.
func (o *DynamicSyndicationRequest) GetNotificationOk() (*DynamicSyndicationRequestNotification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Notification, true
}

// SetNotification sets field value
func (o *DynamicSyndicationRequest) SetNotification(v DynamicSyndicationRequestNotification) {
	o.Notification = v
}

// GetQueryString returns the QueryString field value
func (o *DynamicSyndicationRequest) GetQueryString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.QueryString
}

// GetQueryStringOk returns a tuple with the QueryString field value
// and a boolean to check if the value has been set.
func (o *DynamicSyndicationRequest) GetQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueryString, true
}

// SetQueryString sets field value
func (o *DynamicSyndicationRequest) SetQueryString(v string) {
	o.QueryString = v
}

// GetSyndicationFiles returns the SyndicationFiles field value
func (o *DynamicSyndicationRequest) GetSyndicationFiles() []DynamicSyndicationRequestSyndicationFilesInner {
	if o == nil {
		var ret []DynamicSyndicationRequestSyndicationFilesInner
		return ret
	}

	return o.SyndicationFiles
}

// GetSyndicationFilesOk returns a tuple with the SyndicationFiles field value
// and a boolean to check if the value has been set.
func (o *DynamicSyndicationRequest) GetSyndicationFilesOk() ([]DynamicSyndicationRequestSyndicationFilesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.SyndicationFiles, true
}

// SetSyndicationFiles sets field value
func (o *DynamicSyndicationRequest) SetSyndicationFiles(v []DynamicSyndicationRequestSyndicationFilesInner) {
	o.SyndicationFiles = v
}

func (o DynamicSyndicationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DynamicSyndicationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["archive"] = o.Archive
	toSerialize["correlation_id"] = o.CorrelationId
	toSerialize["manifest_url"] = o.ManifestUrl
	toSerialize["notification"] = o.Notification
	toSerialize["query_string"] = o.QueryString
	toSerialize["syndication_files"] = o.SyndicationFiles
	return toSerialize, nil
}

type NullableDynamicSyndicationRequest struct {
	value *DynamicSyndicationRequest
	isSet bool
}

func (v NullableDynamicSyndicationRequest) Get() *DynamicSyndicationRequest {
	return v.value
}

func (v *NullableDynamicSyndicationRequest) Set(val *DynamicSyndicationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDynamicSyndicationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDynamicSyndicationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDynamicSyndicationRequest(val *DynamicSyndicationRequest) *NullableDynamicSyndicationRequest {
	return &NullableDynamicSyndicationRequest{value: val, isSet: true}
}

func (v NullableDynamicSyndicationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDynamicSyndicationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


