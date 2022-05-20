package utils

import (
	"database/sql"
	"math/big"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/protobuf/ptypes/timestamp"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:121@tcp(localhost:3306)/test?parseTime=true&charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	err = Db.Ping() //连接数据库
	if err != nil {
		fmt.Println("error <======>", err)
		fmt.Println("====>>>数据库连接失败")
	} else {
		fmt.Println("数据库连接成功")
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
) error {
	var err error
	var sqlstate = `SELECT MAX(id) FROM tasklist`
	rows, err := Db.Query(sqlstate)
	var maxid int
	if err != nil {
		panic(err)

	}
	for rows.Next() {
		err = rows.Scan(&maxid)
	}
	fmt.Println(maxid)
	var sqlStr = `INSERT INTO tasklist (id,Taskname,Beneficiary,Sponsor,Category,Amount,Timestamp,State,LaunchTime) VALUES (?,?,?,?,?,?,?,?,?)`
	_, err = Db.Exec(sqlStr, maxid+1, taskname, "0xxx000", add, category, amount, timestamp.String(), state, launchTime)
	if err != nil {
		return err
	}
	fmt.Println("总数据插入成功")
	return err
}

//查询任务数据
type tasklist struct {
	taskid      int // 结果集，参数名需大写
	taskname    string
	add         string
	beneficiary string
	category    string
	amount      int
	timestamp   string
	state       string
	launchTime  string
}

//查询自身发布的任务
func Select(add string) []tasklist {
	Tlist := make([]tasklist, 0)
	var sqlStr = "SELECT * FROM tasklist where Sponsor=?"
	rows, err := Db.Query(sqlStr, add)
	var list tasklist
	fmt.Println("rows=====>", rows)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err := rows.Scan(&list.taskid, &list.taskname, &list.beneficiary, &list.add, &list.category, &list.amount, &list.timestamp, &list.state, &list.launchTime)
		if err != nil {
			panic(err)
		}
		Tlist = append(Tlist, list)
		fmt.Println("自身发布的任务列表为", Tlist)
	}
	return Tlist
}

//自身接受任务的结构体
type accepttasklist struct {
	taskid      int    // 结果集，参数名需大写
	taskname    string //名字
	add         string
	beneficiary string
	category    string //分类
	amount      int    //佣金
	timestamp   string
	state       string
	launchTime  string //任务时间
}

type Tasklist struct {
	Taskid      int    // 结果集，参数名需大写
	Taskname    string //名字
	Add         string
	Beneficiary string
	Category    string //分类
	Amount      int    //佣金
	Timestamp   string
	State       string
	LaunchTime  string //任务时间
}

func Selectaccept(add string) []accepttasklist {
	Tlist := make([]accepttasklist, 0)
	var sqlStr = "SELECT * FROM tasklist where Beneficiary=?"
	rows, err := Db.Query(sqlStr, add)
	var list accepttasklist
	fmt.Println("rows=====>", rows)
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		err := rows.Scan(&list.taskid, &list.taskname, &list.beneficiary, &list.add, &list.category, &list.amount, &list.timestamp, &list.state, &list.launchTime)
		if err != nil {
			panic(err)
		}
		Tlist = append(Tlist, list)
		fmt.Println("自身接受任务列表为", Tlist)
	}
	return Tlist
}

func Showdata() []Tasklist {
	var task []Tasklist
	sql_task := "SELECT * FROM tasklist"
	rows, err := Db.Query(sql_task)
	if err != nil {
		fmt.Println("展示数据出错了", err)
	}
	fmt.Println("rows------------------------------------===", rows)
	for rows.Next() {
		var t Tasklist
		err := rows.Scan(&t.Taskid, &t.Taskname, &t.Add, &t.Beneficiary, &t.Category, &t.Amount, &t.Timestamp, &t.State, &t.LaunchTime)
		if err != nil {
			fmt.Println("tasklist error================================>>>>>>>>>", err)
			return nil
		}
		task = append(task, t)
	}
	return task
}

func Query_bytime(timestamp int) Tasklist {
	var task Tasklist
	err := Db.QueryRow("SELECT * FROM tasklist WHERE timestamp = ?", timestamp).Scan(&task.Taskid, &task.Taskname, &task.Add, &task.Beneficiary, &task.Category, &task.Amount, &task.Timestamp, &task.State, &task.LaunchTime)
	if err != nil {
		fmt.Println("根据时间戳查询出错了")
		fmt.Println("查询错误是====>>>>>>>>>>>>>>", err)
	}
	fmt.Println("task=====================>>>>>>>>>>>>>>>>>>>>>>>>>>>>", task)
	return task

}

func DetailData(timestamp int) Tasklist {
	var task Tasklist
	err := Db.QueryRow("SELECT * FROM tasklist WHERE timestamp = ?", timestamp).Scan(&task.Taskid, &task.Taskname, &task.Add, &task.Beneficiary, &task.Category, &task.Amount, &task.Timestamp, &task.State, &task.LaunchTime)
	if err != nil {
		fmt.Println("信息详情展示出错了")
		fmt.Println("展示错误是====>>>>>>>>>>>>>>", err)
	}
	fmt.Println("task=====================>>>>>>>>>>>>>>>>>>>>>>>>>>>>", task)
	return task
}
func DeletTask(timestamp int){
	var sqlStr = "DELETE FROM tasklist WHERE tasklist.`Timestamp`=?"
	Db.Exec(sqlStr,timestamp)
	updateid(timestamp)
	
}
func updateid(timestamp int){
    var sqlStr = "SELECT id FROM tasklist WHERE tasklist.`Timestamp`=?"
	rows, err := Db.Query(sqlStr)
	if err!= nil {
		panic(err)
	}
	for i := 0; i < count; i++ {
		
	}
}
