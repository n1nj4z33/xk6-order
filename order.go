package order

import (
	"encoding/json"
	"strconv"

	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/order", New())
}

type Module struct{}

func New() *Module {
	return &Module{}
}

type Order struct {
	OrderID int64 `json:"orderId"`
}

func (m *Module) GetOrderID(body string) string {
	var order Order
	err := json.Unmarshal([]byte(body), &order)
	if err != nil {
		println(err)
	}
	return strconv.FormatInt(int64(order.OrderID), 10)
}
