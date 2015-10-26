package csvparser

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
)

type CsvValue string

func (this CsvValue) Int(v int) int {
	value, err := strconv.ParseInt(string(this), 10, 32)
	if err != nil {
		return v
	}

	return int(value)
}

func (this CsvValue) Uint(v uint) uint {
	value, err := strconv.ParseUint(string(this), 10, 32)
	if err != nil {
		return v
	}

	return uint(value)
}

func (this CsvValue) Float(v float32) float32 {
	value, err := strconv.ParseFloat(string(this), 32)
	if err != nil {
		return v
	}

	return float32(value)
}

func (this CsvValue) Str() string {
	return string(this)
}

type Result map[string]map[string]CsvValue

func Parse(fileName string) (result Result, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}

	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		return
	}

	if len(records) < 3 {
		err = errors.New("record empty")
		return
	}

	heads := records[1]
	if len(heads) < 1 {
		err = errors.New("head empty")
		return
	}

	result = make(Result)
	for _, row := range records[2:] {
		if len(row) < 1 {
			result = nil
			err = errors.New("cols empty")
			return
		}

		resultRow := make(map[string]CsvValue)
		result[row[0]] = resultRow
		for i, value := range row {
			resultRow[heads[i]] = CsvValue(value)
		}
	}

	return
}
