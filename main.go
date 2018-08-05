package main

import (
	"fmt"
	"uibencher/uibench"
)

func main() {
	fmt.Println("Running UIBencher...")
	ub := uibench.New()
	ub.LoadTestsConfig("./benchmarks.json")
	ub.LoadDataConfig("./data.json")
	ub.Automate()
	ub.Export("csv")
	fmt.Printf("\n\nDone!")
}
