package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func writeOutput(o Output, name string) {
	f, _ := os.Create(name + "out.txt")
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

func gettext(filename string) []string {
	data, _ := ioutil.ReadFile(filename)
	txtlines := strings.Split(string(data), "\n")
	/*for _, s := range txtlines {
		fmt.Println(s)
	}*/
	return txtlines
}

func parsetext(txtlines []string) Dataset {
	params := strings.Split(txtlines[0], " ")
	books, _ := strconv.Atoi(params[0])
	fmt.Println(books)
	numlibs, _ := strconv.Atoi(params[1])
	fmt.Println(numlibs)
	days, _ := strconv.Atoi(params[2])

	scores := strings.Split(txtlines[1], " ")
	scoresInts := make([]int, books)
	for i, s := range scores {
		scoresInts[i], _ = strconv.Atoi(s)
	}
	libarray := make([]Library, 0)
	librarynum := 0

	for i := 2; i < len(txtlines)-1; i += 2 {
		librarystats := strings.Split(txtlines[i], " ")
		if len(librarystats) < 3 {
			break
		}
		numBooks, _ := strconv.Atoi(librarystats[0])
		//fmt.Println(numBooks)

		signuptime, _ := strconv.Atoi(librarystats[1])

		booksperday, _ := strconv.Atoi(librarystats[2])

		librarybooks := strings.Split(txtlines[i+1], " ")
		bookslist := make([]Book, 0)
		for _, b := range librarybooks {
			booknum, _ := strconv.Atoi(b)
			bookscore := scoresInts[booknum]
			currentbook := Book{bookID: booknum, score: bookscore}
			bookslist = append(bookslist, currentbook)
		}
		newLib := Library{libraryID: librarynum, numBooks: numBooks, signupTime: signuptime, booksPerDay: booksperday, books: bookslist}
		librarynum++

		libarray = append(libarray, newLib)
	}
	data := Dataset{books: books, days: days, numLibraries: numlibs, libraries: libarray}
	return data
}
func readIn() Dataset {
	/*file, err := os.Open("b_read_on.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()*/

	//first values
	data, _ := ioutil.ReadFile("a_example.txt")
	txtlines := strings.Split(string(data), "\n")
	params := make([]string, 3)
	for i, num := range strings.Split(txtlines[0], " ") {
		params[i] = num
	}
	books, _ := strconv.Atoi(params[0])
	numlibs, _ := strconv.Atoi(params[1])
	days, _ := strconv.Atoi(params[2])

	//the line which contains the book scores
	bookList := make([]Book, books)
	//fmt.Println(txtlines)
	for i, num := range strings.Split(txtlines[1], " ") {
		x, _ := strconv.Atoi(num)
		bookList[i] = Book{bookID: i, score: x}
	}
	//fmt.Println(bookList)

	librarys := make([]Library, numlibs)

	for i := 2; i < len(txtlines)-2; i = i + 2 {
		params := strings.Split(txtlines[i], " ")
		numbooks, _ := strconv.Atoi(params[0])
		signtime, _ := strconv.Atoi(params[1])
		bookspday, _ := strconv.Atoi(params[2])
		booklist := make([]Book, numbooks)

		for k, num := range strings.Split(txtlines[i+1], " ") {
			res, _ := strconv.Atoi(num)
			booklist[k] = bookList[res]
		}
		librarys[(i-2)/2] = Library{libraryID: (i - 2) / 2, numBooks: numbooks, signupTime: signtime, booksPerDay: bookspday, books: booklist}
	}

	dataset := Dataset{books: books, numLibraries: numlibs, days: days, libraries: librarys}
	//fmt.Println(dataset)
	return dataset
}
