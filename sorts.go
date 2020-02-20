package main

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

func radixSortLibrariesNoGo(libraries []Library, i int) []Library {
	if i < 0 {
		return libraries
	}
	var ones = make([]Library, 0)
	var zeros = make([]Library, 0)
	for _, n := range libraries {
		if (n.sumScore>>i)&0x1 == 0x1 {
			ones = append(ones, n)
		} else {
			zeros = append(zeros, n)
		}
	}
	if i%8 == 0 {
		c1 := make(chan []Library)
		c2 := make(chan []Library)
		go radixSortLibraries(ones, i-1, c1)
		go radixSortLibraries(zeros, i-1, c2)
		return append(<-c1, (<-c2)...)
	} else {
		ones = radixSortLibrariesNoGo(ones, i-1)
		zeros = radixSortLibrariesNoGo(zeros, i-1)
		return append(ones, zeros...)
	}
}

func radixSortLibraries(libraries []Library, i int, out chan []Library) {
	if i < 0 {
		out <- libraries
		return
	}
	var ones = make([]Library, 0)
	var zeros = make([]Library, 0)
	for _, n := range libraries {
		if (n.sumScore>>i)&0x1 == 0x1 {
			ones = append(ones, n)
		} else {
			zeros = append(zeros, n)
		}
	}
	if i%8 == 0 {
		c1 := make(chan []Library)
		c2 := make(chan []Library)
		go radixSortLibraries(ones, i-1, c1)
		go radixSortLibraries(zeros, i-1, c2)
		out <- append(<-c1, (<-c2)...)
	} else {
		ones = radixSortLibrariesNoGo(ones, i-1)
		zeros = radixSortLibrariesNoGo(zeros, i-1)
		out <- append(ones, zeros...)
	}
}
