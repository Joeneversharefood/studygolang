package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type family_member struct {
	name string
	age  int
}

type dbexe interface {
	Insert()
}

var DB *sqlx.DB

func (fm *family_member) Insert() {
	_, err := DB.Exec("insert into my_family(name, age)values(?, ?)", fm.name, fm.age)

	if nil != err {
		fmt.Printf("insert to database failed\n")
	}
	fmt.Printf("insert into database success")
}

func main() {

	database, err := sqlx.Open("mysql", "yuanye:182550yyw@tcp(localhost:3306)/demo0")

	DB = database

	if nil != err {
		fmt.Printf("open mysql failed, err = %s\n", err)
	}

	fm := family_member{"yuanzhen", 21}

	fm.Insert()

	defer DB.Close()
}
