package main

type Dataset struct {
	books        int
	days         int
	numLibraries int
	libraries    []Library
}

type Library struct {
	libraryID    int
	numBooks     int
	signupTime   int
	booksPerDay  int
	books        []Book
	sumScore     int
	sortedBooks  []Book
	timeSignedUp int
}

type Book struct {
	bookID int
	score  int
}

type libraryScan struct {
	libraryID int
	numBooks  int
	books     []int
}
type Output struct {
	numLibraries int
	libraryScans []libraryScan
}

func main() {
	data := parsetext(gettext("a_example.txt"))
	writeOutput(formOutput(data), "A")
	data = parsetext(gettext("b_read_on.txt"))
	writeOutput(formOutput(data), "B")
	data = parsetext(gettext("c_incunabula.txt"))
	writeOutput(formOutput(data), "C")
	data = parsetext(gettext("d_tough_choices.txt"))
	writeOutput(formOutput(data), "D")
	data = parsetext(gettext("e_so_many_books.txt"))
	writeOutput(formOutput(data), "E")
	data = parsetext(gettext("f_libraries_of_the_world.txt"))
	writeOutput(formOutput(data), "F")

}

func formOutput(data Dataset) Output {
	//data := readIn()
	ls := data.libraries
	//scannedBooks := make([]Book, 0)
	//signupOrder := make([]Library, len(ls))
	channels := make([]chan []Book, len(ls))
	for i, l := range ls {
		c := make(chan []Book)
		channels[i] = c
		go radixSortBooks(l.books, 14, c)
	}
	for i := 0; i < len(ls); i++ {
		booksSorted := <-channels[i]
		ls[i].sortedBooks = append(ls[i].sortedBooks, booksSorted...)
		//fmt.Println(ls[i])
	}
	for _, l := range ls {

		//fmt.Println(l)
		constraint := 50

		if l.numBooks < constraint {
			constraint = l.numBooks
			for i := 0; i < constraint; i++ {
				l.sumScore += l.sortedBooks[i].score
			}
		} else if l.numBooks == 0 {
			l.sumScore = 0
		} else {
			for i := 0; i < constraint; i++ {
				l.sumScore += l.sortedBooks[i].score
			}
		}

	}
	sortedLibraries := radixSortLibrariesNoGo(ls, 14)
	currentday := 0
	out := Output{numLibraries: 0, libraryScans: make([]libraryScan, 0)}
	for _, l := range sortedLibraries {
		currentday += l.signupTime
		if currentday > data.days {
			break
		}
		out.numLibraries++
		l.timeSignedUp = currentday
		perday := l.booksPerDay
		numBooksToAdd := (data.days - currentday) * perday
		var booksToAdd []Book
		if numBooksToAdd > len(l.sortedBooks) {
			booksToAdd = append(booksToAdd, l.sortedBooks[:len(l.sortedBooks)]...)
		} else {
			booksToAdd = append(booksToAdd, l.sortedBooks[:numBooksToAdd]...)

		}
		l.numBooks = len(booksToAdd)
		bookstoaddids := make([]int, 0)
		for _, b := range booksToAdd {
			bookstoaddids = append(bookstoaddids, b.bookID)
		}
		scanoutput := libraryScan{libraryID: l.libraryID, books: bookstoaddids, numBooks: len(booksToAdd)}
		out.libraryScans = append(out.libraryScans, scanoutput)
	}
	return out
}
