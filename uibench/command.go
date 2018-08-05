package uibench

import (
	"os/exec"
	"strconv"
)

func execUibenchCommand(class string, method string, iterations int) (string, error) {
	output, err := exec.Command("adb", "shell", "am", "instrument", "-e", "iterations", strconv.Itoa(iterations), "-w", "-e", "class", "com.android.uibench.janktests."+class+"#"+method, "com.android.uibench.janktests/android.test.InstrumentationTestRunner").CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(output), nil
}

func isAdbPresent() bool {
	_, err := exec.LookPath("adb")
	return err == nil
}
