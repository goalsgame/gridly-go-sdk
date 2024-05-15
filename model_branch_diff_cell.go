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

// BranchDiffCell struct for BranchDiffCell
type BranchDiffCell struct {
	SourceDependencyStatus *string `json:"sourceDependencyStatus,omitempty"`
	DestinationDependencyStatus *string `json:"destinationDependencyStatus,omitempty"`
	SourceSourceStatus *string `json:"sourceSourceStatus,omitempty"`
	DestinationSourceStatus *string `json:"destinationSourceStatus,omitempty"`
	SourceLengthSetting *int32 `json:"sourceLengthSetting,omitempty"`
	DestinationLengthSetting *int32 `json:"destinationLengthSetting,omitempty"`
	Status *string `json:"status,omitempty"`
	ColumnId *string `json:"columnId,omitempty"`
	SourceValue map[string]interface{} `json:"sourceValue,omitempty"`
	DestinationValue map[string]interface{} `json:"destinationValue,omitempty"`
}

// NewBranchDiffCell instantiates a new BranchDiffCell object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBranchDiffCell() *BranchDiffCell {
	this := BranchDiffCell{}
	return &this
}

// NewBranchDiffCellWithDefaults instantiates a new BranchDiffCell object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBranchDiffCellWithDefaults() *BranchDiffCell {
	this := BranchDiffCell{}
	return &this
}

// GetSourceDependencyStatus returns the SourceDependencyStatus field value if set, zero value otherwise.
func (o *BranchDiffCell) GetSourceDependencyStatus() string {
	if o == nil || isNil(o.SourceDependencyStatus) {
		var ret string
		return ret
	}
	return *o.SourceDependencyStatus
}

// GetSourceDependencyStatusOk returns a tuple with the SourceDependencyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchDiffCell) GetSourceDependencyStatusOk() (*string, bool) {
	if o == nil || isNil(o.SourceDependencyStatus) {
    return nil, false
	}
	return o.SourceDependencyStatus, true
}

// HasSourceDependencyStatus returns a boolean if a field has been set.
func (o *BranchDiffCell) HasSourceDependencyStatus() bool {
	if o != nil && !isNil(o.SourceDependencyStatus) {
		return true
	}

	return false
}

// SetSourceDependencyStatus gets a reference to the given string and assigns it to the SourceDependencyStatus field.
func (o *BranchDiffCell) SetSourceDependencyStatus(v string) {
	o.SourceDependencyStatus = &v
}

// GetDestinationDependencyStatus returns the DestinationDependencyStatus field value if set, zero value otherwise.
func (o *BranchDiffCell) GetDestinationDependencyStatus() string {
	if o == nil || isNil(o.DestinationDependencyStatus) {
		var ret string
		return ret
	}
	return *o.DestinationDependencyStatus
}

// GetDestinationDependencyStatusOk returns a tuple with the DestinationDependencyStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchDiffCell) GetDestinationDependencyStatusOk() (*string, bool) {
	if o == nil || isNil(o.DestinationDependencyStatus) {
    return nil, false
	}
	return o.DestinationDependencyStatus, true
}

// HasDestinationDependencyStatus returns a boolean if a field has been set.
func (o *BranchDiffCell) HasDestinationDependencyStatus() bool {
	if o != nil && !isNil(o.DestinationDependencyStatus) {
		return true
	}

	return false
}

// SetDestinationDependencyStatus gets a reference to the given string and assigns it to the DestinationDependencyStatus field.
func (o *BranchDiffCell) SetDestinationDependencyStatus(v string) {
	o.DestinationDependencyStatus = &v
}

// GetSourceSourceStatus returns the SourceSourceStatus field value if set, zero value otherwise.
func (o *BranchDiffCell) GetSourceSourceStatus() string {
	if o == nil || isNil(o.SourceSourceStatus) {
		var ret string
		return ret
	}
	return *o.SourceSourceStatus
}

// GetSourceSourceStatusOk returns a tuple with the SourceSourceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchDiffCell) GetSourceSourceStatusOk() (*string, bool) {
	if o == nil || isNil(o.SourceSourceStatus) {
    return nil, false
	}
	return o.SourceSourceStatus, true
}

