function Login(){
    var Sid=$("#Account").val();
    var Password=$("#password").val();
    // console.log(Sid,Password)
    $.ajax({
        method:"post",
        url: "http://localhost:8080/dapp/login",
        data: { Sid:Sid, Password: Password },
        success: function (data) {
            console.log(data)
            if (data.data==1){
                console.log("success login");
                swal("Good!", "登陆成功", "success");
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