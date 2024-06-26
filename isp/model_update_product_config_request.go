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

// checks if the UpdateProductConfigRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProductConfigRequest{}

// UpdateProductConfigRequest struct for UpdateProductConfigRequest
type UpdateProductConfigRequest struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string `json:"$schema,omitempty" format:"uri" doc:"An optional URL to a JSON Schema document describing this resource"`
	ArchiveSettings *GetProductConfigResponseProductConfigInnerArchiveSettings `json:"archive_settings,omitempty"`
	CollapseConfig *GetProductConfigResponseProductConfigInnerCollapseConfig `json:"collapse_config,omitempty"`
	CollapseTriggerConfig *GetProductConfigResponseProductConfigInnerCollapseTriggerConfig `json:"collapse_trigger_config,omitempty"`
	EdcPartialPresentations *GetProductConfigResponseProductConfigInnerEdcPartialPresentations `json:"edc_partial_presentations,omitempty"`
	// Template to transcode mp4 to hls
	EdcTranscodeTemplate *map[string]interface{} `json:"edc_transcode_template,omitempty" doc:"Template to transcode mp4 to hls"`
	// Notifiaction settings for collapses
	Notifications []GetProductConfigResponseProductConfigInnerNotificationsInner `json:"notifications,omitempty" doc:"Notifiaction settings for collapses"`
	// Region represents the general geolocation the product is in.
	Region *string `json:"region,omitempty" enum:"US_WEST,US_EAST" doc:"Region represents the general geolocation the product is in."`
	// store for product
	Store *string `json:"store,omitempty" doc:"store for product"`
	WorkflowConfig *GetProductConfigResponseProductConfigInnerWorkflowConfig `json:"workflow_config,omitempty"`
}

// NewUpdateProductConfigRequest instantiates a new UpdateProductConfigRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProductConfigRequest() *UpdateProductConfigRequest {
	this := UpdateProductConfigRequest{}
	return &this
}

// NewUpdateProductConfigRequestWithDefaults instantiates a new UpdateProductConfigRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProductConfigRequestWithDefaults() *UpdateProductConfigRequest {
	this := UpdateProductConfigRequest{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *UpdateProductConfigRequest) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequest) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *UpdateProductConfigRequest) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *UpdateProductConfigRequest) SetSchema(v string) {
	o.Schema = &v
}

// GetArchiveSettings returns the ArchiveSettings field value if set, zero value otherwise.
func (o *UpdateProductConfigRequest) GetArchiveSettings() GetProductConfigResponseProductConfigInnerArchiveSettings {
	if o == nil || IsNil(o.ArchiveSettings) {
		var ret GetProductConfigResponseProductConfigInnerArchiveSettings
		return ret
	}
	return *o.ArchiveSettings
}

// GetArchiveSettingsOk returns a tuple with the ArchiveSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequest) GetArchiveSettingsOk() (*GetProductConfigResponseProductConfigInnerArchiveSettings, bool) {
	if o == nil || IsNil(o.ArchiveSettings) {
		return nil, false
	}
	return o.ArchiveSettings, true
}

// HasArchiveSettings returns a boolean if a field has been set.
func (o *UpdateProductConfigRequest) HasArchiveSettings() bool {
	if o != nil && !IsNil(o.ArchiveSettings) {
		return true
	}

	return false
}

// SetArchiveSettings gets a reference to the given GetProductConfigResponseProductConfigInnerArchiveSettings and assigns it to the ArchiveSettings field.
func (o *UpdateProductConfigRequest) SetArchiveSettings(v GetProductConfigResponseProductConfigInnerArchiveSettings) {
	o.ArchiveSettings = &v
}

// GetCollapseConfig returns the CollapseConfig field value if set, zero value otherwise.
func (o *UpdateProductConfigRequest) GetCollapseConfig() GetProductConfigResponseProductConfigInnerCollapseConfig {
	if o == nil || IsNil(o.CollapseConfig) {
		var ret GetProductConfigResponseProductConfigInnerCollapseConfig
		return ret
	}
	return *o.CollapseConfig
}

// GetCollapseConfigOk returns a tuple with the CollapseConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequest) GetCollapseConfigOk() (*GetProductConfigResponseProductConfigInnerCollapseConfig, bool) {
	if o == nil || IsNil(o.CollapseConfig) {
		return nil, false
	}
	return o.CollapseConfig, true
}

// HasCollapseConfig returns a boolean if a field has been set.
func (o *UpdateProductConfigRequest) HasCollapseConfig() bool {
	if o != nil && !IsNil(o.CollapseConfig) {
		return true
	}

	return false
}

// SetCollapseConfig gets a reference to the given GetProductConfigResponseProductConfigInnerCollapseConfig and assigns it to the CollapseConfig field.
func (o *UpdateProductConfigRequest) SetCollapseConfig(v GetProductConfigResponseProductConfigInnerCollapseConfig) {
	o.CollapseConfig = &v
}

