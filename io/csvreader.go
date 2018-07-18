package io

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// reads a CSV file
func ReadCSV(fileName string) {
	// open the file
	//f, err := os.Open("../data/data1.csv")
  tempDir := os.TempDir()
	println("tempDir: ", tempDir)

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// create a new CSV reader
	fmt.Printf("Reading the file...\n")

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1

	// read all the records
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("number of records: ", len(records))
  // fmt.Println("records: ", records)
  // fmt.Println("records: ", records[0])
  for i, j := range records {
		fmt.Printf("\nline %d => %s values =>", i, j)
		for k, z := range j {
			if k == (len(j) - 1) {
			    fmt.Printf(" %s", z)
			} else {
					fmt.Printf(" %s,", z)
			}
		}
	}
	fmt.Printf("\n")
}
