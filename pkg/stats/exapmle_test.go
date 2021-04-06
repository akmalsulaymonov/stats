package stats

import (
	"fmt"

	"github.com/aki8787/bank/v2/pkg/types"
)

func ExamplePayment() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10,
			Category: "El1",
			Status: "FAIL",
		},
		{
			ID:       2,
			Amount:   10,
			Category: "El2",
			Status: "OK",
		},
		{
			ID:       3,
			Amount:   10,
			Category: "El1",
			Status: "INPROGRESS",
		},
	}

	fmt.Println(Avg(payments))
	fmt.Println(TotalInCategory(payments, "El1"))

	//Output:
	// 10
	// 10
}
