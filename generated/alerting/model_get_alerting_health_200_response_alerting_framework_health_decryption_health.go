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

// checks if the GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth{}

// GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth The timestamp and status of the rule decryption.
type GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth struct {
	Status    *string    `json:"status,omitempty"`
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// NewGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth instantiates a new GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth() *GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth {
	this := GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth{}
	return &this
}

// NewGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealthWithDefaults instantiates a new GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealthWithDefaults() *GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth {
	this := GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) SetStatus(v string) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

func (o GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return toSerialize, nil
}

type NullableGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth struct {
	value *GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth
	isSet bool
}

func (v NullableGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) Get() *GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth {
	return v.value
}

func (v *NullableGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) Set(val *GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) {
	v.value = val
	v.isSet = true
}

func (v NullableGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) IsSet() bool {
	return v.isSet
}

func (v *NullableGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth(val *GetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) *NullableGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth {
	return &NullableGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth{value: val, isSet: true}
}

func (v NullableGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetAlertingHealth200ResponseAlertingFrameworkHealthDecryptionHealth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
