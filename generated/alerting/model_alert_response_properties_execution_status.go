/*
Alerting

OpenAPI schema for alerting endpoints

API version: 0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package alerting

import (
	"encoding/json"
	"time"
)

// checks if the AlertResponsePropertiesExecutionStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AlertResponsePropertiesExecutionStatus{}

// AlertResponsePropertiesExecutionStatus struct for AlertResponsePropertiesExecutionStatus
type AlertResponsePropertiesExecutionStatus struct {
	LastExecutionDate *time.Time `json:"lastExecutionDate,omitempty"`
	Status            *string    `json:"status,omitempty"`
}

// NewAlertResponsePropertiesExecutionStatus instantiates a new AlertResponsePropertiesExecutionStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAlertResponsePropertiesExecutionStatus() *AlertResponsePropertiesExecutionStatus {
	this := AlertResponsePropertiesExecutionStatus{}
	return &this
}

// NewAlertResponsePropertiesExecutionStatusWithDefaults instantiates a new AlertResponsePropertiesExecutionStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlertResponsePropertiesExecutionStatusWithDefaults() *AlertResponsePropertiesExecutionStatus {
	this := AlertResponsePropertiesExecutionStatus{}
	return &this
}

// GetLastExecutionDate returns the LastExecutionDate field value if set, zero value otherwise.
func (o *AlertResponsePropertiesExecutionStatus) GetLastExecutionDate() time.Time {
	if o == nil || IsNil(o.LastExecutionDate) {
		var ret time.Time
		return ret
	}
	return *o.LastExecutionDate
}

// GetLastExecutionDateOk returns a tuple with the LastExecutionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertResponsePropertiesExecutionStatus) GetLastExecutionDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastExecutionDate) {
		return nil, false
	}
	return o.LastExecutionDate, true
}

// HasLastExecutionDate returns a boolean if a field has been set.
func (o *AlertResponsePropertiesExecutionStatus) HasLastExecutionDate() bool {
	if o != nil && !IsNil(o.LastExecutionDate) {
		return true
	}

	return false
}

// SetLastExecutionDate gets a reference to the given time.Time and assigns it to the LastExecutionDate field.
func (o *AlertResponsePropertiesExecutionStatus) SetLastExecutionDate(v time.Time) {
	o.LastExecutionDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AlertResponsePropertiesExecutionStatus) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AlertResponsePropertiesExecutionStatus) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AlertResponsePropertiesExecutionStatus) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *AlertResponsePropertiesExecutionStatus) SetStatus(v string) {
	o.Status = &v
}

func (o AlertResponsePropertiesExecutionStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AlertResponsePropertiesExecutionStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LastExecutionDate) {
		toSerialize["lastExecutionDate"] = o.LastExecutionDate
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableAlertResponsePropertiesExecutionStatus struct {
	value *AlertResponsePropertiesExecutionStatus
	isSet bool
}

func (v NullableAlertResponsePropertiesExecutionStatus) Get() *AlertResponsePropertiesExecutionStatus {
	return v.value
}

func (v *NullableAlertResponsePropertiesExecutionStatus) Set(val *AlertResponsePropertiesExecutionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAlertResponsePropertiesExecutionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAlertResponsePropertiesExecutionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlertResponsePropertiesExecutionStatus(val *AlertResponsePropertiesExecutionStatus) *NullableAlertResponsePropertiesExecutionStatus {
	return &NullableAlertResponsePropertiesExecutionStatus{value: val, isSet: true}
}

func (v NullableAlertResponsePropertiesExecutionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlertResponsePropertiesExecutionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
