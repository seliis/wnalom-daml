package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"math"
	"os"
	"os/exec"
	"strings"
)

func GetNewCSV(name string) (*os.File, *csv.Writer, error) {
	if file, err := os.Create(name + ".csv"); err != nil {
		return nil, nil, err
	} else {
		return file, csv.NewWriter(file), nil
	}
}

func GetConfirm() (bool, error) {
	var response string
	fmt.Println("Continue?: (Y/N)")
	if _, err := fmt.Scan(&response); err != nil {
		return false, err
	}

	response = strings.TrimSpace(response)
	response = strings.ToLower(response)

	if response == "y" {
		return true, nil
	} else {
		return false, errors.New("Cancel")
	}
}

func PrintExpectation(genesis string, duration string, interval string) (float64, error) {
	unixGenesis, err := GetUnixGenesis(genesis)
	if err != nil {
		return 0, err
	}

	unixDuration, err := GetUnixDuration(duration)
	if err != nil {
		return 0, err
	}

	unixInterval, err := GetUnixInterval(interval)
	if err != nil {
		return 0, err
	}

	unixNow := GetUnixNow()

	expectedRequest := float64((unixNow-unixGenesis)/unixDuration) + 1
	expectedContain := unixDuration / unixInterval

	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

	fmt.Println("Expected Request:", expectedRequest)
	fmt.Println("Expected Contain:", expectedContain)
	fmt.Println("Expected Downloading Time:", math.Floor(expectedRequest*0.105), "Seconds")

	return expectedRequest, nil
}
