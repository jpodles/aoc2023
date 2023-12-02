package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetFileScanner(fileName string) *bufio.Scanner {
	readFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	return fileScanner
}
