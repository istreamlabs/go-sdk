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

// checks if the GetProductResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetProductResponse{}

// GetProductResponse struct for GetProductResponse
type GetProductResponse struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string `json:"$schema,omitempty"`
	// Errors when fetching product
	Errors []ErrorModelErrorsInner `json:"errors"`
	// Product Array
	Products []GetProductResponseProductsInner `json:"products"`
}

// NewGetProductResponse instantiates a new GetProductResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetProductResponse(errors []ErrorModelErrorsInner, products []GetProductResponseProductsInner) *GetProductResponse {
	this := GetProductResponse{}
	this.Errors = errors
	this.Products = products
	return &this
}

// NewGetProductResponseWithDefaults instantiates a new GetProductResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetProductResponseWithDefaults() *GetProductResponse {
	this := GetProductResponse{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *GetProductResponse) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *GetProductResponse) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *GetProductResponse) SetSchema(v string) {
	o.Schema = &v
}

// GetErrors returns the Errors field value
func (o *GetProductResponse) GetErrors() []ErrorModelErrorsInner {
	if o == nil {
		var ret []ErrorModelErrorsInner
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetErrorsOk() ([]ErrorModelErrorsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Errors, true
}

// SetErrors sets field value
func (o *GetProductResponse) SetErrors(v []ErrorModelErrorsInner) {
	o.Errors = v
}

// GetProducts returns the Products field value
func (o *GetProductResponse) GetProducts() []GetProductResponseProductsInner {
	if o == nil {
		var ret []GetProductResponseProductsInner
		return ret
	}

	return o.Products
}

// GetProductsOk returns a tuple with the Products field value
// and a boolean to check if the value has been set.
func (o *GetProductResponse) GetProductsOk() ([]GetProductResponseProductsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Products, true
}

// SetProducts sets field value
func (o *GetProductResponse) SetProducts(v []GetProductResponseProductsInner) {
	o.Products = v
}

func (o GetProductResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetProductResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["errors"] = o.Errors
	toSerialize["products"] = o.Products
	return toSerialize, nil
}

type NullableGetProductResponse struct {
	value *GetProductResponse
	isSet bool
}

func (v NullableGetProductResponse) Get() *GetProductResponse {
	return v.value
}

func (v *NullableGetProductResponse) Set(val *GetProductResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetProductResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetProductResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetProductResponse(val *GetProductResponse) *NullableGetProductResponse {
	return &NullableGetProductResponse{value: val, isSet: true}
}

func (v NullableGetProductResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetProductResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

