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

// checks if the ChannelTranscodeVideoEncodersInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ChannelTranscodeVideoEncodersInner{}

// ChannelTranscodeVideoEncodersInner struct for ChannelTranscodeVideoEncodersInner
type ChannelTranscodeVideoEncodersInner struct {
	// Bit rate specifies the number in bits used per second. Higher values result in better video quality but bigger file sizes. For H.264 this value is the target of the constrained variable bit rate.
	BitRate *int32 `json:"bit_rate,omitempty" format:"int32" exclusiveMinimum:"1000" doc:"Bit rate specifies the number in bits used per second. Higher values result in better video quality but bigger file sizes. For H.264 this value is the target of the constrained variable bit rate."`
	// Frame rate specifies the number of images that are shown per second when playing back the video. For the best quality playback, this should match or be a multiple of the input source video stream.
	FrameRate *string `json:"frame_rate,omitempty" enum:"FR_23_98,FR_24,FR_25,FR_29_97,FR_30,FR_50,FR_59_94,FR_60" doc:"Frame rate specifies the number of images that are shown per second when playing back the video. For the best quality playback, this should match or be a multiple of the input source video stream."`
	H264 *ChannelTranscodeVideoEncodersInnerH264 `json:"h264,omitempty"`
	H265 *ChannelTranscodeVideoEncodersInnerH265 `json:"h265,omitempty"`
	// Height specifies the video height in pixels. Must be a multiple of two. For H.265 encoders the minimum height is 34.
	Height *int32 `json:"height,omitempty" format:"int32" exclusiveMinimum:"1" doc:"Height specifies the video height in pixels. Must be a multiple of two. For H.265 encoders the minimum height is 34."`
	// Encoder ID. IDs must be unique for all encoders. This ID is referenced when setting up playlist publishing.
	Id *string `json:"id,omitempty" minLength:"1" doc:"Encoder ID. IDs must be unique for all encoders. This ID is referenced when setting up playlist publishing."`
	// Width specifies the video width in pixels. Must be a multiple of two. For H.265 encoders the minimum width is 130.
	Width *int32 `json:"width,omitempty" format:"int32" exclusiveMinimum:"1" doc:"Width specifies the video width in pixels. Must be a multiple of two. For H.265 encoders the minimum width is 130."`
}

// NewChannelTranscodeVideoEncodersInner instantiates a new ChannelTranscodeVideoEncodersInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChannelTranscodeVideoEncodersInner() *ChannelTranscodeVideoEncodersInner {
	this := ChannelTranscodeVideoEncodersInner{}
	return &this
}

// NewChannelTranscodeVideoEncodersInnerWithDefaults instantiates a new ChannelTranscodeVideoEncodersInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChannelTranscodeVideoEncodersInnerWithDefaults() *ChannelTranscodeVideoEncodersInner {
	this := ChannelTranscodeVideoEncodersInner{}
	return &this
}

// GetBitRate returns the BitRate field value if set, zero value otherwise.
func (o *ChannelTranscodeVideoEncodersInner) GetBitRate() int32 {
	if o == nil || IsNil(o.BitRate) {
		var ret int32
		return ret
	}
	return *o.BitRate
}

// GetBitRateOk returns a tuple with the BitRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeVideoEncodersInner) GetBitRateOk() (*int32, bool) {
	if o == nil || IsNil(o.BitRate) {
		return nil, false
	}
	return o.BitRate, true
}

// HasBitRate returns a boolean if a field has been set.
func (o *ChannelTranscodeVideoEncodersInner) HasBitRate() bool {
	if o != nil && !IsNil(o.BitRate) {
		return true
	}

	return false
}

// SetBitRate gets a reference to the given int32 and assigns it to the BitRate field.
func (o *ChannelTranscodeVideoEncodersInner) SetBitRate(v int32) {
	o.BitRate = &v
}

// GetFrameRate returns the FrameRate field value if set, zero value otherwise.
func (o *ChannelTranscodeVideoEncodersInner) GetFrameRate() string {
	if o == nil || IsNil(o.FrameRate) {
		var ret string
		return ret
	}
	return *o.FrameRate
}

// GetFrameRateOk returns a tuple with the FrameRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeVideoEncodersInner) GetFrameRateOk() (*string, bool) {
	if o == nil || IsNil(o.FrameRate) {
		return nil, false
	}
	return o.FrameRate, true
}

// HasFrameRate returns a boolean if a field has been set.
func (o *ChannelTranscodeVideoEncodersInner) HasFrameRate() bool {
	if o != nil && !IsNil(o.FrameRate) {
		return true
	}

	return false
}

