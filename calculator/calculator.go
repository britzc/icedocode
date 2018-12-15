package calculator

type Product interface {
	GetUnitPrice() (price float64)
}

func GetDiscount(product Product, quantity int) (discount float64) {
	if quantity < 10 {
		return 0.0
	}

	return product.GetUnitPrice() * 0.15 * float64(quantity)
}
