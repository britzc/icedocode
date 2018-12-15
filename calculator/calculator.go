package calculator

import "icedo/sandbox/domain"

type DiscountAgent interface {
	GetDiscountParams(productID int) (params []*domain.DiscountParam)
}

type Calculator struct {
	da DiscountAgent
}

func New(da DiscountAgent) (calculator *Calculator) {
	z := &Calculator{
		da: da,
	}

	return z
}
