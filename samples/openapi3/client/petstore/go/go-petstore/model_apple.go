/*
OpenAPI Petstore

This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"encoding/json"
)

// Apple struct for Apple
type Apple struct {
	Cultivar *string `json:"cultivar,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Apple Apple

// NewApple instantiates a new Apple object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApple() *Apple {
	this := Apple{}
	return &this
}

// NewAppleWithDefaults instantiates a new Apple object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppleWithDefaults() *Apple {
	this := Apple{}
	return &this
}

// GetCultivar returns the Cultivar field value if set, zero value otherwise.
func (o *Apple) GetCultivar() string {
	if o == nil || o.Cultivar == nil {
		var ret string
		return ret
	}
	return *o.Cultivar
}

// GetCultivarOk returns a tuple with the Cultivar field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Apple) GetCultivarOk() (*string, bool) {
	if o == nil || o.Cultivar == nil {
		return nil, false
	}
	return o.Cultivar, true
}

// HasCultivar returns a boolean if a field has been set.
func (o *Apple) HasCultivar() bool {
	if o != nil && o.Cultivar != nil {
		return true
	}

	return false
}

// SetCultivar gets a reference to the given string and assigns it to the Cultivar field.
func (o *Apple) SetCultivar(v string) {
	o.Cultivar = &v
}

func (o Apple) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cultivar != nil  {
		toSerialize["cultivar"] = o.Cultivar
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Apple) UnmarshalJSON(bytes []byte) (err error) {
	varApple := _Apple{}

	if err = json.Unmarshal(bytes, &varApple); err == nil {
		*o = Apple(varApple)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "cultivar")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableApple struct {
	value *Apple
	isSet bool
}

func (v NullableApple) Get() *Apple {
	return v.value
}

func (v *NullableApple) Set(val *Apple) {
	v.value = val
	v.isSet = true
}

func (v NullableApple) IsSet() bool {
	return v.isSet
}

func (v *NullableApple) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApple(val *Apple) *NullableApple {
	return &NullableApple{value: val, isSet: true}
}

func (v NullableApple) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApple) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


