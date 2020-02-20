package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readIn(){
	file, err := os.Open("a_example.txt")
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
	//data =: Dataset{books = }

}

func write(o Output) {
	f, _ := os.Create("out.txt")
	w := bufio.NewWriter(f)

	w.WriteString(strconv.Itoa(o.numLibraries) + "\n")

	for _, l := range o.libraryScans {
		w.WriteString(strconv.Itoa(l.libraryID) + " " + strconv.Itoa(l.numBooks) + "\n")
		for _,b:=l.books{
			w.WriteString(strconf.Itoa(b))
		}
		w.WriteString("\n")
	}
	w.Flush()
}

func readIn(){
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


	books, _ := strconv.Atoi(string(txtlines[0][0]))
	numlibs, _:= strconv.Atoi(string(txtlines[0][2]))
	days, _:= strconv.Atoi(string(txtlines[0][4]))

	scores := string(txtlines[1])
	bookList := make([]Book,books)
	id := 0
	s := ""
	for _,c := range scores{
		if unicode.IsSpace(c){
			theScore,_ := strconv.Atoi(s)
			newBook := Book{bookID:id,score:theScore}
			bookList[id] = newBook
			s = ""
			id++
		}else{
			s = s +string(c)
		}	
	}

	librarys := make([]Library,numlibs)
	for i:=0; i < numlibs; i++{
		numbooks,_:= strconv.Atoi(string(txtlines[2*i+2][0]))
		signtime,_ := strconv.Atoi(string(txtlines[2*i+2][2]))
		bookspday,_:= strconv.Atoi(string(txtlines[2*i+2][4]))
		booklist:= make([]int,numbooks)
		id := 0
		s := ""
		for _, c := range string(txtlines[2*i+3]){
			if unicode.IsSpace(c){
				theScore,_:=strconv.Atoi(s)
				booklist[id] = theScore
				s = ""
				id++
			}else{
				s = s +string(c)
			}
		}
		librarys[i] = Library{libraryID:i,numBooks:numbooks,signupTime:signtime,booksPerDay:bookspday,books:booklist}
	}
 
	dataset := Dataset{books:books,numLibraries:numlibs,day:days,libraries:librarys}
	fmt.Println(dataset)
	
}
