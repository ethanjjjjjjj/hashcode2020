package main

type Dataset struct {
	books        int
	day          int
	numLibraries int
	libraries    []Library
}

type Library struct {
	libraryID   int
	numBooks    int
	signupTime  int
	booksPerDay int
	books       []int
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
	//var bookScores []int
	//Dataset data=readIn()
	out := Output{numLibraries: 1, libraryScans: [...]libraryScan{libraryScan{}}}

	write()

}
