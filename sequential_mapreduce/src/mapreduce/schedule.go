package mapreduce

/*
cd /home/s23710660/sequential_mapreduce/src/mapreduce
go test -run Failure mapreduce/...
*/

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

// Failure 실행 안 됨	
/*
	// https://github.com/webglider/mapreduce/blob/master/schedule.go
	idleTasks :=  make(chan DoTaskArgs)

	taskSuccess := make(chan bool)

	idleWorkers := mr.registerChannel

	quitDispatch := make(chan bool)

	go dispatchTasks(idleTasks, idleWorkers, quitDispatch, taskSuccess)
	
	for i := 0; i < ntasks; i++ {
		var taskFile string
		if phase == mapPhase {
			taskFile = mr.files[i]
		}
		taskArgs := DoTaskArgs {
			JobName: mr.jobName,
			File: taskFile,
			Phase: phase,
			TaskNumber: i,
			NumOtherPhase: nios,
		}
		idleTasks <- taskArgs
	}

	for i := 0; i < ntasks; i++ {
		<- taskSuccess
	}

	quitDispatch <- true

	fmt.Printf("Schedule: %v phase done\n", phase)
}

func dispatchTasks(idleTasks chan DoTaskArgs, idleWorkers chan string, 
	quitDispatch chan bool, taskSuccess chan bool) {

	taskQueue := make([]DoTaskArgs, 0)
	workerQueue := make([]string, 0)
	for {
		select {
		case task := <- idleTasks:
			taskQueue = append(taskQueue, task)
		case worker := <- idleWorkers:
			workerQueue = append(workerQueue, worker)
		case <- quitDispatch:
			go func() {
				for i := 0; i < len(workerQueue); i++ {
					idleWorkers <- workerQueue[i]
				}
			}()
			return
		}

		for len(taskQueue) > 0 && len(workerQueue) > 0 {
			go doTask(taskQueue[0], workerQueue[0], idleTasks, idleWorkers, 
				taskSuccess)
			taskQueue = taskQueue[1:]
			workerQueue = workerQueue[1:]
		}
	}
}


func doTask(task DoTaskArgs, worker string, idleTasks chan DoTaskArgs, 
	idleWorkers chan string, taskSuccess chan bool) {

	ok := call(worker, "Worker.DoTask", &task, new(struct{}))

	idleWorkers <- worker
	if ! ok {
		idleTasks <- task
	} else {
		taskSuccess <- true
	}
}
*/


// Failure 실행 안 됨
/*
	// https://blog.csdn.net/freedom1523646952/article/details/108355990
	var wg sync.WaitGroup  
	wg.Add(ntasks)

	for i := 0; i < ntasks; i++ {
		var args DoTaskArgs
		if phase == mapPhase {
			args.JobName = mr.jobName
			args.File = mr.files[i] // File: mapFiles[i]
			args.Phase = phase
			args.TaskNumber = i     // i
			args.NumOtherPhase = nios
		} else {
			args.JobName = mr.jobName
			args.File = ""
			args.Phase = phase
			args.TaskNumber = i     // i
			args.NumOtherPhase = nios

		}

		//go func(taskNum int, nios int, phase jobPhase) {
		go func(args DoTaskArgs, registerChannel chan string){
			res := false
			var worker string
			for res == false {
				worker := <-mr.registerChannel
				// ok := call(worker, "Worker.DoTask", &args, ))
				res = call(worker, "Worker.DoTask", &args, new(struct{}) )
			}			
			go func() {
				mr.registerChannel <- worker
			}()
			wg.Done()
		}(args, mr.registerChannel)
	}
	wg.Wait()
	fmt.Printf("Schedule: %v phase done\n", phase)
}
*/




// Failure 실행 안 됨
///*
	// https://github.com/goldknife6/MIT-6.824-2016/blob/master/src/mapreduce/schedule.go
	var wg sync.WaitGroup  //

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
//*/




// Failure - closed channel
/*
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
*/





