import (
    "encoding/csv"
    "fmt"
    "os"
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