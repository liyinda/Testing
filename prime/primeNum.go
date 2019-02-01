package main

import (
  "fmt"
)


func main(){

  var flag bool
  for i := 1; i <= 80000; i++ {
    //prime
    flag = true
    for n := 2; n < i; n++ {
      if i % n == 0 {
        flag = false
        break
      }
    }
    if flag {
      fmt.Println("素数=",i)

    }

  }

}
