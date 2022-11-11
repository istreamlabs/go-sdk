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

// PutChannelRequest struct for PutChannelRequest
type PutChannelRequest struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string `json:"$schema,omitempty"`
	// Date and time the channel was created.
	Created *time.Time `json:"created,omitempty"`
	// Desired running state for a channel.
	DesiredState *string `json:"desired_state,omitempty"`
	// External Channel ID provided at channel creation time
	Id *string `json:"id,omitempty"`
	Ingest PutChannelRequestIngest `json:"ingest"`
	// Optional labels for a channel. Any included labels must be at least 1 character long, but no greater than 256 characters. The maximum number of labels is 10.
	Labels []string `json:"labels,omitempty"`
	// Date and time the channel was last modified.
	Modified *time.Time `json:"modified,omitempty"`
	// A friendly human-readable name for the channel. This will get displayed in user interfaces.
	Name *string `json:"name,omitempty"`
	Organization *string `json:"organization,omitempty"`
	Packaging *ChannelPackaging `json:"packaging,omitempty"`
	Publishing *ChannelPublishing `json:"publishing,omitempty"`
	// Region represents the general geolocation for transcoding and stream egress from iStreamPlanet. If no region is provided at channel creation time, then 'US_WEST' is used.
	Region *string `json:"region,omitempty"`
	// If the ResourceClass is unspecified the channel will default to run in the 'DYNAMIC' ResourceClass. Note that changing the ResourceClass for a running channel is supported and will be performed with no downtime.
	ResourceClass *string `json:"resource_class,omitempty"`
	// Self link for the channel.
	Self *string `json:"self,omitempty"`
	Signaling *ChannelSignaling `json:"signaling,omitempty"`
	Tags *ChannelTags `json:"tags,omitempty"`
	Transcode *ChannelTranscode `json:"transcode,omitempty"`
}

// NewPutChannelRequest instantiates a new PutChannelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPutChannelRequest(ingest PutChannelRequestIngest) *PutChannelRequest {
	this := PutChannelRequest{}
	this.Ingest = ingest
	return &this
}

