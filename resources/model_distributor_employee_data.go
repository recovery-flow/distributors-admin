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

// checks if the DistributorEmployeeData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DistributorEmployeeData{}

// DistributorEmployeeData struct for DistributorEmployeeData
type DistributorEmployeeData struct {
	// employee ID
	Id string `json:"id"`
	Type string `json:"type"`
	Attributes DistributorEmployeeDataAttributes `json:"attributes"`
}

type _DistributorEmployeeData DistributorEmployeeData

// NewDistributorEmployeeData instantiates a new DistributorEmployeeData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDistributorEmployeeData(id string, type_ string, attributes DistributorEmployeeDataAttributes) *DistributorEmployeeData {
	this := DistributorEmployeeData{}
	this.Id = id
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewDistributorEmployeeDataWithDefaults instantiates a new DistributorEmployeeData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDistributorEmployeeDataWithDefaults() *DistributorEmployeeData {
	this := DistributorEmployeeData{}
	return &this
}

// GetId returns the Id field value
func (o *DistributorEmployeeData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *DistributorEmployeeData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *DistributorEmployeeData) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *DistributorEmployeeData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DistributorEmployeeData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DistributorEmployeeData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *DistributorEmployeeData) GetAttributes() DistributorEmployeeDataAttributes {
	if o == nil {
		var ret DistributorEmployeeDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *DistributorEmployeeData) GetAttributesOk() (*DistributorEmployeeDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *DistributorEmployeeData) SetAttributes(v DistributorEmployeeDataAttributes) {
	o.Attributes = v
}

func (o DistributorEmployeeData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DistributorEmployeeData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

func (o *DistributorEmployeeData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
		"attributes",
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

	varDistributorEmployeeData := _DistributorEmployeeData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDistributorEmployeeData)

	if err != nil {
		return err
	}

	*o = DistributorEmployeeData(varDistributorEmployeeData)

	return err
}

type NullableDistributorEmployeeData struct {
	value *DistributorEmployeeData
	isSet bool
}

func (v NullableDistributorEmployeeData) Get() *DistributorEmployeeData {
	return v.value
}

func (v *NullableDistributorEmployeeData) Set(val *DistributorEmployeeData) {
	v.value = val
	v.isSet = true
}

func (v NullableDistributorEmployeeData) IsSet() bool {
	return v.isSet
}

func (v *NullableDistributorEmployeeData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistributorEmployeeData(val *DistributorEmployeeData) *NullableDistributorEmployeeData {
	return &NullableDistributorEmployeeData{value: val, isSet: true}
}

func (v NullableDistributorEmployeeData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistributorEmployeeData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

