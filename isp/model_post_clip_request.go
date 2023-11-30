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

// checks if the PostClipRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostClipRequest{}

// PostClipRequest struct for PostClipRequest
type PostClipRequest struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string `json:"$schema,omitempty" format:"uri" doc:"An optional URL to a JSON Schema document describing this resource"`
	// Identifer that is carried through archive and collapse notifications for the clip creation
	CorrelationId *string `json:"correlation_id,omitempty" doc:"Identifer that is carried through archive and collapse notifications for the clip creation"`
	// If true, creates the mp4. Default: false
	CreateMp4 *bool `json:"create_mp4,omitempty" doc:"If true, creates the mp4. Default: false"`
	// Description of the clip being created
	Description *string `json:"description,omitempty" maxLength:"80" doc:"Description of the clip being created"`
	// End timestamp in RFC3339Nano format
	End string `json:"end" doc:"End timestamp in RFC3339Nano format"`
	// If set, overrides the mp4 file path for archiving.
	Mp4FilePath *string `json:"mp4_file_path,omitempty" maxLength:"1024" doc:"If set, overrides the mp4 file path for archiving."`
	// Overwrite existing clip. Default: false
	Overwrite *bool `json:"overwrite,omitempty" doc:"Overwrite existing clip. Default: false"`
	// Start timestamp in RFC3339Nano format
	Start string `json:"start" doc:"Start timestamp in RFC3339Nano format"`
}

// NewPostClipRequest instantiates a new PostClipRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostClipRequest(end string, start string) *PostClipRequest {
	this := PostClipRequest{}
	this.End = end
	this.Start = start
	return &this
}

// NewPostClipRequestWithDefaults instantiates a new PostClipRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostClipRequestWithDefaults() *PostClipRequest {
	this := PostClipRequest{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *PostClipRequest) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostClipRequest) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *PostClipRequest) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *PostClipRequest) SetSchema(v string) {
	o.Schema = &v
}

// GetCorrelationId returns the CorrelationId field value if set, zero value otherwise.
func (o *PostClipRequest) GetCorrelationId() string {
	if o == nil || IsNil(o.CorrelationId) {
		var ret string
		return ret
	}
	return *o.CorrelationId
}

// GetCorrelationIdOk returns a tuple with the CorrelationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostClipRequest) GetCorrelationIdOk() (*string, bool) {
	if o == nil || IsNil(o.CorrelationId) {
		return nil, false
	}
	return o.CorrelationId, true
}

// HasCorrelationId returns a boolean if a field has been set.
func (o *PostClipRequest) HasCorrelationId() bool {
	if o != nil && !IsNil(o.CorrelationId) {
		return true
	}

	return false
}

// SetCorrelationId gets a reference to the given string and assigns it to the CorrelationId field.
func (o *PostClipRequest) SetCorrelationId(v string) {
	o.CorrelationId = &v
}

// GetCreateMp4 returns the CreateMp4 field value if set, zero value otherwise.
func (o *PostClipRequest) GetCreateMp4() bool {
	if o == nil || IsNil(o.CreateMp4) {
		var ret bool
		return ret
	}
	return *o.CreateMp4
}

// GetCreateMp4Ok returns a tuple with the CreateMp4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostClipRequest) GetCreateMp4Ok() (*bool, bool) {
	if o == nil || IsNil(o.CreateMp4) {
		return nil, false
	}
	return o.CreateMp4, true
}

// HasCreateMp4 returns a boolean if a field has been set.
func (o *PostClipRequest) HasCreateMp4() bool {
	if o != nil && !IsNil(o.CreateMp4) {
		return true
	}

	return false
}

// SetCreateMp4 gets a reference to the given bool and assigns it to the CreateMp4 field.
func (o *PostClipRequest) SetCreateMp4(v bool) {
	o.CreateMp4 = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PostClipRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostClipRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PostClipRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PostClipRequest) SetDescription(v string) {
	o.Description = &v
}

// GetEnd returns the End field value
func (o *PostClipRequest) GetEnd() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.End
}

// GetEndOk returns a tuple with the End field value
// and a boolean to check if the value has been set.
func (o *PostClipRequest) GetEndOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.End, true
}

// SetEnd sets field value
func (o *PostClipRequest) SetEnd(v string) {
	o.End = v
}

// GetMp4FilePath returns the Mp4FilePath field value if set, zero value otherwise.
func (o *PostClipRequest) GetMp4FilePath() string {
	if o == nil || IsNil(o.Mp4FilePath) {
		var ret string
		return ret
	}
	return *o.Mp4FilePath
}

// GetMp4FilePathOk returns a tuple with the Mp4FilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostClipRequest) GetMp4FilePathOk() (*string, bool) {
	if o == nil || IsNil(o.Mp4FilePath) {
		return nil, false
	}
	return o.Mp4FilePath, true
}

// HasMp4FilePath returns a boolean if a field has been set.
func (o *PostClipRequest) HasMp4FilePath() bool {
	if o != nil && !IsNil(o.Mp4FilePath) {
		return true
	}

	return false
}

// SetMp4FilePath gets a reference to the given string and assigns it to the Mp4FilePath field.
func (o *PostClipRequest) SetMp4FilePath(v string) {
	o.Mp4FilePath = &v
}

// GetOverwrite returns the Overwrite field value if set, zero value otherwise.
func (o *PostClipRequest) GetOverwrite() bool {
	if o == nil || IsNil(o.Overwrite) {
		var ret bool
		return ret
	}
	return *o.Overwrite
}

// GetOverwriteOk returns a tuple with the Overwrite field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PostClipRequest) GetOverwriteOk() (*bool, bool) {
	if o == nil || IsNil(o.Overwrite) {
		return nil, false
	}
	return o.Overwrite, true
}

// HasOverwrite returns a boolean if a field has been set.
func (o *PostClipRequest) HasOverwrite() bool {
	if o != nil && !IsNil(o.Overwrite) {
		return true
	}

	return false
}

// SetOverwrite gets a reference to the given bool and assigns it to the Overwrite field.
func (o *PostClipRequest) SetOverwrite(v bool) {
	o.Overwrite = &v
}

// GetStart returns the Start field value
func (o *PostClipRequest) GetStart() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Start
}

// GetStartOk returns a tuple with the Start field value
// and a boolean to check if the value has been set.
func (o *PostClipRequest) GetStartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Start, true
}

// SetStart sets field value
func (o *PostClipRequest) SetStart(v string) {
	o.Start = v
}

func (o PostClipRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostClipRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	if !IsNil(o.CorrelationId) {
		toSerialize["correlation_id"] = o.CorrelationId
	}
	if !IsNil(o.CreateMp4) {
		toSerialize["create_mp4"] = o.CreateMp4
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["end"] = o.End
	if !IsNil(o.Mp4FilePath) {
		toSerialize["mp4_file_path"] = o.Mp4FilePath
	}
	if !IsNil(o.Overwrite) {
		toSerialize["overwrite"] = o.Overwrite
	}
	toSerialize["start"] = o.Start
	return toSerialize, nil
}

type NullablePostClipRequest struct {
	value *PostClipRequest
	isSet bool
}

func (v NullablePostClipRequest) Get() *PostClipRequest {
	return v.value
}

func (v *NullablePostClipRequest) Set(val *PostClipRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePostClipRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePostClipRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostClipRequest(val *PostClipRequest) *NullablePostClipRequest {
	return &NullablePostClipRequest{value: val, isSet: true}
}

func (v NullablePostClipRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostClipRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


