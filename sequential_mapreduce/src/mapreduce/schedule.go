package mapreduce

import (
	"fmt"
	"sync"
)

// schedule starts and waits for all tasks in the given phase (Map or Reduce).
// Master 구조체의 메서드로 정의 - jobPhase(작업의 종류 - mapPhase, reducePhase) 타입의 phase 인자를 받음
func (mr *Master) schedule(phase jobPhase) {
	
	var ntasks int // 실행할 작업의 수를 저장
	var nios int   // I/O를 저장할 변수로 reduce면 입력 파일의 수, map이면 출력 파일의 수
	
	switch phase {
		case mapPhase:
			ntasks = len(mr.files) // 맵 작업의 수 = 입력 파일의 수
			nios = mr.nReduce      // 각 맵 작업이 생성할 리듀스 작업의 수
		case reducePhase:
			ntasks = mr.nReduce    // 리듀스 작업의 수 
			nios = len(mr.files)   // 리듀스 작업의 입력이 되는 중간 파일의 수
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
	var wg sync.WaitGroup                 // 고루틴이 모두 종료될 때까지 기다리려고 사용 

	/*
	ntasks만큼의 고루틴을 생성하지 않고,  
	taskChannel 채널의 버퍼 크기를 ntasks로 설정한 뒤 고루틴을 실행 = 한 번에 최대 ntasks개의 고루틴이 실행될 수 있음
	각 고루틴이 시작되기 전에 작업 수를 추가함
	*/

	//wg.Add(ntasks)

	// 고루틴이 생성됨 
	taskChannel := make(chan int, ntasks) // ntasks 크기의 버퍼가 있는 정수 채널 taskChannel을 생성(이 채널로 작업 관리)

	// taskChannel에서 작업 번호를 가져와 작업을 처리
	for i := 0; i < ntasks; i++ {         // ntasks만큼의 숫자를 채널에 보냄 각 작업의 번호를 나타내고
	    taskChannel <- i
	}
	close(taskChannel)                    // 모든 작업 번호를 채널에 보낸 후 채널을 닫음

	// registerChannel는 사용 가능한 워커들을 담고 있는 채널, 하나의 워커를 받아옴
	for i := 0; i < ntasks; i++ {         // ntasks만큼의 고루틴을 생성
        
		/*
		각 고루틴마다 한 번만 호출하여 WaitGroup을 올바르게 초기화
		모든 고루틴이 종료될 때까지 기다리는 작업의 수를 추가하는 것
		
		각 고루틴이 시작될 때마다 작업수를 추가
		ntasks 만큼의 고루틴을 생성하고, 각 고루틴이 taskChannel에서 작업을 수행 
	        = 채널 버퍼에 따른 제한 없이 ntasks 만큼의 고루틴이 동시에 실행 가능
		*/
		
		wg.Add(1)                                  
        
		go func() {
            defer wg.Done()                        // 각 고루틴이 작업 완료 후 WaitGroup에 알림

            for taskNumber := range taskChannel {  // 각 고루틴은 taskChannel에서 작업을 가져옴(작업번호를 가져와 작업 수행)
                worker := <-mr.registerChannel     // 워커를 가져와 작업을 할당

                args := DoTaskArgs{
                    JobName:       mr.jobName,
                    Phase:         phase,
                    TaskNumber:    taskNumber,
                    NumOtherPhase: nios,
                    File:          mr.files[taskNumber],
                }

				// 작업이 완료될 때까지 대기
                ok := call(worker, "Worker.DoTask", &args, new(struct{}))
				
                if !ok {
                    fmt.Printf("Schedule: RPC %s DoTask error\n", worker)
                    taskChannel <- taskNumber     // 작업이 실패하면 작업을 taskChannel에 넣어 재할당
                }

                go func() {
                    mr.registerChannel <- worker // 작업이 완료되면 워커를 registerChannel로 반환
                }()
            }
        }()
    }

    wg.Wait() // 모든 고루틴이 작업을 마치면 고루틴이 종료될 때까지 대기

    // END
    debug("Schedule: %v phase done\n", phase)
}
