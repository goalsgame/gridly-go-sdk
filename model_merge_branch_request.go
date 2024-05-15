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

// MergeBranchRequest struct for MergeBranchRequest
type MergeBranchRequest struct {
	MergeRecordOptions []string `json:"mergeRecordOptions,omitempty"`
	MergeRecordConflicts []MergeRecordConflict `json:"mergeRecordConflicts,omitempty"`
	UseLastMergeResolve *bool `json:"useLastMergeResolve,omitempty"`
	Query []FilterField `json:"query,omitempty"`
}

// NewMergeBranchRequest instantiates a new MergeBranchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMergeBranchRequest() *MergeBranchRequest {
	this := MergeBranchRequest{}
	return &this
}

// NewMergeBranchRequestWithDefaults instantiates a new MergeBranchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMergeBranchRequestWithDefaults() *MergeBranchRequest {
	this := MergeBranchRequest{}
	return &this
}

// GetMergeRecordOptions returns the MergeRecordOptions field value if set, zero value otherwise.
func (o *MergeBranchRequest) GetMergeRecordOptions() []string {
	if o == nil || isNil(o.MergeRecordOptions) {
		var ret []string
		return ret
	}
	return o.MergeRecordOptions
}

// GetMergeRecordOptionsOk returns a tuple with the MergeRecordOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeBranchRequest) GetMergeRecordOptionsOk() ([]string, bool) {
	if o == nil || isNil(o.MergeRecordOptions) {
    return nil, false
	}
	return o.MergeRecordOptions, true
}

// HasMergeRecordOptions returns a boolean if a field has been set.
func (o *MergeBranchRequest) HasMergeRecordOptions() bool {
	if o != nil && !isNil(o.MergeRecordOptions) {
		return true
	}

	return false
}

// SetMergeRecordOptions gets a reference to the given []string and assigns it to the MergeRecordOptions field.
func (o *MergeBranchRequest) SetMergeRecordOptions(v []string) {
	o.MergeRecordOptions = v
}

// GetMergeRecordConflicts returns the MergeRecordConflicts field value if set, zero value otherwise.
func (o *MergeBranchRequest) GetMergeRecordConflicts() []MergeRecordConflict {
	if o == nil || isNil(o.MergeRecordConflicts) {
		var ret []MergeRecordConflict
		return ret
	}
	return o.MergeRecordConflicts
}

// GetMergeRecordConflictsOk returns a tuple with the MergeRecordConflicts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeBranchRequest) GetMergeRecordConflictsOk() ([]MergeRecordConflict, bool) {
	if o == nil || isNil(o.MergeRecordConflicts) {
    return nil, false
	}
	return o.MergeRecordConflicts, true
}

// HasMergeRecordConflicts returns a boolean if a field has been set.
func (o *MergeBranchRequest) HasMergeRecordConflicts() bool {
	if o != nil && !isNil(o.MergeRecordConflicts) {
		return true
	}

	return false
}

// SetMergeRecordConflicts gets a reference to the given []MergeRecordConflict and assigns it to the MergeRecordConflicts field.
func (o *MergeBranchRequest) SetMergeRecordConflicts(v []MergeRecordConflict) {
	o.MergeRecordConflicts = v
}

// GetUseLastMergeResolve returns the UseLastMergeResolve field value if set, zero value otherwise.
func (o *MergeBranchRequest) GetUseLastMergeResolve() bool {
	if o == nil || isNil(o.UseLastMergeResolve) {
		var ret bool
		return ret
	}
	return *o.UseLastMergeResolve
}

// GetUseLastMergeResolveOk returns a tuple with the UseLastMergeResolve field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeBranchRequest) GetUseLastMergeResolveOk() (*bool, bool) {
	if o == nil || isNil(o.UseLastMergeResolve) {
    return nil, false
	}
	return o.UseLastMergeResolve, true
}

// HasUseLastMergeResolve returns a boolean if a field has been set.
func (o *MergeBranchRequest) HasUseLastMergeResolve() bool {
	if o != nil && !isNil(o.UseLastMergeResolve) {
		return true
	}

	return false
}

// SetUseLastMergeResolve gets a reference to the given bool and assigns it to the UseLastMergeResolve field.
func (o *MergeBranchRequest) SetUseLastMergeResolve(v bool) {
	o.UseLastMergeResolve = &v
}

// GetQuery returns the Query field value if set, zero value otherwise.
func (o *MergeBranchRequest) GetQuery() []FilterField {
	if o == nil || isNil(o.Query) {
		var ret []FilterField
		return ret
	}
	return o.Query
}

// GetQueryOk returns a tuple with the Query field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MergeBranchRequest) GetQueryOk() ([]FilterField, bool) {
	if o == nil || isNil(o.Query) {
    return nil, false
	}
	return o.Query, true
}

// HasQuery returns a boolean if a field has been set.
func (o *MergeBranchRequest) HasQuery() bool {
	if o != nil && !isNil(o.Query) {
		return true
	}

	return false
}

// SetQuery gets a reference to the given []FilterField and assigns it to the Query field.
func (o *MergeBranchRequest) SetQuery(v []FilterField) {
	o.Query = v
}

func (o MergeBranchRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MergeRecordOptions) {
		toSerialize["mergeRecordOptions"] = o.MergeRecordOptions
	}
	if !isNil(o.MergeRecordConflicts) {
		toSerialize["mergeRecordConflicts"] = o.MergeRecordConflicts
	}
	if !isNil(o.UseLastMergeResolve) {
		toSerialize["useLastMergeResolve"] = o.UseLastMergeResolve
	}
	if !isNil(o.Query) {
		toSerialize["query"] = o.Query
	}
	return json.Marshal(toSerialize)
}

type NullableMergeBranchRequest struct {
	value *MergeBranchRequest
	isSet bool
}

func (v NullableMergeBranchRequest) Get() *MergeBranchRequest {
	return v.value
}

func (v *NullableMergeBranchRequest) Set(val *MergeBranchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMergeBranchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMergeBranchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMergeBranchRequest(val *MergeBranchRequest) *NullableMergeBranchRequest {
	return &NullableMergeBranchRequest{value: val, isSet: true}
}

func (v NullableMergeBranchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMergeBranchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