// HasSourceSourceStatus returns a boolean if a field has been set.
func (o *BranchDiffCell) HasSourceSourceStatus() bool {
	if o != nil && !isNil(o.SourceSourceStatus) {
		return true
	}

	return false
}

// SetSourceSourceStatus gets a reference to the given string and assigns it to the SourceSourceStatus field.
func (o *BranchDiffCell) SetSourceSourceStatus(v string) {
	o.SourceSourceStatus = &v
}

// GetDestinationSourceStatus returns the DestinationSourceStatus field value if set, zero value otherwise.
func (o *BranchDiffCell) GetDestinationSourceStatus() string {
	if o == nil || isNil(o.DestinationSourceStatus) {
		var ret string
		return ret
	}
	return *o.DestinationSourceStatus
}

// GetDestinationSourceStatusOk returns a tuple with the DestinationSourceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchDiffCell) GetDestinationSourceStatusOk() (*string, bool) {
	if o == nil || isNil(o.DestinationSourceStatus) {
    return nil, false
	}
	return o.DestinationSourceStatus, true
}

// HasDestinationSourceStatus returns a boolean if a field has been set.
func (o *BranchDiffCell) HasDestinationSourceStatus() bool {
	if o != nil && !isNil(o.DestinationSourceStatus) {
		return true
	}

	return false
}

// SetDestinationSourceStatus gets a reference to the given string and assigns it to the DestinationSourceStatus field.
func (o *BranchDiffCell) SetDestinationSourceStatus(v string) {
	o.DestinationSourceStatus = &v
}

// GetSourceLengthSetting returns the SourceLengthSetting field value if set, zero value otherwise.
func (o *BranchDiffCell) GetSourceLengthSetting() int32 {
	if o == nil || isNil(o.SourceLengthSetting) {
		var ret int32
		return ret
	}
	return *o.SourceLengthSetting
}

// GetSourceLengthSettingOk returns a tuple with the SourceLengthSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchDiffCell) GetSourceLengthSettingOk() (*int32, bool) {
	if o == nil || isNil(o.SourceLengthSetting) {
    return nil, false
	}
	return o.SourceLengthSetting, true
}

// HasSourceLengthSetting returns a boolean if a field has been set.
func (o *BranchDiffCell) HasSourceLengthSetting() bool {
	if o != nil && !isNil(o.SourceLengthSetting) {
		return true
	}

	return false
}

// SetSourceLengthSetting gets a reference to the given int32 and assigns it to the SourceLengthSetting field.
func (o *BranchDiffCell) SetSourceLengthSetting(v int32) {
	o.SourceLengthSetting = &v
}

// GetDestinationLengthSetting returns the DestinationLengthSetting field value if set, zero value otherwise.
func (o *BranchDiffCell) GetDestinationLengthSetting() int32 {
	if o == nil || isNil(o.DestinationLengthSetting) {
		var ret int32
		return ret
	}
	return *o.DestinationLengthSetting
}

// GetDestinationLengthSettingOk returns a tuple with the DestinationLengthSetting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchDiffCell) GetDestinationLengthSettingOk() (*int32, bool) {
	if o == nil || isNil(o.DestinationLengthSetting) {
    return nil, false
	}
	return o.DestinationLengthSetting, true
}

// HasDestinationLengthSetting returns a boolean if a field has been set.
func (o *BranchDiffCell) HasDestinationLengthSetting() bool {
	if o != nil && !isNil(o.DestinationLengthSetting) {
		return true
	}

	return false
}

// SetDestinationLengthSetting gets a reference to the given int32 and assigns it to the DestinationLengthSetting field.
func (o *BranchDiffCell) SetDestinationLengthSetting(v int32) {
	o.DestinationLengthSetting = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BranchDiffCell) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchDiffCell) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BranchDiffCell) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *BranchDiffCell) SetStatus(v string) {
	o.Status = &v
}

// GetColumnId returns the ColumnId field value if set, zero value otherwise.
func (o *BranchDiffCell) GetColumnId() string {
	if o == nil || isNil(o.ColumnId) {
		var ret string
		return ret
	}
	return *o.ColumnId
}

