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

// checks if the UpdateProductConfigRequestNotificationsInnerNotification type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateProductConfigRequestNotificationsInnerNotification{}

// UpdateProductConfigRequestNotificationsInnerNotification struct for UpdateProductConfigRequestNotificationsInnerNotification
type UpdateProductConfigRequestNotificationsInnerNotification struct {
	NotificationHostname *string `json:"notification_hostname,omitempty"`
	NotificationSettings string `json:"notification_settings"`
	NotificationType string `json:"notification_type"`
	UseChannelApi *bool `json:"use_channel_api,omitempty"`
}

// NewUpdateProductConfigRequestNotificationsInnerNotification instantiates a new UpdateProductConfigRequestNotificationsInnerNotification object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateProductConfigRequestNotificationsInnerNotification(notificationSettings string, notificationType string) *UpdateProductConfigRequestNotificationsInnerNotification {
	this := UpdateProductConfigRequestNotificationsInnerNotification{}
	this.NotificationSettings = notificationSettings
	this.NotificationType = notificationType
	return &this
}

// NewUpdateProductConfigRequestNotificationsInnerNotificationWithDefaults instantiates a new UpdateProductConfigRequestNotificationsInnerNotification object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateProductConfigRequestNotificationsInnerNotificationWithDefaults() *UpdateProductConfigRequestNotificationsInnerNotification {
	this := UpdateProductConfigRequestNotificationsInnerNotification{}
	return &this
}

// GetNotificationHostname returns the NotificationHostname field value if set, zero value otherwise.
func (o *UpdateProductConfigRequestNotificationsInnerNotification) GetNotificationHostname() string {
	if o == nil || IsNil(o.NotificationHostname) {
		var ret string
		return ret
	}
	return *o.NotificationHostname
}

// GetNotificationHostnameOk returns a tuple with the NotificationHostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestNotificationsInnerNotification) GetNotificationHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.NotificationHostname) {
		return nil, false
	}
	return o.NotificationHostname, true
}

// HasNotificationHostname returns a boolean if a field has been set.
func (o *UpdateProductConfigRequestNotificationsInnerNotification) HasNotificationHostname() bool {
	if o != nil && !IsNil(o.NotificationHostname) {
		return true
	}

	return false
}

// SetNotificationHostname gets a reference to the given string and assigns it to the NotificationHostname field.
func (o *UpdateProductConfigRequestNotificationsInnerNotification) SetNotificationHostname(v string) {
	o.NotificationHostname = &v
}

// GetNotificationSettings returns the NotificationSettings field value
func (o *UpdateProductConfigRequestNotificationsInnerNotification) GetNotificationSettings() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotificationSettings
}

// GetNotificationSettingsOk returns a tuple with the NotificationSettings field value
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestNotificationsInnerNotification) GetNotificationSettingsOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationSettings, true
}

// SetNotificationSettings sets field value
func (o *UpdateProductConfigRequestNotificationsInnerNotification) SetNotificationSettings(v string) {
	o.NotificationSettings = v
}

// GetNotificationType returns the NotificationType field value
func (o *UpdateProductConfigRequestNotificationsInnerNotification) GetNotificationType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestNotificationsInnerNotification) GetNotificationTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NotificationType, true
}

// SetNotificationType sets field value
func (o *UpdateProductConfigRequestNotificationsInnerNotification) SetNotificationType(v string) {
	o.NotificationType = v
}

// GetUseChannelApi returns the UseChannelApi field value if set, zero value otherwise.
func (o *UpdateProductConfigRequestNotificationsInnerNotification) GetUseChannelApi() bool {
	if o == nil || IsNil(o.UseChannelApi) {
		var ret bool
		return ret
	}
	return *o.UseChannelApi
}

// GetUseChannelApiOk returns a tuple with the UseChannelApi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateProductConfigRequestNotificationsInnerNotification) GetUseChannelApiOk() (*bool, bool) {
	if o == nil || IsNil(o.UseChannelApi) {
		return nil, false
	}
	return o.UseChannelApi, true
}

// HasUseChannelApi returns a boolean if a field has been set.
func (o *UpdateProductConfigRequestNotificationsInnerNotification) HasUseChannelApi() bool {
	if o != nil && !IsNil(o.UseChannelApi) {
		return true
	}

	return false
}

// SetUseChannelApi gets a reference to the given bool and assigns it to the UseChannelApi field.
func (o *UpdateProductConfigRequestNotificationsInnerNotification) SetUseChannelApi(v bool) {
	o.UseChannelApi = &v
}

func (o UpdateProductConfigRequestNotificationsInnerNotification) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateProductConfigRequestNotificationsInnerNotification) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NotificationHostname) {
		toSerialize["notification_hostname"] = o.NotificationHostname
	}
	toSerialize["notification_settings"] = o.NotificationSettings
	toSerialize["notification_type"] = o.NotificationType
	if !IsNil(o.UseChannelApi) {
		toSerialize["use_channel_api"] = o.UseChannelApi
	}
	return toSerialize, nil
}

type NullableUpdateProductConfigRequestNotificationsInnerNotification struct {
	value *UpdateProductConfigRequestNotificationsInnerNotification
	isSet bool
}

func (v NullableUpdateProductConfigRequestNotificationsInnerNotification) Get() *UpdateProductConfigRequestNotificationsInnerNotification {
	return v.value
}

func (v *NullableUpdateProductConfigRequestNotificationsInnerNotification) Set(val *UpdateProductConfigRequestNotificationsInnerNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateProductConfigRequestNotificationsInnerNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateProductConfigRequestNotificationsInnerNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateProductConfigRequestNotificationsInnerNotification(val *UpdateProductConfigRequestNotificationsInnerNotification) *NullableUpdateProductConfigRequestNotificationsInnerNotification {
	return &NullableUpdateProductConfigRequestNotificationsInnerNotification{value: val, isSet: true}
}

func (v NullableUpdateProductConfigRequestNotificationsInnerNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateProductConfigRequestNotificationsInnerNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


