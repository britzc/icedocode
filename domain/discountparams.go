package domain

import "time"

// DiscountParam provide the period of discount
type DiscountParam struct {
	Percentage float64
	StartTime  *time.Time
	EndTime    *time.Time
}
