package uibench

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const DEBUG = false

type Uibench struct {
	config Config // Test cases to run
}

func New() Uibench {
	return Uibench{}
}

func (u *Uibench) LoadTestsConfig(filePath string) {
	config := loadConfig(filePath)
	u.config.Benchmarks = config.Benchmarks
}

func (u *Uibench) LoadDataConfig(filePath string) {
	config := loadConfig(filePath)
	u.config.Data = config.Data
}

func (u *Uibench) Automate() {
	// Check if adb is present
	if !isAdbPresent() {
		abort("ADB cannot be found!")
	}
	// TODO: Check if device connected
	// TODO: Check if USB debugging permissions have been granted
	// for each config test case, execute "adb shell am"
	for _, benchmarks := range u.config.Benchmarks {
		for idxTests, tests := range benchmarks.Tests {
			fmt.Printf("Executing: %s [%d time(s)]\n", tests.Name, tests.Iterations)
			output, err := execUibenchCommand(benchmarks.Class, tests.Method, tests.Iterations)
			if err != nil {
				fmt.Println("Failed to execute test case: [", benchmarks.Class, ",", tests.Method, ",", tests.Iterations, "]")
			} else {
				// Iterate through each line of adb shell command output
				scanner := bufio.NewScanner(strings.NewReader(output))
				for scanner.Scan() {
					line := scanner.Text()
					for _, data := range u.config.Data {
						// Inefficient. Figure out a way to avoid checking for pattern if we already have a result for it.
						// If line contains one of the patterns to look out for, get the result
						if strings.Contains(line, data.Pattern) {
							resultStr := substringBetween(line, data.Prefix, data.Suffix)
							resultFloat, err := strconv.ParseFloat(resultStr, 64)
							if err != nil {
								fmt.Println(resultStr, "cannot be converted to type float.")
								continue
							}

							// Store result
							var result Result = Result{name: data.Name, value: resultFloat}
							benchmarks.Tests[idxTests].results = append(benchmarks.Tests[idxTests].results, result)
						}
					}
				}
			}
		}
	}

	return
}

func (u *Uibench) Export(format string) {
	switch format {
	case "csv":
		u.exportCsv()
	default:
		fmt.Println("Invalid export format!")
	}
}

func abort(message string) {
	fmt.Println(message)
	fmt.Println("Aborting!")
	os.Exit(1)
}

type Config struct {
	Benchmarks []Benchmarks `json:"benchmarks"`
	Data       []Data       `json:"data"`
}

type Benchmarks struct {
	Class string  `json:"class"`
	Tests []Tests `json:"tests"`
}

type Tests struct {
	Name       string `json:"name"`
	Method     string `json:"method"`
	Iterations int    `json:"iterations"`
	results    []Result
}

type Data struct {
	Key     string `json:"key"`
	Name    string `json:"name"`
	Pattern string `json:"pattern"`
	Prefix  string `json:"prefix"`
	Suffix  string `json:"suffix"`
}

type Result struct {
	name  string
	value float64
}
