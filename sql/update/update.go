package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:Mysql#2510@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Update
	stmt, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	stmt.Exec("Uoxiton Istive", 1)
	stmt.Exec("Sheristone Valeska", 2)

	//delete
	stmt2, _ := db.Prepare("delete from usuarios where id = ?")
	stmt2.Exec(3)
}
