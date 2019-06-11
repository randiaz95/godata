package main

import (
	"fmt"
	"io/ioutil"
)

func ReadCSV(filename string) (records [][]string, err error) {
    f, err := os.Open(filename)
    if err != nil {
        return nil, fmt.Errorf("error opening file: %v", err)
    }
    defer f.Close()

    records, err = csv.NewReader(f).ReadAll()
    if err != nil {
        return nil, fmt.Errorf("error reading file contents as csv: %v", err)
    }

    return records, nil
}

func SplitMeasurements(records [][]string, label string) []History {

	var labelIndex int
	for index, header := range records[0] {
		if header == label {
			labelIndex = index
			break
		}
	}

	var observation History
	observations := make([]History, 0, len(records)-1)
	for recordIndex, row := range records[1:] {
		
		for columnIndex, value := range row {

			if columnIndex == labelIndex {

				observation.Label = row
				continue

			}

			observation.Measurements = append(observation.Measurements, value)
		}

		observations = append(observations, observation)
		observation = History{}

	}

	return observations
}