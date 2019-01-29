package main

import (
  "fmt"
  "bufio"
  "os"
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
  reader := bufio.NewReader(os.Stdin)
  button, _, _ := reader.ReadLine()

  switch button {
    case "1":
      dadaAccount.Query("888888")
    case "2":
      dadaAccount.Withdraw(150, "888888")
    case "3":
      dadaAccount.Deposite(200, "888888")
    case "4":
      dadaAccount.ChangePwd("888888","123123")
    default:
      fmt.Println("1查询，2取款，3存款，4更改密码")
  }
}
