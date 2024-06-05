package main

import (
	"fmt"
	"mapreduce"
	"os"
 
    "strconv"
	"strings"
	"unicode"
)

func mapF(document string, value string) (res []mapreduce.KeyValue) {
	// TODO: you have to write this function

	results := strings.FieldsFunc(value, func(ch rune) bool{return !unicode.IsLetter(ch)})
	for _,v := range results{
		res = append(res, mapreduce.KeyValue{v, "1"})
	}
	return
}

func reduceF(key string, values []string) string {
	// TODO: you also have to write this function

	var res int
	res = len(values) 
	return strconv.Itoa(res)
}

// Can be run in 3 ways:
// 1) Sequential (e.g., go run wc.go master sequential x1.txt .. xN.txt)
// 2) Master (e.g., go run wc.go master localhost:7777 x1.txt .. xN.txt)
// 3) Worker (e.g., go run wc.go worker localhost:7777 localhost:7778 &)
func main() {
	
	if len(os.Args) < 4 {
		fmt.Printf("%s: see usage comments in file\n", os.Args[0])
	} else if os.Args[1] == "master" {
		var mr *mapreduce.Master
		if os.Args[2] == "sequential" {
			mr = mapreduce.Sequential("wcseq", os.Args[3:], 3, mapF, reduceF)
		} else {
			mr = mapreduce.Distributed("wcseq", os.Args[3:], 3, os.Args[2])
		}
		mr.Wait()
	} else {
		mapreduce.RunWorker(os.Args[2], os.Args[3], mapF, reduceF, 100)
	}

}
