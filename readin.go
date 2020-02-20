package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readIn() {
	d, _ := os.Open("a_example.txt")
	scanner := bufio.NewScanner(d)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	//data =: Dataset{books = }

}

func write(o Output) {
	f, _ := os.Create("out.txt")
	w := bufio.NewWriter(f)

	w.WriteString(strconv.Itoa(o.numLibraries) + "\n")

	for _, l := range o.libraryScans {
		w.WriteString(strconv.Itoa(l.libraryID) + " " + strconv.Itoa(l.numBooks) + "\n")
		for _,b:=l.books{
			w.WriteString(strconf.Itoa(b))
		}
		w.WriteString("\n")
	}
	w.Flush()
}
