/*=================================================================

Program name:
	selpg (SELect PaGes)

Purpose:
	Sometimes one needs to extract only a specified range of
pages from an input text file. This program allows the user to do
that.

Author: Yubei Xiao

===================================================================*/
/*================================= package ======================*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math"
	"os"
	"os/exec"
)

/*================================= import ======================*/

/*================================= types =========================*/

type selpgArgs struct {
	startPage  int
	endPage    int
	inFilename string
	pageLen    int
	pageType   bool
	printDest  string
}

/*================================= globals =======================*/

var progname string /* program name, for error messages */
const MAXUInt = ^uint(0)
const INT_MAX = math.MaxInt64

/*================================= main ====================*/
// func main must have no arguments and return values
// we use os.Args as command-line args and use flag to parse args
func main() {
	var sa selpgArgs

	/* save name by which program is invoked, for error messages */
	progname = os.Args[0]

	// initial args
	sa.startPage = -1
	sa.endPage = -1
	sa.pageLen = 72
	sa.pageType = false

	processArgs(len(os.Args), os.Args, &sa)
	processInput(sa)

	return
}

/*================================= process_args() ================*/

func processArgs(ac int, av []string, psa *selpgArgs) {
	/* arg # currently being processed */
	/* arg at index 0 is the command name itself (selpg),
	   first actual arg is at index 1,
	   last arg is at index (ac - 1) */

	/* check the command-line arguments for validity */
	if ac < 3 {
		fmt.Fprint(os.Stderr, "%s: %s\n", progname, "not enough arguments")
		usage()
		os.Exit(1)
	}

	/* handle mandatory args first */
	// var temp selpgArgs
	flag.IntVar(&psa.startPage, "s", -1, "startPage")
	flag.IntVar(&psa.endPage, "e", -1, "endPage")
	flag.IntVar(&psa.pageLen, "l", -1, "pageLen")
	flag.BoolVar(&psa.pageType, "f", false, "pageType")
	flag.StringVar(&psa.printDest, "d", "", "printDest.")
	flag.Parse()

	/* handle 1st arg - start page */
	if psa.startPage < 0 || psa.startPage > (INT_MAX-1) {
		fmt.Fprint(os.Stderr, "%s: %s\n", progname, "invalid start page")
		usage()
		os.Exit(2)
	}

	/* handle 2nd arg - end page */
	if psa.endPage < 0 || psa.endPage > (INT_MAX-1) || psa.endPage < psa.startPage {
		fmt.Fprint(os.Stderr, "%s: %s\n", progname, "invalid end page")
		usage()
		os.Exit(3)
	}

	/* now handle optional args */
	if psa.pageLen < 0 || psa.pageLen > (INT_MAX-1) {
		fmt.Fprint(os.Stderr, "%s: %s\n", progname, "invalid page length")
		usage()
		os.Exit(4)
	}

	if psa.pageType == true {
		if psa.pageLen != -1 {
			fmt.Fprint(os.Stderr, "%s: %s\n", progname, "option should be \"-f\"")
		}
	} else {
		if psa.pageLen < 1 {
			psa.pageLen = 72
		}
	}

	if flag.NArg() > 0 {
		psa.inFilename = flag.Arg(0)
	}
}

/*================================= process_input() ===============*/

func processInput(sa selpgArgs) {
	/* process the input source */
	if sa.inFilename != "" {
		// input from cmd
		var inputReader = bufio.NewReader(os.Stdin)
		//processOutput(inputReader, sa)
		if sa.pageType {
			readByPage(inputReader, sa)
		} else {
			readByLine(inputReader, sa)
		}
	} else {
		// input from file
		inputFile, err := os.Open(sa.inFilename)
		if err != nil {
			panic(err)
		}
		var inputReader = bufio.NewReader(inputFile)
		defer inputFile.Close()
		//processOutput(inputReader, sa)
		if sa.pageType {
			readByPage(inputReader, sa)
		} else {
			readByLine(inputReader, sa)
		}
	}
}

/*================================= process_output() ===============*/

// func processOutput(inputReader *bufio.Reader, sa selpgArgs) {
// 	/* process the output source and to different output type*/
// 	if sa.pageType {
// 		readByPage()
// 	} else {
// 		readByLine()
// 	}
// }
func readByPage(inputReader *bufio.Reader, myArgus selpgArgs) {
	// record pageCount
	pageCount := 1
	// read all pages
	for {
		page, err := inputReader.ReadString('\f')
		if err != nil {
			panic(err)
		}
		// when page number in the chosen range
		if pageCount >= myArgus.startPage && pageCount <= myArgus.endPage {
			// if output type is Stdout
			if myArgus.printDest == "" {
				fmt.Printf(page)
			} else {
				// open ./go input pipe, and output to pipe
				cmd := exec.Command("./out")
				echoInPipe, err := cmd.StdinPipe()
				if err != nil {
					panic(err)
				}
				echoInPipe.Write([]byte(page + "\n"))
				echoInPipe.Close()
				cmd.Stdout = os.Stdout
				cmd.Run()
			}
		}
		if err == io.EOF {
			break
		}
		pageCount++
	}

	// when startPage bigger than pageNumber, output null
	if myArgus.startPage > pageCount {
		fmt.Printf("Warning:\n\tSTARTPAGE(%d) is greater than number of total pages(%d).\noutput will be empty.\n", myArgus.startPage, pageCount)
	}
	// when endPage bigger than pageNumber
	if myArgus.endPage > pageCount {
		fmt.Printf("Warning:\n\tENDPAGE(%d) is greater than number of total pages(%d).\nthere will be less output than expected.\n", myArgus.endPage, pageCount)
	}
}

func readByLine(inputReader *bufio.Reader, myArgus selpgArgs) {
	lineCount := 1
	for {
		line, err := inputReader.ReadString('\n')
		if err != nil {
			panic(err)
		}

		if lineCount > myArgus.pageLen*(myArgus.startPage-1) && lineCount <= myArgus.pageLen*myArgus.endPage {
			if myArgus.printDest == "" {
				fmt.Printf(line)
			} else {
				cmd := exec.Command("./out")
				echoInPipe, err := cmd.StdinPipe()
				if err != nil {
					panic(err)
				}
				echoInPipe.Write([]byte(line))
				echoInPipe.Close()
				cmd.Stdout = os.Stdout
				cmd.Run()
			}
		}
		if err == io.EOF {
			break
		}
		lineCount++
	}
	if myArgus.startPage > lineCount/myArgus.pageLen+1 {
		fmt.Printf("Warning:\n\tSTARTPAGE(%d) is greater than number of total pages(%d).\noutput will be empty.\n", myArgus.startPage, lineCount/myArgus.pageLen+1)
	}
	if myArgus.endPage > lineCount/myArgus.pageLen+1 {
		fmt.Printf("Warning:\n\tENDPAGE(%d) is greater than number of total pages(%d).\nthere will be less output than expected.\n", myArgus.endPage, lineCount/myArgus.pageLen+1)
	}
}

/*================================= usage() =======================*/

func usage() {
	fmt.Fprint(os.Stderr, "%s: %s\n", progname, "\n[USAGE] -sstart_page -eend_page [ -f | -llines_per_page ] [ -ddest ] [ in_filename ]\n")
}

/*================================= EOF ===========================*/
