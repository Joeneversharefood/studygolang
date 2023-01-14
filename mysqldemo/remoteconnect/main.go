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

func mySelect() {

}

type Idgen struct {
	id            int    `db:"id"`
	biz           string `db:"biz"`
	partition     int    `db:"partition"`
	minWorkId     uint64 `db:"min_worker_id"`
	maxWorkId     uint64 `db:"max_work_id"`
	currentWorkId uint64 `db:"current_work_id"`
	step          int    `db:"step"`
	insertTime    uint64 `db:"insert_time"`
	updateTime    uint64 `db:"update_time"`
}

func main() {

	database, err := sqlx.Open("mysql", "uat_svr_app_wr:9e8d#A6f2b64cc5dc8@tcp(rm-f8z101r37v8997aiq.mysql.rds.aliyuncs.com:3306)/997_idgen_iris")

	DB = database

	if nil != err {
		fmt.Printf("open mysql failed, err = %s\n", err)
	} else {
		fmt.Println("connect success")
	}

	var results []Idgen

	err = database.Select(&results, "SELECT *FROM idgen_iris_0 ORDER BY `min_worker_id` ASC")
	if err != nil {
		fmt.Println("select err")
		return
	}
	fmt.Printf("result size = %d\n", len(results))
	fmt.Printf("result = %v\n", results)

	//fm := family_member{"yuanzhen", 21}
	//
	//fm.Insert()

	defer database.Close()
}
