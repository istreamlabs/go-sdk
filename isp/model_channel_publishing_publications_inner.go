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

// checks if the ChannelPublishingPublicationsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelPublishingPublicationsInner{}

// ChannelPublishingPublicationsInner struct for ChannelPublishingPublicationsInner
type ChannelPublishingPublicationsInner struct {
	// Optionally specify which audio encoders should be used for this publication. If none are specified, all audio encoders configured for the transcoder will be used.
	AudioEncoderIds []string `json:"audio_encoder_ids,omitempty" uniqueItems:"true" doc:"Optionally specify which audio encoders should be used for this publication. If none are specified, all audio encoders configured for the transcoder will be used."`
	// Create VODs for all publish points in this publication. Note that Live2VOD must also be configured for the parent |Channel|.
	CreateVods *bool `json:"create_vods,omitempty" doc:"Create VODs for all publish points in this publication. Note that Live2VOD must also be configured for the parent |Channel|."`
	Dash *ChannelPublishingPublicationsInnerDash `json:"dash,omitempty"`
	// Optionally specify which DRMs to advertise in the playlist. If specified, this must be a subset of the DRMs specified by the packager associated with this publication. If omitted or empty, all DRMs specified by the packager will be advertised. This setting can only be used for HLS playlists.
	Drms []string `json:"drms,omitempty" uniqueItems:"true" enum:"WIDEVINE,FAIRPLAY,PRIMETIME,PLAYREADY,W3C_COMMON_CLEAR_KEY,BULK_FILE" doc:"Optionally specify which DRMs to advertise in the playlist. If specified, this must be a subset of the DRMs specified by the packager associated with this publication. If omitted or empty, all DRMs specified by the packager will be advertised. This setting can only be used for HLS playlists."`
	// DVR window is the max sum(duration of media segments) that will be kept in a manifest at a given time in seconds. The max supported DVR window is 10 hours.
	DvrWindowSecs *int32 `json:"dvr_window_secs,omitempty" format:"int32" minimum:"0" maximum:"36000" doc:"DVR window is the max sum(duration of media segments) that will be kept in a manifest at a given time in seconds. The max supported DVR window is 10 hours."`
	// Optionally specify which audio encoders should be used when generating the FER of this Presentation, this overrides the 'audio_encoder_ids' used during the live portion. If none are specified, the 'audio_encoder_ids' field will be used.
	FerAudioEncoderIds []string `json:"fer_audio_encoder_ids,omitempty" uniqueItems:"true" doc:"Optionally specify which audio encoders should be used when generating the FER of this Presentation, this overrides the 'audio_encoder_ids' used during the live portion. If none are specified, the 'audio_encoder_ids' field will be used."`
	Hls *ChannelPublishingPublicationsInnerHls `json:"hls,omitempty"`
	// List of video encoder IDs that should have I-Frame only playlists generated for them.
	IframeOnlyEncoderIds []string `json:"iframe_only_encoder_ids,omitempty" uniqueItems:"true" doc:"List of video encoder IDs that should have I-Frame only playlists generated for them."`
	// Optional master manifest name. When not supplied a default of 'master' will be used.
	MasterPlaylistName *string `json:"master_playlist_name,omitempty" doc:"Optional master manifest name. When not supplied a default of 'master' will be used."`
	Origin *ChannelPublishingPublicationsInnerOrigin `json:"origin,omitempty"`
	// Determines how segments in this publication are packaged. Must reference a packager in 'packaging.packagers'. However, if this is a playlist-only publication (i.e. contains publish points that specify 'playlist_only_for'), this must remain unset as the packager will be inferred from the publication this one is providing playlists for.
	PackagerId *string `json:"packager_id,omitempty" doc:"Determines how segments in this publication are packaged. Must reference a packager in 'packaging.packagers'. However, if this is a playlist-only publication (i.e. contains publish points that specify 'playlist_only_for'), this must remain unset as the packager will be inferred from the publication this one is providing playlists for."`
	// Publish points specify where to output.
	PublishPoints []ChannelPublishingPublicationsInnerPublishPointsInner `json:"publish_points,omitempty" doc:"Publish points specify where to output."`
	// When redundant publishing is enabled succeeding to publish a given media segment to at least one HTTPPublishPoint in publish_points will result in that segment showing up in manifests as playable content. Will require at least two publish_points defined within the same publication.
	RedundantPublishing *bool `json:"redundant_publishing,omitempty" doc:"When redundant publishing is enabled succeeding to publish a given media segment to at least one HTTPPublishPoint in publish_points will result in that segment showing up in manifests as playable content. Will require at least two publish_points defined within the same publication."`
	Startover *ChannelPublishingPublicationsInnerStartover `json:"startover,omitempty"`
	// Optional: Specify what thumbnail_encoders should be in this Publication
	ThumbnailEncoderIds []string `json:"thumbnail_encoder_ids,omitempty" uniqueItems:"true" doc:"Optional: Specify what thumbnail_encoders should be in this Publication"`
	// Optional, indicates whether we should pad the bitrate (false) or use what is explicitly provided (true)
	UseStrictBitrate *bool `json:"use_strict_bitrate,omitempty" doc:"Optional, indicates whether we should pad the bitrate (false) or use what is explicitly provided (true)"`
	// Optionally specify which video encoders should be used for this publication. If none are specified, all video encoders configured for the transcoder will be used.
	VideoEncoderIds []string `json:"video_encoder_ids,omitempty" uniqueItems:"true" doc:"Optionally specify which video encoders should be used for this publication. If none are specified, all video encoders configured for the transcoder will be used."`
}

