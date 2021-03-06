package utils

import (
	"database/sql"
	"math/big"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	// "github.com/golang/protobuf/ptypes/timestamp"
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
	block string,
) error {
	var err error
	var sqlstate = `SELECT MAX(Taskid) FROM tasklist`
	rows, err := Db.Query(sqlstate)
	var maxid int
	if err != nil {
		panic(err)

	}
	for rows.Next() {
		err = rows.Scan(&maxid)
	}
	fmt.Println(maxid)
	var sqlStr = `INSERT INTO tasklist (Taskid,Taskname,Beneficiary,Sponsor,Category,Amount,Timestamp,State,LaunchTime,block) VALUES (?,?,?,?,?,?,?,?,?,?)`
	_, err = Db.Exec(sqlStr, maxid+1, taskname, "0xxx000", add, category, amount, timestamp.String(), state, launchTime, block)
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
	Block       string
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
		err := rows.Scan(&list.taskid, &list.taskname, &list.beneficiary, &list.add, &list.category, &list.amount, &list.timestamp, &list.state, &list.launchTime, &list.Block)
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
	Block       string
}

type CollectTasklist struct {
	Taskid      int    // 结果集，参数名需大写
	Taskname    string //名字
	Add         string
	Beneficiary string
	Category    string //分类
	Amount      int    //佣金
	Timestamp   string
	State       string
	LaunchTime  string //任务时间
	Block       string
	Account     string
}

