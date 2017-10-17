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
  var myArgus argus

  setArgus(&myArgus)
  parseArgus(&myArgus)
  fileProcess(&myArgus)
}

// input and set argus
func setArgus(myArgus *argus)  {
  myArgus.programName = os.Args[0] // os
  flag.Usage = func() { // flag
		fmt.Printf("Usage: %s -s=STARTPAGE -e=ENDPAGE [OPTION]... [FILE]...\n", myArgus.programName)
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

// parse argus
func parseArgus(myArgus * argus) {
  var programName = myArgus.programName
  // startPage and endPage required
  if myArgus.startPage < 0 || myArgus.endPage < 0 {
		processError(proName, fmt.Sprintf("not enough arguments, -s and -e is indispensable."))
	}
  // startPage must be smaller than endPage
  if myArgus.startPage > myArgus.endPage {
    processError(proName, fmt.Sprintf("STARTPAGE(%d) should not be greater than ENDPAGE(%d).", myArgus.startPage, myArgus.endPage))
  }

  myArgus.srcFile = flag.Arg(0)

  var defaultLength int = 72
  if myArgus.pageType == true {
    // ensure exclusion between '-f' and '-l'
    if myArgus.pageLen != -1 {
      processError(proName, "-f and -l=PAGELENGTH are mutually-exclusive.")
    }
  } else {
    // assign a default value
    if myArgus.pageLen < 1 {
      myArgus.pageLen = defaultLength
    }
  }
}
