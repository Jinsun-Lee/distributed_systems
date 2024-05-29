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

	// TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO
	
	var wg sync.WaitGroup 

	taskChannel := make(chan int, ntasks) 
    for i := 0; i < ntasks; i++ { 
        taskChannel <- i
    }
    close(taskChannel)


	for i := 0; i < ntasks; i++ {
        
		wg.Add(1)

        go func() {
            defer wg.Done()  
            for taskNumber := range taskChannel {   
                worker := <-mr.registerChannel  

                args := DoTaskArgs{
                    JobName:       mr.jobName,
                    Phase:         phase,
                    TaskNumber:    taskNumber,
                    NumOtherPhase: nios,
                    File:          mr.files[taskNumber],
                }

                ok := call(worker, "Worker.DoTask", &args, new(struct{}))
                if !ok {
                    fmt.Printf("Schedule: RPC %s DoTask error\n", worker)
                    taskChannel <- taskNumber  
                }

                go func() {
                    mr.registerChannel <- worker  
                }()
            }
        }()
    }

    wg.Wait()
    debug("Schedule: %v phase done\n", phase)
}