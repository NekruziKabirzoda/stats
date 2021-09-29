package stats

import (
	"github.com/NekruziKabirzoda/bank/pkg/types"
)

func Avg(payments []types.Payment) types.Money {

	sum := types.Money(0)

	for i:=0; i<len(payments); i++{
    sum+=payments[i].Amount
	}
	sum = sum/types.Money(len(payments))
return sum
}

func TotalInCategory (payments []types.Payment, category types.Category) types.Money{
	sum := types.Money(0)

	for i:=0; i<len(payments); i++{
		if payments[i].Category == category {
    sum+=payments[i].Amount
	}}
  return sum
}