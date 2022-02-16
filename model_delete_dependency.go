/*
Gridly API

Gridly API documentation

API version: 3.21.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gridly

import (
	"encoding/json"
)

// DeleteDependency struct for DeleteDependency
type DeleteDependency struct {
	Ids *[]string `json:"ids,omitempty"`
}

// NewDeleteDependency instantiates a new DeleteDependency object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteDependency() *DeleteDependency {
	this := DeleteDependency{}
	return &this
}

// NewDeleteDependencyWithDefaults instantiates a new DeleteDependency object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteDependencyWithDefaults() *DeleteDependency {
	this := DeleteDependency{}
	return &this
}

// GetIds returns the Ids field value if set, zero value otherwise.
func (o *DeleteDependency) GetIds() []string {
	if o == nil || o.Ids == nil {
		var ret []string
		return ret
	}
	return *o.Ids
}

// GetIdsOk returns a tuple with the Ids field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteDependency) GetIdsOk() (*[]string, bool) {
	if o == nil || o.Ids == nil {
		return nil, false
	}
	return o.Ids, true
}

// HasIds returns a boolean if a field has been set.
func (o *DeleteDependency) HasIds() bool {
	if o != nil && o.Ids != nil {
		return true
	}

	return false
}

// SetIds gets a reference to the given []string and assigns it to the Ids field.
func (o *DeleteDependency) SetIds(v []string) {
	o.Ids = &v
}

func (o DeleteDependency) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Ids != nil {
		toSerialize["ids"] = o.Ids
	}
	return json.Marshal(toSerialize)
}

type NullableDeleteDependency struct {
	value *DeleteDependency
	isSet bool
}

func (v NullableDeleteDependency) Get() *DeleteDependency {
	return v.value
}

func (v *NullableDeleteDependency) Set(val *DeleteDependency) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteDependency) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteDependency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteDependency(val *DeleteDependency) *NullableDeleteDependency {
	return &NullableDeleteDependency{value: val, isSet: true}
}

func (v NullableDeleteDependency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteDependency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


