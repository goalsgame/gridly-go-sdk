/*
Gridly API

Gridly API documentation

API version: 3.31.0
Contact: 
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gridly

import (
	"encoding/json"
)

// UpdateGrid struct for UpdateGrid
type UpdateGrid struct {
	Name *string `json:"name,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
}

// NewUpdateGrid instantiates a new UpdateGrid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateGrid() *UpdateGrid {
	this := UpdateGrid{}
	return &this
}

// NewUpdateGridWithDefaults instantiates a new UpdateGrid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateGridWithDefaults() *UpdateGrid {
	this := UpdateGrid{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateGrid) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGrid) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateGrid) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateGrid) SetName(v string) {
	o.Name = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UpdateGrid) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateGrid) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UpdateGrid) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *UpdateGrid) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

func (o UpdateGrid) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateGrid struct {
	value *UpdateGrid
	isSet bool
}

func (v NullableUpdateGrid) Get() *UpdateGrid {
	return v.value
}

func (v *NullableUpdateGrid) Set(val *UpdateGrid) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateGrid) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateGrid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateGrid(val *UpdateGrid) *NullableUpdateGrid {
	return &NullableUpdateGrid{value: val, isSet: true}
}

func (v NullableUpdateGrid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateGrid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


