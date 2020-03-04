package main

import (
	"bufio"
	"flag"
	"fmt"
	_ "image/png"
	"os"
)

// ErrCheck - obsługa błedów
// ================================================================================================
func ErrCheck(errNr error) bool {
	if errNr != nil {
		fmt.Println(errNr)
		return false
	}
	return true
}

// main - program entry point
// ================================================================================================
func main() {

	fmt.Println("===============================")
	fmt.Println("= Copy settings between files =")
	fmt.Println("===============================")

	oldFilename := flag.String("o", "", "old configuration file")
	newFilename := flag.String("n", "", "new configuration file")

	flag.Parse()

	if *oldFilename != "" && *newFilename != "" {

		fmt.Println("Old configuration file: ", *oldFilename)
		fmt.Println("New configuration file: ", *newFilename)

		oldFile, err := os.Open(*oldFilename)
		ErrCheck(err)
		defer oldFile.Close()

		scanner1 := bufio.NewScanner(oldFile)
		for scanner1.Scan() {
			fmt.Println(scanner1.Text()) // Println will add back the final '\n'
		}

		newFile, err := os.Open(*newFilename)
		ErrCheck(err)
		defer newFile.Close()

		scanner2 := bufio.NewScanner(oldFile)
		for scanner2.Scan() {
			fmt.Println(scanner2.Text()) // Println will add back the final '\n'
		}

	} else {
		fmt.Println("Use -h flag to read usage")
	}
}
