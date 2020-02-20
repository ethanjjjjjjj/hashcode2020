package main

import "fmt"

type Dataset struct {
	books        int
	days         int
	numLibraries int
	libraries    []Library
}

type Library struct {
	libraryID   int
	numBooks    int
	signupTime  int
	booksPerDay int
	books       []Book
	sumScore    int
	sortedBooks []Book
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

func radixSortBooksNoGo(books []Book, i int) []Book {
	if i < 0 {
		return books
	}
	var ones = make([]Book, 0)
	var zeros = make([]Book, 0)
	for _, n := range books {
		if (n.bookID>>i)&0x1 == 0x1 {
			ones = append(ones, n)
		} else {
			zeros = append(zeros, n)
		}
	}
	if i%8 == 0 {
		c1 := make(chan []Book)
		c2 := make(chan []Book)
		go radixSortBooks(ones, i-1, c1)
		go radixSortBooks(zeros, i-1, c2)
		return append(<-c1, (<-c2)...)
	} else {
		ones = radixSortBooksNoGo(ones, i-1)
		zeros = radixSortBooksNoGo(zeros, i-1)
		return append(ones, zeros...)
	}
}

func radixSortBooks(books []Book, i int, out chan []Book) {
	if i < 0 {
		out <- books
		return
	}
	var ones = make([]Book, 0)
	var zeros = make([]Book, 0)
	for _, n := range books {
		if (n.bookID>>i)&0x1 == 0x1 {
			ones = append(ones, n)
		} else {
			zeros = append(zeros, n)
		}
	}
	if i%8 == 0 {
		c1 := make(chan []Book)
		c2 := make(chan []Book)
		go radixSortBooks(ones, i-1, c1)
		go radixSortBooks(zeros, i-1, c2)
		out <- append(<-c1, (<-c2)...)
	} else {
		ones = radixSortBooksNoGo(ones, i-1)
		zeros = radixSortBooksNoGo(zeros, i-1)
		out <- append(ones, zeros...)
	}
}

func main() {
	c := make(chan []Book)
	go radixSortBooks([]Book{Book{4, 4}, Book{6, 6}, Book{2, 2}, Book{9, 9}, Book{3, 3}}, 5, c)
	fmt.Printf("%v", <-c)
}

func formOutput() Output {
	data := readIn()

	scannedBooks := make([]Book, 0)

	ls := data.libraries
	channels := make([]chan []Book, len(ls))
	for i, l := range ls {
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

}
