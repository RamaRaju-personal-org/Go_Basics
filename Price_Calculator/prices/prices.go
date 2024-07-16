package prices

import (
	"fmt"

	"examples.com/FileManager"
	"examples.com/conversion"
)

//struct
type TaxIncludedPriceJob struct {
	TaxRate float64
	InputPrices []float64
	TaxIncludedPrices map[string]string // key : input price converted to string,  Value: TaxIncludedPrices
}

//method which loads prices from price.txt file  
func (job *TaxIncludedPriceJob) LoadData() {

	// Outsourcing File Access Into A Package
	lines, err := FileManager.ReadLines("prices.txt")
	if err != nil {
		fmt.Println(err)
	}
	
	prices, err := conversion.StringsToFloats(lines)
	// convert prices string to float64 using  StringsToFloats func 
	
		if err != nil {
		    fmt.Println(err)
		}

	job.InputPrices =  prices
	

}

//method which calculates tax included price 
func (job *TaxIncludedPriceJob) Process() {

	job.LoadData()

	result := make(map[string]string) //empty map ---> key :input price converted to string,  Value: TaxIncludedPrices converted to string
		for _, price := range job.InputPrices {
			taxIncludedPrice := price * (1+job.TaxRate)
			result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
			// key: price i.e job.InputPrices converted to string using Sprintf  to 2 decimal values
			// value : Tax Included Prices ---> price * (1+job.TaxRate)
		}
		
	 job.TaxIncludedPrices = result

	 FileManager.WriteJSON(fmt.Sprintf("result_%.0f.json",job.TaxRate*100), job)

	}


//constructor returns a struct and uses pointer 
func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
  return &TaxIncludedPriceJob{
	InputPrices: []float64{10, 20, 30},
	TaxRate: taxRate,
  }
}
