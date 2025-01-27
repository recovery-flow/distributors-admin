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

// checks if the DistributorEmployeeCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DistributorEmployeeCollection{}

// DistributorEmployeeCollection struct for DistributorEmployeeCollection
type DistributorEmployeeCollection struct {
	Data DistributorEmployeeCollectionData `json:"data"`
}

type _DistributorEmployeeCollection DistributorEmployeeCollection

// NewDistributorEmployeeCollection instantiates a new DistributorEmployeeCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistributorEmployeeCollection(data DistributorEmployeeCollectionData) *DistributorEmployeeCollection {
	this := DistributorEmployeeCollection{}
	this.Data = data
	return &this
}

// NewDistributorEmployeeCollectionWithDefaults instantiates a new DistributorEmployeeCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistributorEmployeeCollectionWithDefaults() *DistributorEmployeeCollection {
	this := DistributorEmployeeCollection{}
	return &this
}

// GetData returns the Data field value
func (o *DistributorEmployeeCollection) GetData() DistributorEmployeeCollectionData {
	if o == nil {
		var ret DistributorEmployeeCollectionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DistributorEmployeeCollection) GetDataOk() (*DistributorEmployeeCollectionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DistributorEmployeeCollection) SetData(v DistributorEmployeeCollectionData) {
	o.Data = v
}

func (o DistributorEmployeeCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DistributorEmployeeCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *DistributorEmployeeCollection) UnmarshalJSON(data []byte) (err error) {
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

	varDistributorEmployeeCollection := _DistributorEmployeeCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDistributorEmployeeCollection)

	if err != nil {
		return err
	}

	*o = DistributorEmployeeCollection(varDistributorEmployeeCollection)

	return err
}

type NullableDistributorEmployeeCollection struct {
	value *DistributorEmployeeCollection
	isSet bool
}

func (v NullableDistributorEmployeeCollection) Get() *DistributorEmployeeCollection {
	return v.value
}

func (v *NullableDistributorEmployeeCollection) Set(val *DistributorEmployeeCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributorEmployeeCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributorEmployeeCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributorEmployeeCollection(val *DistributorEmployeeCollection) *NullableDistributorEmployeeCollection {
	return &NullableDistributorEmployeeCollection{value: val, isSet: true}
}

func (v NullableDistributorEmployeeCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributorEmployeeCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


