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

// checks if the PlaceEmployeeCollectionDataAttributesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaceEmployeeCollectionDataAttributesInner{}

// PlaceEmployeeCollectionDataAttributesInner struct for PlaceEmployeeCollectionDataAttributesInner
type PlaceEmployeeCollectionDataAttributesInner struct {
	Data PlaceEmployeeData `json:"data"`
}

type _PlaceEmployeeCollectionDataAttributesInner PlaceEmployeeCollectionDataAttributesInner

// NewPlaceEmployeeCollectionDataAttributesInner instantiates a new PlaceEmployeeCollectionDataAttributesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaceEmployeeCollectionDataAttributesInner(data PlaceEmployeeData) *PlaceEmployeeCollectionDataAttributesInner {
	this := PlaceEmployeeCollectionDataAttributesInner{}
	this.Data = data
	return &this
}

// NewPlaceEmployeeCollectionDataAttributesInnerWithDefaults instantiates a new PlaceEmployeeCollectionDataAttributesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaceEmployeeCollectionDataAttributesInnerWithDefaults() *PlaceEmployeeCollectionDataAttributesInner {
	this := PlaceEmployeeCollectionDataAttributesInner{}
	return &this
}

// GetData returns the Data field value
func (o *PlaceEmployeeCollectionDataAttributesInner) GetData() PlaceEmployeeData {
	if o == nil {
		var ret PlaceEmployeeData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PlaceEmployeeCollectionDataAttributesInner) GetDataOk() (*PlaceEmployeeData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PlaceEmployeeCollectionDataAttributesInner) SetData(v PlaceEmployeeData) {
	o.Data = v
}

func (o PlaceEmployeeCollectionDataAttributesInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaceEmployeeCollectionDataAttributesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *PlaceEmployeeCollectionDataAttributesInner) UnmarshalJSON(data []byte) (err error) {
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

	varPlaceEmployeeCollectionDataAttributesInner := _PlaceEmployeeCollectionDataAttributesInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPlaceEmployeeCollectionDataAttributesInner)

	if err != nil {
		return err
	}

	*o = PlaceEmployeeCollectionDataAttributesInner(varPlaceEmployeeCollectionDataAttributesInner)

	return err
}

type NullablePlaceEmployeeCollectionDataAttributesInner struct {
	value *PlaceEmployeeCollectionDataAttributesInner
	isSet bool
}

func (v NullablePlaceEmployeeCollectionDataAttributesInner) Get() *PlaceEmployeeCollectionDataAttributesInner {
	return v.value
}

func (v *NullablePlaceEmployeeCollectionDataAttributesInner) Set(val *PlaceEmployeeCollectionDataAttributesInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceEmployeeCollectionDataAttributesInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceEmployeeCollectionDataAttributesInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceEmployeeCollectionDataAttributesInner(val *PlaceEmployeeCollectionDataAttributesInner) *NullablePlaceEmployeeCollectionDataAttributesInner {
	return &NullablePlaceEmployeeCollectionDataAttributesInner{value: val, isSet: true}
}

func (v NullablePlaceEmployeeCollectionDataAttributesInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceEmployeeCollectionDataAttributesInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