// NewChannelPublishingPublicationsInner instantiates a new ChannelPublishingPublicationsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelPublishingPublicationsInner() *ChannelPublishingPublicationsInner {
	this := ChannelPublishingPublicationsInner{}
	return &this
}

// NewChannelPublishingPublicationsInnerWithDefaults instantiates a new ChannelPublishingPublicationsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelPublishingPublicationsInnerWithDefaults() *ChannelPublishingPublicationsInner {
	this := ChannelPublishingPublicationsInner{}
	return &this
}

// GetAudioEncoderIds returns the AudioEncoderIds field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInner) GetAudioEncoderIds() []string {
	if o == nil || IsNil(o.AudioEncoderIds) {
		var ret []string
		return ret
	}
	return o.AudioEncoderIds
}

// GetAudioEncoderIdsOk returns a tuple with the AudioEncoderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInner) GetAudioEncoderIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.AudioEncoderIds) {
		return nil, false
	}
	return o.AudioEncoderIds, true
}

// HasAudioEncoderIds returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInner) HasAudioEncoderIds() bool {
	if o != nil && !IsNil(o.AudioEncoderIds) {
		return true
	}

	return false
}

// SetAudioEncoderIds gets a reference to the given []string and assigns it to the AudioEncoderIds field.
func (o *ChannelPublishingPublicationsInner) SetAudioEncoderIds(v []string) {
	o.AudioEncoderIds = v
}

// GetCreateVods returns the CreateVods field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInner) GetCreateVods() bool {
	if o == nil || IsNil(o.CreateVods) {
		var ret bool
		return ret
	}
	return *o.CreateVods
}

// GetCreateVodsOk returns a tuple with the CreateVods field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInner) GetCreateVodsOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateVods) {
		return nil, false
	}
	return o.CreateVods, true
}

// HasCreateVods returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInner) HasCreateVods() bool {
	if o != nil && !IsNil(o.CreateVods) {
		return true
	}

	return false
}

// SetCreateVods gets a reference to the given bool and assigns it to the CreateVods field.
func (o *ChannelPublishingPublicationsInner) SetCreateVods(v bool) {
	o.CreateVods = &v
}

// GetDash returns the Dash field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInner) GetDash() ChannelPublishingPublicationsInnerDash {
	if o == nil || IsNil(o.Dash) {
		var ret ChannelPublishingPublicationsInnerDash
		return ret
	}
	return *o.Dash
}

// GetDashOk returns a tuple with the Dash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInner) GetDashOk() (*ChannelPublishingPublicationsInnerDash, bool) {
	if o == nil || IsNil(o.Dash) {
		return nil, false
	}
	return o.Dash, true
}

