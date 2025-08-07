package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println("Hello world!")

	file, err := os.Open("test.csv")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	csvReader := csv.NewReader(file)

	for {
		line, err := csvReader.Read()

		if err != nil {
			if err == io.EOF {
				break
			}

			log.Println(err)
			continue
		}

		for i, element := range line {
			if i < len(line)-1 {
				fmt.Printf("%s - ", element)
			} else {
				fmt.Print(element)
			}
		}
		fmt.Println()
	}

}