// GetColumnIdOk returns a tuple with the ColumnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchDiffCell) GetColumnIdOk() (*string, bool) {
	if o == nil || isNil(o.ColumnId) {
    return nil, false
	}
	return o.ColumnId, true
}

// HasColumnId returns a boolean if a field has been set.
func (o *BranchDiffCell) HasColumnId() bool {
	if o != nil && !isNil(o.ColumnId) {
		return true
	}

	return false
}

// SetColumnId gets a reference to the given string and assigns it to the ColumnId field.
func (o *BranchDiffCell) SetColumnId(v string) {
	o.ColumnId = &v
}

// GetSourceValue returns the SourceValue field value if set, zero value otherwise.
func (o *BranchDiffCell) GetSourceValue() map[string]interface{} {
	if o == nil || isNil(o.SourceValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.SourceValue
}

// GetSourceValueOk returns a tuple with the SourceValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchDiffCell) GetSourceValueOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.SourceValue) {
    return map[string]interface{}{}, false
	}
	return o.SourceValue, true
}

// HasSourceValue returns a boolean if a field has been set.
func (o *BranchDiffCell) HasSourceValue() bool {
	if o != nil && !isNil(o.SourceValue) {
		return true
	}

	return false
}

// SetSourceValue gets a reference to the given map[string]interface{} and assigns it to the SourceValue field.
func (o *BranchDiffCell) SetSourceValue(v map[string]interface{}) {
	o.SourceValue = v
}

// GetDestinationValue returns the DestinationValue field value if set, zero value otherwise.
func (o *BranchDiffCell) GetDestinationValue() map[string]interface{} {
	if o == nil || isNil(o.DestinationValue) {
		var ret map[string]interface{}
		return ret
	}
	return o.DestinationValue
}

// GetDestinationValueOk returns a tuple with the DestinationValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BranchDiffCell) GetDestinationValueOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.DestinationValue) {
    return map[string]interface{}{}, false
	}
	return o.DestinationValue, true
}

// HasDestinationValue returns a boolean if a field has been set.
func (o *BranchDiffCell) HasDestinationValue() bool {
	if o != nil && !isNil(o.DestinationValue) {
		return true
	}

	return false
}

// SetDestinationValue gets a reference to the given map[string]interface{} and assigns it to the DestinationValue field.
func (o *BranchDiffCell) SetDestinationValue(v map[string]interface{}) {
	o.DestinationValue = v
}

func (o BranchDiffCell) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SourceDependencyStatus) {
		toSerialize["sourceDependencyStatus"] = o.SourceDependencyStatus
	}
	if !isNil(o.DestinationDependencyStatus) {
		toSerialize["destinationDependencyStatus"] = o.DestinationDependencyStatus
	}
	if !isNil(o.SourceSourceStatus) {
		toSerialize["sourceSourceStatus"] = o.SourceSourceStatus
	}
	if !isNil(o.DestinationSourceStatus) {
		toSerialize["destinationSourceStatus"] = o.DestinationSourceStatus
	}
	if !isNil(o.SourceLengthSetting) {
		toSerialize["sourceLengthSetting"] = o.SourceLengthSetting
	}
	if !isNil(o.DestinationLengthSetting) {
		toSerialize["destinationLengthSetting"] = o.DestinationLengthSetting
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.ColumnId) {
		toSerialize["columnId"] = o.ColumnId
	}
	if !isNil(o.SourceValue) {
		toSerialize["sourceValue"] = o.SourceValue
	}
	if !isNil(o.DestinationValue) {
		toSerialize["destinationValue"] = o.DestinationValue
	}
	return json.Marshal(toSerialize)
}

type NullableBranchDiffCell struct {
	value *BranchDiffCell
	isSet bool
}

func (v NullableBranchDiffCell) Get() *BranchDiffCell {
	return v.value
}

func (v *NullableBranchDiffCell) Set(val *BranchDiffCell) {
	v.value = val
	v.isSet = true
}

func (v NullableBranchDiffCell) IsSet() bool {
	return v.isSet
}

func (v *NullableBranchDiffCell) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBranchDiffCell(val *BranchDiffCell) *NullableBranchDiffCell {
	return &NullableBranchDiffCell{value: val, isSet: true}
}

func (v NullableBranchDiffCell) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBranchDiffCell) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