// HasDash returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInner) HasDash() bool {
	if o != nil && !IsNil(o.Dash) {
		return true
	}

	return false
}

// SetDash gets a reference to the given ChannelPublishingPublicationsInnerDash and assigns it to the Dash field.
func (o *ChannelPublishingPublicationsInner) SetDash(v ChannelPublishingPublicationsInnerDash) {
	o.Dash = &v
}

// GetDrms returns the Drms field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInner) GetDrms() []string {
	if o == nil || IsNil(o.Drms) {
		var ret []string
		return ret
	}
	return o.Drms
}

// GetDrmsOk returns a tuple with the Drms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInner) GetDrmsOk() ([]string, bool) {
	if o == nil || IsNil(o.Drms) {
		return nil, false
	}
	return o.Drms, true
}

// HasDrms returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInner) HasDrms() bool {
	if o != nil && !IsNil(o.Drms) {
		return true
	}

	return false
}

// SetDrms gets a reference to the given []string and assigns it to the Drms field.
func (o *ChannelPublishingPublicationsInner) SetDrms(v []string) {
	o.Drms = v
}

// GetDvrWindowSecs returns the DvrWindowSecs field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInner) GetDvrWindowSecs() int32 {
	if o == nil || IsNil(o.DvrWindowSecs) {
		var ret int32
		return ret
	}
	return *o.DvrWindowSecs
}

// GetDvrWindowSecsOk returns a tuple with the DvrWindowSecs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInner) GetDvrWindowSecsOk() (*int32, bool) {
	if o == nil || IsNil(o.DvrWindowSecs) {
		return nil, false
	}
	return o.DvrWindowSecs, true
}

// HasDvrWindowSecs returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInner) HasDvrWindowSecs() bool {
	if o != nil && !IsNil(o.DvrWindowSecs) {
		return true
	}

	return false
}

// SetDvrWindowSecs gets a reference to the given int32 and assigns it to the DvrWindowSecs field.
func (o *ChannelPublishingPublicationsInner) SetDvrWindowSecs(v int32) {
	o.DvrWindowSecs = &v
}

// GetFerAudioEncoderIds returns the FerAudioEncoderIds field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInner) GetFerAudioEncoderIds() []string {
	if o == nil || IsNil(o.FerAudioEncoderIds) {
		var ret []string
		return ret
	}
	return o.FerAudioEncoderIds
}

// GetFerAudioEncoderIdsOk returns a tuple with the FerAudioEncoderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInner) GetFerAudioEncoderIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.FerAudioEncoderIds) {
		return nil, false
	}
	return o.FerAudioEncoderIds, true
}

// HasFerAudioEncoderIds returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInner) HasFerAudioEncoderIds() bool {
	if o != nil && !IsNil(o.FerAudioEncoderIds) {
		return true
	}

	return false
}

// SetFerAudioEncoderIds gets a reference to the given []string and assigns it to the FerAudioEncoderIds field.
func (o *ChannelPublishingPublicationsInner) SetFerAudioEncoderIds(v []string) {
	o.FerAudioEncoderIds = v
}

// GetHls returns the Hls field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInner) GetHls() ChannelPublishingPublicationsInnerHls {
	if o == nil || IsNil(o.Hls) {
		var ret ChannelPublishingPublicationsInnerHls
		return ret
	}
	return *o.Hls
}

// GetHlsOk returns a tuple with the Hls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInner) GetHlsOk() (*ChannelPublishingPublicationsInnerHls, bool) {
	if o == nil || IsNil(o.Hls) {
		return nil, false
	}
	return o.Hls, true
}

// HasHls returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInner) HasHls() bool {
	if o != nil && !IsNil(o.Hls) {
		return true
	}

	return false
}

// SetHls gets a reference to the given ChannelPublishingPublicationsInnerHls and assigns it to the Hls field.
func (o *ChannelPublishingPublicationsInner) SetHls(v ChannelPublishingPublicationsInnerHls) {
	o.Hls = &v
}

