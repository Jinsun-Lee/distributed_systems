package mapreduce

import (
	"encoding/json"
	"os"
	"sort"
)

func doReduce(
	jobName string, 
	reduceTaskNumber int, 
	nMap int, 
	reduceF func(key string, values []string) string,
) {
	// TODO:

	KeyValues := make(map[string][]string, 0)

	for i := 0; i < nMap; i++ {
		filename := reduceName(jobName, i, reduceTaskNumber) 
		
		file, err := os.Open(filename)
		if err != nil {
			debug("doReduce: open file error %s, error %s", filename, err)
		}
		defer file.Close()
		dec := json.NewDecoder(file)
		for {
			var kv KeyValue
			err := dec.Decode(&kv)
			if err != nil {
				break 
			}

			_, ok := KeyValues[kv.Key] 
			if !ok {                
				KeyValues[kv.Key] = make([]string, 0)
			}
			KeyValues[kv.Key] = append(KeyValues[kv.Key], kv.Value)
		}
	}

	var keys []string
	for k, _ := range KeyValues {
		keys = append(keys, k)
	}
	sort.Strings(keys) 

	mergeFileName := mergeName(jobName, reduceTaskNumber) 
	mergeFile, err := os.Create(mergeFileName)
	if err != nil {
		debug("doReduce: create merge file error %s, %s", mergeFileName, err)
	}
	defer mergeFile.Close()

	enc := json.NewEncoder(mergeFile) 

	for _, k := range keys { 
		res := reduceF(k, KeyValues[k])
		err := enc.Encode(&KeyValue{k, res})
		if err != nil {
			debug("doReduce: encode error")
		}
	}
}