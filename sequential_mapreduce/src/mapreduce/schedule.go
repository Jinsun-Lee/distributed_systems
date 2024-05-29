package mapreduce

import (
	"fmt"
	"sync"
)

// schedule starts and waits for all tasks in the given phase (Map or Reduce).
func (mr *Master) schedule(phase jobPhase) {
	var ntasks int
	var nios int // number of inputs (for reduce) or outputs (for map)
	switch phase {
	case mapPhase:
		ntasks = len(mr.files)
		nios = mr.nReduce
	case reducePhase:
		ntasks = mr.nReduce
		nios = len(mr.files)
	}

	debug("Schedule: %v %v tasks (%d I/Os)\n", ntasks, phase, nios)

	// All ntasks tasks have to be scheduled on workers, and only once all of
	// them have been completed successfully should the function return.
	// Remember that workers may fail, and that any given worker may finish
	// multiple tasks.
	//
	// TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO
	//
	// create new go rutine -> assign each task to worker
	// each task have DoTaskArgs(TaskNumber, File)
    // call worker.DoTask()


	// Use a WaitGroup to wait for all tasks to complete
	var wg sync.WaitGroup
	wg.Add(ntasks)

    for i := 0; i < ntasks; i++ {
        wg.Add(1)
        go func(taskNumber int) {
            defer wg.Done()

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
            }
            
			// task fin > add worker to channel
            go func() {
                mr.registerChannel <- worker
            }()

        }(i)
    }

    // wait all task done
    wg.Wait()


	// END
	debug("Schedule: %v phase done\n", phase)
}
