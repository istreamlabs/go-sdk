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

// checks if the GetTaskResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetTaskResponse{}

// GetTaskResponse struct for GetTaskResponse
type GetTaskResponse struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string `json:"$schema,omitempty" format:"uri" doc:"An optional URL to a JSON Schema document describing this resource"`
	// when the task was created in RFC3339Nano format
	Created time.Time `json:"created" format:"date-time" doc:"when the task was created in RFC3339Nano format"`
	// indicates if the task is done or not
	Done bool `json:"done" doc:"indicates if the task is done or not"`
	// when the task ended in RFC3339Nano format
	Ended time.Time `json:"ended" format:"date-time" doc:"when the task ended in RFC3339Nano format"`
	// error message if any for this task
	Error string `json:"error" doc:"error message if any for this task"`
	// number of times the task failed
	FailureCount int32 `json:"failure_count" format:"int32" doc:"number of times the task failed"`
	// ID for the task
	Id string `json:"id" doc:"ID for the task"`
	// params sent to task
	Params string `json:"params" doc:"params sent to task"`
	// the progress of the task
	Progress float64 `json:"progress" format:"double" doc:"the progress of the task"`
	// Region represents the general geolocation the task is in.
	Region *string `json:"region,omitempty" enum:"US_WEST,US_EAST" doc:"Region represents the general geolocation the task is in."`
	// status for the task
	Status string `json:"status" doc:"status for the task"`
	// indicates if the task succeeded or not
	Succeeded bool `json:"succeeded" doc:"indicates if the task succeeded or not"`
	// type of task
	Type int32 `json:"type" format:"int32" doc:"type of task"`
	// priority
	Weight int32 `json:"weight" format:"int32" doc:"priority"`
	// the version of the worker for this task
	WorkerVersion string `json:"workerVersion" doc:"the version of the worker for this task"`
	// the id of the worker for the task
	WorkerId int64 `json:"worker_id" format:"int64" doc:"the id of the worker for the task"`
}

// NewGetTaskResponse instantiates a new GetTaskResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetTaskResponse(created time.Time, done bool, ended time.Time, error_ string, failureCount int32, id string, params string, progress float64, status string, succeeded bool, type_ int32, weight int32, workerVersion string, workerId int64) *GetTaskResponse {
	this := GetTaskResponse{}
	this.Created = created
	this.Done = done
	this.Ended = ended
	this.Error = error_
	this.FailureCount = failureCount
	this.Id = id
	this.Params = params
	this.Progress = progress
	this.Status = status
	this.Succeeded = succeeded
	this.Type = type_
	this.Weight = weight
	this.WorkerVersion = workerVersion
	this.WorkerId = workerId
	return &this
}

// NewGetTaskResponseWithDefaults instantiates a new GetTaskResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetTaskResponseWithDefaults() *GetTaskResponse {
	this := GetTaskResponse{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *GetTaskResponse) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTaskResponse) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *GetTaskResponse) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *GetTaskResponse) SetSchema(v string) {
	o.Schema = &v
}

// GetCreated returns the Created field value
func (o *GetTaskResponse) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *GetTaskResponse) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *GetTaskResponse) SetCreated(v time.Time) {
	o.Created = v
}

// GetDone returns the Done field value
func (o *GetTaskResponse) GetDone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Done
}

// GetDoneOk returns a tuple with the Done field value
// and a boolean to check if the value has been set.
func (o *GetTaskResponse) GetDoneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Done, true
}

// SetDone sets field value
func (o *GetTaskResponse) SetDone(v bool) {
	o.Done = v
}

// GetEnded returns the Ended field value
func (o *GetTaskResponse) GetEnded() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Ended
}

// GetEndedOk returns a tuple with the Ended field value
// and a boolean to check if the value has been set.
func (o *GetTaskResponse) GetEndedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ended, true
}

// SetEnded sets field value
func (o *GetTaskResponse) SetEnded(v time.Time) {
	o.Ended = v
}

// GetError returns the Error field value
func (o *GetTaskResponse) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *GetTaskResponse) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *GetTaskResponse) SetError(v string) {
	o.Error = v
}

// GetFailureCount returns the FailureCount field value
func (o *GetTaskResponse) GetFailureCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FailureCount
}

