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

![image](https://github.com/Jinsun-Lee/distributed_systems/assets/68187536/1fae7236-8c3f-432f-b915-66923fe13f48)

잘 동작함(4.515s 목표)

</br></br>




```
go test -v -run Sequential
go test -v -run Sequential mapreduce/...
```

![image](https://github.com/Jinsun-Lee/distributed_systems/assets/68187536/7bf9315d-836e-4131-b9a1-e65830211d0e)  
...  
![image](https://github.com/Jinsun-Lee/distributed_systems/assets/68187536/5a728f6c-a90a-4d54-b64a-d9471dfdcea7)

잘 동작함(4.635s 목표)  

</br></br>




## Part 2: Single-worker Word Count

```
cd "$GOPATH/src/main"
go run wc.go master sequential pg-*.txt
```

![image](https://github.com/Jinsun-Lee/distributed_systems/assets/68187536/bad6f327-5c51-4d34-bf52-0de9fd5e9ee1)

잘 동작함

</br></br>




```
sort -n -k2 mrtmp.wcseq | tail -10
```

![image](https://github.com/Jinsun-Lee/distributed_systems/assets/68187536/e70b4f37-14a0-4e9f-afd6-4ba440c1bcbf)

잘 동작함

</br></br>




```
rm mrtmp.*
sh ./test-wc.sh
```

![image](https://github.com/Jinsun-Lee/distributed_systems/assets/68187536/06a021cf-d801-4afd-adb9-57e110a34a15)

잘 동작함

</br></br>




## Part 3: Distributing MapReduce tasks

```
cd /home/s23710660/sequential_mapreduce/src/mapreduce
go test -count=1 -run TestBasic -timeout 150s mapreduce/...
```

![image](https://github.com/Jinsun-Lee/distributed_systems/assets/68187536/6e7d70d3-1048-4b38-a103-8d0642f2f101)

잘 동작함(25.613s 목표)

</br></br>




## Part 4: Handling worker failures

```
# swji에서 실행
ulimit -n 16384
go clean -testcache
go test -run Failure mapreduce/...
```

![image](https://github.com/Jinsun-Lee/distributed_systems/assets/68187536/a4909f9b-348e-4197-b067-7bcb1d4eaf49)

잘 동작함

</br></br>




## Part 5: Inverted index generation

```
cd $GOPATH/src/main
go run ii.go master sequential pg-*.txt
head -n5 mrtmp.iiseq
```

![image](https://github.com/Jinsun-Lee/distributed_systems/assets/68187536/cf36ad1f-afba-45a0-b09a-4efb90f839f7)

잘 동작함

</br></br>




```
sort -k1,1 mrtmp.iiseq | sort -snk2,2 mrtmp.iiseq | grep -v '16' | tail -10
```

![image](https://github.com/Jinsun-Lee/distributed_systems/assets/68187536/9d318348-810a-46db-83ea-c505b8586cc2)
 
잘 동작함

</br></br>




```
sh ./test-ii.sh
```

![image](https://github.com/Jinsun-Lee/distributed_systems/assets/68187536/a768bdcf-59d6-43bb-8526-a08e542dd6ed)

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

![image](https://github.com/Jinsun-Lee/distributed_systems/assets/68187536/484de3a1-c3e7-4b75-8679-a381c3e8860f)

</br></br>




## Demo Video Submission
```
cd $GOPATH/src/main
ulimit -n 16384
go clean -testcache
sh test-mr.sh
```

![image](https://github.com/Jinsun-Lee/distributed_systems/assets/68187536/adc895f9-3dc4-4723-bee2-cc15a4aeb481)
  
https://youtu.be/urxOYMrTiCg

잘 동작

</br></br>




## Version

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