// SetFrameRate gets a reference to the given string and assigns it to the FrameRate field.
func (o *ChannelTranscodeVideoEncodersInner) SetFrameRate(v string) {
	o.FrameRate = &v
}

// GetH264 returns the H264 field value if set, zero value otherwise.
func (o *ChannelTranscodeVideoEncodersInner) GetH264() ChannelTranscodeVideoEncodersInnerH264 {
	if o == nil || IsNil(o.H264) {
		var ret ChannelTranscodeVideoEncodersInnerH264
		return ret
	}
	return *o.H264
}

// GetH264Ok returns a tuple with the H264 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeVideoEncodersInner) GetH264Ok() (*ChannelTranscodeVideoEncodersInnerH264, bool) {
	if o == nil || IsNil(o.H264) {
		return nil, false
	}
	return o.H264, true
}

// HasH264 returns a boolean if a field has been set.
func (o *ChannelTranscodeVideoEncodersInner) HasH264() bool {
	if o != nil && !IsNil(o.H264) {
		return true
	}

	return false
}

// SetH264 gets a reference to the given ChannelTranscodeVideoEncodersInnerH264 and assigns it to the H264 field.
func (o *ChannelTranscodeVideoEncodersInner) SetH264(v ChannelTranscodeVideoEncodersInnerH264) {
	o.H264 = &v
}

// GetH265 returns the H265 field value if set, zero value otherwise.
func (o *ChannelTranscodeVideoEncodersInner) GetH265() ChannelTranscodeVideoEncodersInnerH265 {
	if o == nil || IsNil(o.H265) {
		var ret ChannelTranscodeVideoEncodersInnerH265
		return ret
	}
	return *o.H265
}

// GetH265Ok returns a tuple with the H265 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeVideoEncodersInner) GetH265Ok() (*ChannelTranscodeVideoEncodersInnerH265, bool) {
	if o == nil || IsNil(o.H265) {
		return nil, false
	}
	return o.H265, true
}

// HasH265 returns a boolean if a field has been set.
func (o *ChannelTranscodeVideoEncodersInner) HasH265() bool {
	if o != nil && !IsNil(o.H265) {
		return true
	}

	return false
}

// SetH265 gets a reference to the given ChannelTranscodeVideoEncodersInnerH265 and assigns it to the H265 field.
func (o *ChannelTranscodeVideoEncodersInner) SetH265(v ChannelTranscodeVideoEncodersInnerH265) {
	o.H265 = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *ChannelTranscodeVideoEncodersInner) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeVideoEncodersInner) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *ChannelTranscodeVideoEncodersInner) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *ChannelTranscodeVideoEncodersInner) SetHeight(v int32) {
	o.Height = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ChannelTranscodeVideoEncodersInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeVideoEncodersInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ChannelTranscodeVideoEncodersInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ChannelTranscodeVideoEncodersInner) SetId(v string) {
	o.Id = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *ChannelTranscodeVideoEncodersInner) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChannelTranscodeVideoEncodersInner) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *ChannelTranscodeVideoEncodersInner) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *ChannelTranscodeVideoEncodersInner) SetWidth(v int32) {
	o.Width = &v
}

func (o ChannelTranscodeVideoEncodersInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ChannelTranscodeVideoEncodersInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BitRate) {
		toSerialize["bit_rate"] = o.BitRate
	}
	if !IsNil(o.FrameRate) {
		toSerialize["frame_rate"] = o.FrameRate
	}
	if !IsNil(o.H264) {
		toSerialize["h264"] = o.H264
	}
	if !IsNil(o.H265) {
		toSerialize["h265"] = o.H265
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	return toSerialize, nil
}

type NullableChannelTranscodeVideoEncodersInner struct {
	value *ChannelTranscodeVideoEncodersInner
	isSet bool
}

func (v NullableChannelTranscodeVideoEncodersInner) Get() *ChannelTranscodeVideoEncodersInner {
	return v.value
}

func (v *NullableChannelTranscodeVideoEncodersInner) Set(val *ChannelTranscodeVideoEncodersInner) {
	v.value = val
	v.isSet = true
}

func (v NullableChannelTranscodeVideoEncodersInner) IsSet() bool {
	return v.isSet
}

func (v *NullableChannelTranscodeVideoEncodersInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChannelTranscodeVideoEncodersInner(val *ChannelTranscodeVideoEncodersInner) *NullableChannelTranscodeVideoEncodersInner {
	return &NullableChannelTranscodeVideoEncodersInner{value: val, isSet: true}
}

func (v NullableChannelTranscodeVideoEncodersInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChannelTranscodeVideoEncodersInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