// GetIframeOnlyEncoderIds returns the IframeOnlyEncoderIds field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInner) GetIframeOnlyEncoderIds() []string {
	if o == nil || IsNil(o.IframeOnlyEncoderIds) {
		var ret []string
		return ret
	}
	return o.IframeOnlyEncoderIds
}

// GetIframeOnlyEncoderIdsOk returns a tuple with the IframeOnlyEncoderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInner) GetIframeOnlyEncoderIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.IframeOnlyEncoderIds) {
		return nil, false
	}
	return o.IframeOnlyEncoderIds, true
}

// HasIframeOnlyEncoderIds returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInner) HasIframeOnlyEncoderIds() bool {
	if o != nil && !IsNil(o.IframeOnlyEncoderIds) {
		return true
	}

	return false
}

// SetIframeOnlyEncoderIds gets a reference to the given []string and assigns it to the IframeOnlyEncoderIds field.
func (o *ChannelPublishingPublicationsInner) SetIframeOnlyEncoderIds(v []string) {
	o.IframeOnlyEncoderIds = v
}

// GetMasterPlaylistName returns the MasterPlaylistName field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInner) GetMasterPlaylistName() string {
	if o == nil || IsNil(o.MasterPlaylistName) {
		var ret string
		return ret
	}
	return *o.MasterPlaylistName
}

// GetMasterPlaylistNameOk returns a tuple with the MasterPlaylistName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInner) GetMasterPlaylistNameOk() (*string, bool) {
	if o == nil || IsNil(o.MasterPlaylistName) {
		return nil, false
	}
	return o.MasterPlaylistName, true
}

// HasMasterPlaylistName returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInner) HasMasterPlaylistName() bool {
	if o != nil && !IsNil(o.MasterPlaylistName) {
		return true
	}

	return false
}

// SetMasterPlaylistName gets a reference to the given string and assigns it to the MasterPlaylistName field.
func (o *ChannelPublishingPublicationsInner) SetMasterPlaylistName(v string) {
	o.MasterPlaylistName = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInner) GetOrigin() ChannelPublishingPublicationsInnerOrigin {
	if o == nil || IsNil(o.Origin) {
		var ret ChannelPublishingPublicationsInnerOrigin
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInner) GetOriginOk() (*ChannelPublishingPublicationsInnerOrigin, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInner) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given ChannelPublishingPublicationsInnerOrigin and assigns it to the Origin field.
func (o *ChannelPublishingPublicationsInner) SetOrigin(v ChannelPublishingPublicationsInnerOrigin) {
	o.Origin = &v
}

// GetPackagerId returns the PackagerId field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInner) GetPackagerId() string {
	if o == nil || IsNil(o.PackagerId) {
		var ret string
		return ret
	}
	return *o.PackagerId
}

// GetPackagerIdOk returns a tuple with the PackagerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInner) GetPackagerIdOk() (*string, bool) {
	if o == nil || IsNil(o.PackagerId) {
		return nil, false
	}
	return o.PackagerId, true
}

// HasPackagerId returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInner) HasPackagerId() bool {
	if o != nil && !IsNil(o.PackagerId) {
		return true
	}

	return false
}

// SetPackagerId gets a reference to the given string and assigns it to the PackagerId field.
func (o *ChannelPublishingPublicationsInner) SetPackagerId(v string) {
	o.PackagerId = &v
}

// GetPublishPoints returns the PublishPoints field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInner) GetPublishPoints() []ChannelPublishingPublicationsInnerPublishPointsInner {
	if o == nil || IsNil(o.PublishPoints) {
		var ret []ChannelPublishingPublicationsInnerPublishPointsInner
		return ret
	}
	return o.PublishPoints
}

// GetPublishPointsOk returns a tuple with the PublishPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInner) GetPublishPointsOk() ([]ChannelPublishingPublicationsInnerPublishPointsInner, bool) {
	if o == nil || IsNil(o.PublishPoints) {
		return nil, false
	}
	return o.PublishPoints, true
}

