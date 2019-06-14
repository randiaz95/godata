package main

import (
	"os"
	"fmt"
	"strconv"
	"encoding/csv"
	"github.com/hlts2/gohot"
)

type Column struct {
	Header string
	Tipe string
	StringValues []string
	FloatValues [][]float64
}

func (c *Column) Process() {

	c.Header = c.StringValues[0]
	_, err := strconv.ParseFloat(c.StringValues[1], 64)

	if err == nil {
		c.Tipe = "quantitative"
		c.StringValues = c.StringValues[1:]
		c.FloatValues = FloatMapper(c.StringValues[1:])
	} else {
		c.Tipe = "categorical"
		c.StringValues = c.StringValues[1:]
		hotmap := gohot.CreateOneHotVectorFromTokens(c.StringValues[1:])
		c.FloatValues = FloatVectorizer(c.StringValues[1:], hotmap)
	}

}

func FloatMapper(input []string) [][]float64 {

	output := [][]float64{}
	for index, field := range input {
		floatfield, err := strconv.ParseFloat(field, 64)
		if err != nil {
			panic("Column has non numeric type: " + field)
		}
		output[index] = []float64{floatfield}
	}
	return output

}

func FloatVectorizer(input []string, hotmap map[string]string) [][]float64{
	
	output := [][]float64{}
	for _, field := range input[1:] {
		vector := []float64{}

		for _, number := range hotmap[field] {
			floatfield, err := strconv.ParseFloat(string(number), 64)
			if err != nil {
				panic("Failed to process categorical field: " + field)
			}
			vector = append(vector, floatfield)
		}

		output = append(output, vector)
	}
	return output
}


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

	allcolumns := []Column{}

	for col:=0; col<len(records[0]); col++ {
		
		if col == labelIndex {
			continue
		}

		supercolumn := Column{}
		
		for row:=0; row<len(records); row++ {
			
			supercolumn.StringValues = append(supercolumn.StringValues, records[row][col])

		}
		
		supercolumn.Process()
		allcolumns = append(allcolumns, supercolumn)
	}
	
	fmt.Println("LENGTH OF COLUMNS: ", len(allcolumns))
	for index, column := range allcolumns {
		fmt.Printf("Length of float values in column %s is %s\n", string(index), len(column.FloatValues))
		fmt.Printf("Length of string values in column %s is %s\n", string(index), len(column.StringValues))
	
	}

	observations := make([]History, len(records)-1)

	for col:=0; col<len(allcolumns); col++ {
		observations = append(observations, History{Label: allcolumns[col].Header,
												   Measurements: []float64{},
												  })
	}

	for rowIndex, obs := range observations {
		for _, column := range allcolumns {

				obs.Measurements = append(obs.Measurements, column.FloatValues[rowIndex]...)
			
		}
	}

	return observations

}
