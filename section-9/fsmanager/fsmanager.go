package fsmanager

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

type FileManager struct {
	ReadFileName string
}

func New(readFileName string) *FileManager {
	return &FileManager{
		ReadFileName: readFileName,
	}
}

func (fm FileManager) FromDataToFloat() ([]float64, error) {
	file, err := os.Open(fm.ReadFileName)

	if err != nil {
		return nil, errors.New("error opening file: " + err.Error())
	}

	inputPrices := []float64{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		price, err := strconv.ParseFloat(line, 64)

		if err != nil {
			return nil, errors.New("error parsing file: " + err.Error())
		}

		inputPrices = append(inputPrices, price)
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.New("error reading file: " + err.Error())
	}

	defer file.Close()
	return inputPrices, nil
}
