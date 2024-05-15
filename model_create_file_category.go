/*
Gridly API

Gridly API documentation

API version: 4.33.0
Contact: support@gridly.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gridly

import (
	"encoding/json"
)

// CreateFileCategory struct for CreateFileCategory
type CreateFileCategory struct {
	Name string `json:"name"`
}

// NewCreateFileCategory instantiates a new CreateFileCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateFileCategory(name string) *CreateFileCategory {
	this := CreateFileCategory{}
	this.Name = name
	return &this
}

// NewCreateFileCategoryWithDefaults instantiates a new CreateFileCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateFileCategoryWithDefaults() *CreateFileCategory {
	this := CreateFileCategory{}
	return &this
}

// GetName returns the Name field value
func (o *CreateFileCategory) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateFileCategory) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateFileCategory) SetName(v string) {
	o.Name = v
}

func (o CreateFileCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableCreateFileCategory struct {
	value *CreateFileCategory
	isSet bool
}

func (v NullableCreateFileCategory) Get() *CreateFileCategory {
	return v.value
}

func (v *NullableCreateFileCategory) Set(val *CreateFileCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateFileCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateFileCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateFileCategory(val *CreateFileCategory) *NullableCreateFileCategory {
	return &NullableCreateFileCategory{value: val, isSet: true}
}

func (v NullableCreateFileCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateFileCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


