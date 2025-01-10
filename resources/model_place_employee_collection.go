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

// checks if the PlaceEmployeeCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PlaceEmployeeCollection{}

// PlaceEmployeeCollection struct for PlaceEmployeeCollection
type PlaceEmployeeCollection struct {
	Data PlaceEmployeeCollectionData `json:"data"`
}

type _PlaceEmployeeCollection PlaceEmployeeCollection

// NewPlaceEmployeeCollection instantiates a new PlaceEmployeeCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlaceEmployeeCollection(data PlaceEmployeeCollectionData) *PlaceEmployeeCollection {
	this := PlaceEmployeeCollection{}
	this.Data = data
	return &this
}

// NewPlaceEmployeeCollectionWithDefaults instantiates a new PlaceEmployeeCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlaceEmployeeCollectionWithDefaults() *PlaceEmployeeCollection {
	this := PlaceEmployeeCollection{}
	return &this
}

// GetData returns the Data field value
func (o *PlaceEmployeeCollection) GetData() PlaceEmployeeCollectionData {
	if o == nil {
		var ret PlaceEmployeeCollectionData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PlaceEmployeeCollection) GetDataOk() (*PlaceEmployeeCollectionData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PlaceEmployeeCollection) SetData(v PlaceEmployeeCollectionData) {
	o.Data = v
}

func (o PlaceEmployeeCollection) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PlaceEmployeeCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *PlaceEmployeeCollection) UnmarshalJSON(data []byte) (err error) {
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

	varPlaceEmployeeCollection := _PlaceEmployeeCollection{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPlaceEmployeeCollection)

	if err != nil {
		return err
	}

	*o = PlaceEmployeeCollection(varPlaceEmployeeCollection)

	return err
}

type NullablePlaceEmployeeCollection struct {
	value *PlaceEmployeeCollection
	isSet bool
}

func (v NullablePlaceEmployeeCollection) Get() *PlaceEmployeeCollection {
	return v.value
}

func (v *NullablePlaceEmployeeCollection) Set(val *PlaceEmployeeCollection) {
	v.value = val
	v.isSet = true
}

func (v NullablePlaceEmployeeCollection) IsSet() bool {
	return v.isSet
}

func (v *NullablePlaceEmployeeCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlaceEmployeeCollection(val *PlaceEmployeeCollection) *NullablePlaceEmployeeCollection {
	return &NullablePlaceEmployeeCollection{value: val, isSet: true}
}

func (v NullablePlaceEmployeeCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlaceEmployeeCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

