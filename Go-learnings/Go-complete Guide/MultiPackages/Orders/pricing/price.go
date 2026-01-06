package pricing

const TaxRate = 0.18 // 18% GST

func CalculateTotal(basePrice float64) float64 {
	return basePrice + (basePrice * TaxRate)
}
