package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*Antes de Usar o Banco deve-se Subir o servidor
service mysqld start
*/

func exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}
	return result
}

func main() {
	db, err := sql.Open("mysql", "root:Mysql#2510@/")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Conectado ao Banco")
	exec(db, "create database if not exists cursogo")
	exec(db, "use cursogo")
	exec(db, "drop table if exists usuarios")
	exec(db, `create table usuarios (
		id integer auto_increment,
		nome varchar(80),
		PRIMARY KEY (id)
		)`)

	fmt.Println("FIM do Banco")

}
