package main

import (
    "fmt"
	"os"
	"log"
	"bufio"
	"strconv"
)

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

func main() {
	

}
