package orderEntities

import (
	"github.com/wiraphatys/shop-management-go/database"
)

type OrderData struct {
	CID        string               `json:"cid"`
	OrderLines []database.OrderLine `json:"order_lines"`
}

type OrderLineData struct {
	OID string `json:"oid"`
	PID string `json:"pid"`
}

// type OrderLineData struct {
// 	PID     string `json:"pid"`
// 	Quatity int    `json:"quantity"`
// }