// GetCollapseTriggerConfig returns the CollapseTriggerConfig field value if set, zero value otherwise.
func (o *UpdateProductConfigRequest) GetCollapseTriggerConfig() GetProductConfigResponseProductConfigInnerCollapseTriggerConfig {
	if o == nil || IsNil(o.CollapseTriggerConfig) {
		var ret GetProductConfigResponseProductConfigInnerCollapseTriggerConfig
		return ret
	}
	return *o.CollapseTriggerConfig
}

// GetCollapseTriggerConfigOk returns a tuple with the CollapseTriggerConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequest) GetCollapseTriggerConfigOk() (*GetProductConfigResponseProductConfigInnerCollapseTriggerConfig, bool) {
	if o == nil || IsNil(o.CollapseTriggerConfig) {
		return nil, false
	}
	return o.CollapseTriggerConfig, true
}

// HasCollapseTriggerConfig returns a boolean if a field has been set.
func (o *UpdateProductConfigRequest) HasCollapseTriggerConfig() bool {
	if o != nil && !IsNil(o.CollapseTriggerConfig) {
		return true
	}

	return false
}

// SetCollapseTriggerConfig gets a reference to the given GetProductConfigResponseProductConfigInnerCollapseTriggerConfig and assigns it to the CollapseTriggerConfig field.
func (o *UpdateProductConfigRequest) SetCollapseTriggerConfig(v GetProductConfigResponseProductConfigInnerCollapseTriggerConfig) {
	o.CollapseTriggerConfig = &v
}

// GetEdcPartialPresentations returns the EdcPartialPresentations field value if set, zero value otherwise.
func (o *UpdateProductConfigRequest) GetEdcPartialPresentations() GetProductConfigResponseProductConfigInnerEdcPartialPresentations {
	if o == nil || IsNil(o.EdcPartialPresentations) {
		var ret GetProductConfigResponseProductConfigInnerEdcPartialPresentations
		return ret
	}
	return *o.EdcPartialPresentations
}

// GetEdcPartialPresentationsOk returns a tuple with the EdcPartialPresentations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequest) GetEdcPartialPresentationsOk() (*GetProductConfigResponseProductConfigInnerEdcPartialPresentations, bool) {
	if o == nil || IsNil(o.EdcPartialPresentations) {
		return nil, false
	}
	return o.EdcPartialPresentations, true
}

// HasEdcPartialPresentations returns a boolean if a field has been set.
func (o *UpdateProductConfigRequest) HasEdcPartialPresentations() bool {
	if o != nil && !IsNil(o.EdcPartialPresentations) {
		return true
	}

	return false
}

// SetEdcPartialPresentations gets a reference to the given GetProductConfigResponseProductConfigInnerEdcPartialPresentations and assigns it to the EdcPartialPresentations field.
func (o *UpdateProductConfigRequest) SetEdcPartialPresentations(v GetProductConfigResponseProductConfigInnerEdcPartialPresentations) {
	o.EdcPartialPresentations = &v
}

// GetEdcTranscodeTemplate returns the EdcTranscodeTemplate field value if set, zero value otherwise.
func (o *UpdateProductConfigRequest) GetEdcTranscodeTemplate() map[string]interface{} {
	if o == nil || IsNil(o.EdcTranscodeTemplate) {
		var ret map[string]interface{}
		return ret
	}
	return *o.EdcTranscodeTemplate
}

// GetEdcTranscodeTemplateOk returns a tuple with the EdcTranscodeTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequest) GetEdcTranscodeTemplateOk() (*map[string]interface{}, bool) {
	if o == nil || IsNil(o.EdcTranscodeTemplate) {
		return nil, false
	}
	return o.EdcTranscodeTemplate, true
}

// HasEdcTranscodeTemplate returns a boolean if a field has been set.
func (o *UpdateProductConfigRequest) HasEdcTranscodeTemplate() bool {
	if o != nil && !IsNil(o.EdcTranscodeTemplate) {
		return true
	}

	return false
}

// SetEdcTranscodeTemplate gets a reference to the given map[string]interface{} and assigns it to the EdcTranscodeTemplate field.
func (o *UpdateProductConfigRequest) SetEdcTranscodeTemplate(v map[string]interface{}) {
	o.EdcTranscodeTemplate = &v
}

// GetNotifications returns the Notifications field value if set, zero value otherwise.
func (o *UpdateProductConfigRequest) GetNotifications() []GetProductConfigResponseProductConfigInnerNotificationsInner {
	if o == nil || IsNil(o.Notifications) {
		var ret []GetProductConfigResponseProductConfigInnerNotificationsInner
		return ret
	}
	return o.Notifications
}

