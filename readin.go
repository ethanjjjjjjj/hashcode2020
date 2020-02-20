package main

import (
	"bufio"
	"os"
	"strconv"
)

func writeOutput(o Output) {
	f, _ := os.Create("out.txt")
	w := bufio.NewWriter(f)

	w.WriteString(strconv.Itoa(o.numLibraries) + "\n")

	for _, l := range o.libraryScans {
		w.WriteString(strconv.Itoa(l.libraryID) + " " + strconv.Itoa(l.numBooks) + "\n")
		for _, b := range l.books {
			w.WriteString(strconv.Itoa(b) + " ")
		}
		w.WriteString("\n")
	}
	w.Flush()
}

/*func readIn(){
	file, err := os.Open("a_example.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	books, _ := strconv.Atoi(string(txtlines[0][0]))
	numlibs, _:= strconv.Atoi(string(txtlines[0][2]))
	days, _:= strconv.Atoi(string(txtlines[0][4]))

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	books, _ := strconv.Atoi(string(txtlines[0][0]))
	numlibs, _:= strconv.Atoi(string(txtlines[0][2]))
	days, _:= strconv.Atoi(string(txtlines[0][4]))

	dataset := Dataset{books:books,numLibraries:numlibs,day:days}
	fmt.Println(dataset)

	fmt.Println(string(txtlines[0][0]))
	for _, eachline := range txtlines {
		fmt.Println(eachline)
	}

}*/
