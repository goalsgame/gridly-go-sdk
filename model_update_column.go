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

// UpdateColumn struct for UpdateColumn
type UpdateColumn struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Type *string `json:"type,omitempty"`
	LanguageCode *string `json:"languageCode,omitempty"`
	LocalizationType *string `json:"localizationType,omitempty"`
	SelectionOptions []string `json:"selectionOptions,omitempty"`
	NumberFormat *NumberFormat `json:"numberFormat,omitempty"`
	Reference *Reference `json:"reference,omitempty"`
	Formula *Formula `json:"formula,omitempty"`
	DateTimeFormat *DateTimeFormat `json:"dateTimeFormat,omitempty"`
	Viewable *bool `json:"viewable,omitempty"`
	Editable *bool `json:"editable,omitempty"`
	NewId *string `json:"newId,omitempty"`
}

// NewUpdateColumn instantiates a new UpdateColumn object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateColumn() *UpdateColumn {
	this := UpdateColumn{}
	return &this
}

// NewUpdateColumnWithDefaults instantiates a new UpdateColumn object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateColumnWithDefaults() *UpdateColumn {
	this := UpdateColumn{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateColumn) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateColumn) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *UpdateColumn) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateColumn) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateColumn) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateColumn) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
    return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *UpdateColumn) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateColumn) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UpdateColumn) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateColumn) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
    return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UpdateColumn) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UpdateColumn) SetType(v string) {
	o.Type = &v
}

// GetLanguageCode returns the LanguageCode field value if set, zero value otherwise.
func (o *UpdateColumn) GetLanguageCode() string {
	if o == nil || isNil(o.LanguageCode) {
		var ret string
		return ret
	}
	return *o.LanguageCode
}

// GetLanguageCodeOk returns a tuple with the LanguageCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateColumn) GetLanguageCodeOk() (*string, bool) {
	if o == nil || isNil(o.LanguageCode) {
    return nil, false
	}
	return o.LanguageCode, true
}

// HasLanguageCode returns a boolean if a field has been set.
func (o *UpdateColumn) HasLanguageCode() bool {
	if o != nil && !isNil(o.LanguageCode) {
		return true
	}

	return false
}

// SetLanguageCode gets a reference to the given string and assigns it to the LanguageCode field.
func (o *UpdateColumn) SetLanguageCode(v string) {
	o.LanguageCode = &v
}

// GetLocalizationType returns the LocalizationType field value if set, zero value otherwise.
func (o *UpdateColumn) GetLocalizationType() string {
	if o == nil || isNil(o.LocalizationType) {
		var ret string
		return ret
	}
	return *o.LocalizationType
}

// GetLocalizationTypeOk returns a tuple with the LocalizationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateColumn) GetLocalizationTypeOk() (*string, bool) {
	if o == nil || isNil(o.LocalizationType) {
    return nil, false
	}
	return o.LocalizationType, true
}

// HasLocalizationType returns a boolean if a field has been set.
func (o *UpdateColumn) HasLocalizationType() bool {
	if o != nil && !isNil(o.LocalizationType) {
		return true
	}

	return false
}

// SetLocalizationType gets a reference to the given string and assigns it to the LocalizationType field.
func (o *UpdateColumn) SetLocalizationType(v string) {
	o.LocalizationType = &v
}

// GetSelectionOptions returns the SelectionOptions field value if set, zero value otherwise.
func (o *UpdateColumn) GetSelectionOptions() []string {
	if o == nil || isNil(o.SelectionOptions) {
		var ret []string
		return ret
	}
	return o.SelectionOptions
}

// GetSelectionOptionsOk returns a tuple with the SelectionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateColumn) GetSelectionOptionsOk() ([]string, bool) {
	if o == nil || isNil(o.SelectionOptions) {
    return nil, false
	}
	return o.SelectionOptions, true
}

// HasSelectionOptions returns a boolean if a field has been set.
func (o *UpdateColumn) HasSelectionOptions() bool {
	if o != nil && !isNil(o.SelectionOptions) {
		return true
	}

	return false
}

// SetSelectionOptions gets a reference to the given []string and assigns it to the SelectionOptions field.
func (o *UpdateColumn) SetSelectionOptions(v []string) {
	o.SelectionOptions = v
}

// GetNumberFormat returns the NumberFormat field value if set, zero value otherwise.
func (o *UpdateColumn) GetNumberFormat() NumberFormat {
	if o == nil || isNil(o.NumberFormat) {
		var ret NumberFormat
		return ret
	}
	return *o.NumberFormat
}

// GetNumberFormatOk returns a tuple with the NumberFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateColumn) GetNumberFormatOk() (*NumberFormat, bool) {
	if o == nil || isNil(o.NumberFormat) {
    return nil, false
	}
	return o.NumberFormat, true
}

// HasNumberFormat returns a boolean if a field has been set.
func (o *UpdateColumn) HasNumberFormat() bool {
	if o != nil && !isNil(o.NumberFormat) {
		return true
	}

	return false
}

