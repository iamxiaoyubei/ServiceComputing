package main

import (
  "bufio"
  "flag"
  "fmt"
  "io"
  "os"
  "os/exec"
)

type argus struct {
  programName string
  startPage int
  endPage int
  srcFile string
  pageLen int
  pageType bool // true for -f, false for -l
  desPrint string
}

func main()  {
  var myArgus argus;

  setArgus(&myArgus);
  parseArgus(&myArgus);
  fileProcess(&myArgs);
}

func setArgus(myArgus *argus)  {
  myArgus.programName = os.Args[0]; // os
  flag.Usage = func() { // flag
		fmt.Printf("Usage: %s -s=STARTPAGE -e=ENDPAGE [OPTION]... [FILE]...\n", myArgs.programName)
		fmt.Printf("Select specified pages from file or standard input.\n\n")
		fmt.Printf("With no FILE, or when FILE is -, read standard input.\n\n")
		fmt.Printf("\t-s=STARTPAGE\tPages number starts at STARTPAGE\n")
		fmt.Printf("\t-e=ENDPAGE\tPages number ends at ENDPAGE\n")
		fmt.Printf("\t-l=PAGELENGTH\tThe number of lines of each page\n")
		fmt.Printf("\t-f\tInput file use 'f' to seperate two pages\n")
		fmt.Printf("\t-d\tThe destination of output\n\n")
	}
  flag.IntVar(&myArgus.startPage, "s", -1, "specify start page.")
	flag.IntVar(&myArgus.endPage, "e", -1, "specify end page.")
	flag.IntVar(&myArgus.pageLen, "l", -1, "specify page length(number of lines).")
	flag.BoolVar(&myArgus.pageType, "f", false, "specify type of input file.")
	flag.StringVar(&myArgus.desPrint, "d", "", "specify the destination program.")

	flag.Parse()

}
