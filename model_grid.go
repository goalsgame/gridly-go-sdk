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

// Grid struct for Grid
type Grid struct {
	Columns *[]ViewColumn `json:"columns,omitempty"`
	DefaultAccessViewId *string `json:"defaultAccessViewId,omitempty"`
	Description *string `json:"description,omitempty"`
	Id *string `json:"id,omitempty"`
	Metadata *map[string]string `json:"metadata,omitempty"`
	Name *string `json:"name,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewGrid instantiates a new Grid object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGrid() *Grid {
	this := Grid{}
	return &this
}

// NewGridWithDefaults instantiates a new Grid object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGridWithDefaults() *Grid {
	this := Grid{}
	return &this
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *Grid) GetColumns() []ViewColumn {
	if o == nil || o.Columns == nil {
		var ret []ViewColumn
		return ret
	}
	return *o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grid) GetColumnsOk() (*[]ViewColumn, bool) {
	if o == nil || o.Columns == nil {
		return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *Grid) HasColumns() bool {
	if o != nil && o.Columns != nil {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []ViewColumn and assigns it to the Columns field.
func (o *Grid) SetColumns(v []ViewColumn) {
	o.Columns = &v
}

// GetDefaultAccessViewId returns the DefaultAccessViewId field value if set, zero value otherwise.
func (o *Grid) GetDefaultAccessViewId() string {
	if o == nil || o.DefaultAccessViewId == nil {
		var ret string
		return ret
	}
	return *o.DefaultAccessViewId
}

// GetDefaultAccessViewIdOk returns a tuple with the DefaultAccessViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grid) GetDefaultAccessViewIdOk() (*string, bool) {
	if o == nil || o.DefaultAccessViewId == nil {
		return nil, false
	}
	return o.DefaultAccessViewId, true
}

// HasDefaultAccessViewId returns a boolean if a field has been set.
func (o *Grid) HasDefaultAccessViewId() bool {
	if o != nil && o.DefaultAccessViewId != nil {
		return true
	}

	return false
}

// SetDefaultAccessViewId gets a reference to the given string and assigns it to the DefaultAccessViewId field.
func (o *Grid) SetDefaultAccessViewId(v string) {
	o.DefaultAccessViewId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Grid) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grid) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Grid) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Grid) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Grid) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grid) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Grid) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Grid) SetId(v string) {
	o.Id = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *Grid) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grid) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *Grid) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *Grid) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Grid) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grid) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Grid) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Grid) SetName(v string) {
	o.Name = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Grid) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Grid) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Grid) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Grid) SetStatus(v string) {
	o.Status = &v
}

func (o Grid) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Columns != nil {
		toSerialize["columns"] = o.Columns
	}
	if o.DefaultAccessViewId != nil {
		toSerialize["defaultAccessViewId"] = o.DefaultAccessViewId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableGrid struct {
	value *Grid
	isSet bool
}

func (v NullableGrid) Get() *Grid {
	return v.value
}

func (v *NullableGrid) Set(val *Grid) {
	v.value = val
	v.isSet = true
}

func (v NullableGrid) IsSet() bool {
	return v.isSet
}

func (v *NullableGrid) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGrid(val *Grid) *NullableGrid {
	return &NullableGrid{value: val, isSet: true}
}

func (v NullableGrid) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGrid) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


