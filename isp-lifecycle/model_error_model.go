/*
 * Channel Lifecycle State API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package isp

import (
	"encoding/json"
)

// checks if the ErrorModel type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorModel{}

// ErrorModel struct for ErrorModel
type ErrorModel struct {
	// A URL to the JSON Schema for this object.
	Schema interface{} `json:"$schema,omitempty" format:"uri" doc:"A URL to the JSON Schema for this object."`
	// A human-readable explanation specific to this occurrence of the problem.
	Detail interface{} `json:"detail,omitempty" doc:"A human-readable explanation specific to this occurrence of the problem."`
	// Optional list of individual error details
	Errors interface{} `json:"errors,omitempty" doc:"Optional list of individual error details"`
	// A URI reference that identifies the specific occurrence of the problem.
	Instance interface{} `json:"instance,omitempty" format:"uri" doc:"A URI reference that identifies the specific occurrence of the problem."`
	// HTTP status code
	Status interface{} `json:"status,omitempty" format:"int64" doc:"HTTP status code"`
	// A short, human-readable summary of the problem type. This value should not change between occurrences of the error.
	Title interface{} `json:"title,omitempty" doc:"A short, human-readable summary of the problem type. This value should not change between occurrences of the error."`
	// A URI reference to human-readable documentation for the error.
	Type interface{} `json:"type,omitempty" format:"uri" default:"about:blank" doc:"A URI reference to human-readable documentation for the error."`
}

// NewErrorModel instantiates a new ErrorModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorModel() *ErrorModel {
	this := ErrorModel{}
	return &this
}

// NewErrorModelWithDefaults instantiates a new ErrorModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorModelWithDefaults() *ErrorModel {
	this := ErrorModel{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErrorModel) GetSchema() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorModel) GetSchemaOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return &o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *ErrorModel) HasSchema() bool {
	if o != nil && IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given interface{} and assigns it to the Schema field.
func (o *ErrorModel) SetSchema(v interface{}) {
	o.Schema = v
}

// GetDetail returns the Detail field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErrorModel) GetDetail() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorModel) GetDetailOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Detail) {
		return nil, false
	}
	return &o.Detail, true
}

// HasDetail returns a boolean if a field has been set.
func (o *ErrorModel) HasDetail() bool {
	if o != nil && IsNil(o.Detail) {
		return true
	}

	return false
}

// SetDetail gets a reference to the given interface{} and assigns it to the Detail field.
func (o *ErrorModel) SetDetail(v interface{}) {
	o.Detail = v
}

// GetErrors returns the Errors field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErrorModel) GetErrors() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorModel) GetErrorsOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return &o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ErrorModel) HasErrors() bool {
	if o != nil && IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given interface{} and assigns it to the Errors field.
func (o *ErrorModel) SetErrors(v interface{}) {
	o.Errors = v
}

// GetInstance returns the Instance field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErrorModel) GetInstance() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorModel) GetInstanceOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Instance) {
		return nil, false
	}
	return &o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *ErrorModel) HasInstance() bool {
	if o != nil && IsNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given interface{} and assigns it to the Instance field.
func (o *ErrorModel) SetInstance(v interface{}) {
	o.Instance = v
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErrorModel) GetStatus() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorModel) GetStatusOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return &o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ErrorModel) HasStatus() bool {
	if o != nil && IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given interface{} and assigns it to the Status field.
func (o *ErrorModel) SetStatus(v interface{}) {
	o.Status = v
}

// GetTitle returns the Title field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErrorModel) GetTitle() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorModel) GetTitleOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return &o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *ErrorModel) HasTitle() bool {
	if o != nil && IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given interface{} and assigns it to the Title field.
func (o *ErrorModel) SetTitle(v interface{}) {
	o.Title = v
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ErrorModel) GetType() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ErrorModel) GetTypeOk() (*interface{}, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return &o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *ErrorModel) HasType() bool {
	if o != nil && IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given interface{} and assigns it to the Type field.
func (o *ErrorModel) SetType(v interface{}) {
	o.Type = v
}

func (o ErrorModel) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorModel) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Schema != nil {
		toSerialize["$schema"] = o.Schema
	}
	if o.Detail != nil {
		toSerialize["detail"] = o.Detail
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	if o.Instance != nil {
		toSerialize["instance"] = o.Instance
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableErrorModel struct {
	value *ErrorModel
	isSet bool
}

func (v NullableErrorModel) Get() *ErrorModel {
	return v.value
}

func (v *NullableErrorModel) Set(val *ErrorModel) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorModel) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorModel(val *ErrorModel) *NullableErrorModel {
	return &NullableErrorModel{value: val, isSet: true}
}

func (v NullableErrorModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
