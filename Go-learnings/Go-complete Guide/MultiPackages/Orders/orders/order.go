package orders

import (
	"errors"

	"github.com/prajna/orderapp/pricing"
	"github.com/prajna/orderapp/utils"
)

type Order struct {
	ID    string
	Price float64
}

func NewOrder(id string, price float64) *Order {
	utils.Print("Creating order " + id)
	return &Order{
		ID:    id,
		Price: price,
	}
}

func (o *Order) FinalPrice() float64 {
	return pricing.CalculateTotal(o.Price)
}

func (o *Order) Validate() error {
	if o.Price <= 0 {
		return errors.New("order price must be greater than zero")
	}
	return nil
}
