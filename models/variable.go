package models

import (
	"time"
)

type Size struct {
	Size_ID      int    `json:"size_id"`
	Size_name_th string `json:"size_name_th"`
	Size_name_en string `json:"size_name_en"`
	Size_price   int    `json:"size_price"`
	Size_Stock   int    `json:"size_stock"`
}

type Flavor struct {
	Flavor_ID      int    `json:"flavor_id"`
	Flavor_name_th string `json:"f_name_th"`
	Flavor_name_en string `json:"f_name_en"`
	Flavor_price   int    `json:"f_price"`
	Flavor_Stock   int    `json:"f_stock"`
}

type Topping struct {
	Topping_ID      int    `json:"topping_id"`
	Topping_name_th string `json:"tp_name_th"`
	Topping_name_en string `json:"tp_name_en"`
	Topping_price   int    `json:"tp_price"`
	Topping_Stock   int    `json:"tp_stock"`
}

type Sauce struct {
	Sauce_ID      int    `json:"sauce_id"`
	Sauce_name_th string `json:"s_name_th"`
	Sauce_name_en string `json:"s_name_en"`
	Sauce_price   int    `json:"s_price"`
	Sauce_Stock   int    `json:"s_stock"`
}

type Order_detail struct {
	Order_id   int `json:"detail_id"`
	Size_ID    int `json:"size_id"`
	Flavor_ID  int `json:"flavor_id"`
	Topping_ID int `json:"topping_id"`
	Sauce_ID   int `json:"sauce_id"`
	Price      int `json:"price"`
}

type Payment struct {
	Payment_id     int       `json:"pm_id"`
	Payment_method string    `json:"pm_method"`
	Payment_amount int       `json:"pm_amount"`
	Order_id       int       `json:"id"`
	Payment_date   time.Time `json:"payment_date"`
	Payment_time   time.Time `json:"payment_time"`
}
