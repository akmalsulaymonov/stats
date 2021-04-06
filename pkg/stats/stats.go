package stats

import "github.com/aki8787/bank/pkg/types"

func Avg(payment []types.Payment) types.Money {
	sum := 0
	for _, i := range payment {
		sum += int(i.Amount)
	}
	return types.Money(sum / len(payment))
}

func TotalInCategory(payment []types.Payment, category types.Category) types.Money {
	sum := 0
	for _, i := range payment {
		if i.Category == category {
			sum += int(i.Amount)
		}
	}
	return types.Money(sum)
}
