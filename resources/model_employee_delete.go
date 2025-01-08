/*
Title

Title

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the EmployeeDelete type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmployeeDelete{}

// EmployeeDelete struct for EmployeeDelete
type EmployeeDelete struct {
	Data EmployeeDeleteData `json:"data"`
}

type _EmployeeDelete EmployeeDelete

// NewEmployeeDelete instantiates a new EmployeeDelete object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployeeDelete(data EmployeeDeleteData) *EmployeeDelete {
	this := EmployeeDelete{}
	this.Data = data
	return &this
}

// NewEmployeeDeleteWithDefaults instantiates a new EmployeeDelete object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployeeDeleteWithDefaults() *EmployeeDelete {
	this := EmployeeDelete{}
	return &this
}

// GetData returns the Data field value
func (o *EmployeeDelete) GetData() EmployeeDeleteData {
	if o == nil {
		var ret EmployeeDeleteData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *EmployeeDelete) GetDataOk() (*EmployeeDeleteData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *EmployeeDelete) SetData(v EmployeeDeleteData) {
	o.Data = v
}

func (o EmployeeDelete) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmployeeDelete) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *EmployeeDelete) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varEmployeeDelete := _EmployeeDelete{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmployeeDelete)

	if err != nil {
		return err
	}

	*o = EmployeeDelete(varEmployeeDelete)

	return err
}

type NullableEmployeeDelete struct {
	value *EmployeeDelete
	isSet bool
}

func (v NullableEmployeeDelete) Get() *EmployeeDelete {
	return v.value
}

func (v *NullableEmployeeDelete) Set(val *EmployeeDelete) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployeeDelete) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployeeDelete) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployeeDelete(val *EmployeeDelete) *NullableEmployeeDelete {
	return &NullableEmployeeDelete{value: val, isSet: true}
}

func (v NullableEmployeeDelete) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployeeDelete) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

