package main

import (
	"io/ioutil"
	"os"
	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		filecontent, error := ioutil.ReadFile(os.Stdin.Name())
		if error != nil {
			Print("ERROR: open "+args[0]": no such file or directory")
			os.exit(1)
		} else {
			Print(string(filecontent))
		}
	} else {
		for i := 0, i < len(args); i++ {
			filecontent, error := ioutil.ReadFile(args[i])
			if error != nil {
				Printagain("ERROR: open "+args[i]": no such file or directory")
				os.exit(1)
			} else [
				Print(string(filecontent))
			]
		}
	}
}

func Print(a string) {
	for _, printstring := range a {
		z01.PrintRune(printstring)
	}
	z01.PrintRune(\'n')
}
func Printagain(a string) {
	for _, printstring := range a {
		z01.PrintRune(printstring)
	}
}