package main

import (
	"fmt"

	"github.com/lfsc09/udemy-go-complete-guide/section-9/fsmanager"
	"github.com/lfsc09/udemy-go-complete-guide/section-9/prices"
)

func main() {
	fileManager := fsmanager.New("prices.txt")
	taxRates := []float64{0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		priceJob := prices.NewTaxIncludedPriceJob(fileManager, taxRate)
		err := priceJob.Process()

		if err != nil {
			fmt.Printf("Error processing prices for taxRate[%.2f]: %v\n", taxRate, err)
		}
	}
}
