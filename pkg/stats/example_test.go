package stats

import (
	"github.com/NekruziKabirzoda/bank/v2/pkg/types"
	"fmt"
)

func ExampleAvg() {
	payments := []types.Payment {
		{ID: 1, Amount:10, Status: "FAIL"},
		{ID: 2, Amount:400},
		{ID: 3, Amount:10, Status: "FAIL"},
	}


    fmt.Println(Avg (payments))
	//Output: 400
}

func ExampleTotalInCategory() {
	payments := []types.Payment {
		{ID: 1, Amount:10, Category: "a"},
		{ID: 2, Amount:400, Category: "b"},
		{ID: 3, Amount:10, Category: "a", Status: "FAIL"},
	}


    fmt.Println(TotalInCategory(payments, "a"))
	//Output: 10
}