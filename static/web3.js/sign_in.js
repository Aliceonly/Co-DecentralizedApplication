function demo7() {
    var
        closeInSeconds = 5,
        displayText = ' #1 秒后将自动跳转首页，点击取消即可留下',
        timer;
    swal({
        title: "登陆成功!",
        text: displayText.replace(/#1/, closeInSeconds),
        imageUrl: "../static/image/check.png",
        timer: closeInSeconds * 1000,
        showCancelButton: true,
    }, function () {
            window.location.href = "/account"
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
function Login(){
    var Account=$("#Account").val();
    var Password=$("#password").val();
    var sign_in_up = document.querySelector("#sign_in_up");
    console.log(Account,Password)
    $.ajax({
        method:"post",
        url: "http://localhost:8080/dapp/login",
        data: { Account:Account, Password: Password },
        success: function (data) {
            console.log(data)
            if (data.data==1){
                window.sessionStorage.setItem("Global_Account",Account);
                // Global_Account = Account
                console.log("success login");
                demo7()
            }
            if (data.data==2){
                swal("Sorry!", "密码为空，请重试！", "error");
            }
            if (data.data==3){
                swal("Sorry!", "账号为空，请重试！", "error");
            }
            if (data.data==4){
                swal("Sorry!", "输入为空，请重试！", "error");
            }
            if (data.data==0){
                swal("OMG!", "登录失败，密码错误请重试！", "error");
                console.log("faile login")
            }
        },
        error: function (data) {
            swal("OMG!", "登录失败，请重试！", "error");
            console.log("error====>", error)
            console.log("error data===>", data)
        }
    })
}