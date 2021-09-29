package stats

import (
	"github.com/NekruziKabirzoda/bank/v2/pkg/types"
)

func Avg(payments []types.Payment) (types.Money) {

	sum := types.Money(0)
	//a :=[]types.Payment{}
	b:=1

	for i:=0; i<len(payments); i++{
		if payments[i].Status==types.StatusFail{
			continue
			
		}
			 
		
    sum+=payments[i].Amount
	b++
	}
	
	
	sum = sum/types.Money(b-1)
return sum
}

func TotalInCategory (payments []types.Payment, category types.Category) types.Money{
	sum := types.Money(0)

	for i:=0; i<len(payments); i++{
		if payments[i].Category == category{
			if   payments[i].Status== types.StatusFail{
				continue
			}
    sum+=payments[i].Amount
	}}
  return sum
} 