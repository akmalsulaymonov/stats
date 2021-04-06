package stats

import "github.com/aki8787/bank/v2/pkg/types"

func Avg(payment []types.Payment) types.Money {
	sum := 0
	k := 0
	for _, i := range payment {
		if i.Status != "FAIL" {
			sum += int(i.Amount)
			k ++
		}
	}
	return types.Money(sum / k)
}

func TotalInCategory(payment []types.Payment, category types.Category) types.Money {
	sum := 0
	for _, i := range payment {
		if i.Category == category && i.Status != "FAIL" {
			sum += int(i.Amount)
		}
	}
	return types.Money(sum)
}
