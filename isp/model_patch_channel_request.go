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

// checks if the PatchChannelRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchChannelRequest{}

// PatchChannelRequest struct for PatchChannelRequest
type PatchChannelRequest struct {
	// An optional URL to a JSON Schema document describing this resource
	Schema *string `json:"$schema,omitempty" format:"uri" doc:"An optional URL to a JSON Schema document describing this resource"`
	// Date and time the channel was created.
	Created *time.Time `json:"created,omitempty" format:"date-time" doc:"Date and time the channel was created."`
	// Desired running state for a channel.
	DesiredState *string `json:"desired_state,omitempty" enum:"ON,OFF" doc:"Desired running state for a channel."`
	// External Channel ID provided at channel creation time
	Id *string `json:"id,omitempty" minLength:"1" pattern:"/^([a-z0-9]+(-*[a-z0-9]+)*)$/" doc:"External Channel ID provided at channel creation time"`
	Ingest *PatchChannelRequestIngest `json:"ingest,omitempty"`
	// Optional labels for a channel. Any included labels must be at least 1 character long, but no greater than 256 characters. The maximum number of labels is 10.
	Labels []string `json:"labels,omitempty" maxItems:"10" doc:"Optional labels for a channel. Any included labels must be at least 1 character long, but no greater than 256 characters. The maximum number of labels is 10."`
	// Date and time the channel was last modified.
	Modified *time.Time `json:"modified,omitempty" format:"date-time" doc:"Date and time the channel was last modified."`
	// A friendly human-readable name for the channel. This will get displayed in user interfaces.
	Name *string `json:"name,omitempty" doc:"A friendly human-readable name for the channel. This will get displayed in user interfaces."`
	Organization *string `json:"organization,omitempty" minLength:"1"`
	Packaging *ChannelPackaging `json:"packaging,omitempty"`
	Publishing *ChannelPublishing `json:"publishing,omitempty"`
	// Region represents the general geolocation for transcoding and stream egress from iStreamPlanet. If no region is provided at channel creation time, then 'US_WEST' is used.
	Region *string `json:"region,omitempty" enum:"US_WEST,US_EAST" doc:"Region represents the general geolocation for transcoding and stream egress from iStreamPlanet. If no region is provided at channel creation time, then 'US_WEST' is used."`
	// If the ResourceClass is unspecified the channel will default to run in the 'DYNAMIC' ResourceClass. Note that changing the ResourceClass for a running channel is supported and will be performed with no downtime.
	ResourceClass *string `json:"resource_class,omitempty" enum:"DYNAMIC,STATIC" doc:"If the ResourceClass is unspecified the channel will default to run in the 'DYNAMIC' ResourceClass. Note that changing the ResourceClass for a running channel is supported and will be performed with no downtime."`
	// Self link for the channel.
	Self *string `json:"self,omitempty" format:"uri-reference" doc:"Self link for the channel."`
	Signaling *ChannelSignaling `json:"signaling,omitempty"`
	Tags *ChannelTags `json:"tags,omitempty"`
	Transcode *ChannelTranscode `json:"transcode,omitempty"`
}

// NewPatchChannelRequest instantiates a new PatchChannelRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchChannelRequest() *PatchChannelRequest {
	this := PatchChannelRequest{}
	return &this
}

// NewPatchChannelRequestWithDefaults instantiates a new PatchChannelRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchChannelRequestWithDefaults() *PatchChannelRequest {
	this := PatchChannelRequest{}
	return &this
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *PatchChannelRequest) GetSchema() string {
	if o == nil || IsNil(o.Schema) {
		var ret string
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchChannelRequest) GetSchemaOk() (*string, bool) {
	if o == nil || IsNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *PatchChannelRequest) HasSchema() bool {
	if o != nil && !IsNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given string and assigns it to the Schema field.
func (o *PatchChannelRequest) SetSchema(v string) {
	o.Schema = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *PatchChannelRequest) GetCreated() time.Time {
	if o == nil || IsNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchChannelRequest) GetCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *PatchChannelRequest) HasCreated() bool {
	if o != nil && !IsNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *PatchChannelRequest) SetCreated(v time.Time) {
	o.Created = &v
}

// GetDesiredState returns the DesiredState field value if set, zero value otherwise.
func (o *PatchChannelRequest) GetDesiredState() string {
	if o == nil || IsNil(o.DesiredState) {
		var ret string
		return ret
	}
	return *o.DesiredState
}

// GetDesiredStateOk returns a tuple with the DesiredState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchChannelRequest) GetDesiredStateOk() (*string, bool) {
	if o == nil || IsNil(o.DesiredState) {
		return nil, false
	}
	return o.DesiredState, true
}

// HasDesiredState returns a boolean if a field has been set.
func (o *PatchChannelRequest) HasDesiredState() bool {
	if o != nil && !IsNil(o.DesiredState) {
		return true
	}

	return false
}

