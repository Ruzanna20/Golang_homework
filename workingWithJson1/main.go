package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Record struct {
	Time    string  `json:"time"` 
	Station string  `json:"station"`
	Temp    float64 `json:"temperature"`
	Rain    float64 `json:"rain"`
}

func readRecord(r io.Reader) (Record, error) {
	var record Record
	decoder := json.NewDecoder(r)
	decoder.Decode(&record)
	return record, nil
}

func main() {

	file, _ := os.Open("data.json")
	defer file.Close()

	record, _ := readRecord(file)
	fmt.Println(record)

}
