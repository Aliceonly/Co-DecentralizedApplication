package utils

import (
	"database/sql"
	"math/big"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql","root:121@tcp(localhost:3306)/test?parseTime=true&charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	err = Db.Ping()                                                    //连接数据库
	if err!=nil {
		fmt.Println("数据库连接失败")
}
}
//存放任务信息
func Insert(taskname string,
	add string,
	category string,
	amount int,
	timestamp *big.Int,
	state string,
	launchTime string,
	)  error {
	var err error
	var sqlstate = `SELECT MAX(id) FROM tasklist`
	rows,err :=Db.Query(sqlstate)
	var maxid int
	if err!=nil {
		panic(err)

	}
	for rows.Next() {
		err=rows.Scan(&maxid)
	}
	fmt.Println(maxid)
	var sqlStr = `INSERT INTO tasklist (id,Taskname,Beneficiary,Sponsor,Category,Amount,Timestamp,State,LaunchTime) VALUES (?,?,?,?,?,?,?,?,?)`
	_, err = Db.Exec(sqlStr,maxid+1,taskname,"0xxx000",add,category,amount,timestamp.String(),state,launchTime)
	if err != nil {
		return err
	}
 fmt.Println("总数据插入成功")
	return err
}
//查询任务数据
type tasklist struct { 
	taskid int ;// 结果集，参数名需大写
	taskname string;
	add string;
	beneficiary string;
	category string;
	amount int;
	timestamp string;
	state string;
	launchTime string;
}
func Select(add string) []tasklist {
	Tlist := make([]tasklist, 0)
	var sqlStr="SELECT * FROM tasklist where add="
	rows,err :=Db.Query(sqlStr)
	var list tasklist 
	fmt.Println("rows=====>",rows)
	if err!=nil {
		panic(err)
	}
	for rows.Next() {
		err :=rows.Scan(&list.taskid,&list.taskname,&list.beneficiary,&list.add,&list.category,&list.amount,&list.timestamp,&list.state,&list.launchTime)
		if err!=nil {
			 panic(err)
		}
		Tlist=append(Tlist, list)
		fmt.Println("任务列表为",Tlist)
	}
	 return Tlist
}