// SetDesiredState gets a reference to the given string and assigns it to the DesiredState field.
func (o *PatchChannelRequest) SetDesiredState(v string) {
	o.DesiredState = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PatchChannelRequest) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchChannelRequest) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PatchChannelRequest) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PatchChannelRequest) SetId(v string) {
	o.Id = &v
}

// GetIngest returns the Ingest field value if set, zero value otherwise.
func (o *PatchChannelRequest) GetIngest() PatchChannelRequestIngest {
	if o == nil || IsNil(o.Ingest) {
		var ret PatchChannelRequestIngest
		return ret
	}
	return *o.Ingest
}

// GetIngestOk returns a tuple with the Ingest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchChannelRequest) GetIngestOk() (*PatchChannelRequestIngest, bool) {
	if o == nil || IsNil(o.Ingest) {
		return nil, false
	}
	return o.Ingest, true
}

// HasIngest returns a boolean if a field has been set.
func (o *PatchChannelRequest) HasIngest() bool {
	if o != nil && !IsNil(o.Ingest) {
		return true
	}

	return false
}

// SetIngest gets a reference to the given PatchChannelRequestIngest and assigns it to the Ingest field.
func (o *PatchChannelRequest) SetIngest(v PatchChannelRequestIngest) {
	o.Ingest = &v
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *PatchChannelRequest) GetLabels() []string {
	if o == nil || IsNil(o.Labels) {
		var ret []string
		return ret
	}
	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchChannelRequest) GetLabelsOk() ([]string, bool) {
	if o == nil || IsNil(o.Labels) {
		return nil, false
	}
	return o.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *PatchChannelRequest) HasLabels() bool {
	if o != nil && !IsNil(o.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given []string and assigns it to the Labels field.
func (o *PatchChannelRequest) SetLabels(v []string) {
	o.Labels = v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *PatchChannelRequest) GetModified() time.Time {
	if o == nil || IsNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchChannelRequest) GetModifiedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *PatchChannelRequest) HasModified() bool {
	if o != nil && !IsNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *PatchChannelRequest) SetModified(v time.Time) {
	o.Modified = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchChannelRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchChannelRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchChannelRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchChannelRequest) SetName(v string) {
	o.Name = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *PatchChannelRequest) GetOrganization() string {
	if o == nil || IsNil(o.Organization) {
		var ret string
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchChannelRequest) GetOrganizationOk() (*string, bool) {
	if o == nil || IsNil(o.Organization) {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *PatchChannelRequest) HasOrganization() bool {
	if o != nil && !IsNil(o.Organization) {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given string and assigns it to the Organization field.
func (o *PatchChannelRequest) SetOrganization(v string) {
	o.Organization = &v
}

// GetPackaging returns the Packaging field value if set, zero value otherwise.
func (o *PatchChannelRequest) GetPackaging() ChannelPackaging {
	if o == nil || IsNil(o.Packaging) {
		var ret ChannelPackaging
		return ret
	}
	return *o.Packaging
}

// GetPackagingOk returns a tuple with the Packaging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchChannelRequest) GetPackagingOk() (*ChannelPackaging, bool) {
	if o == nil || IsNil(o.Packaging) {
		return nil, false
	}
	return o.Packaging, true
}

// HasPackaging returns a boolean if a field has been set.
func (o *PatchChannelRequest) HasPackaging() bool {
	if o != nil && !IsNil(o.Packaging) {
		return true
	}

	return false
}

// SetPackaging gets a reference to the given ChannelPackaging and assigns it to the Packaging field.
func (o *PatchChannelRequest) SetPackaging(v ChannelPackaging) {
	o.Packaging = &v
}

// GetPublishing returns the Publishing field value if set, zero value otherwise.
func (o *PatchChannelRequest) GetPublishing() ChannelPublishing {
	if o == nil || IsNil(o.Publishing) {
		var ret ChannelPublishing
		return ret
	}
	return *o.Publishing
}

// GetPublishingOk returns a tuple with the Publishing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchChannelRequest) GetPublishingOk() (*ChannelPublishing, bool) {
	if o == nil || IsNil(o.Publishing) {
		return nil, false
	}
	return o.Publishing, true
}

// HasPublishing returns a boolean if a field has been set.
func (o *PatchChannelRequest) HasPublishing() bool {
	if o != nil && !IsNil(o.Publishing) {
		return true
	}

	return false
}

// SetPublishing gets a reference to the given ChannelPublishing and assigns it to the Publishing field.
func (o *PatchChannelRequest) SetPublishing(v ChannelPublishing) {
	o.Publishing = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *PatchChannelRequest) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchChannelRequest) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// HasRegion returns a boolean if a field has been set.
func (o *PatchChannelRequest) HasRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *PatchChannelRequest) SetRegion(v string) {
	o.Region = &v
}

// GetResourceClass returns the ResourceClass field value if set, zero value otherwise.
func (o *PatchChannelRequest) GetResourceClass() string {
	if o == nil || IsNil(o.ResourceClass) {
		var ret string
		return ret
	}
	return *o.ResourceClass
}

// GetResourceClassOk returns a tuple with the ResourceClass field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchChannelRequest) GetResourceClassOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceClass) {
		return nil, false
	}
	return o.ResourceClass, true
}

