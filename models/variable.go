package models

type Size struct {
	Size_ID      int    `json:"size_id"`
	Size_name_th string `json:"size_name_th"`
	Size_name_en string `json:"size_name_en"`
	Size_price   int    `json:"size_price"`
	Size_Stock   int    `json:"size_stock"`
}

type Flavor struct {
	Flavor_ID      int    `json:"flavor_id"`
	Flavor_name_th string `json:"flavor_name_th"`
	Flavor_name_en string `json:"flavor_name_en"`
	Flavor_price   int    `json:"flavor_price"`
	Flavor_Stock   int    `json:"flavor_stock"`
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
	Sauce_name_th string `json:"sauce_name_th"`
	Sauce_name_en string `json:"sauce_name_en"`
	Sauce_price   int    `json:"sauce_price"`
	Sauce_Stock   int    `json:"sauce_stock"`
}

type Order_detail struct {
	Order_id   int `json:"order_id"`
	Size_name_en string `json:"size_name_en"`
	Flavor_name_en string `json:"flavor_name_en"`
	Topping_name_en []string `json:"tp_name_en"`
	Sauce_name_en string `json:"sauce_name_en"`
	Sum_Price      int `json:"price"`
}
