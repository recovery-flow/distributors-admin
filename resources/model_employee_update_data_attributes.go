/*
Title

Title

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resources

import (
	"encoding/json"
	"time"
)

// checks if the EmployeeUpdateDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmployeeUpdateDataAttributes{}

// EmployeeUpdateDataAttributes struct for EmployeeUpdateDataAttributes
type EmployeeUpdateDataAttributes struct {
	// User email
	Email *string `json:"email,omitempty"`
	// New user password
	Password *string `json:"password,omitempty"`
	// Distributor ID
	DistributorId *string `json:"distributor_id,omitempty"`
	// Creation date
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// NewEmployeeUpdateDataAttributes instantiates a new EmployeeUpdateDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmployeeUpdateDataAttributes() *EmployeeUpdateDataAttributes {
	this := EmployeeUpdateDataAttributes{}
	return &this
}

// NewEmployeeUpdateDataAttributesWithDefaults instantiates a new EmployeeUpdateDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmployeeUpdateDataAttributesWithDefaults() *EmployeeUpdateDataAttributes {
	this := EmployeeUpdateDataAttributes{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *EmployeeUpdateDataAttributes) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeUpdateDataAttributes) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *EmployeeUpdateDataAttributes) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *EmployeeUpdateDataAttributes) SetEmail(v string) {
	o.Email = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *EmployeeUpdateDataAttributes) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeUpdateDataAttributes) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// HasPassword returns a boolean if a field has been set.
func (o *EmployeeUpdateDataAttributes) HasPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *EmployeeUpdateDataAttributes) SetPassword(v string) {
	o.Password = &v
}

// GetDistributorId returns the DistributorId field value if set, zero value otherwise.
func (o *EmployeeUpdateDataAttributes) GetDistributorId() string {
	if o == nil || IsNil(o.DistributorId) {
		var ret string
		return ret
	}
	return *o.DistributorId
}

// GetDistributorIdOk returns a tuple with the DistributorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeUpdateDataAttributes) GetDistributorIdOk() (*string, bool) {
	if o == nil || IsNil(o.DistributorId) {
		return nil, false
	}
	return o.DistributorId, true
}

// HasDistributorId returns a boolean if a field has been set.
func (o *EmployeeUpdateDataAttributes) HasDistributorId() bool {
	if o != nil && !IsNil(o.DistributorId) {
		return true
	}

	return false
}

// SetDistributorId gets a reference to the given string and assigns it to the DistributorId field.
func (o *EmployeeUpdateDataAttributes) SetDistributorId(v string) {
	o.DistributorId = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *EmployeeUpdateDataAttributes) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmployeeUpdateDataAttributes) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *EmployeeUpdateDataAttributes) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *EmployeeUpdateDataAttributes) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

func (o EmployeeUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmployeeUpdateDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
	}
	if !IsNil(o.DistributorId) {
		toSerialize["distributor_id"] = o.DistributorId
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	return toSerialize, nil
}

type NullableEmployeeUpdateDataAttributes struct {
	value *EmployeeUpdateDataAttributes
	isSet bool
}

func (v NullableEmployeeUpdateDataAttributes) Get() *EmployeeUpdateDataAttributes {
	return v.value
}

func (v *NullableEmployeeUpdateDataAttributes) Set(val *EmployeeUpdateDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableEmployeeUpdateDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableEmployeeUpdateDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmployeeUpdateDataAttributes(val *EmployeeUpdateDataAttributes) *NullableEmployeeUpdateDataAttributes {
	return &NullableEmployeeUpdateDataAttributes{value: val, isSet: true}
}

func (v NullableEmployeeUpdateDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmployeeUpdateDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

