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

// checks if the GetProductConfigResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProductConfigResponse{}

// GetProductConfigResponse struct for GetProductConfigResponse
type GetProductConfigResponse struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string `json:"$schema,omitempty" format:"uri" doc:"An optional URL to a JSON Schema document describing this resource"`
	// Errors when fetching product config
	Errors []ErrorModelErrorsInner `json:"errors" doc:"Errors when fetching product config"`
	// Product Config for a given org and product id
	ProductConfig []GetProductConfigResponseProductConfigInner `json:"product_config" doc:"Product Config for a given org and product id"`
}

// NewGetProductConfigResponse instantiates a new GetProductConfigResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductConfigResponse(errors []ErrorModelErrorsInner, productConfig []GetProductConfigResponseProductConfigInner) *GetProductConfigResponse {
	this := GetProductConfigResponse{}
	this.Errors = errors
	this.ProductConfig = productConfig
	return &this
}

// NewGetProductConfigResponseWithDefaults instantiates a new GetProductConfigResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductConfigResponseWithDefaults() *GetProductConfigResponse {
	this := GetProductConfigResponse{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *GetProductConfigResponse) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponse) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *GetProductConfigResponse) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *GetProductConfigResponse) SetSchema(v string) {
	o.Schema = &v
}

// GetErrors returns the Errors field value
func (o *GetProductConfigResponse) GetErrors() []ErrorModelErrorsInner {
	if o == nil {
		var ret []ErrorModelErrorsInner
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponse) GetErrorsOk() ([]ErrorModelErrorsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *GetProductConfigResponse) SetErrors(v []ErrorModelErrorsInner) {
	o.Errors = v
}

// GetProductConfig returns the ProductConfig field value
func (o *GetProductConfigResponse) GetProductConfig() []GetProductConfigResponseProductConfigInner {
	if o == nil {
		var ret []GetProductConfigResponseProductConfigInner
		return ret
	}

	return o.ProductConfig
}

// GetProductConfigOk returns a tuple with the ProductConfig field value
// and a boolean to check if the value has been set.
func (o *GetProductConfigResponse) GetProductConfigOk() ([]GetProductConfigResponseProductConfigInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.ProductConfig, true
}

// SetProductConfig sets field value
func (o *GetProductConfigResponse) SetProductConfig(v []GetProductConfigResponseProductConfigInner) {
	o.ProductConfig = v
}

func (o GetProductConfigResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProductConfigResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["errors"] = o.Errors
	toSerialize["product_config"] = o.ProductConfig
	return toSerialize, nil
}

type NullableGetProductConfigResponse struct {
	value *GetProductConfigResponse
	isSet bool
}

func (v NullableGetProductConfigResponse) Get() *GetProductConfigResponse {
	return v.value
}

func (v *NullableGetProductConfigResponse) Set(val *GetProductConfigResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductConfigResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductConfigResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductConfigResponse(val *GetProductConfigResponse) *NullableGetProductConfigResponse {
	return &NullableGetProductConfigResponse{value: val, isSet: true}
}

func (v NullableGetProductConfigResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProductConfigResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


