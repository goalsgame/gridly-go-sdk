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

// CreateDatabase body
type CreateDatabase struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	EnableGuidRecord *bool `json:"enableGuidRecord,omitempty"`
	Id *string `json:"id,omitempty"`
}

// NewCreateDatabase instantiates a new CreateDatabase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateDatabase(name string) *CreateDatabase {
	this := CreateDatabase{}
	this.Name = name
	return &this
}

// NewCreateDatabaseWithDefaults instantiates a new CreateDatabase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateDatabaseWithDefaults() *CreateDatabase {
	this := CreateDatabase{}
	return &this
}

// GetName returns the Name field value
func (o *CreateDatabase) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateDatabase) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateDatabase) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateDatabase) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatabase) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateDatabase) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateDatabase) SetDescription(v string) {
	o.Description = &v
}

// GetEnableGuidRecord returns the EnableGuidRecord field value if set, zero value otherwise.
func (o *CreateDatabase) GetEnableGuidRecord() bool {
	if o == nil || isNil(o.EnableGuidRecord) {
		var ret bool
		return ret
	}
	return *o.EnableGuidRecord
}

// GetEnableGuidRecordOk returns a tuple with the EnableGuidRecord field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatabase) GetEnableGuidRecordOk() (*bool, bool) {
	if o == nil || isNil(o.EnableGuidRecord) {
    return nil, false
	}
	return o.EnableGuidRecord, true
}

// HasEnableGuidRecord returns a boolean if a field has been set.
func (o *CreateDatabase) HasEnableGuidRecord() bool {
	if o != nil && !isNil(o.EnableGuidRecord) {
		return true
	}

	return false
}

// SetEnableGuidRecord gets a reference to the given bool and assigns it to the EnableGuidRecord field.
func (o *CreateDatabase) SetEnableGuidRecord(v bool) {
	o.EnableGuidRecord = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CreateDatabase) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateDatabase) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
    return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CreateDatabase) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CreateDatabase) SetId(v string) {
	o.Id = &v
}

func (o CreateDatabase) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.EnableGuidRecord) {
		toSerialize["enableGuidRecord"] = o.EnableGuidRecord
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableCreateDatabase struct {
	value *CreateDatabase
	isSet bool
}

func (v NullableCreateDatabase) Get() *CreateDatabase {
	return v.value
}

func (v *NullableCreateDatabase) Set(val *CreateDatabase) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateDatabase) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateDatabase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateDatabase(val *CreateDatabase) *NullableCreateDatabase {
	return &NullableCreateDatabase{value: val, isSet: true}
}

func (v NullableCreateDatabase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateDatabase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


