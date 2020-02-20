package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func readIn() Dataset {
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

	//first values

	params := make([]string, 3)
	for i, num := range strings.Split(txtlines[0], " ") {
		params[i] = num
	}
	books, _ := strconv.Atoi(params[0])
	numlibs, _ := strconv.Atoi(params[0])
	days, _ := strconv.Atoi(params[0])

	//the line which contains the book scores
	bookList := make([]Book, books)

	for i, num := range strings.Split(txtlines[1], " ") {
		x, _ := strconv.Atoi(num)
		bookList[i] = Book{bookID: i, score: x}
	}
	fmt.Println(bookList)

	librarys := make([]Library, numlibs)

	for i := 2; i < len(txtlines); i = i + 2 {
		params := strings.Split(txtlines[i], " ")
		numbooks, _ := strconv.Atoi(params[0])
		signtime, _ := strconv.Atoi(params[1])
		bookspday, _ := strconv.Atoi(params[2])
		booklist := make([]Book, numbooks)

		for k, num := range strings.Split(txtlines[i+1], " ") {
			res, _ := strconv.Atoi(num)
			booklist[k] = bookList[res]
		}
		librarys[i] = Library{libraryID: i, numBooks: numbooks, signupTime: signtime, booksPerDay: bookspday, books: booklist}
	}

	dataset := Dataset{books: books, numLibraries: numlibs, days: days, libraries: librarys}
	fmt.Println(dataset)
	return dataset
}