// GetNotificationsOk returns a tuple with the Notifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequest) GetNotificationsOk() ([]GetProductConfigResponseProductConfigInnerNotificationsInner, bool) {
	if o == nil || IsNil(o.Notifications) {
		return nil, false
	}
	return o.Notifications, true
}

// HasNotifications returns a boolean if a field has been set.
func (o *UpdateProductConfigRequest) HasNotifications() bool {
	if o != nil && !IsNil(o.Notifications) {
		return true
	}

	return false
}

// SetNotifications gets a reference to the given []GetProductConfigResponseProductConfigInnerNotificationsInner and assigns it to the Notifications field.
func (o *UpdateProductConfigRequest) SetNotifications(v []GetProductConfigResponseProductConfigInnerNotificationsInner) {
	o.Notifications = v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *UpdateProductConfigRequest) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequest) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *UpdateProductConfigRequest) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *UpdateProductConfigRequest) SetRegion(v string) {
	o.Region = &v
}

// GetStore returns the Store field value if set, zero value otherwise.
func (o *UpdateProductConfigRequest) GetStore() string {
	if o == nil || IsNil(o.Store) {
		var ret string
		return ret
	}
	return *o.Store
}

// GetStoreOk returns a tuple with the Store field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequest) GetStoreOk() (*string, bool) {
	if o == nil || IsNil(o.Store) {
		return nil, false
	}
	return o.Store, true
}

// HasStore returns a boolean if a field has been set.
func (o *UpdateProductConfigRequest) HasStore() bool {
	if o != nil && !IsNil(o.Store) {
		return true
	}

	return false
}

// SetStore gets a reference to the given string and assigns it to the Store field.
func (o *UpdateProductConfigRequest) SetStore(v string) {
	o.Store = &v
}

// GetWorkflowConfig returns the WorkflowConfig field value if set, zero value otherwise.
func (o *UpdateProductConfigRequest) GetWorkflowConfig() GetProductConfigResponseProductConfigInnerWorkflowConfig {
	if o == nil || IsNil(o.WorkflowConfig) {
		var ret GetProductConfigResponseProductConfigInnerWorkflowConfig
		return ret
	}
	return *o.WorkflowConfig
}

// GetWorkflowConfigOk returns a tuple with the WorkflowConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequest) GetWorkflowConfigOk() (*GetProductConfigResponseProductConfigInnerWorkflowConfig, bool) {
	if o == nil || IsNil(o.WorkflowConfig) {
		return nil, false
	}
	return o.WorkflowConfig, true
}

// HasWorkflowConfig returns a boolean if a field has been set.
func (o *UpdateProductConfigRequest) HasWorkflowConfig() bool {
	if o != nil && !IsNil(o.WorkflowConfig) {
		return true
	}

	return false
}

// SetWorkflowConfig gets a reference to the given GetProductConfigResponseProductConfigInnerWorkflowConfig and assigns it to the WorkflowConfig field.
func (o *UpdateProductConfigRequest) SetWorkflowConfig(v GetProductConfigResponseProductConfigInnerWorkflowConfig) {
	o.WorkflowConfig = &v
}

func (o UpdateProductConfigRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProductConfigRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	if !IsNil(o.ArchiveSettings) {
		toSerialize["archive_settings"] = o.ArchiveSettings
	}
	if !IsNil(o.CollapseConfig) {
		toSerialize["collapse_config"] = o.CollapseConfig
	}
	if !IsNil(o.CollapseTriggerConfig) {
		toSerialize["collapse_trigger_config"] = o.CollapseTriggerConfig
	}
	if !IsNil(o.EdcPartialPresentations) {
		toSerialize["edc_partial_presentations"] = o.EdcPartialPresentations
	}
	if !IsNil(o.EdcTranscodeTemplate) {
		toSerialize["edc_transcode_template"] = o.EdcTranscodeTemplate
	}
	if !IsNil(o.Notifications) {
		toSerialize["notifications"] = o.Notifications
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.Store) {
		toSerialize["store"] = o.Store
	}
	if !IsNil(o.WorkflowConfig) {
		toSerialize["workflow_config"] = o.WorkflowConfig
	}
	return toSerialize, nil
}

type NullableUpdateProductConfigRequest struct {
	value *UpdateProductConfigRequest
	isSet bool
}

func (v NullableUpdateProductConfigRequest) Get() *UpdateProductConfigRequest {
	return v.value
}

func (v *NullableUpdateProductConfigRequest) Set(val *UpdateProductConfigRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProductConfigRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProductConfigRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProductConfigRequest(val *UpdateProductConfigRequest) *NullableUpdateProductConfigRequest {
	return &NullableUpdateProductConfigRequest{value: val, isSet: true}
}

func (v NullableUpdateProductConfigRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProductConfigRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


