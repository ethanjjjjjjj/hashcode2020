package main

import "fmt"

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

func radixSortBooksNoGo(books []int, i int) []int {
	if i < 0 {
		return books
	}
	var ones = make([]int, 0)
	var zeros = make([]int, 0)
	for _, n := range books {
		if (n>>i)&0x1 == 0x1 {
			ones = append(ones, n)
		} else {
			zeros = append(zeros, n)
		}
	}
	if i%8 == 0 {
		c1 := make(chan []int)
		c2 := make(chan []int)
		go radixSortBooks(ones, i-1, c1)
		go radixSortBooks(zeros, i-1, c2)
		return append(<-c1, (<-c2)...)
	} else {
		ones = radixSortBooksNoGo(ones, i-1)
		zeros = radixSortBooksNoGo(zeros, i-1)
		return append(ones, zeros...)
	}
}

func radixSortBooks(books []int, i int, out chan []int) {
	if i < 0 {
		out <- books
		return
	}
	var ones = make([]int, 0)
	var zeros = make([]int, 0)
	for _, n := range books {
		if (n>>i)&0x1 == 0x1 {
			ones = append(ones, n)
		} else {
			zeros = append(zeros, n)
		}
	}
	if i%8 == 0 {
		c1 := make(chan []int)
		c2 := make(chan []int)
		go radixSortBooks(ones, i-1, c1)
		go radixSortBooks(zeros, i-1, c2)
		out <- append(<-c1, (<-c2)...)
	} else {
		ones = radixSortBooksNoGo(ones, i-1)
		zeros = radixSortBooksNoGo(zeros, i-1)
		out <- append(ones, zeros...)
	}
}

func libary(l Library, scoreBack chan int) {

}

func main() {
	c := make(chan []int)
	go radixSortBooks([]int{4, 6, 2, 9, 3}, 5, c)
	fmt.Printf("%v", <-c)
}
