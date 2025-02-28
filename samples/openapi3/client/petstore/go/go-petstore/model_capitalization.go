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

// Capitalization struct for Capitalization
type Capitalization struct {
	SmallCamel *string `json:"smallCamel,omitempty"`
	CapitalCamel *string `json:"CapitalCamel,omitempty"`
	SmallSnake *string `json:"small_Snake,omitempty"`
	CapitalSnake *string `json:"Capital_Snake,omitempty"`
	SCAETHFlowPoints *string `json:"SCA_ETH_Flow_Points,omitempty"`
	// Name of the pet 
	ATT_NAME *string `json:"ATT_NAME,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _Capitalization Capitalization

// NewCapitalization instantiates a new Capitalization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCapitalization() *Capitalization {
	this := Capitalization{}
	return &this
}

// NewCapitalizationWithDefaults instantiates a new Capitalization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCapitalizationWithDefaults() *Capitalization {
	this := Capitalization{}
	return &this
}

// GetSmallCamel returns the SmallCamel field value if set, zero value otherwise.
func (o *Capitalization) GetSmallCamel() string {
	if o == nil || o.SmallCamel == nil {
		var ret string
		return ret
	}
	return *o.SmallCamel
}

// GetSmallCamelOk returns a tuple with the SmallCamel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capitalization) GetSmallCamelOk() (*string, bool) {
	if o == nil || o.SmallCamel == nil {
		return nil, false
	}
	return o.SmallCamel, true
}

// HasSmallCamel returns a boolean if a field has been set.
func (o *Capitalization) HasSmallCamel() bool {
	if o != nil && o.SmallCamel != nil {
		return true
	}

	return false
}

// SetSmallCamel gets a reference to the given string and assigns it to the SmallCamel field.
func (o *Capitalization) SetSmallCamel(v string) {
	o.SmallCamel = &v
}

// GetCapitalCamel returns the CapitalCamel field value if set, zero value otherwise.
func (o *Capitalization) GetCapitalCamel() string {
	if o == nil || o.CapitalCamel == nil {
		var ret string
		return ret
	}
	return *o.CapitalCamel
}

// GetCapitalCamelOk returns a tuple with the CapitalCamel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capitalization) GetCapitalCamelOk() (*string, bool) {
	if o == nil || o.CapitalCamel == nil {
		return nil, false
	}
	return o.CapitalCamel, true
}

// HasCapitalCamel returns a boolean if a field has been set.
func (o *Capitalization) HasCapitalCamel() bool {
	if o != nil && o.CapitalCamel != nil {
		return true
	}

	return false
}

// SetCapitalCamel gets a reference to the given string and assigns it to the CapitalCamel field.
func (o *Capitalization) SetCapitalCamel(v string) {
	o.CapitalCamel = &v
}

// GetSmallSnake returns the SmallSnake field value if set, zero value otherwise.
func (o *Capitalization) GetSmallSnake() string {
	if o == nil || o.SmallSnake == nil {
		var ret string
		return ret
	}
	return *o.SmallSnake
}

// GetSmallSnakeOk returns a tuple with the SmallSnake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capitalization) GetSmallSnakeOk() (*string, bool) {
	if o == nil || o.SmallSnake == nil {
		return nil, false
	}
	return o.SmallSnake, true
}

// HasSmallSnake returns a boolean if a field has been set.
func (o *Capitalization) HasSmallSnake() bool {
	if o != nil && o.SmallSnake != nil {
		return true
	}

	return false
}

// SetSmallSnake gets a reference to the given string and assigns it to the SmallSnake field.
func (o *Capitalization) SetSmallSnake(v string) {
	o.SmallSnake = &v
}

// GetCapitalSnake returns the CapitalSnake field value if set, zero value otherwise.
func (o *Capitalization) GetCapitalSnake() string {
	if o == nil || o.CapitalSnake == nil {
		var ret string
		return ret
	}
	return *o.CapitalSnake
}

// GetCapitalSnakeOk returns a tuple with the CapitalSnake field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capitalization) GetCapitalSnakeOk() (*string, bool) {
	if o == nil || o.CapitalSnake == nil {
		return nil, false
	}
	return o.CapitalSnake, true
}

// HasCapitalSnake returns a boolean if a field has been set.
func (o *Capitalization) HasCapitalSnake() bool {
	if o != nil && o.CapitalSnake != nil {
		return true
	}

	return false
}

// SetCapitalSnake gets a reference to the given string and assigns it to the CapitalSnake field.
func (o *Capitalization) SetCapitalSnake(v string) {
	o.CapitalSnake = &v
}

// GetSCAETHFlowPoints returns the SCAETHFlowPoints field value if set, zero value otherwise.
func (o *Capitalization) GetSCAETHFlowPoints() string {
	if o == nil || o.SCAETHFlowPoints == nil {
		var ret string
		return ret
	}
	return *o.SCAETHFlowPoints
}

// GetSCAETHFlowPointsOk returns a tuple with the SCAETHFlowPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capitalization) GetSCAETHFlowPointsOk() (*string, bool) {
	if o == nil || o.SCAETHFlowPoints == nil {
		return nil, false
	}
	return o.SCAETHFlowPoints, true
}

// HasSCAETHFlowPoints returns a boolean if a field has been set.
func (o *Capitalization) HasSCAETHFlowPoints() bool {
	if o != nil && o.SCAETHFlowPoints != nil {
		return true
	}

	return false
}

// SetSCAETHFlowPoints gets a reference to the given string and assigns it to the SCAETHFlowPoints field.
func (o *Capitalization) SetSCAETHFlowPoints(v string) {
	o.SCAETHFlowPoints = &v
}

// GetATT_NAME returns the ATT_NAME field value if set, zero value otherwise.
func (o *Capitalization) GetATT_NAME() string {
	if o == nil || o.ATT_NAME == nil {
		var ret string
		return ret
	}
	return *o.ATT_NAME
}

// GetATT_NAMEOk returns a tuple with the ATT_NAME field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Capitalization) GetATT_NAMEOk() (*string, bool) {
	if o == nil || o.ATT_NAME == nil {
		return nil, false
	}
	return o.ATT_NAME, true
}

// HasATT_NAME returns a boolean if a field has been set.
func (o *Capitalization) HasATT_NAME() bool {
	if o != nil && o.ATT_NAME != nil {
		return true
	}

	return false
}

// SetATT_NAME gets a reference to the given string and assigns it to the ATT_NAME field.
func (o *Capitalization) SetATT_NAME(v string) {
	o.ATT_NAME = &v
}

func (o Capitalization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SmallCamel != nil  {
		toSerialize["smallCamel"] = o.SmallCamel
	}
	if o.CapitalCamel != nil  {
		toSerialize["CapitalCamel"] = o.CapitalCamel
	}
	if o.SmallSnake != nil  {
		toSerialize["small_Snake"] = o.SmallSnake
	}
	if o.CapitalSnake != nil  {
		toSerialize["Capital_Snake"] = o.CapitalSnake
	}
	if o.SCAETHFlowPoints != nil  {
		toSerialize["SCA_ETH_Flow_Points"] = o.SCAETHFlowPoints
	}
	if o.ATT_NAME != nil  {
		toSerialize["ATT_NAME"] = o.ATT_NAME
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *Capitalization) UnmarshalJSON(bytes []byte) (err error) {
	varCapitalization := _Capitalization{}

	if err = json.Unmarshal(bytes, &varCapitalization); err == nil {
		*o = Capitalization(varCapitalization)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "smallCamel")
		delete(additionalProperties, "CapitalCamel")
		delete(additionalProperties, "small_Snake")
		delete(additionalProperties, "Capital_Snake")
		delete(additionalProperties, "SCA_ETH_Flow_Points")
		delete(additionalProperties, "ATT_NAME")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCapitalization struct {
	value *Capitalization
	isSet bool
}

func (v NullableCapitalization) Get() *Capitalization {
	return v.value
}

func (v *NullableCapitalization) Set(val *Capitalization) {
	v.value = val
	v.isSet = true
}

func (v NullableCapitalization) IsSet() bool {
	return v.isSet
}

func (v *NullableCapitalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCapitalization(val *Capitalization) *NullableCapitalization {
	return &NullableCapitalization{value: val, isSet: true}
}

func (v NullableCapitalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCapitalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