// NewPutChannelRequestWithDefaults instantiates a new PutChannelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPutChannelRequestWithDefaults() *PutChannelRequest {
	this := PutChannelRequest{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *PutChannelRequest) GetSchema() string {
	if o == nil || o.Schema == nil {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutChannelRequest) GetSchemaOk() (*string, bool) {
	if o == nil || o.Schema == nil {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *PutChannelRequest) HasSchema() bool {
	if o != nil && o.Schema != nil {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *PutChannelRequest) SetSchema(v string) {
	o.Schema = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *PutChannelRequest) GetCreated() time.Time {
	if o == nil || o.Created == nil {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutChannelRequest) GetCreatedOk() (*time.Time, bool) {
	if o == nil || o.Created == nil {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *PutChannelRequest) HasCreated() bool {
	if o != nil && o.Created != nil {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *PutChannelRequest) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDesiredState returns the DesiredState field value if set, zero value otherwise.
func (o *PutChannelRequest) GetDesiredState() string {
	if o == nil || o.DesiredState == nil {
		var ret string
		return ret
	}
	return *o.DesiredState
}

// GetDesiredStateOk returns a tuple with the DesiredState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutChannelRequest) GetDesiredStateOk() (*string, bool) {
	if o == nil || o.DesiredState == nil {
		return nil, false
	}
	return o.DesiredState, true
}

// HasDesiredState returns a boolean if a field has been set.
func (o *PutChannelRequest) HasDesiredState() bool {
	if o != nil && o.DesiredState != nil {
		return true
	}

	return false
}

// SetDesiredState gets a reference to the given string and assigns it to the DesiredState field.
func (o *PutChannelRequest) SetDesiredState(v string) {
	o.DesiredState = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PutChannelRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutChannelRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PutChannelRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PutChannelRequest) SetId(v string) {
	o.Id = &v
}

// GetIngest returns the Ingest field value
func (o *PutChannelRequest) GetIngest() PutChannelRequestIngest {
	if o == nil {
		var ret PutChannelRequestIngest
		return ret
	}

	return o.Ingest
}

// GetIngestOk returns a tuple with the Ingest field value
// and a boolean to check if the value has been set.
func (o *PutChannelRequest) GetIngestOk() (*PutChannelRequestIngest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ingest, true
}

// SetIngest sets field value
func (o *PutChannelRequest) SetIngest(v PutChannelRequestIngest) {
	o.Ingest = v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *PutChannelRequest) GetLabels() []string {
	if o == nil || o.Labels == nil {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutChannelRequest) GetLabelsOk() ([]string, bool) {
	if o == nil || o.Labels == nil {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *PutChannelRequest) HasLabels() bool {
	if o != nil && o.Labels != nil {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *PutChannelRequest) SetLabels(v []string) {
	o.Labels = v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *PutChannelRequest) GetModified() time.Time {
	if o == nil || o.Modified == nil {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutChannelRequest) GetModifiedOk() (*time.Time, bool) {
	if o == nil || o.Modified == nil {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *PutChannelRequest) HasModified() bool {
	if o != nil && o.Modified != nil {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *PutChannelRequest) SetModified(v time.Time) {
	o.Modified = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PutChannelRequest) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutChannelRequest) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PutChannelRequest) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PutChannelRequest) SetName(v string) {
	o.Name = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *PutChannelRequest) GetOrganization() string {
	if o == nil || o.Organization == nil {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutChannelRequest) GetOrganizationOk() (*string, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *PutChannelRequest) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *PutChannelRequest) SetOrganization(v string) {
	o.Organization = &v
}

// GetPackaging returns the Packaging field value if set, zero value otherwise.
func (o *PutChannelRequest) GetPackaging() ChannelPackaging {
	if o == nil || o.Packaging == nil {
		var ret ChannelPackaging
		return ret
	}
	return *o.Packaging
}

// GetPackagingOk returns a tuple with the Packaging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutChannelRequest) GetPackagingOk() (*ChannelPackaging, bool) {
	if o == nil || o.Packaging == nil {
		return nil, false
	}
	return o.Packaging, true
}

// HasPackaging returns a boolean if a field has been set.
func (o *PutChannelRequest) HasPackaging() bool {
	if o != nil && o.Packaging != nil {
		return true
	}

	return false
}

// SetPackaging gets a reference to the given ChannelPackaging and assigns it to the Packaging field.
func (o *PutChannelRequest) SetPackaging(v ChannelPackaging) {
	o.Packaging = &v
}

// GetPublishing returns the Publishing field value if set, zero value otherwise.
func (o *PutChannelRequest) GetPublishing() ChannelPublishing {
	if o == nil || o.Publishing == nil {
		var ret ChannelPublishing
		return ret
	}
	return *o.Publishing
}

// GetPublishingOk returns a tuple with the Publishing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutChannelRequest) GetPublishingOk() (*ChannelPublishing, bool) {
	if o == nil || o.Publishing == nil {
		return nil, false
	}
	return o.Publishing, true
}

// HasPublishing returns a boolean if a field has been set.
func (o *PutChannelRequest) HasPublishing() bool {
	if o != nil && o.Publishing != nil {
		return true
	}

	return false
}

// SetPublishing gets a reference to the given ChannelPublishing and assigns it to the Publishing field.
func (o *PutChannelRequest) SetPublishing(v ChannelPublishing) {
	o.Publishing = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *PutChannelRequest) GetRegion() string {
	if o == nil || o.Region == nil {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutChannelRequest) GetRegionOk() (*string, bool) {
	if o == nil || o.Region == nil {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *PutChannelRequest) HasRegion() bool {
	if o != nil && o.Region != nil {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *PutChannelRequest) SetRegion(v string) {
	o.Region = &v
}

// GetResourceClass returns the ResourceClass field value if set, zero value otherwise.
func (o *PutChannelRequest) GetResourceClass() string {
	if o == nil || o.ResourceClass == nil {
		var ret string
		return ret
	}
	return *o.ResourceClass
}

// GetResourceClassOk returns a tuple with the ResourceClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutChannelRequest) GetResourceClassOk() (*string, bool) {
	if o == nil || o.ResourceClass == nil {
		return nil, false
	}
	return o.ResourceClass, true
}

// HasResourceClass returns a boolean if a field has been set.
func (o *PutChannelRequest) HasResourceClass() bool {
	if o != nil && o.ResourceClass != nil {
		return true
	}

	return false
}

// SetResourceClass gets a reference to the given string and assigns it to the ResourceClass field.
func (o *PutChannelRequest) SetResourceClass(v string) {
	o.ResourceClass = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *PutChannelRequest) GetSelf() string {
	if o == nil || o.Self == nil {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutChannelRequest) GetSelfOk() (*string, bool) {
	if o == nil || o.Self == nil {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *PutChannelRequest) HasSelf() bool {
	if o != nil && o.Self != nil {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *PutChannelRequest) SetSelf(v string) {
	o.Self = &v
}

// GetSignaling returns the Signaling field value if set, zero value otherwise.
func (o *PutChannelRequest) GetSignaling() ChannelSignaling {
	if o == nil || o.Signaling == nil {
		var ret ChannelSignaling
		return ret
	}
	return *o.Signaling
}

// GetSignalingOk returns a tuple with the Signaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutChannelRequest) GetSignalingOk() (*ChannelSignaling, bool) {
	if o == nil || o.Signaling == nil {
		return nil, false
	}
	return o.Signaling, true
}

// HasSignaling returns a boolean if a field has been set.
func (o *PutChannelRequest) HasSignaling() bool {
	if o != nil && o.Signaling != nil {
		return true
	}

	return false
}

// SetSignaling gets a reference to the given ChannelSignaling and assigns it to the Signaling field.
func (o *PutChannelRequest) SetSignaling(v ChannelSignaling) {
	o.Signaling = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PutChannelRequest) GetTags() ChannelTags {
	if o == nil || o.Tags == nil {
		var ret ChannelTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutChannelRequest) GetTagsOk() (*ChannelTags, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PutChannelRequest) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given ChannelTags and assigns it to the Tags field.
func (o *PutChannelRequest) SetTags(v ChannelTags) {
	o.Tags = &v
}

// GetTranscode returns the Transcode field value if set, zero value otherwise.
func (o *PutChannelRequest) GetTranscode() ChannelTranscode {
	if o == nil || o.Transcode == nil {
		var ret ChannelTranscode
		return ret
	}
	return *o.Transcode
}

// GetTranscodeOk returns a tuple with the Transcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PutChannelRequest) GetTranscodeOk() (*ChannelTranscode, bool) {
	if o == nil || o.Transcode == nil {
		return nil, false
	}
	return o.Transcode, true
}

// HasTranscode returns a boolean if a field has been set.
func (o *PutChannelRequest) HasTranscode() bool {
	if o != nil && o.Transcode != nil {
		return true
	}

	return false
}

// SetTranscode gets a reference to the given ChannelTranscode and assigns it to the Transcode field.
func (o *PutChannelRequest) SetTranscode(v ChannelTranscode) {
	o.Transcode = &v
}

func (o PutChannelRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Schema != nil {
		toSerialize["$schema"] = o.Schema
	}
	if o.Created != nil {
		toSerialize["created"] = o.Created
	}
	if o.DesiredState != nil {
		toSerialize["desired_state"] = o.DesiredState
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["ingest"] = o.Ingest
	}
	if o.Labels != nil {
		toSerialize["labels"] = o.Labels
	}
	if o.Modified != nil {
		toSerialize["modified"] = o.Modified
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Organization != nil {
		toSerialize["organization"] = o.Organization
	}
	if o.Packaging != nil {
		toSerialize["packaging"] = o.Packaging
	}
	if o.Publishing != nil {
		toSerialize["publishing"] = o.Publishing
	}
	if o.Region != nil {
		toSerialize["region"] = o.Region
	}
	if o.ResourceClass != nil {
		toSerialize["resource_class"] = o.ResourceClass
	}
	if o.Self != nil {
		toSerialize["self"] = o.Self
	}
	if o.Signaling != nil {
		toSerialize["signaling"] = o.Signaling
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.Transcode != nil {
		toSerialize["transcode"] = o.Transcode
	}
	return json.Marshal(toSerialize)
}

type NullablePutChannelRequest struct {
	value *PutChannelRequest
	isSet bool
}

func (v NullablePutChannelRequest) Get() *PutChannelRequest {
	return v.value
}

func (v *NullablePutChannelRequest) Set(val *PutChannelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePutChannelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePutChannelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePutChannelRequest(val *PutChannelRequest) *NullablePutChannelRequest {
	return &NullablePutChannelRequest{value: val, isSet: true}
}

func (v NullablePutChannelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePutChannelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


