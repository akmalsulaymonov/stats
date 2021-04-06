package stats

import (
	"fmt"

	"github.com/aki8787/bank/pkg/types"
)

func ExamplePayment() {
	payments := []types.Payment{
		{
			ID:       1,
			Amount:   10,
			Category: "El1",
		},
		{
			ID:       2,
			Amount:   10,
			Category: "El2",
		},
		{
			ID:       3,
			Amount:   10,
			Category: "El1",
		},
	}

	fmt.Println(TotalInCategory(payments, "El1"))

	//Output:
	// 20
}
