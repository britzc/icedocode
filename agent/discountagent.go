package agent

import (
	"icedo/sandbox/domain"

	nats "github.com/nats-io/go-nats"
)

func New() (discountAgent *DiscountAgent) {
	nc, _ := nats.Connect(nats.DefaultURL)

	z := &DiscountAgent{
		nc: nc,
	}

	return z
}

type DiscountAgent struct {
	nc *nats.Conn
}

func GetDiscountParams(productID int) (params []*domain.DiscountParam) {
	return nil
}
