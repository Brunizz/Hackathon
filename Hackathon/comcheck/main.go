package main

import (
	"fmt"
	"os"
)

func main() {
	result := ""
	args := os.Args[1:]
	trigger := []string{"01", "galaxy", "galaxy 01"}
	for _, res := range args {
		for _, word := range trigger {
			if res == word {
				result += "Alert!!!\n"
				break
			}
		}
	}
	fmt.Print(result)
}
