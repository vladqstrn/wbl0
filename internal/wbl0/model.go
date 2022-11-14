package wbl0

import (
	"time"
)

type Order struct {
	Order_uid          string    `json:"order_uid" validate:"min=36"`
	Track_number       string    `json:"track_number" validate:"min=1"`
	Entry              string    `json:"entry" validate:"min=1"`
	Locale             string    `json:"locale" validate:"min=1"`
	Internal_signature string    `json:"internal_signature"`
	Customer_id        string    `json:"customer_id" validate:"min=1"`
	Delivery_service   string    `json:"delivery_service" validate:"min=1"`
	Shardkey           string    `json:"shardkey" validate:"min=1"`
	Sm_id              int       `json:"sm_id" validate:"min=1"`
	Date_created       time.Time `json:"date_created"`
	Oof_shard          string    `json:"oof_shard" validate:"min=1"`
	Delivery           Delivery  `json:"delivery"`
	Payment            Payment   `json:"payment"`
	Item               []Item    `json:"items"`
}

type Delivery struct {
	Name    string `json:"name" validate:"min=1"`
	Phone   string `json:"phone" validate:"e164"`
	Zip     string `json:"zip" validate:"min=1"`
	City    string `json:"city" validate:"min=1"`
	Address string `json:"address" validate:"min=1"`
	Region  string `json:"region" validate:"min=1"`
	Email   string `json:"email" validate:"email"`
}

type Payment struct {
	Transaction   string `json:"transaction" validate:"min=1"`
	Request_id    string `json:"request_id"`
	Currency      string `json:"currency" validate:"min=1"`
	Provider      string `json:"provider" validate:"min=1"`
	Amount        int    `json:"amount" validate:"min=1"`
	Payment_dt    int    `json:"payment_dt" validate:"min=1"`
	Bank          string `json:"bank" validate:"min=1"`
	Delivery_cost int    `json:"delivery_cost" validate:"min=1"`
	Goods_total   int    `json:"goods_total" validate:"min=1"`
	Custom_fee    int    `json:"custom_fee" validate:"min=0"`
}

type Item struct {
	Chrt_id      int    `json:"chrt_id"`
	Track_number string `json:"track_number"`
	Price        int    `json:"price"`
	Rid          string `json:"rid"`
	Name         string `json:"name"`
	Sale         int    `json:"sale"`
	Size         string `json:"size"`
	Total_price  int    `json:"total_price"`
	Nm_id        int    `json:"nm_id"`
	Brand        string `json:"brand"`
	Status       int    `json:"status"`
}
