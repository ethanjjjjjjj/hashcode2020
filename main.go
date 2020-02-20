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
	writeOutput(formOutput())
}

func formOutput() Output {
	data := readIn()
	ls := data.libraries
	//scannedBooks := make([]Book, 0)
	//signupOrder := make([]Library, len(ls))
	channels := make([]chan []Book, len(ls))
	for _, l := range ls {
		c := make(chan []Book)
		go radixSortBooks(l.books, 10, c)
	}
	for i, l := range ls {
		l.sortedBooks = <-channels[i]
	}
	for _, l := range ls {
		for i := 0; i < 5; i++ {
			l.sumScore += l.sortedBooks[i].score
		}
	}
	sortedLibraries := radixSortLibrariesNoGo(ls, 10)
	currentday := 0
	out := Output{numLibraries: 0, libraryScans: make([]libraryScan, 0)}
	for _, l := range sortedLibraries {
		currentday += l.signupTime
		if currentday > data.days {
			break
		}
		l.timeSignedUp = currentday
		perday := l.booksPerDay
		numBooksToAdd := (data.days - currentday) * perday
		booksToAdd := l.sortedBooks[:numBooksToAdd]
		bookstoaddids := make([]int, 0)
		for _, b := range booksToAdd {
			bookstoaddids = append(bookstoaddids, b.bookID)
		}
		scanoutput := libraryScan{libraryID: l.libraryID, books: bookstoaddids}
		out.libraryScans = append(out.libraryScans, scanoutput)
	}
	return out
}
