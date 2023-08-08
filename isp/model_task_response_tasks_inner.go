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

// checks if the TaskResponseTasksInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskResponseTasksInner{}

// TaskResponseTasksInner struct for TaskResponseTasksInner
type TaskResponseTasksInner struct {
	// when the task was created in RFC3339Nano format
	Created time.Time `json:"created"`
	// indicates if the task is done or not
	Done bool `json:"done"`
	// when the task ended in RFC3339Nano format
	Ended time.Time `json:"ended"`
	// error message if any for this task
	Error string `json:"error"`
	// number of times the task failed
	FailureCount int32 `json:"failure_count"`
	// ID for the task
	Id string `json:"id"`
	// params sent to task
	Params string `json:"params"`
	// the progress of the task
	Progress float64 `json:"progress"`
	// Region represents the general geolocation the task is in.
	Region *string `json:"region,omitempty"`
	// Link to this resource
	Self *string `json:"self,omitempty"`
	// status for the task
	Status string `json:"status"`
	// indicates if the task succeeded or not
	Succeeded bool `json:"succeeded"`
	// type of task
	Type int32 `json:"type"`
	// priority
	Weight int32 `json:"weight"`
	// the version of the worker for this task
	WorkerVersion string `json:"workerVersion"`
	// the id of the worker for the task
	WorkerId int64 `json:"worker_id"`
}

// NewTaskResponseTasksInner instantiates a new TaskResponseTasksInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskResponseTasksInner(created time.Time, done bool, ended time.Time, error_ string, failureCount int32, id string, params string, progress float64, status string, succeeded bool, type_ int32, weight int32, workerVersion string, workerId int64) *TaskResponseTasksInner {
	this := TaskResponseTasksInner{}
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

// NewTaskResponseTasksInnerWithDefaults instantiates a new TaskResponseTasksInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskResponseTasksInnerWithDefaults() *TaskResponseTasksInner {
	this := TaskResponseTasksInner{}
	return &this
}

// GetCreated returns the Created field value
func (o *TaskResponseTasksInner) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *TaskResponseTasksInner) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *TaskResponseTasksInner) SetCreated(v time.Time) {
	o.Created = v
}

// GetDone returns the Done field value
func (o *TaskResponseTasksInner) GetDone() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Done
}

// GetDoneOk returns a tuple with the Done field value
// and a boolean to check if the value has been set.
func (o *TaskResponseTasksInner) GetDoneOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Done, true
}

// SetDone sets field value
func (o *TaskResponseTasksInner) SetDone(v bool) {
	o.Done = v
}

// GetEnded returns the Ended field value
func (o *TaskResponseTasksInner) GetEnded() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Ended
}

// GetEndedOk returns a tuple with the Ended field value
// and a boolean to check if the value has been set.
func (o *TaskResponseTasksInner) GetEndedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ended, true
}

// SetEnded sets field value
func (o *TaskResponseTasksInner) SetEnded(v time.Time) {
	o.Ended = v
}

// GetError returns the Error field value
func (o *TaskResponseTasksInner) GetError() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Error
}

// GetErrorOk returns a tuple with the Error field value
// and a boolean to check if the value has been set.
func (o *TaskResponseTasksInner) GetErrorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Error, true
}

// SetError sets field value
func (o *TaskResponseTasksInner) SetError(v string) {
	o.Error = v
}

// GetFailureCount returns the FailureCount field value
func (o *TaskResponseTasksInner) GetFailureCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FailureCount
}

// GetFailureCountOk returns a tuple with the FailureCount field value
// and a boolean to check if the value has been set.
func (o *TaskResponseTasksInner) GetFailureCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FailureCount, true
}

// SetFailureCount sets field value
func (o *TaskResponseTasksInner) SetFailureCount(v int32) {
	o.FailureCount = v
}

// GetId returns the Id field value
func (o *TaskResponseTasksInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *TaskResponseTasksInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *TaskResponseTasksInner) SetId(v string) {
	o.Id = v
}

// GetParams returns the Params field value
func (o *TaskResponseTasksInner) GetParams() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Params
}

// GetParamsOk returns a tuple with the Params field value
// and a boolean to check if the value has been set.
func (o *TaskResponseTasksInner) GetParamsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Params, true
}

