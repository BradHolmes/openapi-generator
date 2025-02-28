/*
OpenAPI Petstore

This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package petstore

import (
	"encoding/json"
	"time"
)

// Order struct for Order
type Order struct {
	Id *int64 `json:"id,omitempty"`
	PetId *int64 `json:"petId,omitempty"`
	Quantity *int32 `json:"quantity,omitempty"`
	ShipDate *time.Time `json:"shipDate,omitempty"`
	// Order Status
	Status *string `json:"status,omitempty"`
	Complete *bool `json:"complete,omitempty"`
}

// NewOrder instantiates a new Order object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrder() *Order {
	this := Order{}
	var complete bool = false
	this.Complete = &complete
	return &this
}

// NewOrderWithDefaults instantiates a new Order object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderWithDefaults() *Order {
	this := Order{}
	var complete bool = false
	this.Complete = &complete
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Order) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Order) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Order) SetId(v int64) {
	o.Id = &v
}

// GetPetId returns the PetId field value if set, zero value otherwise.
func (o *Order) GetPetId() int64 {
	if o == nil || o.PetId == nil {
		var ret int64
		return ret
	}
	return *o.PetId
}

// GetPetIdOk returns a tuple with the PetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetPetIdOk() (*int64, bool) {
	if o == nil || o.PetId == nil {
		return nil, false
	}
	return o.PetId, true
}

// HasPetId returns a boolean if a field has been set.
func (o *Order) HasPetId() bool {
	if o != nil && o.PetId != nil {
		return true
	}

	return false
}

// SetPetId gets a reference to the given int64 and assigns it to the PetId field.
func (o *Order) SetPetId(v int64) {
	o.PetId = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *Order) GetQuantity() int32 {
	if o == nil || o.Quantity == nil {
		var ret int32
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetQuantityOk() (*int32, bool) {
	if o == nil || o.Quantity == nil {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *Order) HasQuantity() bool {
	if o != nil && o.Quantity != nil {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int32 and assigns it to the Quantity field.
func (o *Order) SetQuantity(v int32) {
	o.Quantity = &v
}

// GetShipDate returns the ShipDate field value if set, zero value otherwise.
func (o *Order) GetShipDate() time.Time {
	if o == nil || o.ShipDate == nil {
		var ret time.Time
		return ret
	}
	return *o.ShipDate
}

// GetShipDateOk returns a tuple with the ShipDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetShipDateOk() (*time.Time, bool) {
	if o == nil || o.ShipDate == nil {
		return nil, false
	}
	return o.ShipDate, true
}

// HasShipDate returns a boolean if a field has been set.
func (o *Order) HasShipDate() bool {
	if o != nil && o.ShipDate != nil {
		return true
	}

	return false
}

// SetShipDate gets a reference to the given time.Time and assigns it to the ShipDate field.
func (o *Order) SetShipDate(v time.Time) {
	o.ShipDate = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Order) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Order) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Order) SetStatus(v string) {
	o.Status = &v
}

// GetComplete returns the Complete field value if set, zero value otherwise.
func (o *Order) GetComplete() bool {
	if o == nil || o.Complete == nil {
		var ret bool
		return ret
	}
	return *o.Complete
}

// GetCompleteOk returns a tuple with the Complete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Order) GetCompleteOk() (*bool, bool) {
	if o == nil || o.Complete == nil {
		return nil, false
	}
	return o.Complete, true
}

// HasComplete returns a boolean if a field has been set.
func (o *Order) HasComplete() bool {
	if o != nil && o.Complete != nil {
		return true
	}

	return false
}

// SetComplete gets a reference to the given bool and assigns it to the Complete field.
func (o *Order) SetComplete(v bool) {
	o.Complete = &v
}

func (o Order) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil  {
		toSerialize["id"] = o.Id
	}
	if o.PetId != nil  {
		toSerialize["petId"] = o.PetId
	}
	if o.Quantity != nil  {
		toSerialize["quantity"] = o.Quantity
	}
	if o.ShipDate != nil  {
		toSerialize["shipDate"] = o.ShipDate
	}
	if o.Status != nil  {
		toSerialize["status"] = o.Status
	}
	if o.Complete != nil  {
		toSerialize["complete"] = o.Complete
	}
	return json.Marshal(toSerialize)
}

type NullableOrder struct {
	value *Order
	isSet bool
}

func (v NullableOrder) Get() *Order {
	return v.value
}

func (v *NullableOrder) Set(val *Order) {
	v.value = val
	v.isSet = true
}

func (v NullableOrder) IsSet() bool {
	return v.isSet
}

func (v *NullableOrder) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrder(val *Order) *NullableOrder {
	return &NullableOrder{value: val, isSet: true}
}

func (v NullableOrder) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrder) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