// HasPublishPoints returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInner) HasPublishPoints() bool {
	if o != nil && !IsNil(o.PublishPoints) {
		return true
	}

	return false
}

// SetPublishPoints gets a reference to the given []ChannelPublishingPublicationsInnerPublishPointsInner and assigns it to the PublishPoints field.
func (o *ChannelPublishingPublicationsInner) SetPublishPoints(v []ChannelPublishingPublicationsInnerPublishPointsInner) {
	o.PublishPoints = v
}

// GetRedundantPublishing returns the RedundantPublishing field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInner) GetRedundantPublishing() bool {
	if o == nil || IsNil(o.RedundantPublishing) {
		var ret bool
		return ret
	}
	return *o.RedundantPublishing
}

// GetRedundantPublishingOk returns a tuple with the RedundantPublishing field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInner) GetRedundantPublishingOk() (*bool, bool) {
	if o == nil || IsNil(o.RedundantPublishing) {
		return nil, false
	}
	return o.RedundantPublishing, true
}

// HasRedundantPublishing returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInner) HasRedundantPublishing() bool {
	if o != nil && !IsNil(o.RedundantPublishing) {
		return true
	}

	return false
}

// SetRedundantPublishing gets a reference to the given bool and assigns it to the RedundantPublishing field.
func (o *ChannelPublishingPublicationsInner) SetRedundantPublishing(v bool) {
	o.RedundantPublishing = &v
}

// GetStartover returns the Startover field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInner) GetStartover() ChannelPublishingPublicationsInnerStartover {
	if o == nil || IsNil(o.Startover) {
		var ret ChannelPublishingPublicationsInnerStartover
		return ret
	}
	return *o.Startover
}

// GetStartoverOk returns a tuple with the Startover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInner) GetStartoverOk() (*ChannelPublishingPublicationsInnerStartover, bool) {
	if o == nil || IsNil(o.Startover) {
		return nil, false
	}
	return o.Startover, true
}

// HasStartover returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInner) HasStartover() bool {
	if o != nil && !IsNil(o.Startover) {
		return true
	}

	return false
}

// SetStartover gets a reference to the given ChannelPublishingPublicationsInnerStartover and assigns it to the Startover field.
func (o *ChannelPublishingPublicationsInner) SetStartover(v ChannelPublishingPublicationsInnerStartover) {
	o.Startover = &v
}

// GetThumbnailEncoderIds returns the ThumbnailEncoderIds field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInner) GetThumbnailEncoderIds() []string {
	if o == nil || IsNil(o.ThumbnailEncoderIds) {
		var ret []string
		return ret
	}
	return o.ThumbnailEncoderIds
}

// GetThumbnailEncoderIdsOk returns a tuple with the ThumbnailEncoderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInner) GetThumbnailEncoderIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ThumbnailEncoderIds) {
		return nil, false
	}
	return o.ThumbnailEncoderIds, true
}

// HasThumbnailEncoderIds returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInner) HasThumbnailEncoderIds() bool {
	if o != nil && !IsNil(o.ThumbnailEncoderIds) {
		return true
	}

	return false
}

// SetThumbnailEncoderIds gets a reference to the given []string and assigns it to the ThumbnailEncoderIds field.
func (o *ChannelPublishingPublicationsInner) SetThumbnailEncoderIds(v []string) {
	o.ThumbnailEncoderIds = v
}

// GetUseStrictBitrate returns the UseStrictBitrate field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInner) GetUseStrictBitrate() bool {
	if o == nil || IsNil(o.UseStrictBitrate) {
		var ret bool
		return ret
	}
	return *o.UseStrictBitrate
}

// GetUseStrictBitrateOk returns a tuple with the UseStrictBitrate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInner) GetUseStrictBitrateOk() (*bool, bool) {
	if o == nil || IsNil(o.UseStrictBitrate) {
		return nil, false
	}
	return o.UseStrictBitrate, true
}

// HasUseStrictBitrate returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInner) HasUseStrictBitrate() bool {
	if o != nil && !IsNil(o.UseStrictBitrate) {
		return true
	}

	return false
}