// SetParams sets field value
func (o *TaskResponseTasksInner) SetParams(v string) {
	o.Params = v
}

// GetProgress returns the Progress field value
func (o *TaskResponseTasksInner) GetProgress() float64 {
	if o == nil {
		var ret float64
		return ret
	}

	return o.Progress
}

// GetProgressOk returns a tuple with the Progress field value
// and a boolean to check if the value has been set.
func (o *TaskResponseTasksInner) GetProgressOk() (*float64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Progress, true
}

// SetProgress sets field value
func (o *TaskResponseTasksInner) SetProgress(v float64) {
	o.Progress = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *TaskResponseTasksInner) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResponseTasksInner) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *TaskResponseTasksInner) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *TaskResponseTasksInner) SetRegion(v string) {
	o.Region = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *TaskResponseTasksInner) GetSelf() string {
	if o == nil || IsNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskResponseTasksInner) GetSelfOk() (*string, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *TaskResponseTasksInner) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *TaskResponseTasksInner) SetSelf(v string) {
	o.Self = &v
}

// GetStatus returns the Status field value
func (o *TaskResponseTasksInner) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *TaskResponseTasksInner) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *TaskResponseTasksInner) SetStatus(v string) {
	o.Status = v
}

// GetSucceeded returns the Succeeded field value
func (o *TaskResponseTasksInner) GetSucceeded() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Succeeded
}

// GetSucceededOk returns a tuple with the Succeeded field value
// and a boolean to check if the value has been set.
func (o *TaskResponseTasksInner) GetSucceededOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Succeeded, true
}

// SetSucceeded sets field value
func (o *TaskResponseTasksInner) SetSucceeded(v bool) {
	o.Succeeded = v
}

// GetType returns the Type field value
func (o *TaskResponseTasksInner) GetType() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TaskResponseTasksInner) GetTypeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TaskResponseTasksInner) SetType(v int32) {
	o.Type = v
}

// GetWeight returns the Weight field value
func (o *TaskResponseTasksInner) GetWeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Weight
}

// GetWeightOk returns a tuple with the Weight field value
// and a boolean to check if the value has been set.
func (o *TaskResponseTasksInner) GetWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Weight, true
}

// SetWeight sets field value
func (o *TaskResponseTasksInner) SetWeight(v int32) {
	o.Weight = v
}

// GetWorkerVersion returns the WorkerVersion field value
func (o *TaskResponseTasksInner) GetWorkerVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WorkerVersion
}

// GetWorkerVersionOk returns a tuple with the WorkerVersion field value
// and a boolean to check if the value has been set.
func (o *TaskResponseTasksInner) GetWorkerVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkerVersion, true
}

// SetWorkerVersion sets field value
func (o *TaskResponseTasksInner) SetWorkerVersion(v string) {
	o.WorkerVersion = v
}

// GetWorkerId returns the WorkerId field value
func (o *TaskResponseTasksInner) GetWorkerId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.WorkerId
}

// GetWorkerIdOk returns a tuple with the WorkerId field value
// and a boolean to check if the value has been set.
func (o *TaskResponseTasksInner) GetWorkerIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WorkerId, true
}

// SetWorkerId sets field value
func (o *TaskResponseTasksInner) SetWorkerId(v int64) {
	o.WorkerId = v
}

func (o TaskResponseTasksInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskResponseTasksInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
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
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	toSerialize["status"] = o.Status
	toSerialize["succeeded"] = o.Succeeded
	toSerialize["type"] = o.Type
	toSerialize["weight"] = o.Weight
	toSerialize["workerVersion"] = o.WorkerVersion
	toSerialize["worker_id"] = o.WorkerId
	return toSerialize, nil
}

type NullableTaskResponseTasksInner struct {
	value *TaskResponseTasksInner
	isSet bool
}

func (v NullableTaskResponseTasksInner) Get() *TaskResponseTasksInner {
	return v.value
}

func (v *NullableTaskResponseTasksInner) Set(val *TaskResponseTasksInner) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskResponseTasksInner) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskResponseTasksInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskResponseTasksInner(val *TaskResponseTasksInner) *NullableTaskResponseTasksInner {
	return &NullableTaskResponseTasksInner{value: val, isSet: true}
}

func (v NullableTaskResponseTasksInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskResponseTasksInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


