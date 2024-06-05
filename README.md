> [!NOTE]  
> Distributed Systems (COS-418)
  
    
### Part 1: Map/Reduce input and output

```
cd /home/s23710660/sequential_mapreduce
export GOPATH="$PWD"
echo $GOPATH
cd src

cd mapreduce 
go mod init mapreduce
go mod tidy

cd ..
export GO111MODULE=off

go test -count=1 -run Sequential mapreduce/...
```

사진  
잘 동작함(4.515s 목표)

</br></br>




```
go test -v -run Sequential
go test -v -run Sequential mapreduce/...
```

사진  
...  
사진  
잘 동작함(4.635s 목표)  

</br></br>




## Part 2: Single-worker Word Count

```
cd "$GOPATH/src/main"
go run wc.go master sequential pg-*.txt
```

사진  
잘 동작함

</br></br>




```
sort -n -k2 mrtmp.wcseq | tail -10
```

사진  
잘 동작함

</br></br>




```
rm mrtmp.*
sh ./test-wc.sh
```

사진  
잘 동작함

</br></br>




## Part 3: Distributing MapReduce tasks

```
cd /home/s23710660/sequential_mapreduce/src/mapreduce
go test -count=1 -run TestBasic -timeout 150s mapreduce/...
```

사진  
잘 동작함(25.613s 목표)

</br></br>




## Part 4: Handling worker failures

```
# swji에서 실행
ulimit -n 16384
go test -run Failure mapreduce/...
```

잘 동작함

</br></br>




## Part 5: Inverted index generation

```
cd $GOPATH/src/main
go run ii.go master sequential pg-*.txt
head -n5 mrtmp.iiseq
```

사진  
잘 동작함

</br></br>




```
sort -k1,1 mrtmp.iiseq | sort -snk2,2 mrtmp.iiseq | grep -v '16' | tail -10
```

사진  
잘 동작함

</br></br>




```
sh ./test-ii.sh
```

사진  
잘 동작함

</br></br>



## Source Code Submission

```
# 제출
cd /home/s23710660/sequential_mapreduce/src
distsys_submit mapreduce ./

# 제출 확인
distsys_check_submission mapreduce
```

</br></br>




## Demo Video Submission
```
cd $GOPATH/src/main
ulimit -n 16384
go clean -testcache
sh test-mr.sh
```
</br></br>




# Version

https://github.com/Jinsun-Lee/distributed_systems/blob/master/sequential_mapreduce/src/mapreduce/schedule.go

```
v.3.1.6 - "schedule() - "
v.6.0.3 - " "
```


```
import (
	"time"
)

func main() {

	//start := time.Now()

	//end := time.Now()
	//elapsed := end.Sub(start)
	//fmt.Printf("Execution time: %s\n", elapsed)
}
```

</br></br>




<details>
<summary>기타</summary>
<div markdown="1">       

- https://www.cs.princeton.edu/courses/archive/fall17/cos418/a1-2.html
- https://www.cnblogs.com/lizhaolong/p/16437276.html
- https://github.com/lovesh/COS-418
- https://blog.51cto.com/u_15703183/5443253
- https://github.com/webglider/mapreduce/blob/master/schedule.go
- https://blog.csdn.net/freedom1523646952/article/details/108355990

- (참고) https://blog.csdn.net/freedom1523646952/article/details/108355990
- (참고) https://github.com/goldknife6/MIT-6.824-2016/blob/master/src/mapreduce/schedule.go

</div>
</details>
