package main

import (
  "fmt"
  "bufio"
  "os"
  "io"
  "strings"
  "strconv"
)
type Account struct {
  AccountNo string
  Pwd string
  Balance float64
}
//存款
func (account *Account) Deposite(money float64, pwd string) {
  if pwd != account.Pwd {
    fmt.Println("输入密码不正确")
    return
  }
  if money <= 0 {
    fmt.Println("输入金额不正确")
    return
  }
  account.Balance += money
  fmt.Println("存款金额为",account.Balance)

}
//取款
func (account *Account) Withdraw(money float64, pwd string) {
  if pwd != account.Pwd {
    fmt.Println("输入密码不正确")
    return
  }
  if money <= 0 || money > account.Balance {
    fmt.Println("输入金额不正确")
    return
  }
  account.Balance -= money
  fmt.Println("取款余额为",account.Balance)

}
//查询
func (account *Account) Query(pwd string) {
  if pwd != account.Pwd {
    fmt.Println("输入密码不正确")
    return
  }
  fmt.Println("查询余额为",account.Balance)

}

//更改密码
func (account *Account) ChangePwd(pwd string, newpwd string) {
  if pwd != account.Pwd {
    fmt.Println("输入密码不正确")
    return
  }
  account.Pwd = newpwd
  fmt.Println("新密码为",account.Pwd)

}


func main() {
  dadaAccount := &Account {
    AccountNo : "rm111111" ,
    Pwd : "888888" ,
    Balance : 100.0 ,
  }
  fmt.Println("1查询，2取款，3存款，4更改秘密")
  buff := bufio.NewReader(os.Stdin)
  for {
    line, err := buff.ReadString('\n')
    if err != nil || io.EOF == err {
      break
    }
    switch line {
    case "1\n":
      fmt.Println("请输入查询密码")
      input, err:= buff.ReadString('\n')
      if err != nil {
        return
      }
      input = strings.Replace(input,"\n","",-1)
      dadaAccount.Query(input)
    case "2\n":
      fmt.Println("请输入账号密码")
      inputpwd, err:= buff.ReadString('\n')
      if err != nil {
        return
      }
      inputpwd = strings.Replace(inputpwd,"\n","",-1)
      fmt.Println("请输入取款金额")
      inputwithdraw, err:= buff.ReadString('\n')
      if err != nil {
        return
      }
      inputwithdraw = strings.Replace(inputwithdraw,"\n","",-1)
      inputfloat,err := strconv.ParseFloat(inputwithdraw,64)
      dadaAccount.Withdraw(inputfloat, inputpwd)
    case "3\n":
      fmt.Println("请输入账号密码")
      inputpwd, err:= buff.ReadString('\n')
      if err != nil {
        return
      }
      inputpwd = strings.Replace(inputpwd,"\n","",-1)
      fmt.Println("请输入存款金额")
      inputdeposite, err:= buff.ReadString('\n')
      if err != nil {
        return
      }
      inputdeposite = strings.Replace(inputdeposite,"\n","",-1)
      inputfloat,err := strconv.ParseFloat(inputdeposite,64)
      dadaAccount.Deposite(inputfloat, inputpwd)
    case "4\n":
      fmt.Println("请输入密码")
      inputpwd, err:= buff.ReadString('\n')
      if err != nil {
        return
      }
      inputpwd = strings.Replace(inputpwd,"\n","",-1)
      fmt.Println("请输入新密码")
      inputnewpwd, err:= buff.ReadString('\n')
      if err != nil {
        return
      }
      inputnewpwd = strings.Replace(inputnewpwd,"\n","",-1)
      dadaAccount.ChangePwd(inputpwd,inputnewpwd)
    default:
      fmt.Println("1查询，2取款，3存款，4更改密码")
    }
  }
}


