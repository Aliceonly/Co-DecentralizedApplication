package controller

import (
	"fmt"
	"dapp/dao"
	"test/model"
	"test/utils"
	"html/template"
	"net/http"
	_ "strconv"
)

var Evcode string

//解析首页
func GetHome(w http.ResponseWriter, r *http.Request) {
	//解析首页接口
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, "")
}

//Logout //处理用户注销的函数
func Logout(w http.ResponseWriter, r *http.Request) {
	//获取Cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//获取cookie的value值
		cookieValue := cookie.Value
		//删除数据库中与之对应的Session
		dao.DeleteSession(cookieValue)
		//设置cookie失效
		cookie.MaxAge = -1
		//将修改之后的cookie发送给浏览器
		http.SetCookie(w, cookie)
	}
	//返回首页
	t := template.Must(template.ParseFiles("views/index.html"))
	t.Execute(w, "")
}

func SendCode(w http.ResponseWriter, r *http.Request) {
	email := r.PostFormValue("email")
	user, _ := dao.CheckEmail(email)
	if user.ID > 0 {
		Evcode = utils.SendVcode(email)
	}
}

//Login 处理用户登录的函数
func Login(w http.ResponseWriter, r *http.Request) {
	//判断是否已经登录
	flag, _ := dao.IsLogin(r)
	if flag {
		//已经登录
		//去首页
		t := template.Must(template.ParseFiles("views/index.html"))
		t.Execute(w, "")
	} else {
		//获取邮箱和验证码
		email := r.PostFormValue("email")
		Vcode := r.PostFormValue("code")
		//调用userdao中验证用户名和密码的方法
		user, _ := dao.CheckEmail(email)
		if user.ID > 0 {
			if Vcode == "" {
				Evcode = utils.SendVcode(email)
			} else {
				//获取邮箱验证码
				strVcode := fmt.Sprintf("%v", Vcode)
				if strVcode == Evcode {
					//验证码匹配正确
					//生成UUID作为Session的id
					uuid := utils.CreateUUID()
					//创建一个Session
					sess := &model.Session_chenjunjie{
						SessionID: uuid,
						UserName:  user.Username,
						UserID:    user.ID,
						GroupID:   user.GroupID,
					}
					//将Session保存到数据库中
					dao.AddSession(sess)
					//创建一个Cookie，让它与Session相关联
					cookie := http.Cookie{
						Name:     "user",
						Value:    uuid,
						HttpOnly: true,
					}
					//将cookie发送给浏览器
					http.SetCookie(w, &cookie)
					//清楚Evcode
					Evcode = ""
					//登陆成功接口
					t := template.Must(template.ParseFiles("views/index.html"))
					t.Execute(w, user)
				} else {
					t := template.Must(template.ParseFiles("views/pages/sign-in.html"))
					t.Execute(w, "登录失败")
				}
			}
		} else {
			t := template.Must(template.ParseFiles("views/pages/sign-up.html"))
			t.Execute(w, "用户不存在")
		}
	}
}

//ELogin 处理用户邮箱登录的函数
func ELogin(w http.ResponseWriter, r *http.Request) {
	//判断是否已经登录
	flag, _ := dao.IsLogin(r)
	if flag {
		//已经登录
		//去首页
		t := template.Must(template.ParseFiles("views/index.html"))
		t.Execute(w, "")
	} else {
		//获取用户名和密码
		email := r.PostFormValue("email")
		password := r.PostFormValue("password")
		//调用userdao中验证用户名和密码的方法
		user, _ := dao.CheckEmailAndPassword(email, password)
		if user.ID > 0 {
			//用户名和密码正确
			//生成UUID作为Session的id
			uuid := utils.CreateUUID()
			//创建一个Session
			sess := &model.Session_chenjunjie{
				SessionID: uuid,
				UserName:  user.Username,
				UserID:    user.ID,
				GroupID:   user.GroupID,
			}
			//将Session保存到数据库中
			dao.AddSession(sess)
			//创建一个Cookie，让它与Session相关联
			cookie := http.Cookie{
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
			}
			//将cookie发送给浏览器
			http.SetCookie(w, &cookie)
			//登陆成功接口
			t := template.Must(template.ParseFiles("views/index.html"))
			t.Execute(w, user)
		} else {
			//用户名或密码不正确
			//登陆失败接口
			t := template.Must(template.ParseFiles("views/pages/sign-in.html"))
			t.Execute(w, "用户名或密码不正确！")
		}
	}
}

//Regist 处理用户的注册函数
func Regist(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	// //获取邮箱验证码
	// Vcode := r.PostFormValue("vcode")
	//调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		//用户名已存在
		//注册失败接口
		t := template.Must(template.ParseFiles("views/pages/sign-up.html"))
		t.Execute(w, "用户名已存在！")
	} else {
		// strVcode := fmt.Sprintf("%v", Vcode)
		// Evcode := utils.SendVcode(email)
		//如果邮箱验证码正确，将用户信息保存到数据库中
		// if strVcode == Evcode {
		dao.SaveUser(username, password, email)
		//注册成功接口
		t := template.Must(template.ParseFiles("views/pages/sign-in.html"))
		t.Execute(w, "")
	// 	// } else {
	// 		//注册失败接口
	// 		t := template.Must(template.ParseFiles("views/pages/sign-up.html"))
	// 		t.Execute(w, "邮箱验证码错误！")
	// 	}
	// }
	}
}

//RegistAsAdministrator 处理超级管理员的注册函数
func RegistAsAdministrator(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	//调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		//用户名已存在
		//注册失败接口
		t := template.Must(template.ParseFiles("views/pages/sign-up.html"))
		t.Execute(w, "用户名已存在！")
	} else {
		//用户名可用，将用户信息保存到数据库中
		dao.SaveAsAdministrator(username, password, email)
		//用户名和密码正确
		//注册成功接口
		t := template.Must(template.ParseFiles("views/pages/sign-in.html"))
		t.Execute(w, "")
	}
}

/*
//CheckUserName 通过发送Ajax验证用户名是否可用
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	//获取用户输入的用户名
	username := r.PostFormValue("username")
	//调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		//用户名已存在
		w.Write([]byte("用户名已存在！"))
	} else {
		//用户名可用
		w.Write([]byte("<font style='color:green'>用户名可用！</font>"))
	}
}

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	usersinfo, _ := dao.GetUserInfo()
	//返回接口
	t := template.Must(template.ParseFiles(""))
	t.Execute(w, usersinfo)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("userId")
	dao.DeleteUser(userID)
	GetUserInfo(w, r)
}

func ToUpdateUser(w http.ResponseWriter, r *http.Request) {
	userID := r.FormValue("userId")
	user, _ := dao.GetUserByID(userID)
	if user.ID > 0 {
		//返回接口
		t := template.Must(template.ParseFiles(""))
		t.Execute(w, user)
	} else {
		//返回接口
		t := template.Must(template.ParseFiles(""))
		t.Execute(w, "")
	}
}

func UpdateOrAddUser(w http.ResponseWriter, r *http.Request) {
	ID := r.FormValue("userId")
	user, _ := dao.GetUserByID(ID)
	username := r.PostFormValue("username")
	password := r.PostFormValue("passwd")
	email := r.PostFormValue("email")
	userID, _ := strconv.ParseInt(ID, 10, 0)
	if user.ID > 0 {
		dao.UpdateUser(username, password, email, userID)
	} else {
		dao.SaveUser(username, password, email)
	}
	GetUserInfo(w, r)
}
*/