// GetFailureCountOk returns a tuple with the FailureCount field value
// and a boolean to check if the value has been set.
func (o *GetTaskResponse) GetFailureCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailureCount, true
}

// SetFailureCount sets field value
func (o *GetTaskResponse) SetFailureCount(v int32) {
	o.FailureCount = v
}

// GetId returns the Id field value
func (o *GetTaskResponse) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GetTaskResponse) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GetTaskResponse) SetId(v string) {
	o.Id = v
}

// GetParams returns the Params field value
func (o *GetTaskResponse) GetParams() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *GetTaskResponse) GetParamsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *GetTaskResponse) SetParams(v string) {
	o.Params = v
}

// GetProgress returns the Progress field value
func (o *GetTaskResponse) GetProgress() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *GetTaskResponse) GetProgressOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Progress, true
}

// SetProgress sets field value
func (o *GetTaskResponse) SetProgress(v float64) {
	o.Progress = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *GetTaskResponse) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetTaskResponse) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *GetTaskResponse) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *GetTaskResponse) SetRegion(v string) {
	o.Region = &v
}

// GetStatus returns the Status field value
func (o *GetTaskResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetTaskResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetTaskResponse) SetStatus(v string) {
	o.Status = v
}

// GetSucceeded returns the Succeeded field value
func (o *GetTaskResponse) GetSucceeded() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Succeeded
}

// GetSucceededOk returns a tuple with the Succeeded field value
// and a boolean to check if the value has been set.
func (o *GetTaskResponse) GetSucceededOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Succeeded, true
}

// SetSucceeded sets field value
func (o *GetTaskResponse) SetSucceeded(v bool) {
	o.Succeeded = v
}

// GetType returns the Type field value
func (o *GetTaskResponse) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GetTaskResponse) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GetTaskResponse) SetType(v int32) {
	o.Type = v
}

// GetWeight returns the Weight field value
func (o *GetTaskResponse) GetWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *GetTaskResponse) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *GetTaskResponse) SetWeight(v int32) {
	o.Weight = v
}

// GetWorkerVersion returns the WorkerVersion field value
func (o *GetTaskResponse) GetWorkerVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkerVersion
}

// GetWorkerVersionOk returns a tuple with the WorkerVersion field value
// and a boolean to check if the value has been set.
func (o *GetTaskResponse) GetWorkerVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkerVersion, true
}

// SetWorkerVersion sets field value
func (o *GetTaskResponse) SetWorkerVersion(v string) {
	o.WorkerVersion = v
}

// GetWorkerId returns the WorkerId field value
func (o *GetTaskResponse) GetWorkerId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.WorkerId
}

// GetWorkerIdOk returns a tuple with the WorkerId field value
// and a boolean to check if the value has been set.
func (o *GetTaskResponse) GetWorkerIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkerId, true
}

// SetWorkerId sets field value
func (o *GetTaskResponse) SetWorkerId(v int64) {
	o.WorkerId = v
}

func (o GetTaskResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetTaskResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	toSerialize["created"] = o.Created
	toSerialize["done"] = o.Done
	toSerialize["ended"] = o.Ended
	toSerialize["error"] = o.Error
	toSerialize["failure_count"] = o.FailureCount
	toSerialize["id"] = o.Id
	toSerialize["params"] = o.Params
	toSerialize["progress"] = o.Progress
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	toSerialize["status"] = o.Status
	toSerialize["succeeded"] = o.Succeeded
	toSerialize["type"] = o.Type
	toSerialize["weight"] = o.Weight
	toSerialize["workerVersion"] = o.WorkerVersion
	toSerialize["worker_id"] = o.WorkerId
	return toSerialize, nil
}

type NullableGetTaskResponse struct {
	value *GetTaskResponse
	isSet bool
}

func (v NullableGetTaskResponse) Get() *GetTaskResponse {
	return v.value
}

func (v *NullableGetTaskResponse) Set(val *GetTaskResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetTaskResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetTaskResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetTaskResponse(val *GetTaskResponse) *NullableGetTaskResponse {
	return &NullableGetTaskResponse{value: val, isSet: true}
}

func (v NullableGetTaskResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetTaskResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
