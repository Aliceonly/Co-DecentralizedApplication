package utils

import (
	"database/sql"
	"time"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
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
//存放任务信息
func Insert(add string,number int)  error {
	var err error
	var sqlstate = `SELECT MAX(id) FROM TASKLIST`
	rows,err :=Db.Query(sqlstate)
	var maxid int
	if err!=nil {
		panic(err)

	}
	for rows.Next() {
		err=rows.Scan(&maxid)
	}
	fmt.Println(maxid)
	var sqlStr = `INSERT INTO TASKLIST (id,Taskname,Sponsor,Category,Amount,Timestamp,S) VALUES (?,?,?,?)`
	_, err = Db.Exec(sqlStr,maxid+1,add,time.Now().Format("2006-01-02 15:04:05"),number)
	if err != nil {
		return err
	}
 fmt.Println("总数据插入成功")
	return err
}


