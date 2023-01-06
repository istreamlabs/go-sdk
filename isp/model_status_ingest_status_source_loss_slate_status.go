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

// StatusIngestStatusSourceLossSlateStatus Source-loss slate status.
type StatusIngestStatusSourceLossSlateStatus struct {
	// Reports whether this is the active source - i.e. the source being transcoded.
	Active *bool `json:"active,omitempty"`
	// Continuity counter error total since the transcoder pod started.
	CcErrors *int64 `json:"cc_errors,omitempty"`
	// The last time data was received for this source.
	LastDataReceived *time.Time `json:"last_data_received,omitempty"`
	// Reports whether this source was pinned to be always preferred (if available).
	Pinned *bool `json:"pinned,omitempty"`
	Pmt *StatusIngestStatusPrimaryStatusPmt `json:"pmt,omitempty"`
	// A measure of the source's quality if available. Zero is the perfect score. The higher the score, the worst the quality.
	QualityScore *float64 `json:"quality_score,omitempty"`
	// If unset (empty string), the source is available, and could be made active if necessary. Otherwise, it contains a message indicating why the source is unavailable. E.g. 'No source AUs received'.
	UnavailableReason *string `json:"unavailable_reason,omitempty"`
}

// NewStatusIngestStatusSourceLossSlateStatus instantiates a new StatusIngestStatusSourceLossSlateStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusIngestStatusSourceLossSlateStatus() *StatusIngestStatusSourceLossSlateStatus {
	this := StatusIngestStatusSourceLossSlateStatus{}
	return &this
}

// NewStatusIngestStatusSourceLossSlateStatusWithDefaults instantiates a new StatusIngestStatusSourceLossSlateStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusIngestStatusSourceLossSlateStatusWithDefaults() *StatusIngestStatusSourceLossSlateStatus {
	this := StatusIngestStatusSourceLossSlateStatus{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *StatusIngestStatusSourceLossSlateStatus) GetActive() bool {
	if o == nil || o.Active == nil {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusIngestStatusSourceLossSlateStatus) GetActiveOk() (*bool, bool) {
	if o == nil || o.Active == nil {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *StatusIngestStatusSourceLossSlateStatus) HasActive() bool {
	if o != nil && o.Active != nil {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *StatusIngestStatusSourceLossSlateStatus) SetActive(v bool) {
	o.Active = &v
}

// GetCcErrors returns the CcErrors field value if set, zero value otherwise.
func (o *StatusIngestStatusSourceLossSlateStatus) GetCcErrors() int64 {
	if o == nil || o.CcErrors == nil {
		var ret int64
		return ret
	}
	return *o.CcErrors
}

// GetCcErrorsOk returns a tuple with the CcErrors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusIngestStatusSourceLossSlateStatus) GetCcErrorsOk() (*int64, bool) {
	if o == nil || o.CcErrors == nil {
		return nil, false
	}
	return o.CcErrors, true
}

// HasCcErrors returns a boolean if a field has been set.
func (o *StatusIngestStatusSourceLossSlateStatus) HasCcErrors() bool {
	if o != nil && o.CcErrors != nil {
		return true
	}

	return false
}

// SetCcErrors gets a reference to the given int64 and assigns it to the CcErrors field.
func (o *StatusIngestStatusSourceLossSlateStatus) SetCcErrors(v int64) {
	o.CcErrors = &v
}

// GetLastDataReceived returns the LastDataReceived field value if set, zero value otherwise.
func (o *StatusIngestStatusSourceLossSlateStatus) GetLastDataReceived() time.Time {
	if o == nil || o.LastDataReceived == nil {
		var ret time.Time
		return ret
	}
	return *o.LastDataReceived
}

// GetLastDataReceivedOk returns a tuple with the LastDataReceived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusIngestStatusSourceLossSlateStatus) GetLastDataReceivedOk() (*time.Time, bool) {
	if o == nil || o.LastDataReceived == nil {
		return nil, false
	}
	return o.LastDataReceived, true
}

// HasLastDataReceived returns a boolean if a field has been set.
func (o *StatusIngestStatusSourceLossSlateStatus) HasLastDataReceived() bool {
	if o != nil && o.LastDataReceived != nil {
		return true
	}

	return false
}

// SetLastDataReceived gets a reference to the given time.Time and assigns it to the LastDataReceived field.
func (o *StatusIngestStatusSourceLossSlateStatus) SetLastDataReceived(v time.Time) {
	o.LastDataReceived = &v
}

// GetPinned returns the Pinned field value if set, zero value otherwise.
func (o *StatusIngestStatusSourceLossSlateStatus) GetPinned() bool {
	if o == nil || o.Pinned == nil {
		var ret bool
		return ret
	}
	return *o.Pinned
}

// GetPinnedOk returns a tuple with the Pinned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusIngestStatusSourceLossSlateStatus) GetPinnedOk() (*bool, bool) {
	if o == nil || o.Pinned == nil {
		return nil, false
	}
	return o.Pinned, true
}

// HasPinned returns a boolean if a field has been set.
func (o *StatusIngestStatusSourceLossSlateStatus) HasPinned() bool {
	if o != nil && o.Pinned != nil {
		return true
	}

	return false
}