// SetNumberFormat gets a reference to the given NumberFormat and assigns it to the NumberFormat field.
func (o *UpdateColumn) SetNumberFormat(v NumberFormat) {
	o.NumberFormat = &v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *UpdateColumn) GetReference() Reference {
	if o == nil || isNil(o.Reference) {
		var ret Reference
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateColumn) GetReferenceOk() (*Reference, bool) {
	if o == nil || isNil(o.Reference) {
    return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *UpdateColumn) HasReference() bool {
	if o != nil && !isNil(o.Reference) {
		return true
	}

	return false
}

// SetReference gets a reference to the given Reference and assigns it to the Reference field.
func (o *UpdateColumn) SetReference(v Reference) {
	o.Reference = &v
}

// GetFormula returns the Formula field value if set, zero value otherwise.
func (o *UpdateColumn) GetFormula() Formula {
	if o == nil || isNil(o.Formula) {
		var ret Formula
		return ret
	}
	return *o.Formula
}

// GetFormulaOk returns a tuple with the Formula field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateColumn) GetFormulaOk() (*Formula, bool) {
	if o == nil || isNil(o.Formula) {
    return nil, false
	}
	return o.Formula, true
}

// HasFormula returns a boolean if a field has been set.
func (o *UpdateColumn) HasFormula() bool {
	if o != nil && !isNil(o.Formula) {
		return true
	}

	return false
}

// SetFormula gets a reference to the given Formula and assigns it to the Formula field.
func (o *UpdateColumn) SetFormula(v Formula) {
	o.Formula = &v
}

// GetDateTimeFormat returns the DateTimeFormat field value if set, zero value otherwise.
func (o *UpdateColumn) GetDateTimeFormat() DateTimeFormat {
	if o == nil || isNil(o.DateTimeFormat) {
		var ret DateTimeFormat
		return ret
	}
	return *o.DateTimeFormat
}

// GetDateTimeFormatOk returns a tuple with the DateTimeFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateColumn) GetDateTimeFormatOk() (*DateTimeFormat, bool) {
	if o == nil || isNil(o.DateTimeFormat) {
    return nil, false
	}
	return o.DateTimeFormat, true
}

// HasDateTimeFormat returns a boolean if a field has been set.
func (o *UpdateColumn) HasDateTimeFormat() bool {
	if o != nil && !isNil(o.DateTimeFormat) {
		return true
	}

	return false
}

// SetDateTimeFormat gets a reference to the given DateTimeFormat and assigns it to the DateTimeFormat field.
func (o *UpdateColumn) SetDateTimeFormat(v DateTimeFormat) {
	o.DateTimeFormat = &v
}

// GetViewable returns the Viewable field value if set, zero value otherwise.
func (o *UpdateColumn) GetViewable() bool {
	if o == nil || isNil(o.Viewable) {
		var ret bool
		return ret
	}
	return *o.Viewable
}

// GetViewableOk returns a tuple with the Viewable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateColumn) GetViewableOk() (*bool, bool) {
	if o == nil || isNil(o.Viewable) {
    return nil, false
	}
	return o.Viewable, true
}

// HasViewable returns a boolean if a field has been set.
func (o *UpdateColumn) HasViewable() bool {
	if o != nil && !isNil(o.Viewable) {
		return true
	}

	return false
}

// SetViewable gets a reference to the given bool and assigns it to the Viewable field.
func (o *UpdateColumn) SetViewable(v bool) {
	o.Viewable = &v
}

// GetEditable returns the Editable field value if set, zero value otherwise.
func (o *UpdateColumn) GetEditable() bool {
	if o == nil || isNil(o.Editable) {
		var ret bool
		return ret
	}
	return *o.Editable
}

// GetEditableOk returns a tuple with the Editable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateColumn) GetEditableOk() (*bool, bool) {
	if o == nil || isNil(o.Editable) {
    return nil, false
	}
	return o.Editable, true
}

// HasEditable returns a boolean if a field has been set.
func (o *UpdateColumn) HasEditable() bool {
	if o != nil && !isNil(o.Editable) {
		return true
	}

	return false
}

// SetEditable gets a reference to the given bool and assigns it to the Editable field.
func (o *UpdateColumn) SetEditable(v bool) {
	o.Editable = &v
}

// GetNewId returns the NewId field value if set, zero value otherwise.
func (o *UpdateColumn) GetNewId() string {
	if o == nil || isNil(o.NewId) {
		var ret string
		return ret
	}
	return *o.NewId
}

// GetNewIdOk returns a tuple with the NewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateColumn) GetNewIdOk() (*string, bool) {
	if o == nil || isNil(o.NewId) {
    return nil, false
	}
	return o.NewId, true
}

// HasNewId returns a boolean if a field has been set.
func (o *UpdateColumn) HasNewId() bool {
	if o != nil && !isNil(o.NewId) {
		return true
	}

	return false
}

// SetNewId gets a reference to the given string and assigns it to the NewId field.
func (o *UpdateColumn) SetNewId(v string) {
	o.NewId = &v
}

func (o UpdateColumn) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.LanguageCode) {
		toSerialize["languageCode"] = o.LanguageCode
	}
	if !isNil(o.LocalizationType) {
		toSerialize["localizationType"] = o.LocalizationType
	}
	if !isNil(o.SelectionOptions) {
		toSerialize["selectionOptions"] = o.SelectionOptions
	}
	if !isNil(o.NumberFormat) {
		toSerialize["numberFormat"] = o.NumberFormat
	}
	if !isNil(o.Reference) {
		toSerialize["reference"] = o.Reference
	}
	if !isNil(o.Formula) {
		toSerialize["formula"] = o.Formula
	}
	if !isNil(o.DateTimeFormat) {
		toSerialize["dateTimeFormat"] = o.DateTimeFormat
	}
	if !isNil(o.Viewable) {
		toSerialize["viewable"] = o.Viewable
	}
	if !isNil(o.Editable) {
		toSerialize["editable"] = o.Editable
	}
	if !isNil(o.NewId) {
		toSerialize["newId"] = o.NewId
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateColumn struct {
	value *UpdateColumn
	isSet bool
}

func (v NullableUpdateColumn) Get() *UpdateColumn {
	return v.value
}

func (v *NullableUpdateColumn) Set(val *UpdateColumn) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateColumn(val *UpdateColumn) *NullableUpdateColumn {
	return &NullableUpdateColumn{value: val, isSet: true}
}

func (v NullableUpdateColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


