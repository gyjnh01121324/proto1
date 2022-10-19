package main

import (
  "fmt"
  "database/sql"

  _ "github.com/go-sql-driver/mysql"

)

func main () {fmt.Println("Работа с MySQL")
db, err := sql.Open("mysql", "mysql:mysql@tcp(127.0.0.1:3306)/golang")
if err != nil {
  panic(err)
}
defer db.Close("Подключено!!!")
}
