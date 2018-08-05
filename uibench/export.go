package uibench

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func (u *Uibench) exportCsv() {
	file, err := os.Create("uibench_results.csv")
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write column heading row
	var headingRow []string
	headingRow = append(headingRow, "Name")
	for _, dataHeading := range u.config.Data {
		headingRow = append(headingRow, dataHeading.Name)
	}
	headingRow = append(headingRow, "Iterations")
	err = writer.Write(headingRow)
	checkError("Failed to write to csv file!", err)

	for _, benchmarks := range u.config.Benchmarks {
		for _, tests := range benchmarks.Tests {
			var resultRow []string

			// Append test name
			resultRow = append(resultRow, tests.Name)

			// Append test results
			for _, results := range tests.results {
				// Assumes that:
				// 1. All results were obtained for data.json in Automate() function
				// 2. Stored results are in the same order as data.json
				resultRow = append(resultRow, strconv.FormatFloat(results.value, 'f', -1, 64))
			}

			// Append number of iterations
			resultRow = append(resultRow, strconv.Itoa(tests.Iterations))

			// Write row to CSV
			err := writer.Write(resultRow)
			checkError("Failed to write to csv file!", err)
		}
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
