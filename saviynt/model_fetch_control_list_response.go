/*
Saviynt Enterprise Identity Cloud (EIC) API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 5.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package saviynt

import (
	"encoding/json"
)

// checks if the FetchControlListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FetchControlListResponse{}

// FetchControlListResponse 
type FetchControlListResponse struct {
	Msg *string `json:"msg,omitempty"`
	Controls []FetchControlListControl `json:"controls,omitempty"`
	DisplayCount *int32 `json:"displayCount,omitempty"`
	ErrorCode *string `json:"errorCode,omitempty"`
	TotalCount *int32 `json:"totalCount,omitempty"`
}

// NewFetchControlListResponse instantiates a new FetchControlListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFetchControlListResponse() *FetchControlListResponse {
	this := FetchControlListResponse{}
	return &this
}

// NewFetchControlListResponseWithDefaults instantiates a new FetchControlListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFetchControlListResponseWithDefaults() *FetchControlListResponse {
	this := FetchControlListResponse{}
	return &this
}

// GetMsg returns the Msg field value if set, zero value otherwise.
func (o *FetchControlListResponse) GetMsg() string {
	if o == nil || IsNil(o.Msg) {
		var ret string
		return ret
	}
	return *o.Msg
}

// GetMsgOk returns a tuple with the Msg field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchControlListResponse) GetMsgOk() (*string, bool) {
	if o == nil || IsNil(o.Msg) {
		return nil, false
	}
	return o.Msg, true
}

// HasMsg returns a boolean if a field has been set.
func (o *FetchControlListResponse) HasMsg() bool {
	if o != nil && !IsNil(o.Msg) {
		return true
	}

	return false
}

// SetMsg gets a reference to the given string and assigns it to the Msg field.
func (o *FetchControlListResponse) SetMsg(v string) {
	o.Msg = &v
}

// GetControls returns the Controls field value if set, zero value otherwise.
func (o *FetchControlListResponse) GetControls() []FetchControlListControl {
	if o == nil || IsNil(o.Controls) {
		var ret []FetchControlListControl
		return ret
	}
	return o.Controls
}

// GetControlsOk returns a tuple with the Controls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchControlListResponse) GetControlsOk() ([]FetchControlListControl, bool) {
	if o == nil || IsNil(o.Controls) {
		return nil, false
	}
	return o.Controls, true
}

// HasControls returns a boolean if a field has been set.
func (o *FetchControlListResponse) HasControls() bool {
	if o != nil && !IsNil(o.Controls) {
		return true
	}

	return false
}

// SetControls gets a reference to the given []FetchControlListControl and assigns it to the Controls field.
func (o *FetchControlListResponse) SetControls(v []FetchControlListControl) {
	o.Controls = v
}

// GetDisplayCount returns the DisplayCount field value if set, zero value otherwise.
func (o *FetchControlListResponse) GetDisplayCount() int32 {
	if o == nil || IsNil(o.DisplayCount) {
		var ret int32
		return ret
	}
	return *o.DisplayCount
}

// GetDisplayCountOk returns a tuple with the DisplayCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchControlListResponse) GetDisplayCountOk() (*int32, bool) {
	if o == nil || IsNil(o.DisplayCount) {
		return nil, false
	}
	return o.DisplayCount, true
}

// HasDisplayCount returns a boolean if a field has been set.
func (o *FetchControlListResponse) HasDisplayCount() bool {
	if o != nil && !IsNil(o.DisplayCount) {
		return true
	}

	return false
}

// SetDisplayCount gets a reference to the given int32 and assigns it to the DisplayCount field.
func (o *FetchControlListResponse) SetDisplayCount(v int32) {
	o.DisplayCount = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *FetchControlListResponse) GetErrorCode() string {
	if o == nil || IsNil(o.ErrorCode) {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchControlListResponse) GetErrorCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCode) {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *FetchControlListResponse) HasErrorCode() bool {
	if o != nil && !IsNil(o.ErrorCode) {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *FetchControlListResponse) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetTotalCount returns the TotalCount field value if set, zero value otherwise.
func (o *FetchControlListResponse) GetTotalCount() int32 {
	if o == nil || IsNil(o.TotalCount) {
		var ret int32
		return ret
	}
	return *o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FetchControlListResponse) GetTotalCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalCount) {
		return nil, false
	}
	return o.TotalCount, true
}

// HasTotalCount returns a boolean if a field has been set.
func (o *FetchControlListResponse) HasTotalCount() bool {
	if o != nil && !IsNil(o.TotalCount) {
		return true
	}

	return false
}

// SetTotalCount gets a reference to the given int32 and assigns it to the TotalCount field.
func (o *FetchControlListResponse) SetTotalCount(v int32) {
	o.TotalCount = &v
}

func (o FetchControlListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FetchControlListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Msg) {
		toSerialize["msg"] = o.Msg
	}
	if !IsNil(o.Controls) {
		toSerialize["controls"] = o.Controls
	}
	if !IsNil(o.DisplayCount) {
		toSerialize["displayCount"] = o.DisplayCount
	}
	if !IsNil(o.ErrorCode) {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if !IsNil(o.TotalCount) {
		toSerialize["totalCount"] = o.TotalCount
	}
	return toSerialize, nil
}

type NullableFetchControlListResponse struct {
	value *FetchControlListResponse
	isSet bool
}

func (v NullableFetchControlListResponse) Get() *FetchControlListResponse {
	return v.value
}

func (v *NullableFetchControlListResponse) Set(val *FetchControlListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFetchControlListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFetchControlListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFetchControlListResponse(val *FetchControlListResponse) *NullableFetchControlListResponse {
	return &NullableFetchControlListResponse{value: val, isSet: true}
}

func (v NullableFetchControlListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFetchControlListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


