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

// ProjectDetail struct for ProjectDetail
type ProjectDetail struct {
	Id *int64 `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Role *Role `json:"role,omitempty"`
	Databases []Database `json:"databases,omitempty"`
	Groups []Group `json:"groups,omitempty"`
}

// NewProjectDetail instantiates a new ProjectDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectDetail() *ProjectDetail {
	this := ProjectDetail{}
	return &this
}

// NewProjectDetailWithDefaults instantiates a new ProjectDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectDetailWithDefaults() *ProjectDetail {
	this := ProjectDetail{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ProjectDetail) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetail) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ProjectDetail) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ProjectDetail) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ProjectDetail) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetail) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ProjectDetail) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ProjectDetail) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ProjectDetail) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetail) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ProjectDetail) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ProjectDetail) SetDescription(v string) {
	o.Description = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *ProjectDetail) GetRole() Role {
	if o == nil || o.Role == nil {
		var ret Role
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetail) GetRoleOk() (*Role, bool) {
	if o == nil || o.Role == nil {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *ProjectDetail) HasRole() bool {
	if o != nil && o.Role != nil {
		return true
	}

	return false
}

// SetRole gets a reference to the given Role and assigns it to the Role field.
func (o *ProjectDetail) SetRole(v Role) {
	o.Role = &v
}

// GetDatabases returns the Databases field value if set, zero value otherwise.
func (o *ProjectDetail) GetDatabases() []Database {
	if o == nil || o.Databases == nil {
		var ret []Database
		return ret
	}
	return o.Databases
}

// GetDatabasesOk returns a tuple with the Databases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetail) GetDatabasesOk() ([]Database, bool) {
	if o == nil || o.Databases == nil {
		return nil, false
	}
	return o.Databases, true
}

// HasDatabases returns a boolean if a field has been set.
func (o *ProjectDetail) HasDatabases() bool {
	if o != nil && o.Databases != nil {
		return true
	}

	return false
}

// SetDatabases gets a reference to the given []Database and assigns it to the Databases field.
func (o *ProjectDetail) SetDatabases(v []Database) {
	o.Databases = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *ProjectDetail) GetGroups() []Group {
	if o == nil || o.Groups == nil {
		var ret []Group
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectDetail) GetGroupsOk() ([]Group, bool) {
	if o == nil || o.Groups == nil {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ProjectDetail) HasGroups() bool {
	if o != nil && o.Groups != nil {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []Group and assigns it to the Groups field.
func (o *ProjectDetail) SetGroups(v []Group) {
	o.Groups = v
}

func (o ProjectDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Role != nil {
		toSerialize["role"] = o.Role
	}
	if o.Databases != nil {
		toSerialize["databases"] = o.Databases
	}
	if o.Groups != nil {
		toSerialize["groups"] = o.Groups
	}
	return json.Marshal(toSerialize)
}

type NullableProjectDetail struct {
	value *ProjectDetail
	isSet bool
}

func (v NullableProjectDetail) Get() *ProjectDetail {
	return v.value
}

func (v *NullableProjectDetail) Set(val *ProjectDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectDetail(val *ProjectDetail) *NullableProjectDetail {
	return &NullableProjectDetail{value: val, isSet: true}
}

func (v NullableProjectDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


