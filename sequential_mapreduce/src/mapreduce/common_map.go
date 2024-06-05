package mapreduce

import (
	"encoding/json"
	"hash/fnv"
	"os"
)


func doMap(
	jobName string, 
	mapTaskNumber int, 
	inFile string,
	nReduce int, 
	mapF func(file string, contents string) []KeyValue,
) {
	// TODO:

	inputfile, err := os.Open(inFile)
	if err != nil {
		debug("doMap: open file error %s, %s", inFile, err)
	}
	defer inputfile.Close()


	fileInfo, err := inputfile.Stat() 
	if err != nil {
		debug("doMap: get state fail %s, %s", inFile, err)
	}

	Content := make([]byte, fileInfo.Size()) 
	ReadBytes, err := inputfile.Read(Content)
	if err != nil {
		debug("doMap: Read file error %s, %s", inFile, err)
	} else if int64(ReadBytes) != fileInfo.Size() {
		debug("doMap: Read file error, don`t have enough bytes %s, %s", inFile, err)
	}

	keyValues := mapF(inFile, string(Content)) 

	for i := 0; i < nReduce; i++ {
		filename := reduceName(jobName, mapTaskNumber, i) 
		reduceFile, err := os.Create(filename)           
		if err != nil {
			debug("doMap: create intermediate file %s, %s", filename, err)
		}
		defer reduceFile.Close() 
		enc := json.NewEncoder(reduceFile) 
		for _, kv := range keyValues {
			if ihash(kv.Key)%uint32(nReduce) == uint32(i) {  
				err := enc.Encode(&kv)
				if err != nil {
					debug("doMap: encode error: %s", err)
				}
			}
		}
	}
}

func ihash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}