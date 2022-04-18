package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:121@tcp(localhost:3306)/test")
	if err != nil {
		panic(err.Error())
	}
	err = Db.Ping()                                                    //连接数据库
	if err!=nil {
		fmt.Println("数据库连接失败")
	
}
}
func SelectTasklist() int {
	var sqlStr=` SELECT COUNT(transationNumber) FROM alldate`
	//alldate,err :=db.Prepare(sqlStr)
	rows,err :=Db.Query(sqlStr)
	if err!=nil {
		panic(err)
	}
	var date int
	for rows.Next() {
		err :=rows.Scan(&date)
		if err!=nil {
			panic(err)
		}
		fmt.Println(date)
	}
	return date
}


