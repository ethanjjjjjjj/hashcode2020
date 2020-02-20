package main

import (
	"main"
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

func readIn(){
	d, err := os.Open("a_example.txt")
	scanner :=  bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
	//data =: Dataset{books = }
	
}