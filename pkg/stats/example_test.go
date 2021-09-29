package stats

import (
	"github.com/NekruziKabirzoda/bank/pkg/types"
	"fmt"
)

func ExampleAvg() {
	payments := []types.Payment {
		{ID: 1, Amount:10},
		{ID: 2, Amount:400},
		{ID: 3, Amount:10},
	}


    fmt.Println(Avg (payments))
	//Output: 140
}

func ExampleTotalInCategory() {
	payments := []types.Payment {
		{ID: 1, Amount:10, Category: "a"},
		{ID: 2, Amount:400, Category: "b"},
		{ID: 3, Amount:10, Category: "a"},
	}


    fmt.Println(TotalInCategory(payments, "a"))
	//Output: 20
}