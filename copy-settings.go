package main

import (
	"bufio"
	"flag"
	"fmt"
	_ "image/png"
	"os"
	"strings"
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
	finFilename := flag.String("f", "", "final configuration file")
	splitSign := flag.String("s", "", "split sign")

	flag.Parse()

	if *oldFilename != "" && *newFilename != "" && *finFilename != "" && *splitSign != "" {

		fmt.Println("Old configuration file: ", *oldFilename)
		fmt.Println("New configuration file: ", *newFilename)
		fmt.Println("Final configuration file: ", *finFilename)
		fmt.Println("Split sign: ", *splitSign)

		oldFile, err := os.Open(*oldFilename)
		ErrCheck(err)
		defer oldFile.Close()
		newFile, err := os.Open(*newFilename)
		ErrCheck(err)
		defer newFile.Close()

		finFileMode, err := os.Stat(*newFilename)
		ErrCheck(err)

		finFile, err := os.Create(*finFilename)
		ErrCheck(err)
		defer finFile.Close()

		finFile.Chmod(finFileMode.Mode())

		scannerNew := bufio.NewScanner(newFile)
		for scannerNew.Scan() {

			keyNew := strings.Split(scannerNew.Text(), *splitSign)

			if len(keyNew) > 1 {

				oldFile.Seek(0, 0)
				scannerOld := bufio.NewScanner(oldFile)
				var keyOld []string
				found := false
				for scannerOld.Scan() {

					keyOld = strings.Split(scannerOld.Text(), *splitSign)

					if len(keyOld) > 1 {
						if keyOld[0] == keyNew[0] && keyOld[1] != keyNew[1] {
							fmt.Println(keyNew[0] + "=" + keyNew[1] + " --> " + keyOld[0] + "=" + keyOld[1])
							found = true
							break
						}
					}
				}
				if found {
					finFile.WriteString(keyOld[0] + *splitSign + keyOld[1] + "\n")
				} else {
					finFile.WriteString(keyNew[0] + *splitSign + keyNew[1] + "\n")
				}

			} else {
				finFile.WriteString(keyNew[0] + "\n")
			}

		}

	} else {
		fmt.Println("Use -h flag to read usage")
	}
}
