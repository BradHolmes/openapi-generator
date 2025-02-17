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

// ReadOnlyWithDefault struct for ReadOnlyWithDefault
type ReadOnlyWithDefault struct {
	Prop1 *string `json:"prop1,omitempty"`
	Prop2 *string `json:"prop2,omitempty"`
	Prop3 *string `json:"prop3,omitempty"`
	BoolProp1 *bool `json:"boolProp1,omitempty"`
	BoolProp2 *bool `json:"boolProp2,omitempty"`
	IntProp1 *float32 `json:"intProp1,omitempty"`
	IntProp2 *float32 `json:"intProp2,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReadOnlyWithDefault ReadOnlyWithDefault

// NewReadOnlyWithDefault instantiates a new ReadOnlyWithDefault object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReadOnlyWithDefault() *ReadOnlyWithDefault {
	this := ReadOnlyWithDefault{}
	var prop3 string = "defaultProp3"
	this.Prop3 = &prop3
	var boolProp2 bool = true
	this.BoolProp2 = &boolProp2
	var intProp2 float32 = 120
	this.IntProp2 = &intProp2
	return &this
}

// NewReadOnlyWithDefaultWithDefaults instantiates a new ReadOnlyWithDefault object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReadOnlyWithDefaultWithDefaults() *ReadOnlyWithDefault {
	this := ReadOnlyWithDefault{}
	var prop3 string = "defaultProp3"
	this.Prop3 = &prop3
	var boolProp2 bool = true
	this.BoolProp2 = &boolProp2
	var intProp2 float32 = 120
	this.IntProp2 = &intProp2
	return &this
}

// GetProp1 returns the Prop1 field value if set, zero value otherwise.
func (o *ReadOnlyWithDefault) GetProp1() string {
	if o == nil || o.Prop1 == nil {
		var ret string
		return ret
	}
	return *o.Prop1
}

// GetProp1Ok returns a tuple with the Prop1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadOnlyWithDefault) GetProp1Ok() (*string, bool) {
	if o == nil || o.Prop1 == nil {
		return nil, false
	}
	return o.Prop1, true
}

// HasProp1 returns a boolean if a field has been set.
func (o *ReadOnlyWithDefault) HasProp1() bool {
	if o != nil && o.Prop1 != nil {
		return true
	}

	return false
}

// SetProp1 gets a reference to the given string and assigns it to the Prop1 field.
func (o *ReadOnlyWithDefault) SetProp1(v string) {
	o.Prop1 = &v
}

// GetProp2 returns the Prop2 field value if set, zero value otherwise.
func (o *ReadOnlyWithDefault) GetProp2() string {
	if o == nil || o.Prop2 == nil {
		var ret string
		return ret
	}
	return *o.Prop2
}

// GetProp2Ok returns a tuple with the Prop2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadOnlyWithDefault) GetProp2Ok() (*string, bool) {
	if o == nil || o.Prop2 == nil {
		return nil, false
	}
	return o.Prop2, true
}

// HasProp2 returns a boolean if a field has been set.
func (o *ReadOnlyWithDefault) HasProp2() bool {
	if o != nil && o.Prop2 != nil {
		return true
	}

	return false
}

// SetProp2 gets a reference to the given string and assigns it to the Prop2 field.
func (o *ReadOnlyWithDefault) SetProp2(v string) {
	o.Prop2 = &v
}

// GetProp3 returns the Prop3 field value if set, zero value otherwise.
func (o *ReadOnlyWithDefault) GetProp3() string {
	if o == nil || o.Prop3 == nil {
		var ret string
		return ret
	}
	return *o.Prop3
}

// GetProp3Ok returns a tuple with the Prop3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadOnlyWithDefault) GetProp3Ok() (*string, bool) {
	if o == nil || o.Prop3 == nil {
		return nil, false
	}
	return o.Prop3, true
}

// HasProp3 returns a boolean if a field has been set.
func (o *ReadOnlyWithDefault) HasProp3() bool {
	if o != nil && o.Prop3 != nil {
		return true
	}

	return false
}

// SetProp3 gets a reference to the given string and assigns it to the Prop3 field.
func (o *ReadOnlyWithDefault) SetProp3(v string) {
	o.Prop3 = &v
}

// GetBoolProp1 returns the BoolProp1 field value if set, zero value otherwise.
func (o *ReadOnlyWithDefault) GetBoolProp1() bool {
	if o == nil || o.BoolProp1 == nil {
		var ret bool
		return ret
	}
	return *o.BoolProp1
}

// GetBoolProp1Ok returns a tuple with the BoolProp1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadOnlyWithDefault) GetBoolProp1Ok() (*bool, bool) {
	if o == nil || o.BoolProp1 == nil {
		return nil, false
	}
	return o.BoolProp1, true
}

// HasBoolProp1 returns a boolean if a field has been set.
func (o *ReadOnlyWithDefault) HasBoolProp1() bool {
	if o != nil && o.BoolProp1 != nil {
		return true
	}

	return false
}

// SetBoolProp1 gets a reference to the given bool and assigns it to the BoolProp1 field.
func (o *ReadOnlyWithDefault) SetBoolProp1(v bool) {
	o.BoolProp1 = &v
}

// GetBoolProp2 returns the BoolProp2 field value if set, zero value otherwise.
func (o *ReadOnlyWithDefault) GetBoolProp2() bool {
	if o == nil || o.BoolProp2 == nil {
		var ret bool
		return ret
	}
	return *o.BoolProp2
}

// GetBoolProp2Ok returns a tuple with the BoolProp2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadOnlyWithDefault) GetBoolProp2Ok() (*bool, bool) {
	if o == nil || o.BoolProp2 == nil {
		return nil, false
	}
	return o.BoolProp2, true
}

// HasBoolProp2 returns a boolean if a field has been set.
func (o *ReadOnlyWithDefault) HasBoolProp2() bool {
	if o != nil && o.BoolProp2 != nil {
		return true
	}

	return false
}

// SetBoolProp2 gets a reference to the given bool and assigns it to the BoolProp2 field.
func (o *ReadOnlyWithDefault) SetBoolProp2(v bool) {
	o.BoolProp2 = &v
}

// GetIntProp1 returns the IntProp1 field value if set, zero value otherwise.
func (o *ReadOnlyWithDefault) GetIntProp1() float32 {
	if o == nil || o.IntProp1 == nil {
		var ret float32
		return ret
	}
	return *o.IntProp1
}

// GetIntProp1Ok returns a tuple with the IntProp1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadOnlyWithDefault) GetIntProp1Ok() (*float32, bool) {
	if o == nil || o.IntProp1 == nil {
		return nil, false
	}
	return o.IntProp1, true
}

// HasIntProp1 returns a boolean if a field has been set.
func (o *ReadOnlyWithDefault) HasIntProp1() bool {
	if o != nil && o.IntProp1 != nil {
		return true
	}

	return false
}

// SetIntProp1 gets a reference to the given float32 and assigns it to the IntProp1 field.
func (o *ReadOnlyWithDefault) SetIntProp1(v float32) {
	o.IntProp1 = &v
}

// GetIntProp2 returns the IntProp2 field value if set, zero value otherwise.
func (o *ReadOnlyWithDefault) GetIntProp2() float32 {
	if o == nil || o.IntProp2 == nil {
		var ret float32
		return ret
	}
	return *o.IntProp2
}

// GetIntProp2Ok returns a tuple with the IntProp2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReadOnlyWithDefault) GetIntProp2Ok() (*float32, bool) {
	if o == nil || o.IntProp2 == nil {
		return nil, false
	}
	return o.IntProp2, true
}

// HasIntProp2 returns a boolean if a field has been set.
func (o *ReadOnlyWithDefault) HasIntProp2() bool {
	if o != nil && o.IntProp2 != nil {
		return true
	}

	return false
}

// SetIntProp2 gets a reference to the given float32 and assigns it to the IntProp2 field.
func (o *ReadOnlyWithDefault) SetIntProp2(v float32) {
	o.IntProp2 = &v
}

func (o ReadOnlyWithDefault) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Prop1 != nil  {
		toSerialize["prop1"] = o.Prop1
	}
	if o.Prop2 != nil  {
		toSerialize["prop2"] = o.Prop2
	}
	if o.Prop3 != nil  {
		toSerialize["prop3"] = o.Prop3
	}
	if o.BoolProp1 != nil  {
		toSerialize["boolProp1"] = o.BoolProp1
	}
	if o.BoolProp2 != nil  {
		toSerialize["boolProp2"] = o.BoolProp2
	}
	if o.IntProp1 != nil  {
		toSerialize["intProp1"] = o.IntProp1
	}
	if o.IntProp2 != nil  {
		toSerialize["intProp2"] = o.IntProp2
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ReadOnlyWithDefault) UnmarshalJSON(bytes []byte) (err error) {
	varReadOnlyWithDefault := _ReadOnlyWithDefault{}

	if err = json.Unmarshal(bytes, &varReadOnlyWithDefault); err == nil {
		*o = ReadOnlyWithDefault(varReadOnlyWithDefault)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "prop1")
		delete(additionalProperties, "prop2")
		delete(additionalProperties, "prop3")
		delete(additionalProperties, "boolProp1")
		delete(additionalProperties, "boolProp2")
		delete(additionalProperties, "intProp1")
		delete(additionalProperties, "intProp2")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReadOnlyWithDefault struct {
	value *ReadOnlyWithDefault
	isSet bool
}

func (v NullableReadOnlyWithDefault) Get() *ReadOnlyWithDefault {
	return v.value
}

func (v *NullableReadOnlyWithDefault) Set(val *ReadOnlyWithDefault) {
	v.value = val
	v.isSet = true
}

func (v NullableReadOnlyWithDefault) IsSet() bool {
	return v.isSet
}

func (v *NullableReadOnlyWithDefault) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReadOnlyWithDefault(val *ReadOnlyWithDefault) *NullableReadOnlyWithDefault {
	return &NullableReadOnlyWithDefault{value: val, isSet: true}
}

func (v NullableReadOnlyWithDefault) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReadOnlyWithDefault) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


