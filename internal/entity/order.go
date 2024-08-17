package entity

import (
	"errors"

)

var msgerros = []string{"Invalid ID", "Invalid price", "Invalid tax" }


type Order struct{
	ID 				string
	Price			float64
	Tax				float64
	FinalPrice		float64

}

func NewOrder(id string, price float64, tax float64) (*Order, error){
	order := &Order{
		ID: id,
		Price: price ,
		Tax: tax,
	}
	err := order.IsValid()
	if err != nil {
		return nil, err
	}

	return order, nil
}


func (o *Order) IsValid() error{

	if o.ID == ""{
		return errors.New(msgerros[0])
	}
	if o.Price <= 0{
		return errors.New(msgerros[1])
	}
	if o.Tax <= 0{
		return errors.New(msgerros[2])
	}
	return nil
}

func (o *Order) CalculateFinalPrice() error{
	o.FinalPrice = o.Price + o.Tax
	err := o.IsValid()
	if err != nil {
		return err
	}
	return nil
}