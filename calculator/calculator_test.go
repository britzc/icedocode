// +build unit_test

package calculator

import (
	"fmt"
	"icedo/sandbox/domain"
	"testing"
)

type MockDiscountAgent struct {
	Params []*domain.DiscountParam
}

func (z *MockDiscountAgent) GetDiscountParams(productID int) (params []*domain.DiscountParam) {
	return z.Params
}

func Test_New(t *testing.T) {
	mda := &MockDiscountAgent{}

	fmt.Println("yyyayayyayayayyayayya")

	actual := New(mda)
	if actual == nil {
		t.Errorf("crap")
	}
}
