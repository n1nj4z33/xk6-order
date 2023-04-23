package order

import (
	"encoding/json"
	"strconv"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/order", New())
}

func New() *Order {
	return &Order{}
}

type Order struct {
	OrderID int64 `json:"orderId"`
}

func (o *Order) GetOrderID(body string) string {
	err := json.Unmarshal([]byte(body), &o)
	if err != nil {
		println(err)
	}
	return strconv.FormatInt(int64(o.OrderID), 10)
}
