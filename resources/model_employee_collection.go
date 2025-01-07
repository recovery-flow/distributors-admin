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

// checks if the EmployeeCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmployeeCollection{}

// EmployeeCollection struct for EmployeeCollection
type EmployeeCollection struct {
	Data EmployeeCollectionData `json:"data"`
}

type _EmployeeCollection EmployeeCollection

// NewEmployeeCollection instantiates a new EmployeeCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployeeCollection(data EmployeeCollectionData) *EmployeeCollection {
	this := EmployeeCollection{}
	this.Data = data
	return &this
}

// NewEmployeeCollectionWithDefaults instantiates a new EmployeeCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployeeCollectionWithDefaults() *EmployeeCollection {
	this := EmployeeCollection{}
	return &this
}

// GetData returns the Data field value
func (o *EmployeeCollection) GetData() EmployeeCollectionData {
	if o == nil {
		var ret EmployeeCollectionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *EmployeeCollection) GetDataOk() (*EmployeeCollectionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *EmployeeCollection) SetData(v EmployeeCollectionData) {
	o.Data = v
}

func (o EmployeeCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmployeeCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *EmployeeCollection) UnmarshalJSON(data []byte) (err error) {
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

	varEmployeeCollection := _EmployeeCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEmployeeCollection)

	if err != nil {
		return err
	}

	*o = EmployeeCollection(varEmployeeCollection)

	return err
}

type NullableEmployeeCollection struct {
	value *EmployeeCollection
	isSet bool
}

func (v NullableEmployeeCollection) Get() *EmployeeCollection {
	return v.value
}

func (v *NullableEmployeeCollection) Set(val *EmployeeCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployeeCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployeeCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployeeCollection(val *EmployeeCollection) *NullableEmployeeCollection {
	return &NullableEmployeeCollection{value: val, isSet: true}
}

func (v NullableEmployeeCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployeeCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

