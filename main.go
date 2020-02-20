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