type User struct {
	Sid       int
	Telephone string
	Passwd    string
	Account   string
	Sname     string
	Sage      string
	Major     string
	Grade     string
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
		err := rows.Scan(&t.Taskid, &t.Taskname, &t.Add, &t.Beneficiary, &t.Category, &t.Amount, &t.Timestamp, &t.State, &t.LaunchTime, &t.Block)
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
	err := Db.QueryRow("SELECT * FROM tasklist WHERE timestamp = ?", timestamp).Scan(&task.Taskid, &task.Taskname, &task.Add, &task.Beneficiary, &task.Category, &task.Amount, &task.Timestamp, &task.State, &task.LaunchTime, &task.Block)
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

func DeletTask(timestamp int) {
	var sqlStr1 = "SELECT Taskid FROM tasklist WHERE tasklist.`Timestamp`=?"
	rows, err := Db.Query(sqlStr1, timestamp)
	if err != nil {
		// print(err)
		fmt.Print(err)
	}
	var a int
	for rows.Next() {
		err := rows.Scan(&a)
		if err != nil {
			// panic(err)
			fmt.Print(err)
		}
	}
	var sqlStr = "DELETE FROM tasklist WHERE tasklist.`Timestamp`=?"
	_, err = Db.Exec(sqlStr, timestamp)
	if err != nil {
		// print(err)
		fmt.Print("delete err---->", err)
	}
	Updateid(a)
	fmt.Print("删除任务成功")
}
func Updateid(a int) {
	var sql = "update tasklist set Taskid=Taskid-1 where Taskid >?"
	Db.Exec(sql, a)
}

//注册

func CreateUser(sid int, tele string, pd string, account string) {
	var sql = `INSERT INTO user(Sid,Telephone,Passwd,Account) VALUES(?,?,?,?)`
	_, err = Db.Exec(sql, sid, tele, pd, account)
	if err != nil {
		panic(err)
		// fmt.Print(err)
	}
}

//登录
func Login(Account string) string {
	var sql = `select Passwd from user where Account = ?`
	var pd string
	err := Db.QueryRow(sql, Account).Scan(&pd)
	if err != nil {
		fmt.Println("登录出错了，错误是====>>>>>>>>>>>>>>", err)
	}
	// fmt.Println("user============>>>>>>>>>>>>>>>>>>>>>>>>>>>>", user)
	return pd
}

//注销账户
func CancleUser(Account string) {
	var sql = `DELETE from user where Account = ?`
	_, err := Db.Exec(sql, Account)
	if err != nil {
		fmt.Println("注销账户出错了，错误是====>>>>>>>>>>>>>>", err)
	}
}

//查询用户信息
func Query_User(Account string) User {
	var user User
	err := Db.QueryRow("SELECT Sid,Telephone,Passwd,Account,IFNULL(Sname,'未填写'),IFNULL(Sage,'未填写'),IFNULL(Major,'未填写'),IFNULL(grade,'未填写')  FROM USER WHERE Account = ?", Account).Scan(&user.Sid, &user.Telephone, &user.Passwd, &user.Account, &user.Sname, &user.Sage, &user.Major, &user.Grade)

	if err != nil {
		fmt.Println("查询用户信息====>>>>>>>>>>>>>>", err)
	}
	fmt.Println("task=====================>>>>>>>>>>>>>>>>>>>>>>>>>>>>", user)
	return user
}

//修改个人信息
func Update_User(Account string, Sid int, Sname string, Sage string, Telephone string, Major string, Grade string) {
	var sql = "UPDATE User SET Sid = ?,Sname = ?,Sage = ?,Telephone = ?,Major = ?,Grade = ? WHERE Account = ?"
	_, err := Db.Exec(sql, Sid, Sname, Sage, Telephone, Major, Grade, Account)
	if err != nil {
		panic(err)
	}
	// fmt.Println(err)
}

//查询个人发布订单的更多信息
func Self_Order_show(Account string) []Tasklist {
	var serach_task []Tasklist
	sql_serach_task := "SELECT * FROM tasklist WHERE Sponsor = ?"
	rows, err := Db.Query(sql_serach_task, Account)
	if err != nil {
		fmt.Println("查询个人订单出错了", err)
	}
	fmt.Println("rows------------------------------------===", rows)
	for rows.Next() {
		var t Tasklist
		err := rows.Scan(&t.Taskid, &t.Taskname, &t.Add, &t.Beneficiary, &t.Category, &t.Amount, &t.Timestamp, &t.State, &t.LaunchTime, &t.Block)
		if err != nil {
			fmt.Println("tasklist error================================>>>>>>>>>", err)
			return nil
		}
		serach_task = append(serach_task, t)
	}
	return serach_task

}

func Update_beneficiary(state string, account string, timestamp string) {
	var sql = "UPDATE tasklist SET beneficiary = ?,state=? WHERE timestamp = ?"
	_, err := Db.Exec(sql, account, state, timestamp)
	if err != nil {
		panic(err)
	}
}

func Query_unget_order(amount string, state string) []Tasklist {
	var serach_task []Tasklist
	sql_serach_task := "select * from tasklist where state=?  and Sponsor=?"
	rows, err := Db.Query(sql_serach_task, state, amount)
	if err != nil {
		fmt.Println("显示未接受订单出错", err)
	}
	fmt.Println("rows------------------------------------===", rows)
	for rows.Next() {
		var t Tasklist
		err := rows.Scan(&t.Taskid, &t.Taskname, &t.Add, &t.Beneficiary, &t.Category, &t.Amount, &t.Timestamp, &t.State, &t.LaunchTime, &t.Block)
		if err != nil {
			fmt.Println("tasklist error================================>>>>>>>>>", err)
			return nil
		}
		serach_task = append(serach_task, t)
	}
	return serach_task
}

//查询校园跑腿类型的订单
func Campus_order(entry string) []Tasklist {
	var serach_task []Tasklist
	sql_serach_task := "select * from tasklist where category = ?"
	rows, err := Db.Query(sql_serach_task, entry)
	if err != nil {
		fmt.Println("显示校园跑腿类型订单出错", err)
	}
	fmt.Println("rows------------------------------------===", rows)
	for rows.Next() {
		var t Tasklist
		err := rows.Scan(&t.Taskid, &t.Taskname, &t.Add, &t.Beneficiary, &t.Category, &t.Amount, &t.Timestamp, &t.State, &t.LaunchTime, &t.Block)
		if err != nil {
			fmt.Println("tasklist error================================>>>>>>>>>", err)
			return nil
		}
		serach_task = append(serach_task, t)
	}
	return serach_task
}

//查询校园兼职类型的订单
func CWork_order(entry string) []Tasklist {
	var serach_task []Tasklist
	sql_serach_task := "select * from tasklist where category = ?"
	rows, err := Db.Query(sql_serach_task, entry)
	if err != nil {
		fmt.Println("显示校园兼职类型订单出错", err)
	}
	fmt.Println("rows------------------------------------===", rows)
	for rows.Next() {
		var t Tasklist
		err := rows.Scan(&t.Taskid, &t.Taskname, &t.Add, &t.Beneficiary, &t.Category, &t.Amount, &t.Timestamp, &t.State, &t.LaunchTime, &t.Block)
		if err != nil {
			fmt.Println("tasklist error================================>>>>>>>>>", err)
			return nil
		}
		serach_task = append(serach_task, t)
	}
	return serach_task
}

//查询共享服务类型的订单
func Shared_order(entry string) []Tasklist {
	var serach_task []Tasklist
	sql_serach_task := "select * from tasklist where category = ?"
	rows, err := Db.Query(sql_serach_task, entry)
	if err != nil {
		fmt.Println("显示共享服务类型订单出错", err)
	}
	fmt.Println("rows------------------------------------===", rows)
	for rows.Next() {
		var t Tasklist
		err := rows.Scan(&t.Taskid, &t.Taskname, &t.Add, &t.Beneficiary, &t.Category, &t.Amount, &t.Timestamp, &t.State, &t.LaunchTime, &t.Block)
		if err != nil {
			fmt.Println("tasklist error================================>>>>>>>>>", err)
			return nil
		}
		serach_task = append(serach_task, t)
	}
	return serach_task
}

//查看校园跑腿类型的订单更多的情况
func Campus_order_Read_more(timestamp int) Tasklist {
	var task Tasklist
	err := Db.QueryRow("SELECT * FROM tasklist WHERE timestamp = ?", timestamp).Scan(&task.Taskid, &task.Taskname, &task.Add, &task.Beneficiary, &task.Category, &task.Amount, &task.Timestamp, &task.State, &task.LaunchTime, &task.Block)
	if err != nil {
		fmt.Println("校园跑腿类型的订单展示出错了")
		fmt.Println("展示错误是====>>>>>>>>>>>>>>", err)
	}
	fmt.Println("task=====================>>>>>>>>>>>>>>>>>>>>>>>>>>>>", task)
	return task
}

//查看校园兼职类型的订单更多的情况
func CWork_order_Read_more(timestamp int) Tasklist {
	var task Tasklist
	err := Db.QueryRow("SELECT * FROM tasklist WHERE timestamp = ?", timestamp).Scan(&task.Taskid, &task.Taskname, &task.Add, &task.Beneficiary, &task.Category, &task.Amount, &task.Timestamp, &task.State, &task.LaunchTime, &task.Block)
	if err != nil {
		fmt.Println("校园兼职类型的订单展示出错了")
		fmt.Println("展示错误是====>>>>>>>>>>>>>>", err)
	}
	fmt.Println("task=====================>>>>>>>>>>>>>>>>>>>>>>>>>>>>", task)
	return task
}

//查看共享服务类型的订单更多的情况
func Shared_order_Read_more(timestamp int) Tasklist {
	var task Tasklist
	err := Db.QueryRow("SELECT * FROM tasklist WHERE timestamp = ?", timestamp).Scan(&task.Taskid, &task.Taskname, &task.Add, &task.Beneficiary, &task.Category, &task.Amount, &task.Timestamp, &task.State, &task.LaunchTime, &task.Block)
	if err != nil {
		fmt.Println("共享服务类型的订单展示出错了")
		fmt.Println("展示错误是====>>>>>>>>>>>>>>", err)
	}
	fmt.Println("task=====================>>>>>>>>>>>>>>>>>>>>>>>>>>>>", task)
	return task
}

func Query_selfacceptOrder(add string) []Tasklist {
	var serach_task []Tasklist
	sql_serach_task := "select * from tasklist where Beneficiary = ?"
	rows, err := Db.Query(sql_serach_task, add)
	if err != nil {
		fmt.Println("显示共享服务类型订单出错", err)
	}
	fmt.Println("rows------------------------------------===", rows)
	for rows.Next() {
		var t Tasklist
		err := rows.Scan(&t.Taskid, &t.Taskname, &t.Add, &t.Beneficiary, &t.Category, &t.Amount, &t.Timestamp, &t.State, &t.LaunchTime, &t.Block)
		if err != nil {
			fmt.Println("tasklist error================================>>>>>>>>>", err)
			return nil
		}
		serach_task = append(serach_task, t)
	}
	return serach_task
}

//取消订单时更新状态
func Update_state(state string, timestamp string) {
	var sql = "UPDATE tasklist SET state=? WHERE timestamp = ?"
	_, err := Db.Exec(sql, state, timestamp)
	if err != nil {
		panic(err)
	}
}

func CreateCollectOrder(Account string, Timestamp string) int {
	var task Tasklist
	err := Db.QueryRow("SELECT * FROM tasklist WHERE timestamp = ?", Timestamp).Scan(&task.Taskid, &task.Taskname, &task.Add, &task.Beneficiary, &task.Category, &task.Amount, &task.Timestamp, &task.State, &task.LaunchTime, &task.Block)
	if err != nil {
		fmt.Println("共享服务类型的订单展示出错了")
		fmt.Println("展示错误是====>>>>>>>>>>>>>>", err)
	}
	fmt.Println("task=====================>>>>>>>>>>>>>>>>>>>>>>>>>>>>", task)

	var sql = `INSERT INTO CollectOrder  VALUES(?,?,?,?,?,?,?,?,?,?,?)`
	_, err = Db.Exec(sql, task.Taskid, task.Taskname, task.Add, task.Beneficiary, task.Category, task.Amount, Timestamp, task.State, task.LaunchTime, task.Block, Account)
	if err != nil {
		panic(err)
	}
	return 1
}

func Query_collect_order(Account string) []CollectTasklist {
	var serach_task []CollectTasklist
	sql_serach_task := "select * from CollectOrder where account = ?"
	rows, err := Db.Query(sql_serach_task, Account)
	if err != nil {
		fmt.Println("显示个人收藏订单出错", err)
	}
	fmt.Println("rows------------------------------------===", rows)
	for rows.Next() {
		var t CollectTasklist
		err := rows.Scan(&t.Taskid, &t.Taskname, &t.Add, &t.Beneficiary, &t.Category, &t.Amount, &t.Timestamp, &t.State, &t.LaunchTime, &t.Block, &t.Account)
		if err != nil {
			fmt.Println("tasklist error================================>>>>>>>>>", err)
			return nil
		}
		serach_task = append(serach_task, t)
	}
	return serach_task
}

func Cancle_CollectOrder(Timestamp string) {
	var sqlStr = "DELETE FROM collectorder  WHERE collectorder.`Timestamp`=?"
	_, err = Db.Exec(sqlStr, Timestamp)
	if err != nil {
		// print(err)
		fmt.Print("delete err---->", err)
	}

	fmt.Print("删除任务成功")
}

func Query_Dim_Order(name string) []Tasklist {
	var serach_task []Tasklist
	sql_serach_task := "select * from tasklist where taskname like ?"
	rows, err := Db.Query(sql_serach_task, name)
	if err != nil {
		fmt.Println("显示个人收藏订单出错", err)
	}
	fmt.Println("rows------------------------------------===", rows)
	for rows.Next() {
		var t Tasklist
		err := rows.Scan(&t.Taskid, &t.Taskname, &t.Add, &t.Beneficiary, &t.Category, &t.Amount, &t.Timestamp, &t.State, &t.LaunchTime, &t.Block)
		if err != nil {
			fmt.Println("tasklist error================================>>>>>>>>>", err)
			return nil
		}
		serach_task = append(serach_task, t)
	}
	return serach_task

}

//查询已接单的订单
func Query_hadAcceptOrder(account string, state string) []Tasklist {
	var serach_task []Tasklist
	sql_serach_task := "select * from tasklist where Sponsor = ?  and state = ? "
	rows, err := Db.Query(sql_serach_task, account, state)
	if err != nil {
		fmt.Println("查询已接单的订单出错", err)
	}
	fmt.Println("rows------------------------------------===", rows)
	for rows.Next() {
		var t Tasklist
		err := rows.Scan(&t.Taskid, &t.Taskname, &t.Add, &t.Beneficiary, &t.Category, &t.Amount, &t.Timestamp, &t.State, &t.LaunchTime, &t.Block)
		if err != nil {
			fmt.Println("tasklist error================================>>>>>>>>>", err)
			return nil
		}
		serach_task = append(serach_task, t)
	}
	return serach_task
}
