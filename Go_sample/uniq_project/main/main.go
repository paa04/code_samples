package main

import (
	. "awesomeProject/uniq_project/files"
	. "awesomeProject/uniq_project/parser"
	. "awesomeProject/uniq_project/uniq"
	"flag"
	"log"
)

func main() {
	var inputFile MyFileR
	inputFile.FileName = ""
	var outputFile MyFileW
	outputFile.FileName = ""
	uniqKey := 0
	var nunF int
	var numS int
	flagU, flagC, flagD, flagI := GetFlags(&nunF, &numS)
	if *flagU && *flagD || *flagC && *flagD || *flagC && *flagU {
		log.Fatal("Can't use 'u','d','c' flags together")
		return
	}

	uniqKey = DefineUnKey(flagU, flagD)

	files := flag.Args()
	GetFiles(files, &inputFile, &outputFile)
	arr := Scan(Reader(inputFile))
	ans := Uniq(arr, len(arr), numS, nunF, uniqKey, *flagI)
	Clear(outputFile)
	Write(outputFile, ans, *flagC)
}
