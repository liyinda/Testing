package main
import (
    "path/filepath"
    "os"
    "fmt"
    "io/ioutil"
    "io"
    "crypto/md5"
    "flag"
    "strings"
)

var (
  filePath = flag.String("path", ".", "File Path")
  logFileName = flag.String("log", "md5sum.log", "Log File Name")
)


func WriteWithIo(name,content string) {
    fileObj,err := os.OpenFile(name,os.O_RDWR|os.O_CREATE|os.O_APPEND,0644)
    if err != nil {
        fmt.Println("Failed to open the file",err.Error())
        os.Exit(2)
    }
    if  _,err := io.WriteString(fileObj,content);err == nil {
        fmt.Println("Successful appending to the file with os.OpenFile and io.WriteString.",content)
    }
}



func md5sum(file string) string {
    data, err := ioutil.ReadFile(file)
    if err != nil {
        return ""
    }

    return fmt.Sprintf("%x", md5.Sum(data))
}

func getFilelist(path string) {
        err := filepath.Walk(path, func(path string, f os.FileInfo, err error) error {
                if ( f == nil ) {return err}
                if f.IsDir() {return nil}
                //write md5sum
                if string([]byte(path)[:1]) == "." {return nil}
                relativepath := strings.TrimLeft(path, *filePath)

                //defer WriteWithIo(*logFileName, path + "\t" + md5sum(path) + "\n")
                defer WriteWithIo(*logFileName, relativepath + "\t" + md5sum(path) + "\n")
                fmt.Println(path)
                fmt.Println(md5sum(path))
                return nil
        })
        if err != nil {
                fmt.Printf("filepath.Walk() returned %v\n", err)
        }
}


func main(){
        flag.Parse()
        getFilelist(*filePath)
}
