package main

import (
  "fmt"
  "os"
  "io"
  "bufio"
//  "io/ioutil"
)

func WriteWithBufio(name,content string) {
    if fileObj,err := os.OpenFile(name,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644);err == nil {
        defer fileObj.Close()
        writeObj := bufio.NewWriterSize(fileObj,4096)
        //
//       if _,err := writeObj.WriteString(content);err == nil {
//              fmt.Println("Successful appending buffer and flush to file with bufio's Writer obj WriteString method",content)
//           }
        //使用Write方法,需要使用Writer对象的Flush方法将buffer中的数据刷到磁盘
        buf := []byte(content)
        if _,err := writeObj.Write(buf);err == nil {
            fmt.Println("Successful appending to the buffer with os.OpenFile and bufio's Writer obj Write method.",content)
            if  err := writeObj.Flush(); err != nil {panic(err)}
            fmt.Println("Successful flush the buffer data to file ",content)
        }
        }
}

func Exists(path string) bool {
        _, err := os.Stat(path)    //os.Stat获取文件信息
        if err != nil {
                if os.IsExist(err) {
                        return true
                }
                return false
        }
        return true
}

func getList(listChan chan string) {
  //将文件中的列表逐行读取并传入listChan
  fi, err := os.Open("html.txt")
  if err != nil {
    fmt.Printf("Error: %s\n", err)
  }

  br := bufio.NewReader(fi)
  for {
    f, _, c := br.ReadLine()
    if c == io.EOF {
      break
    }
    listChan<- string(f)
    fmt.Println(string(f))
  }  
  close(listChan)
}

func traversingFile(listChan chan string, traversingChan chan string, exitChan chan bool) {

  for {
    file, ok:= <-listChan
    if !ok {
      break
    }
    if Exists(file) == false {
      traversingChan<- file
    } 
  }

  fmt.Println("primeChan done")
  exitChan<- true

}

func main(){

  //开启三个协程
  listChan := make(chan string, 1000)
  traversingChan := make(chan string, 2000)
  exitChan := make(chan bool, 100)

  //将文件中的列表逐行读取并传入listChan
  go getList(listChan)

  //开启100个协程，将listChan中的文件名称出放入traversingChan中判断如果不存在放入，
  //traversingChan如果操作完成将传递给exitChan
  for i := 0; i < 100; i++ {
    go traversingFile(listChan, traversingChan, exitChan)
  }

  go func(){
    for i := 0; i < 100; i++ {
      <-exitChan
    }
    close(traversingChan)
  }()

  for {
    res, ok:= <-traversingChan
    if !ok {
      break
    }
    fmt.Printf("不存在文件=%d\n", res)
    //将指定内容写入到文件中
    WriteWithBufio("noExit.txt", res + "\n")

  }

}
