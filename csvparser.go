package csvparser

import (
	"encoding/csv"
	"errors"
	"os"
)

type Result map[string]map[string]string

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

		resultRow := make(map[string]string)
		result[row[0]] = resultRow
		for i, value := range row {
			resultRow[heads[i]] = value
		}
	}

	return
}