// SetUseStrictBitrate gets a reference to the given bool and assigns it to the UseStrictBitrate field.
func (o *ChannelPublishingPublicationsInner) SetUseStrictBitrate(v bool) {
	o.UseStrictBitrate = &v
}

// GetVideoEncoderIds returns the VideoEncoderIds field value if set, zero value otherwise.
func (o *ChannelPublishingPublicationsInner) GetVideoEncoderIds() []string {
	if o == nil || IsNil(o.VideoEncoderIds) {
		var ret []string
		return ret
	}
	return o.VideoEncoderIds
}

// GetVideoEncoderIdsOk returns a tuple with the VideoEncoderIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelPublishingPublicationsInner) GetVideoEncoderIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.VideoEncoderIds) {
		return nil, false
	}
	return o.VideoEncoderIds, true
}

// HasVideoEncoderIds returns a boolean if a field has been set.
func (o *ChannelPublishingPublicationsInner) HasVideoEncoderIds() bool {
	if o != nil && !IsNil(o.VideoEncoderIds) {
		return true
	}

	return false
}

// SetVideoEncoderIds gets a reference to the given []string and assigns it to the VideoEncoderIds field.
func (o *ChannelPublishingPublicationsInner) SetVideoEncoderIds(v []string) {
	o.VideoEncoderIds = v
}

func (o ChannelPublishingPublicationsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelPublishingPublicationsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AudioEncoderIds) {
		toSerialize["audio_encoder_ids"] = o.AudioEncoderIds
	}
	if !IsNil(o.CreateVods) {
		toSerialize["create_vods"] = o.CreateVods
	}
	if !IsNil(o.Dash) {
		toSerialize["dash"] = o.Dash
	}
	if !IsNil(o.Drms) {
		toSerialize["drms"] = o.Drms
	}
	if !IsNil(o.DvrWindowSecs) {
		toSerialize["dvr_window_secs"] = o.DvrWindowSecs
	}
	if !IsNil(o.FerAudioEncoderIds) {
		toSerialize["fer_audio_encoder_ids"] = o.FerAudioEncoderIds
	}
	if !IsNil(o.Hls) {
		toSerialize["hls"] = o.Hls
	}
	if !IsNil(o.IframeOnlyEncoderIds) {
		toSerialize["iframe_only_encoder_ids"] = o.IframeOnlyEncoderIds
	}
	if !IsNil(o.MasterPlaylistName) {
		toSerialize["master_playlist_name"] = o.MasterPlaylistName
	}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !IsNil(o.PackagerId) {
		toSerialize["packager_id"] = o.PackagerId
	}
	if !IsNil(o.PublishPoints) {
		toSerialize["publish_points"] = o.PublishPoints
	}
	if !IsNil(o.RedundantPublishing) {
		toSerialize["redundant_publishing"] = o.RedundantPublishing
	}
	if !IsNil(o.Startover) {
		toSerialize["startover"] = o.Startover
	}
	if !IsNil(o.ThumbnailEncoderIds) {
		toSerialize["thumbnail_encoder_ids"] = o.ThumbnailEncoderIds
	}
	if !IsNil(o.UseStrictBitrate) {
		toSerialize["use_strict_bitrate"] = o.UseStrictBitrate
	}
	if !IsNil(o.VideoEncoderIds) {
		toSerialize["video_encoder_ids"] = o.VideoEncoderIds
	}
	return toSerialize, nil
}

type NullableChannelPublishingPublicationsInner struct {
	value *ChannelPublishingPublicationsInner
	isSet bool
}

func (v NullableChannelPublishingPublicationsInner) Get() *ChannelPublishingPublicationsInner {
	return v.value
}

func (v *NullableChannelPublishingPublicationsInner) Set(val *ChannelPublishingPublicationsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelPublishingPublicationsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelPublishingPublicationsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelPublishingPublicationsInner(val *ChannelPublishingPublicationsInner) *NullableChannelPublishingPublicationsInner {
	return &NullableChannelPublishingPublicationsInner{value: val, isSet: true}
}

func (v NullableChannelPublishingPublicationsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelPublishingPublicationsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


