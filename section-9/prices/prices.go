package prices

import (
	"fmt"
)

type iOManager interface {
	FromDataToFloat() ([]float64, error)
}

type TaxIncludedPriceJob struct {
	IOManager         iOManager
	TaxRate           float64
	InputPrices       []float64
	TaxIncludedPrices map[string]float64
}

func NewTaxIncludedPriceJob(ioManager iOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager:   ioManager,
		TaxRate:     taxRate,
		InputPrices: []float64{},
	}
}

func (job TaxIncludedPriceJob) Process() error {
	err := job.loadData()

	if err != nil {
		return err
	}

	result := map[string]string{}

	for _, price := range job.InputPrices {
		taxIncludedPrice := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedPrice)
	}

	fmt.Println(result)
	return nil
}

func (job *TaxIncludedPriceJob) loadData() error {
	readPrices, err := job.IOManager.FromDataToFloat()

	if err != nil {
		return err
	}

	job.InputPrices = readPrices
	return nil
}
