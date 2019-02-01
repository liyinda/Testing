package main

import (
  "fmt"
)

func putNum(intChan chan int) {
  for i := 1; i <= 80000; i++ {
    intChan<- i
  }
  close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
  //var num int
  var flag bool
  for {
    num, ok:= <-intChan
    if !ok {
      break
    }
    
    flag = true
    //判断是否是素数
    for i := 2; i < num; i++ {
      if num % i == 0 {
        flag = false
        break
      }
    }

    if flag {
      primeChan<- num
    }
 
  }

  fmt.Println("primeChan done")
  exitChan<- true

}

func main(){

  //开启三个协程
  intChan := make(chan int, 1000)
  primeChan := make(chan int, 2000)
  exitChan := make(chan bool, 4)

  //将8000个数放入intChan管道中
  go putNum(intChan)

  //开启4个协程，将intChan中8000个数取出放入primeChan中判断，
  //primeChan如果操作完成将传递给exitChan
  for i := 0; i < 4; i++ {
    go primeNum(intChan, primeChan, exitChan)
  }

  go func(){
    for i := 0; i < 4; i++ {
      <-exitChan
    }
    close(primeChan)
  }()

  for {
    res, ok:= <-primeChan
    if !ok {
      break
    }
    fmt.Printf("素数=%d\n", res)
  }

}
