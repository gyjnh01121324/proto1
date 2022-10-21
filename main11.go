package main

import (
  "fmt"
  "database/sql"

  _ "github.com/go-sql-driver/mysql"

)

func main () {
db, err := sql.Open("mysql", "root:roo@tcp(127.0.0.1:3306)/golang")
if err != nil {
  panic(err)
}
defer db.Close ()
//Установка данных
insert, err := db.Query("INSERT INTO `users` (`name`, `age` ) VALUES('Alex', 25)")
if err != nil {
  panic(err)
}
defer insert.Close()
 fmt.Println("Подключение к MySQL")
}
