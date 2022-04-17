package dao

import (
	"test/model"
	"test/utils"
)

//CheckUserNameAndPassword 根据用户名和密码从数据库中查询一条记录
func CheckUserNameAndPassword(username string, password string) (*model.User_chenjunjie, error) {
	//写sql语句
	sqlStr := "select id,username,password,email,groupid from users where username = ? and password = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, username, password)
	user := &model.User_chenjunjie{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.GroupID)
	return user, nil
}

//CheckEmailAndPassword 根据邮箱和密码从数据库中查询一条记录
func CheckEmailAndPassword(email string, password string) (*model.User_chenjunjie, error) {
	//写sql语句
	sqlStr := "select id,username,password,email,groupid from users where email = ? and password = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, email, password)
	user := &model.User_chenjunjie{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.GroupID)
	return user, nil
}

//CheckEmailAndPassword 根据邮箱和密码从数据库中查询一条记录
func CheckEmail(email string) (*model.User_chenjunjie, error) {
	//写sql语句
	sqlStr := "select id,username,password,email,groupid from users where email = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, email)
	user := &model.User_chenjunjie{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.GroupID)
	return user, nil
}

//CheckUserName 根据用户名和密码从数据库中查询一条记录
func CheckUserName(username string) (*model.User_chenjunjie, error) {
	//写sql语句
	sqlStr := "select id,username,password,email from users where username = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User_chenjunjie{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.GroupID)
	return user, nil
}

//SaveUser 向数据库中插入用户信息
func SaveUser(username string, password string, email string) error {
	//写sql语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, username, password, email)
	if err != nil {
		return err
	}
	return nil
}

func SaveAsAdministrator(username string, password string, email string) error {
	//写sql语句
	sqlStr := "insert into users(username,password,email,groupid) values(?,?,?,1)"
	//执行
	_, err := utils.Db.Exec(sqlStr, username, password, email)
	if err != nil {
		return err
	}
	return nil
}

func GetUserInfo() ([]*model.User_chenjunjie, error) {
	//写sql语句
	sql := "select id,username,PASSWORD,email,groupid from users"
	//执行
	rows, err := utils.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	var users []*model.User_chenjunjie
	for rows.Next() {
		user := &model.User_chenjunjie{}
		rows.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.GroupID)
		users = append(users, user)
	}
	return users, nil
}

/*
func DeleteUser(userID string) error {
	sqlStr := "delete from users where id = ?"
	_, err := utils.Db.Exec(sqlStr, userID)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByID(userID string) (*model.User_chenjunjie, error) {
	sqlStr := "select * from users where id = ?"
	row := utils.Db.QueryRow(sqlStr, userID)
	user := &model.User_chenjunjie{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.GroupID)
	return user, nil
}

func AddUser(b *model.User_chenjunjie) error {
	slqStr := "insert into users(username,PASSWORD,email) values(?,?,?)"
	_, err := utils.Db.Exec(slqStr, b.Username, b.Password, b.Email)
	if err != nil {
		return err
	}
	return nil
}

func UpdateUser(username string, password string, email string, userID int64) error {
	sqlStr := "update users set username=?,PASSWORD=?,email=? where id=?"
	_, err := utils.Db.Exec(sqlStr, username, password, email, userID)
	if err != nil {
		return err
	}
	return nil
}
*/