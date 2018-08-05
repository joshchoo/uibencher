# UIBencher
UIBencher makes it easy to automate the execution of multiple UIBench tests, extract the results, and export them.

## Getting Started

### Prerequisites
* Computer with ADB and Golang installed.
* Enabled "USB debugging" and "Stay Awake" toggles in Developer Options.
* Installed UiBench.apk and UiBenchJankTests.apk (provided in apks folder).

## Running the tests

### Setting up
Clone the UIBencher repo to your $GOPATH or $GOROOT.

To keep the testing environment as consistent as possible, please follow the following steps:
* Connect your phone to your computer.
* Reboot your phone.
* Force close your recent apps.
* Enable Airplane mode, and disable automatic brightness.
* Ensure that CPU frequencies have settled down.

### Customizing tests to run (optional)
* Remove any tests that you do not want to run from _benchmark.json_.
* Change the number of iterations to run for each test. Take note that increasing the iterations beyond the defaults may cause some tests to fail to complete for unknown reasons. It should be safe to decrease the number of iterations.

### Customizing the results to collect (optional)
* Modify _data.json_ only if you know what you are doing.

### Instructions
* From your computer terminal, run: 
```bash
go run main.go
```
* A _uibench_results.csv_ will be produced.
