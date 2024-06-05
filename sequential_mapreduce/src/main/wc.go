package main

import (
	"fmt"
	//"sequential_mapreduce/src/mapreduce"
	"mapreduce"
	"os"
 
    "strconv"
	"strings"
	"unicode"

	//"time"
)


func mapF(document string, value string) (res []mapreduce.KeyValue) {
	// TODO: you have to write this function
 
	/*
	words := strings.FieldsFunc(value, func(c rune) bool {
		return !unicode.IsLetter(c)
	})
	for _, word := range words {
		res = append(res, mapreduce.KeyValue{Key: word, Value: "1"})
	}
	return res
	*/

	results := strings.FieldsFunc(value, func(ch rune) bool{return !unicode.IsLetter(ch)})
	for _,v := range results{
		res = append(res, mapreduce.KeyValue{v, "1"})
	}
	return
}

func reduceF(key string, values []string) string {
	// TODO: you also have to write this function

	/*
	count := 0
	for _, value := range values {
		c, err := strconv.Atoi(value)
		if err == nil {
			count += c
		}
	}
	return strconv.Itoa(count)
	*/

	var res int
	res = len(values) 
	return strconv.Itoa(res)
}

// Can be run in 3 ways:
// 1) Sequential (e.g., go run wc.go master sequential x1.txt .. xN.txt)
// 2) Master (e.g., go run wc.go master localhost:7777 x1.txt .. xN.txt)
// 3) Worker (e.g., go run wc.go worker localhost:7777 localhost:7778 &)
func main() {
	
	//start := time.Now()

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

	//end := time.Now()
	//elapsed := end.Sub(start)
	//fmt.Printf("Execution time: %s\n", elapsed)
}
