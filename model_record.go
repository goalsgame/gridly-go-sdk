/*
Gridly API

Gridly API documentation

API version: 3.21.4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package gridly

import (
	"encoding/json"
	"time"
)

// Record struct for Record
type Record struct {
	Cells *[]Cell `json:"cells,omitempty"`
	Id *string `json:"id,omitempty"`
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`
	LastModifiedTime *time.Time `json:"lastModifiedTime,omitempty"`
	Path *string `json:"path,omitempty"`
}

// NewRecord instantiates a new Record object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRecord() *Record {
	this := Record{}
	return &this
}

// NewRecordWithDefaults instantiates a new Record object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRecordWithDefaults() *Record {
	this := Record{}
	return &this
}

// GetCells returns the Cells field value if set, zero value otherwise.
func (o *Record) GetCells() []Cell {
	if o == nil || o.Cells == nil {
		var ret []Cell
		return ret
	}
	return *o.Cells
}

// GetCellsOk returns a tuple with the Cells field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Record) GetCellsOk() (*[]Cell, bool) {
	if o == nil || o.Cells == nil {
		return nil, false
	}
	return o.Cells, true
}

// HasCells returns a boolean if a field has been set.
func (o *Record) HasCells() bool {
	if o != nil && o.Cells != nil {
		return true
	}

	return false
}

// SetCells gets a reference to the given []Cell and assigns it to the Cells field.
func (o *Record) SetCells(v []Cell) {
	o.Cells = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Record) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Record) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Record) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Record) SetId(v string) {
	o.Id = &v
}

// GetLastModifiedBy returns the LastModifiedBy field value if set, zero value otherwise.
func (o *Record) GetLastModifiedBy() string {
	if o == nil || o.LastModifiedBy == nil {
		var ret string
		return ret
	}
	return *o.LastModifiedBy
}

// GetLastModifiedByOk returns a tuple with the LastModifiedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Record) GetLastModifiedByOk() (*string, bool) {
	if o == nil || o.LastModifiedBy == nil {
		return nil, false
	}
	return o.LastModifiedBy, true
}

// HasLastModifiedBy returns a boolean if a field has been set.
func (o *Record) HasLastModifiedBy() bool {
	if o != nil && o.LastModifiedBy != nil {
		return true
	}

	return false
}

// SetLastModifiedBy gets a reference to the given string and assigns it to the LastModifiedBy field.
func (o *Record) SetLastModifiedBy(v string) {
	o.LastModifiedBy = &v
}

// GetLastModifiedTime returns the LastModifiedTime field value if set, zero value otherwise.
func (o *Record) GetLastModifiedTime() time.Time {
	if o == nil || o.LastModifiedTime == nil {
		var ret time.Time
		return ret
	}
	return *o.LastModifiedTime
}

// GetLastModifiedTimeOk returns a tuple with the LastModifiedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Record) GetLastModifiedTimeOk() (*time.Time, bool) {
	if o == nil || o.LastModifiedTime == nil {
		return nil, false
	}
	return o.LastModifiedTime, true
}

// HasLastModifiedTime returns a boolean if a field has been set.
func (o *Record) HasLastModifiedTime() bool {
	if o != nil && o.LastModifiedTime != nil {
		return true
	}

	return false
}

// SetLastModifiedTime gets a reference to the given time.Time and assigns it to the LastModifiedTime field.
func (o *Record) SetLastModifiedTime(v time.Time) {
	o.LastModifiedTime = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *Record) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Record) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *Record) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *Record) SetPath(v string) {
	o.Path = &v
}

func (o Record) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cells != nil {
		toSerialize["cells"] = o.Cells
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.LastModifiedBy != nil {
		toSerialize["lastModifiedBy"] = o.LastModifiedBy
	}
	if o.LastModifiedTime != nil {
		toSerialize["lastModifiedTime"] = o.LastModifiedTime
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableRecord struct {
	value *Record
	isSet bool
}

func (v NullableRecord) Get() *Record {
	return v.value
}

func (v *NullableRecord) Set(val *Record) {
	v.value = val
	v.isSet = true
}

func (v NullableRecord) IsSet() bool {
	return v.isSet
}

func (v *NullableRecord) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRecord(val *Record) *NullableRecord {
	return &NullableRecord{value: val, isSet: true}
}

func (v NullableRecord) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRecord) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


