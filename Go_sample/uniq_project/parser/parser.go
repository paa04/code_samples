package parser

import (
	. "awesomeProject/uniq_project/files"
	"flag"
	"log"
)

func GetFlags(numF *int, numS *int) (*bool, *bool, *bool, *bool) {
	flagU := flag.Bool("u", false, "only uniq strings")
	flagC := flag.Bool("c", false, "count strings")
	flagD := flag.Bool("d", false, "only not uniq strings")
	flagI := flag.Bool("i", false, "ignore reg")
	flag.IntVar(numF, "f", 0, "ignore words")
	flag.IntVar(numS, "s", 0, "ignore chars")
	flag.Parse()
	return flagU, flagC, flagD, flagI
}
func DefineUnKey(flagU *bool, flagD *bool) int {
	if *flagU {
		return 1
	} else if *flagD {
		return -1
	}
	return 0
}

func GetFiles(files []string, inputFile *MyFileR, outputFile *MyFileW) {
	if len(files) >= 3 {
		log.Fatal("Too many files")
		return
	} else if len(files) == 1 || len(files) == 2 {
		inputFile.FileName = files[0]
		if len(files) == 2 {
			outputFile.FileName = files[1]
		}
	}
}
