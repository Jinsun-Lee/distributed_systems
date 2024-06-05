### Part 1: Map/Reduce input and output

```
cd /home/s23710660/sequential_mapreduce
export GOPATH="$PWD"
echo $GOPATH
cd src

cd mapreduce 
go mod init mapreduce
go mod tidy

# 다운로드한 모듈 전부 제거
# go clean -modcache

cd ..
export GO111MODULE=off

go test -count=1 -run Sequential mapreduce/...
```

![alt text](image-1.png)
잘 동작함(4.515s가 목표)

</br></br>




```
go test -v -run Sequential
go test -v -run Sequential mapreduce/...
```

![alt text](image-4.png)
...
![alt text](image-5.png)
잘 동작함(4.635s가 목표)  

![alt text](image-6.png)

원래는 위처럼 Sequential까지만 입력해도 실행되어야 하는 것 같음

</br></br>




## Part 2: Single-worker Word Count

```
cd "$GOPATH/src/main"
go run wc.go master sequential pg-*.txt
```

![alt text](image-7.png)
잘 동작함

</br></br>




```
sort -n -k2 mrtmp.wcseq | tail -10
```

![alt text](image-8.png)
잘 동작함

</br></br>




```
rm mrtmp.*
sh ./test-wc.sh
```

![alt text](image-9.png)
잘 동작함

</br></br>




## Part 3: Distributing MapReduce tasks

```
cd /home/s23710660/sequential_mapreduce/src/mapreduce
go test -count=1 -run TestBasic -timeout 150s mapreduce/...
```

![alt text](image-10.png)
잘 동작함(25.613s 목표)

</br></br>




## Part 4: Handling worker failures

```
# 지에서 실행
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

![alt text](image-11.png)
잘 동작함

</br></br>




```
sort -k1,1 mrtmp.iiseq | sort -snk2,2 mrtmp.iiseq | grep -v '16' | tail -10
```

![alt text](image-12.png)
잘 동작함

</br></br>




```
sh ./test-ii.sh
```

![alt text](image-13.png)
잘 동작함

</br></br>



## Source Code Submission

```
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

```
v.3.1.6 - "schedule() - "
https://github.com/Jinsun-Lee/distributed_systems/blob/master/sequential_mapreduce/src/mapreduce/schedule.go
```

</br></br>




<details>
<summary>기타</summary>
<div markdown="1">       

- https://www.cnblogs.com/lizhaolong/p/16437276.html
- 

</div>
</details>

---

<details>
<summary>Part1 부분 전체 내용</summary>
<div markdown="1">       

```
cd /home/s23710660/sequential_mapreduce
export GOPATH="$PWD"
echo $GOPATH
cd src

# 한 번만 실행하면 됨
cd mapreduce 
go mod init mapreduce
go mod tidy

# mapreduce 안에서 실행할 때에는 
# go mod edit -replace 
# go mod tidy

# src에서 동작시키려면
cd ..
# export GOPATH="$PWD"
# go env -w GO111MODULE=auto 

# export GO111MODULE=on   # Go Module 기능 사용
export GO111MODULE=off  # GOPATH 방법으로 동작
# export GO111MODULE=auto # src의 하위라면 GOPATH, src가 아닌 외부라면 Go Module로 동작

go test -run Sequential mapreduce/...
```
이거 되는지 확인

```
go test -v -run Sequential
```
이거 되는지 확인

</div>
</details>

</br></br>
