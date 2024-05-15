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

// TranslationCount struct for TranslationCount
type TranslationCount struct {
	All *int64 `json:"all,omitempty"`
	OutOfDate *int64 `json:"outOfDate,omitempty"`
	Unset *int64 `json:"unset,omitempty"`
	UpToDate *int64 `json:"upToDate,omitempty"`
}

// NewTranslationCount instantiates a new TranslationCount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTranslationCount() *TranslationCount {
	this := TranslationCount{}
	return &this
}

// NewTranslationCountWithDefaults instantiates a new TranslationCount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTranslationCountWithDefaults() *TranslationCount {
	this := TranslationCount{}
	return &this
}

// GetAll returns the All field value if set, zero value otherwise.
func (o *TranslationCount) GetAll() int64 {
	if o == nil || isNil(o.All) {
		var ret int64
		return ret
	}
	return *o.All
}

// GetAllOk returns a tuple with the All field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranslationCount) GetAllOk() (*int64, bool) {
	if o == nil || isNil(o.All) {
    return nil, false
	}
	return o.All, true
}

// HasAll returns a boolean if a field has been set.
func (o *TranslationCount) HasAll() bool {
	if o != nil && !isNil(o.All) {
		return true
	}

	return false
}

// SetAll gets a reference to the given int64 and assigns it to the All field.
func (o *TranslationCount) SetAll(v int64) {
	o.All = &v
}

// GetOutOfDate returns the OutOfDate field value if set, zero value otherwise.
func (o *TranslationCount) GetOutOfDate() int64 {
	if o == nil || isNil(o.OutOfDate) {
		var ret int64
		return ret
	}
	return *o.OutOfDate
}

// GetOutOfDateOk returns a tuple with the OutOfDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranslationCount) GetOutOfDateOk() (*int64, bool) {
	if o == nil || isNil(o.OutOfDate) {
    return nil, false
	}
	return o.OutOfDate, true
}

// HasOutOfDate returns a boolean if a field has been set.
func (o *TranslationCount) HasOutOfDate() bool {
	if o != nil && !isNil(o.OutOfDate) {
		return true
	}

	return false
}

// SetOutOfDate gets a reference to the given int64 and assigns it to the OutOfDate field.
func (o *TranslationCount) SetOutOfDate(v int64) {
	o.OutOfDate = &v
}

// GetUnset returns the Unset field value if set, zero value otherwise.
func (o *TranslationCount) GetUnset() int64 {
	if o == nil || isNil(o.Unset) {
		var ret int64
		return ret
	}
	return *o.Unset
}

// GetUnsetOk returns a tuple with the Unset field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranslationCount) GetUnsetOk() (*int64, bool) {
	if o == nil || isNil(o.Unset) {
    return nil, false
	}
	return o.Unset, true
}

// HasUnset returns a boolean if a field has been set.
func (o *TranslationCount) HasUnset() bool {
	if o != nil && !isNil(o.Unset) {
		return true
	}

	return false
}

// SetUnset gets a reference to the given int64 and assigns it to the Unset field.
func (o *TranslationCount) SetUnset(v int64) {
	o.Unset = &v
}

// GetUpToDate returns the UpToDate field value if set, zero value otherwise.
func (o *TranslationCount) GetUpToDate() int64 {
	if o == nil || isNil(o.UpToDate) {
		var ret int64
		return ret
	}
	return *o.UpToDate
}

// GetUpToDateOk returns a tuple with the UpToDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TranslationCount) GetUpToDateOk() (*int64, bool) {
	if o == nil || isNil(o.UpToDate) {
    return nil, false
	}
	return o.UpToDate, true
}

// HasUpToDate returns a boolean if a field has been set.
func (o *TranslationCount) HasUpToDate() bool {
	if o != nil && !isNil(o.UpToDate) {
		return true
	}

	return false
}

// SetUpToDate gets a reference to the given int64 and assigns it to the UpToDate field.
func (o *TranslationCount) SetUpToDate(v int64) {
	o.UpToDate = &v
}

func (o TranslationCount) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.All) {
		toSerialize["all"] = o.All
	}
	if !isNil(o.OutOfDate) {
		toSerialize["outOfDate"] = o.OutOfDate
	}
	if !isNil(o.Unset) {
		toSerialize["unset"] = o.Unset
	}
	if !isNil(o.UpToDate) {
		toSerialize["upToDate"] = o.UpToDate
	}
	return json.Marshal(toSerialize)
}

type NullableTranslationCount struct {
	value *TranslationCount
	isSet bool
}

func (v NullableTranslationCount) Get() *TranslationCount {
	return v.value
}

func (v *NullableTranslationCount) Set(val *TranslationCount) {
	v.value = val
	v.isSet = true
}

func (v NullableTranslationCount) IsSet() bool {
	return v.isSet
}

func (v *NullableTranslationCount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTranslationCount(val *TranslationCount) *NullableTranslationCount {
	return &NullableTranslationCount{value: val, isSet: true}
}

func (v NullableTranslationCount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTranslationCount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


