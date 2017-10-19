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
	"flag"
	"fmt"
	"math"
	"os"
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
	sa.pageType = 'l'

	processArgs(len(os.Args), os.Args, &sa)
	processInput(sa)

	return
}

func usage() {

}

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
}
func processInput(sa selpgArgs) {

}
