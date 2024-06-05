package mapreduce

import (
	"fmt"
	"sync"
)

func (mr *Master) schedule(phase jobPhase) {
	
	var ntasks int 
	var nios int
	
	switch phase {
		case mapPhase:
			ntasks = len(mr.files) 
			nios = mr.nReduce      
		case reducePhase:
			ntasks = mr.nReduce    
			nios = len(mr.files)   
	}

	debug("Schedule: %v %v tasks (%d I/Os)\n", ntasks, phase, nios)

	// TODO
	var wg sync.WaitGroup  

	for i := 0; i < ntasks; i++ {
		wg.Add(1)
		go func(taskNum int, nios int, phase jobPhase) {
			debug("DEBUG: current taskNum: %v, nios: %v, phase: %v\n", taskNum, nios, phase)
			for  {
				worker := <-mr.registerChannel  
				debug("DEBUG: current worker port: %v\n", worker)

				var args DoTaskArgs
				args.JobName = mr.jobName
				args.File = mr.files[taskNum]
				args.Phase = phase
				args.TaskNumber = taskNum
				args.NumOtherPhase = nios
				ok := call(worker, "Worker.DoTask", &args, new(struct{}))
				if ok {
					wg.Done()
					mr.registerChannel <- worker
					break
				}  
			}
		}(i, nios, phase)
	}
	wg.Wait()
	fmt.Printf("Schedule: %v phase done\n", phase)
}