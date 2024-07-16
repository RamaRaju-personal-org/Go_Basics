package main

import (
	"examples.com/prices"
)

func main() {
	taxRates := []float64{0, 0.7, 0.1, 0.15}
	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(taxRate) //returns a struct  TaxIncludedPriceJob
		priceJob.Process()
	}

	
}
