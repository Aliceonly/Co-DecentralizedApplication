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
//
func Insert(add string,number int)  error {
	var err error
	var sqlstate2 = `SELECT MAX(id) FROM alldate`
	rows,err :=db.Query(sqlstate2)
	var maxid int
	if err!=nil {
		panic(err)

	}
	for rows.Next() {
		err=rows.Scan(&maxid)

	}
	fmt.Println(maxid)
	var sqlStr = `INSERT INTO alldate (id,address,transationtime,transationNumber) VALUES (?,?,?,?)`
	_, err = db.Exec(sqlStr,maxid+1,add,time.Now().Format("2006-01-02 15:04:05"),number)
	if err != nil {
		return err
	}
 fmt.Println("总数据插入成功")
	return err
}