// SetPinned gets a reference to the given bool and assigns it to the Pinned field.
func (o *StatusIngestStatusSourceLossSlateStatus) SetPinned(v bool) {
	o.Pinned = &v
}

// GetPmt returns the Pmt field value if set, zero value otherwise.
func (o *StatusIngestStatusSourceLossSlateStatus) GetPmt() StatusIngestStatusPrimaryStatusPmt {
	if o == nil || o.Pmt == nil {
		var ret StatusIngestStatusPrimaryStatusPmt
		return ret
	}
	return *o.Pmt
}

// GetPmtOk returns a tuple with the Pmt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusIngestStatusSourceLossSlateStatus) GetPmtOk() (*StatusIngestStatusPrimaryStatusPmt, bool) {
	if o == nil || o.Pmt == nil {
		return nil, false
	}
	return o.Pmt, true
}

// HasPmt returns a boolean if a field has been set.
func (o *StatusIngestStatusSourceLossSlateStatus) HasPmt() bool {
	if o != nil && o.Pmt != nil {
		return true
	}

	return false
}

// SetPmt gets a reference to the given StatusIngestStatusPrimaryStatusPmt and assigns it to the Pmt field.
func (o *StatusIngestStatusSourceLossSlateStatus) SetPmt(v StatusIngestStatusPrimaryStatusPmt) {
	o.Pmt = &v
}

// GetQualityScore returns the QualityScore field value if set, zero value otherwise.
func (o *StatusIngestStatusSourceLossSlateStatus) GetQualityScore() float64 {
	if o == nil || o.QualityScore == nil {
		var ret float64
		return ret
	}
	return *o.QualityScore
}

// GetQualityScoreOk returns a tuple with the QualityScore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusIngestStatusSourceLossSlateStatus) GetQualityScoreOk() (*float64, bool) {
	if o == nil || o.QualityScore == nil {
		return nil, false
	}
	return o.QualityScore, true
}

// HasQualityScore returns a boolean if a field has been set.
func (o *StatusIngestStatusSourceLossSlateStatus) HasQualityScore() bool {
	if o != nil && o.QualityScore != nil {
		return true
	}

	return false
}

// SetQualityScore gets a reference to the given float64 and assigns it to the QualityScore field.
func (o *StatusIngestStatusSourceLossSlateStatus) SetQualityScore(v float64) {
	o.QualityScore = &v
}

// GetUnavailableReason returns the UnavailableReason field value if set, zero value otherwise.
func (o *StatusIngestStatusSourceLossSlateStatus) GetUnavailableReason() string {
	if o == nil || o.UnavailableReason == nil {
		var ret string
		return ret
	}
	return *o.UnavailableReason
}

// GetUnavailableReasonOk returns a tuple with the UnavailableReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusIngestStatusSourceLossSlateStatus) GetUnavailableReasonOk() (*string, bool) {
	if o == nil || o.UnavailableReason == nil {
		return nil, false
	}
	return o.UnavailableReason, true
}

// HasUnavailableReason returns a boolean if a field has been set.
func (o *StatusIngestStatusSourceLossSlateStatus) HasUnavailableReason() bool {
	if o != nil && o.UnavailableReason != nil {
		return true
	}

	return false
}

// SetUnavailableReason gets a reference to the given string and assigns it to the UnavailableReason field.
func (o *StatusIngestStatusSourceLossSlateStatus) SetUnavailableReason(v string) {
	o.UnavailableReason = &v
}

func (o StatusIngestStatusSourceLossSlateStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Active != nil {
		toSerialize["active"] = o.Active
	}
	if o.CcErrors != nil {
		toSerialize["cc_errors"] = o.CcErrors
	}
	if o.LastDataReceived != nil {
		toSerialize["last_data_received"] = o.LastDataReceived
	}
	if o.Pinned != nil {
		toSerialize["pinned"] = o.Pinned
	}
	if o.Pmt != nil {
		toSerialize["pmt"] = o.Pmt
	}
	if o.QualityScore != nil {
		toSerialize["quality_score"] = o.QualityScore
	}
	if o.UnavailableReason != nil {
		toSerialize["unavailable_reason"] = o.UnavailableReason
	}
	return json.Marshal(toSerialize)
}

type NullableStatusIngestStatusSourceLossSlateStatus struct {
	value *StatusIngestStatusSourceLossSlateStatus
	isSet bool
}

func (v NullableStatusIngestStatusSourceLossSlateStatus) Get() *StatusIngestStatusSourceLossSlateStatus {
	return v.value
}

func (v *NullableStatusIngestStatusSourceLossSlateStatus) Set(val *StatusIngestStatusSourceLossSlateStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusIngestStatusSourceLossSlateStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusIngestStatusSourceLossSlateStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusIngestStatusSourceLossSlateStatus(val *StatusIngestStatusSourceLossSlateStatus) *NullableStatusIngestStatusSourceLossSlateStatus {
	return &NullableStatusIngestStatusSourceLossSlateStatus{value: val, isSet: true}
}

func (v NullableStatusIngestStatusSourceLossSlateStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusIngestStatusSourceLossSlateStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


