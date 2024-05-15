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

// MovePath struct for MovePath
type MovePath struct {
	Names []string `json:"names,omitempty"`
	FromParentPath *string `json:"fromParentPath,omitempty"`
	ToParentPath *string `json:"toParentPath,omitempty"`
	MoveBefore *string `json:"moveBefore,omitempty"`
	MoveAfter *string `json:"moveAfter,omitempty"`
}

// NewMovePath instantiates a new MovePath object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMovePath() *MovePath {
	this := MovePath{}
	return &this
}

// NewMovePathWithDefaults instantiates a new MovePath object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMovePathWithDefaults() *MovePath {
	this := MovePath{}
	return &this
}

// GetNames returns the Names field value if set, zero value otherwise.
func (o *MovePath) GetNames() []string {
	if o == nil || isNil(o.Names) {
		var ret []string
		return ret
	}
	return o.Names
}

// GetNamesOk returns a tuple with the Names field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovePath) GetNamesOk() ([]string, bool) {
	if o == nil || isNil(o.Names) {
    return nil, false
	}
	return o.Names, true
}

// HasNames returns a boolean if a field has been set.
func (o *MovePath) HasNames() bool {
	if o != nil && !isNil(o.Names) {
		return true
	}

	return false
}

// SetNames gets a reference to the given []string and assigns it to the Names field.
func (o *MovePath) SetNames(v []string) {
	o.Names = v
}

// GetFromParentPath returns the FromParentPath field value if set, zero value otherwise.
func (o *MovePath) GetFromParentPath() string {
	if o == nil || isNil(o.FromParentPath) {
		var ret string
		return ret
	}
	return *o.FromParentPath
}

// GetFromParentPathOk returns a tuple with the FromParentPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovePath) GetFromParentPathOk() (*string, bool) {
	if o == nil || isNil(o.FromParentPath) {
    return nil, false
	}
	return o.FromParentPath, true
}

// HasFromParentPath returns a boolean if a field has been set.
func (o *MovePath) HasFromParentPath() bool {
	if o != nil && !isNil(o.FromParentPath) {
		return true
	}

	return false
}

// SetFromParentPath gets a reference to the given string and assigns it to the FromParentPath field.
func (o *MovePath) SetFromParentPath(v string) {
	o.FromParentPath = &v
}

// GetToParentPath returns the ToParentPath field value if set, zero value otherwise.
func (o *MovePath) GetToParentPath() string {
	if o == nil || isNil(o.ToParentPath) {
		var ret string
		return ret
	}
	return *o.ToParentPath
}

// GetToParentPathOk returns a tuple with the ToParentPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovePath) GetToParentPathOk() (*string, bool) {
	if o == nil || isNil(o.ToParentPath) {
    return nil, false
	}
	return o.ToParentPath, true
}

// HasToParentPath returns a boolean if a field has been set.
func (o *MovePath) HasToParentPath() bool {
	if o != nil && !isNil(o.ToParentPath) {
		return true
	}

	return false
}

// SetToParentPath gets a reference to the given string and assigns it to the ToParentPath field.
func (o *MovePath) SetToParentPath(v string) {
	o.ToParentPath = &v
}

// GetMoveBefore returns the MoveBefore field value if set, zero value otherwise.
func (o *MovePath) GetMoveBefore() string {
	if o == nil || isNil(o.MoveBefore) {
		var ret string
		return ret
	}
	return *o.MoveBefore
}

// GetMoveBeforeOk returns a tuple with the MoveBefore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovePath) GetMoveBeforeOk() (*string, bool) {
	if o == nil || isNil(o.MoveBefore) {
    return nil, false
	}
	return o.MoveBefore, true
}

// HasMoveBefore returns a boolean if a field has been set.
func (o *MovePath) HasMoveBefore() bool {
	if o != nil && !isNil(o.MoveBefore) {
		return true
	}

	return false
}

// SetMoveBefore gets a reference to the given string and assigns it to the MoveBefore field.
func (o *MovePath) SetMoveBefore(v string) {
	o.MoveBefore = &v
}

// GetMoveAfter returns the MoveAfter field value if set, zero value otherwise.
func (o *MovePath) GetMoveAfter() string {
	if o == nil || isNil(o.MoveAfter) {
		var ret string
		return ret
	}
	return *o.MoveAfter
}

// GetMoveAfterOk returns a tuple with the MoveAfter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MovePath) GetMoveAfterOk() (*string, bool) {
	if o == nil || isNil(o.MoveAfter) {
    return nil, false
	}
	return o.MoveAfter, true
}

// HasMoveAfter returns a boolean if a field has been set.
func (o *MovePath) HasMoveAfter() bool {
	if o != nil && !isNil(o.MoveAfter) {
		return true
	}

	return false
}

// SetMoveAfter gets a reference to the given string and assigns it to the MoveAfter field.
func (o *MovePath) SetMoveAfter(v string) {
	o.MoveAfter = &v
}

func (o MovePath) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Names) {
		toSerialize["names"] = o.Names
	}
	if !isNil(o.FromParentPath) {
		toSerialize["fromParentPath"] = o.FromParentPath
	}
	if !isNil(o.ToParentPath) {
		toSerialize["toParentPath"] = o.ToParentPath
	}
	if !isNil(o.MoveBefore) {
		toSerialize["moveBefore"] = o.MoveBefore
	}
	if !isNil(o.MoveAfter) {
		toSerialize["moveAfter"] = o.MoveAfter
	}
	return json.Marshal(toSerialize)
}

type NullableMovePath struct {
	value *MovePath
	isSet bool
}

func (v NullableMovePath) Get() *MovePath {
	return v.value
}

func (v *NullableMovePath) Set(val *MovePath) {
	v.value = val
	v.isSet = true
}

func (v NullableMovePath) IsSet() bool {
	return v.isSet
}

func (v *NullableMovePath) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMovePath(val *MovePath) *NullableMovePath {
	return &NullableMovePath{value: val, isSet: true}
}

func (v NullableMovePath) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMovePath) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


