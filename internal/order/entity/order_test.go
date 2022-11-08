package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGivenEmptyId_whenCreateANewOrder_thenShouldReceveAnError(t *testing.T) {
	order := Order{}
	assert.Error(t, order.IsValid(), "invalid id")
}

func TestGivenEmptyPrice_whenCreateANewOrder_thenShouldReceveAnError(t *testing.T) {
	order := Order{Id: "123"}
	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenEmptyTax_whenCreateANewOrder_thenShouldReceveAnError(t *testing.T) {
	order := Order{Id: "123"}
	assert.Error(t, order.IsValid(), "invalid price")
}

func TestGivenAValidParams_WhenICallNewOrder_thenIShouldReceiveCreateOrderWIthAllParams(t *testing.T) {
	order := Order{
		Id:    "123",
		Price: 10.0,
		Tax:   2.0,
	}
	assert.Equal(t, "123", order.Id)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
	assert.Nil(t, order.IsValid())
}

func TestGivenAValidParams_WhenICallNewOrderFunc_thenIShouldReceiveCreateOrderWIthAllParams(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2)
	assert.Nil(t, err)
	assert.Equal(t, "123", order.Id)
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 2.0, order.Tax)
}

func TestGivenAPriceAndTax_WhenIcallCalculatePrice(t *testing.T) {
	order, err := NewOrder("123", 10.0, 2.0)
	assert.Nil(t, err)
	assert.Nil(t, order.CalculateFinalPrice())
	order.CalculateFinalPrice()
	assert.Equal(t, 12.0, order.FinalPrice)
}
