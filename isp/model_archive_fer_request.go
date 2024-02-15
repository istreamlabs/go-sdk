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

// checks if the ArchiveFERRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ArchiveFERRequest{}

// ArchiveFERRequest struct for ArchiveFERRequest
type ArchiveFERRequest struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string `json:"$schema,omitempty" format:"uri" doc:"An optional URL to a JSON Schema document describing this resource"`
	// Correlation ID for this FER archive request
	CorrelationId string `json:"correlation_id" doc:"Correlation ID for this FER archive request"`
	// Portion of the query string that applies to all packages
	GlobalQueryString string `json:"global_query_string" doc:"Portion of the query string that applies to all packages"`
	Notification ArchiveFERRequestNotification `json:"notification"`
	// Packages to be archived as FERs
	Packages []ArchiveFERRequestPackagesInner `json:"packages" doc:"Packages to be archived as FERs"`
}

// NewArchiveFERRequest instantiates a new ArchiveFERRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArchiveFERRequest(correlationId string, globalQueryString string, notification ArchiveFERRequestNotification, packages []ArchiveFERRequestPackagesInner) *ArchiveFERRequest {
	this := ArchiveFERRequest{}
	this.CorrelationId = correlationId
	this.GlobalQueryString = globalQueryString
	this.Notification = notification
	this.Packages = packages
	return &this
}

// NewArchiveFERRequestWithDefaults instantiates a new ArchiveFERRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArchiveFERRequestWithDefaults() *ArchiveFERRequest {
	this := ArchiveFERRequest{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *ArchiveFERRequest) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ArchiveFERRequest) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *ArchiveFERRequest) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *ArchiveFERRequest) SetSchema(v string) {
	o.Schema = &v
}

// GetCorrelationId returns the CorrelationId field value
func (o *ArchiveFERRequest) GetCorrelationId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value
// and a boolean to check if the value has been set.
func (o *ArchiveFERRequest) GetCorrelationIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CorrelationId, true
}

// SetCorrelationId sets field value
func (o *ArchiveFERRequest) SetCorrelationId(v string) {
	o.CorrelationId = v
}

// GetGlobalQueryString returns the GlobalQueryString field value
func (o *ArchiveFERRequest) GetGlobalQueryString() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.GlobalQueryString
}

// GetGlobalQueryStringOk returns a tuple with the GlobalQueryString field value
// and a boolean to check if the value has been set.
func (o *ArchiveFERRequest) GetGlobalQueryStringOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.GlobalQueryString, true
}

// SetGlobalQueryString sets field value
func (o *ArchiveFERRequest) SetGlobalQueryString(v string) {
	o.GlobalQueryString = v
}

// GetNotification returns the Notification field value
func (o *ArchiveFERRequest) GetNotification() ArchiveFERRequestNotification {
	if o == nil {
		var ret ArchiveFERRequestNotification
		return ret
	}

	return o.Notification
}

// GetNotificationOk returns a tuple with the Notification field value
// and a boolean to check if the value has been set.
func (o *ArchiveFERRequest) GetNotificationOk() (*ArchiveFERRequestNotification, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Notification, true
}

// SetNotification sets field value
func (o *ArchiveFERRequest) SetNotification(v ArchiveFERRequestNotification) {
	o.Notification = v
}

// GetPackages returns the Packages field value
func (o *ArchiveFERRequest) GetPackages() []ArchiveFERRequestPackagesInner {
	if o == nil {
		var ret []ArchiveFERRequestPackagesInner
		return ret
	}

	return o.Packages
}

// GetPackagesOk returns a tuple with the Packages field value
// and a boolean to check if the value has been set.
func (o *ArchiveFERRequest) GetPackagesOk() ([]ArchiveFERRequestPackagesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Packages, true
}

// SetPackages sets field value
func (o *ArchiveFERRequest) SetPackages(v []ArchiveFERRequestPackagesInner) {
	o.Packages = v
}

func (o ArchiveFERRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ArchiveFERRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["correlation_id"] = o.CorrelationId
	toSerialize["global_query_string"] = o.GlobalQueryString
	toSerialize["notification"] = o.Notification
	toSerialize["packages"] = o.Packages
	return toSerialize, nil
}

type NullableArchiveFERRequest struct {
	value *ArchiveFERRequest
	isSet bool
}

func (v NullableArchiveFERRequest) Get() *ArchiveFERRequest {
	return v.value
}

func (v *NullableArchiveFERRequest) Set(val *ArchiveFERRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableArchiveFERRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableArchiveFERRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArchiveFERRequest(val *ArchiveFERRequest) *NullableArchiveFERRequest {
	return &NullableArchiveFERRequest{value: val, isSet: true}
}

func (v NullableArchiveFERRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArchiveFERRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

