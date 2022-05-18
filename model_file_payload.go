/*
Elastic Email REST API

This API is based on the REST API architecture, allowing the user to easily manage their data with this resource-based approach.    Every API call is established on which specific request type (GET, POST, PUT, DELETE) will be used.    The API has a limit of 20 concurrent connections and a hard timeout of 600 seconds per request.    To start using this API, you will need your Access Token (available <a target=\"_blank\" href=\"https://elasticemail.com/account#/settings/new/manage-api\">here</a>). Remember to keep it safe. Required access levels are listed in the given request’s description.    This is the documentation for REST API. If you’d like to read our legacy documentation regarding Web API v2 click <a target=\"_blank\" href=\"https://api.elasticemail.com/public/help\">here</a>.    Downloadable library clients can be found in our Github repository <a target=\"_blank\" href=\"https://github.com/ElasticEmail?tab=repositories&q=%22rest+api%22+in%3Areadme\">here</a>

API version: 4.0.0
Contact: support@elasticemail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// FilePayload struct for FilePayload
type FilePayload struct {
	// Content of the file sent as binary data
	BinaryContent string `json:"BinaryContent"`
	// Filename
	Name *string `json:"Name,omitempty"`
	// Type of file's content (e.g. image/jpeg)
	ContentType *string `json:"ContentType,omitempty"`
}

// NewFilePayload instantiates a new FilePayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFilePayload(binaryContent string) *FilePayload {
	this := FilePayload{}
	this.BinaryContent = binaryContent
	return &this
}

// NewFilePayloadWithDefaults instantiates a new FilePayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFilePayloadWithDefaults() *FilePayload {
	this := FilePayload{}
	return &this
}

// GetBinaryContent returns the BinaryContent field value
func (o *FilePayload) GetBinaryContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BinaryContent
}

// GetBinaryContentOk returns a tuple with the BinaryContent field value
// and a boolean to check if the value has been set.
func (o *FilePayload) GetBinaryContentOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.BinaryContent, true
}

// SetBinaryContent sets field value
func (o *FilePayload) SetBinaryContent(v string) {
	o.BinaryContent = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FilePayload) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilePayload) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FilePayload) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FilePayload) SetName(v string) {
	o.Name = &v
}

// GetContentType returns the ContentType field value if set, zero value otherwise.
func (o *FilePayload) GetContentType() string {
	if o == nil || o.ContentType == nil {
		var ret string
		return ret
	}
	return *o.ContentType
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FilePayload) GetContentTypeOk() (*string, bool) {
	if o == nil || o.ContentType == nil {
		return nil, false
	}
	return o.ContentType, true
}

// HasContentType returns a boolean if a field has been set.
func (o *FilePayload) HasContentType() bool {
	if o != nil && o.ContentType != nil {
		return true
	}

	return false
}

// SetContentType gets a reference to the given string and assigns it to the ContentType field.
func (o *FilePayload) SetContentType(v string) {
	o.ContentType = &v
}

func (o FilePayload) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["BinaryContent"] = o.BinaryContent
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.ContentType != nil {
		toSerialize["ContentType"] = o.ContentType
	}
	return json.Marshal(toSerialize)
}

type NullableFilePayload struct {
	value *FilePayload
	isSet bool
}

func (v NullableFilePayload) Get() *FilePayload {
	return v.value
}

func (v *NullableFilePayload) Set(val *FilePayload) {
	v.value = val
	v.isSet = true
}

func (v NullableFilePayload) IsSet() bool {
	return v.isSet
}

func (v *NullableFilePayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFilePayload(val *FilePayload) *NullableFilePayload {
	return &NullableFilePayload{value: val, isSet: true}
}

func (v NullableFilePayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFilePayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

