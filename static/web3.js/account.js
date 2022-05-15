let Account_status = window.sessionStorage.getItem("Global_Account");
if (Account_status !=""){
let Account = window.sessionStorage.getItem("Global_Account")
document.getElementById("account").innerHTML = Account;
function resume() {
    $.ajax({
        method: "post",
        url: "http://localhost:8080/dapp/user_info",
        data: { Account: Account},
        success: function (data) {
            sessionStorage.setItem('QueryUser',JSON.stringify(data.data))
            console.log(data.data)
            window.location.href="/resume"
        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
}
}else{
    swal("Sorry!", "查看失败未登录", "error");
}

function exit_login(){
    console.log("exit");
    $.ajax({
        method: "get",
        url: "/exit_account",
        success: function (data) {
            console.log("目前账户情况：",data)
        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
    
}