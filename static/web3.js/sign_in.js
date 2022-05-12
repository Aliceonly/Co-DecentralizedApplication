var Global_Account;
function Login(){
    var Account=$("#Account").val();
    var Password=$("#password").val();
    var sign_in_up = document.querySelector("#sign_in_up");
    // console.log(Sid,Password)
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
                swal({
                    title: "登录成功",
                    text: '<span style="color:red">点击</span><a style="color:#3b3bf4" href="/account">我的</a><span style="color:red">去添加更多的信息吧！</span><br/>5秒后自动关闭。',
                    imageUrl: "../static/image/check.png",
                    html: true,
                    timer: 10000,
                    showConfirmButton: false
                });
                // in_up.style.display = none
                // swal("Good!", "登陆成功", "success");
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