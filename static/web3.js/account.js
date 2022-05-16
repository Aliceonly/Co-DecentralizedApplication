function demo7() {
    var
        closeInSeconds = 5,
        displayText = ' #1 秒后将自动跳转首页，点击取消即可留下',
        timer;
    swal({
        title: "退出成功!",
        text: displayText.replace(/#1/, closeInSeconds),
        imageUrl: "../static/image/check.png",
        timer: closeInSeconds * 1000,
        showCancelButton: true,
    }, function () {
            window.location.href = "/"
        }
    );
    timer = setInterval(function () {
        closeInSeconds--;
        if (closeInSeconds < 0) {
            clearInterval(timer);
        }
                
        $('.sweet-alert > p').text(displayText.replace(/#1/, closeInSeconds));
                
    }, 1000);
}

let Account = window.sessionStorage.getItem("Global_Account");
// if (Account_status !=""){

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
// }else{
    // swal("Sorry!", "查看失败未登录", "error");
// }
// 
function exit_login(){
    console.log("exit");
    $.ajax({
        method: "get",
        url: "/exit_account",
        success: function (data) {
            console.log("目前账户情况：",data)
            demo7()
        },
        error: function (data) {
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
    
}