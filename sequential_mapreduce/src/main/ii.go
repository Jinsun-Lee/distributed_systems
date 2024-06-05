package main

import (
	"fmt"
	"mapreduce"
	"os"

	"sort"
	"strconv"
	"strings"
	"unicode"
)

type StringList []string

func (s StringList) Len() int {
    return len(s)
}

func (s StringList) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s StringList) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func mapF(document string, value string) (res []mapreduce.KeyValue) {
	// TODO: you should complete this to do the inverted index challenge
	 
	splitFunc := func(r rune) bool {
        return !unicode.IsLetter(r)
    }
    words := strings.FieldsFunc(value, splitFunc)
    for _, word := range words{
        res = append(res, mapreduce.KeyValue{word, document})
	}
	return res
}


func reduceF(key string, values []string) string {
	// TODO: you should complete this to do the inverted index challenge

	tmp := make(map[string]int)
	for _, val := range values{
		tmp[val] = 1
	}

	var newValues [] string

	for v, _ := range tmp{
		newValues = append(newValues, v)
	}

	length := len(newValues)
	res := strconv.Itoa(length)
	res += " "

	sort.Sort(StringList(newValues))

	for i := 0; i < length-1; i++ {
		res += newValues[i] + ","
	}

	res += newValues[length-1]

	return res
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
			mr = mapreduce.Sequential("iiseq", os.Args[3:], 3, mapF, reduceF)
		} else {
			mr = mapreduce.Distributed("iiseq", os.Args[3:], 3, os.Args[2])
		}
		mr.Wait()
	} else {
		mapreduce.RunWorker(os.Args[2], os.Args[3], mapF, reduceF, 100)
	}

}