// HasResourceClass returns a boolean if a field has been set.
func (o *PatchChannelRequest) HasResourceClass() bool {
	if o != nil && !IsNil(o.ResourceClass) {
		return true
	}

	return false
}

// SetResourceClass gets a reference to the given string and assigns it to the ResourceClass field.
func (o *PatchChannelRequest) SetResourceClass(v string) {
	o.ResourceClass = &v
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *PatchChannelRequest) GetSelf() string {
	if o == nil || IsNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchChannelRequest) GetSelfOk() (*string, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *PatchChannelRequest) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *PatchChannelRequest) SetSelf(v string) {
	o.Self = &v
}

// GetSignaling returns the Signaling field value if set, zero value otherwise.
func (o *PatchChannelRequest) GetSignaling() ChannelSignaling {
	if o == nil || IsNil(o.Signaling) {
		var ret ChannelSignaling
		return ret
	}
	return *o.Signaling
}

// GetSignalingOk returns a tuple with the Signaling field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchChannelRequest) GetSignalingOk() (*ChannelSignaling, bool) {
	if o == nil || IsNil(o.Signaling) {
		return nil, false
	}
	return o.Signaling, true
}

// HasSignaling returns a boolean if a field has been set.
func (o *PatchChannelRequest) HasSignaling() bool {
	if o != nil && !IsNil(o.Signaling) {
		return true
	}

	return false
}

// SetSignaling gets a reference to the given ChannelSignaling and assigns it to the Signaling field.
func (o *PatchChannelRequest) SetSignaling(v ChannelSignaling) {
	o.Signaling = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchChannelRequest) GetTags() ChannelTags {
	if o == nil || IsNil(o.Tags) {
		var ret ChannelTags
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchChannelRequest) GetTagsOk() (*ChannelTags, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchChannelRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given ChannelTags and assigns it to the Tags field.
func (o *PatchChannelRequest) SetTags(v ChannelTags) {
	o.Tags = &v
}

// GetTranscode returns the Transcode field value if set, zero value otherwise.
func (o *PatchChannelRequest) GetTranscode() ChannelTranscode {
	if o == nil || IsNil(o.Transcode) {
		var ret ChannelTranscode
		return ret
	}
	return *o.Transcode
}

// GetTranscodeOk returns a tuple with the Transcode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchChannelRequest) GetTranscodeOk() (*ChannelTranscode, bool) {
	if o == nil || IsNil(o.Transcode) {
		return nil, false
	}
	return o.Transcode, true
}

// HasTranscode returns a boolean if a field has been set.
func (o *PatchChannelRequest) HasTranscode() bool {
	if o != nil && !IsNil(o.Transcode) {
		return true
	}

	return false
}

// SetTranscode gets a reference to the given ChannelTranscode and assigns it to the Transcode field.
func (o *PatchChannelRequest) SetTranscode(v ChannelTranscode) {
	o.Transcode = &v
}

func (o PatchChannelRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchChannelRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schema) {
		toSerialize["$schema"] = o.Schema
	}
	if !IsNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !IsNil(o.DesiredState) {
		toSerialize["desired_state"] = o.DesiredState
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Ingest) {
		toSerialize["ingest"] = o.Ingest
	}
	if !IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	if !IsNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Organization) {
		toSerialize["organization"] = o.Organization
	}
	if !IsNil(o.Packaging) {
		toSerialize["packaging"] = o.Packaging
	}
	if !IsNil(o.Publishing) {
		toSerialize["publishing"] = o.Publishing
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.ResourceClass) {
		toSerialize["resource_class"] = o.ResourceClass
	}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	if !IsNil(o.Signaling) {
		toSerialize["signaling"] = o.Signaling
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.Transcode) {
		toSerialize["transcode"] = o.Transcode
	}
	return toSerialize, nil
}

type NullablePatchChannelRequest struct {
	value *PatchChannelRequest
	isSet bool
}

func (v NullablePatchChannelRequest) Get() *PatchChannelRequest {
	return v.value
}

func (v *NullablePatchChannelRequest) Set(val *PatchChannelRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchChannelRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchChannelRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchChannelRequest(val *PatchChannelRequest) *NullablePatchChannelRequest {
	return &NullablePatchChannelRequest{value: val, isSet: true}
}

func (v NullablePatchChannelRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchChannelRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


