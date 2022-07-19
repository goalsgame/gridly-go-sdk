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

// SetRecord struct for SetRecord
type SetRecord struct {
	Id *string `json:"id,omitempty"`
	Cells []SetCell `json:"cells,omitempty"`
	Path *string `json:"path,omitempty"`
}

// NewSetRecord instantiates a new SetRecord object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetRecord() *SetRecord {
	this := SetRecord{}
	return &this
}

// NewSetRecordWithDefaults instantiates a new SetRecord object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetRecordWithDefaults() *SetRecord {
	this := SetRecord{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SetRecord) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetRecord) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SetRecord) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SetRecord) SetId(v string) {
	o.Id = &v
}

// GetCells returns the Cells field value if set, zero value otherwise.
func (o *SetRecord) GetCells() []SetCell {
	if o == nil || o.Cells == nil {
		var ret []SetCell
		return ret
	}
	return o.Cells
}

// GetCellsOk returns a tuple with the Cells field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetRecord) GetCellsOk() ([]SetCell, bool) {
	if o == nil || o.Cells == nil {
		return nil, false
	}
	return o.Cells, true
}

// HasCells returns a boolean if a field has been set.
func (o *SetRecord) HasCells() bool {
	if o != nil && o.Cells != nil {
		return true
	}

	return false
}

// SetCells gets a reference to the given []SetCell and assigns it to the Cells field.
func (o *SetRecord) SetCells(v []SetCell) {
	o.Cells = v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *SetRecord) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetRecord) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SetRecord) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *SetRecord) SetPath(v string) {
	o.Path = &v
}

func (o SetRecord) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Cells != nil {
		toSerialize["cells"] = o.Cells
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableSetRecord struct {
	value *SetRecord
	isSet bool
}

func (v NullableSetRecord) Get() *SetRecord {
	return v.value
}

func (v *NullableSetRecord) Set(val *SetRecord) {
	v.value = val
	v.isSet = true
}

func (v NullableSetRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableSetRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetRecord(val *SetRecord) *NullableSetRecord {
	return &NullableSetRecord{value: val, isSet: true}
}

func (v NullableSetRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


