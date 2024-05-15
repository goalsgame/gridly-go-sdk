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

// CreateView struct for CreateView
type CreateView struct {
	Name *string `json:"name,omitempty"`
	GridId *string `json:"gridId,omitempty"`
	Columns []AddViewColumn `json:"columns,omitempty"`
}

// NewCreateView instantiates a new CreateView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateView() *CreateView {
	this := CreateView{}
	return &this
}

// NewCreateViewWithDefaults instantiates a new CreateView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateViewWithDefaults() *CreateView {
	this := CreateView{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateView) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateView) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateView) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateView) SetName(v string) {
	o.Name = &v
}

// GetGridId returns the GridId field value if set, zero value otherwise.
func (o *CreateView) GetGridId() string {
	if o == nil || isNil(o.GridId) {
		var ret string
		return ret
	}
	return *o.GridId
}

// GetGridIdOk returns a tuple with the GridId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateView) GetGridIdOk() (*string, bool) {
	if o == nil || isNil(o.GridId) {
    return nil, false
	}
	return o.GridId, true
}

// HasGridId returns a boolean if a field has been set.
func (o *CreateView) HasGridId() bool {
	if o != nil && !isNil(o.GridId) {
		return true
	}

	return false
}

// SetGridId gets a reference to the given string and assigns it to the GridId field.
func (o *CreateView) SetGridId(v string) {
	o.GridId = &v
}

// GetColumns returns the Columns field value if set, zero value otherwise.
func (o *CreateView) GetColumns() []AddViewColumn {
	if o == nil || isNil(o.Columns) {
		var ret []AddViewColumn
		return ret
	}
	return o.Columns
}

// GetColumnsOk returns a tuple with the Columns field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateView) GetColumnsOk() ([]AddViewColumn, bool) {
	if o == nil || isNil(o.Columns) {
    return nil, false
	}
	return o.Columns, true
}

// HasColumns returns a boolean if a field has been set.
func (o *CreateView) HasColumns() bool {
	if o != nil && !isNil(o.Columns) {
		return true
	}

	return false
}

// SetColumns gets a reference to the given []AddViewColumn and assigns it to the Columns field.
func (o *CreateView) SetColumns(v []AddViewColumn) {
	o.Columns = v
}

func (o CreateView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.GridId) {
		toSerialize["gridId"] = o.GridId
	}
	if !isNil(o.Columns) {
		toSerialize["columns"] = o.Columns
	}
	return json.Marshal(toSerialize)
}

type NullableCreateView struct {
	value *CreateView
	isSet bool
}

func (v NullableCreateView) Get() *CreateView {
	return v.value
}

func (v *NullableCreateView) Set(val *CreateView) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateView) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateView(val *CreateView) *NullableCreateView {
	return &NullableCreateView{value: val, isSet: true}
}

func (v NullableCreateView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


